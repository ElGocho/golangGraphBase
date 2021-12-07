package helpers

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
			log.Fatalf("%s: %s", msg, err)	
	}
}

func IsNil(s interface{}) bool{
	return s == nil
}

func FindString(array []string, search string) int{
	for index,elem := range array{
		if elem == search{
			return index
		}
	}

	return -1
}
