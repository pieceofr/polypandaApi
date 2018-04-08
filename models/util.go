package models

import (
	"strconv"
)

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
