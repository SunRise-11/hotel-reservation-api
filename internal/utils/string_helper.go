package utils

import (
	"strconv"
)

func ConvertToUint(input interface{}) (uint64, error) {

	if input == nil {
		return 0, nil
	}

	switch input.(type) {
	case string:
		returnValue, err := strconv.ParseUint(input.(string), 10, 64)
		if err != nil {
			return 0, err
		}

		return returnValue, nil
	}

	return 0, nil
}