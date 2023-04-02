package utils

import (
	"errors"
	"fmt"
)

func GetValueFromMap(myMap map[string]float32, key string) (float32, error) {
	value, isPresent := myMap[key]
	if isPresent {
		return value, nil
	} else {
		errMessage := fmt.Sprintf("Err: %s IS NOT PRESENT", key)
		return value, errors.New(errMessage)
	}
}
