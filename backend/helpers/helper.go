package helpers

import (
	
)

func CheckPanic(err error){
	if err != nil {
		panic(err)
	}
}