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

/*ConvtNumGRTZero convert to integer if it is greater than zero */
func ConvtNumGRTZero(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil || num == 0 {
		return -1
	}
	return num
}

/*ConvtNumGRTEqZero convert to integer if number greater or equall to zero*/
func ConvtNumGRTEqZero(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return num
}
