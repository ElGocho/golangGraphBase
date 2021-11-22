package helpers

import (
)


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
