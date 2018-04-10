package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

const defalultMaxRecords = 1000

/*QueryAllPandas return All Pandas*/
func QueryAllPandas() ([]Panda, int, error) {
	SQLConnect()
	stm := fmt.Sprintf("SELECT * from %s ORDER BY %s",
		beego.AppConfig.String("pandatable"), beego.AppConfig.String("sortcolumn"))
	rows, err := sqldb.Query(stm)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	maxrecords, err := beego.AppConfig.Int("maxRecords")
	if err != nil {
		maxrecords = defalultMaxRecords // Default Value
	}
	pandas := make([]Panda, 0, maxrecords)
	num := 0
	for num = 0; rows.Next(); num++ {
		panda := Panda{}
		err := rows.Scan(&panda.PandaIndex, &panda.Genes, &panda.Birthtime, &panda.Cooldown, &panda.Rank, &panda.MotherID,
			&panda.FatherID, &panda.Generation, &panda.Owner, &panda.Ownername, &panda.Photourl, &panda.Snapurl)
		if err != nil {
			log.Fatal(err)
			return nil, num, err
		}
		pandas = append(pandas, panda)
	}
	if err := rows.Err(); err != nil {
		return nil, num, err
	}
	return pandas, num, nil
}

/*QueryPandaByPage return All Pandas*/
func QueryPandaByPage(page int) ([]Panda, int, error) {
	SQLConnect()
	pagesize, err := beego.AppConfig.Int("pagesize")
	if err != nil {
		return nil, 0, err
	}
	stm := fmt.Sprintf("SELECT * from %s WHERE pandaIndex BETWEEN ? AND ? ORDER BY %s",
		beego.AppConfig.String("pandatable"), beego.AppConfig.String("sortcolumn"))
	rows, err := sqldb.Query(stm, (page-1)*pagesize, page*pagesize-1)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	pandas := make([]Panda, 0)
	num := 0
	for num = 0; rows.Next(); num++ {
		panda := Panda{}
		err := rows.Scan(&panda.PandaIndex, &panda.Genes, &panda.Birthtime, &panda.Cooldown, &panda.Rank, &panda.MotherID,
			&panda.FatherID, &panda.Generation, &panda.Owner, &panda.Ownername, &panda.Photourl, &panda.Snapurl)
		if err != nil {
			log.Fatal(err)
			return nil, num, err
		}
		pandas = append(pandas, panda)
	}
	if err := rows.Err(); err != nil {
		return nil, num, err
	}
	return pandas, num, nil
}

/*QueryPandaByIndex return All Pandas*/
func QueryPandaByIndex(idx int) ([]Panda, int, error) {
	SQLConnect()

	stm := fmt.Sprintf("SELECT * from %s WHERE pandaIndex=?", beego.AppConfig.String("pandatable"))
	rows, err := sqldb.Query(stm, idx)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	pandas := make([]Panda, 0)
	num := 0
	for num = 0; rows.Next(); num++ {
		panda := Panda{}
		err := rows.Scan(&panda.PandaIndex, &panda.Genes, &panda.Birthtime, &panda.Cooldown, &panda.Rank, &panda.MotherID,
			&panda.FatherID, &panda.Generation, &panda.Owner, &panda.Ownername, &panda.Photourl, &panda.Snapurl)
		if err != nil {
			log.Fatal(err)
			return nil, 0, err
		}
		pandas = append(pandas, panda)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return pandas, num, nil
}

/*QueryRowCount query number of rows in a table*/
func QueryRowCount(table string) (int, error) {
	SQLConnect()
	var count int
	log.Println(table)
	row := sqldb.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table))
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/*SetNameByOwner set the new name for the owner*/
func SetNameByOwner(name string, owner []byte) error {
	SQLConnect()
	//只執行一次不做 prepare
	_, err := sqldb.Exec(`UPDATE panda SET ownername=? WHERE BINARY owner=?`, name, owner)
	if err != nil {
		return err
	}
	return nil
}

/*SetNumInPage set number of Record in one page*/
func SetNumInPage(num int) bool {
	beego.AppConfig.Set("pagesize", strconv.Itoa(num))
	if beego.AppConfig.String("pagesize") == strconv.Itoa(num) {
		log.Println("pagesize is set to ", beego.AppConfig.String("pagesize"))
		return true
	}
	return false
}

/*SetURLByIndex set the new name for the owner*/
func SetURLByIndex(index int, url string) error {
	SQLConnect()
	//只執行一次不做 prepare
	_, err := sqldb.Exec(`UPDATE panda SET photourl=? WHERE pandaIndex=?`, url, index)
	if err != nil {
		return err
	}
	return nil
}
