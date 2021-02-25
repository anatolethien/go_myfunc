package myfunc

import (
	"errors"
)

func Atoi(s string) (int, error) {

	var f = s[0]
	var n int

	if f == '-' || f == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0, errors.New("invalid string")
		}

	}

	for _, v := range s {
		v -= '0'
		if v > 9 {
			return 0, errors.New("invalid string")
		}
		n = n*10 + int(v)
	}

	if f == '-' {
		n = -n
	}

	return n, nil
}
