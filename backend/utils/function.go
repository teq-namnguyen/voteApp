package utils

import "strconv"

func StrArrContains(s []string, e string) bool {
	for i := range s {
		if s[i] == e {
			return true
		}
	}

	return false
}

func StrToInt(i string) (*int, error) {
	o, err := strconv.Atoi(i)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
