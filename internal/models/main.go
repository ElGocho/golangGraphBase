package models

import (
	"fmt"
	"reflect"
)


func GetID(item interface{}) (interface{}, error){
	var resp interface{}

	switch(reflect.TypeOf(item).String()){
		case "*models.Article": 
			i, OK := item.(*Article) 

			if !OK {
				return nil, fmt.Errorf("Not possible cast to struct") 
			}

			resp = i.ID
	}

	return resp, nil
}
