package models

import (
	"log"
	"regexp"
	"strconv"
)

/*IsValidOwnername check is name the valid ownername in which every single byte is a character */
func IsValidOwnername(name string) bool {
	r, err := regexp.Compile(`\w+`)
	if err != nil {
		log.Println("[Error]ValidOwnername err:", name)
		return false
	}
	return r.MatchString(name)
}

/*ConvertPageNum convert input string to number and if it is a a valid num, return -1*/
func ConvertPageNum(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return num
}

/*IsValidPageSize convert input string to number and if it is a a valid num, return -1*/
func IsValidPageSize(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return num
}
