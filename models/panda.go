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
	numPage, err := beego.AppConfig.Int("numRecordsInPage")
	if err != nil {
		return nil, 0, err
	}
	stm := fmt.Sprintf("SELECT * from %s WHERE pandaIndex BETWEEN ? AND ? ORDER BY %s",
		beego.AppConfig.String("pandatable"), beego.AppConfig.String("sortcolumn"))
	rows, err := sqldb.Query(stm, (page-1)*numPage, page*numPage-1)
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
	beego.AppConfig.Set("numRecordsInPage", strconv.Itoa(num))
	if beego.AppConfig.String("numRecordsInPage") == strconv.Itoa(num) {
		log.Println("numRecordsInPage is set to ", beego.AppConfig.String("numRecordsInPage"))
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
