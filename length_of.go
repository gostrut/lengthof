package lengthof

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/gostrut/strut.v1/invalid"
	. "gopkg.in/nowk/go-calm.v1"
)

// Validator of lengthof validates the length of a field
func Validator(n, t string, v *reflect.Value) (invalid.Field, error) {
	var z int64

	err := Calm(func() {
		z = int64(v.Len())
	})
	if err != nil {
		return nil, err
	}

	var sp []string
	if t != ":" {
		sp = strings.Split(t, ":")
	}

	switch len(sp) {
	case 1:
		i, err := parseInt(sp[0])
		if err != nil {
			return nil, err
		}

		if z != i {
			return LengthOfExactError{&field{n}, i}, nil
		}

	case 2:
		i, err := parseInt(sp[0])
		if err != nil {
			return nil, err
		}

		j, err := parseInt(sp[1])
		if err != nil {
			return nil, err
		}

		if i == 0 {
			if !(z <= j) {
				return LengthOfLessError{&field{n}, j}, nil
			}

			break
		}

		if j == 0 {
			if !(z >= i) {
				return LengthOfGreaterError{&field{n}, i}, nil
			}

			break
		}

		if !(z >= i && z <= j) {
			return LengthOfRangeError{&field{n}, i, j}, nil
		}

	default:
		return nil, fmt.Errorf("error: tag: unprocessable value `%s`", t)
	}

	return nil, nil
}

// parseInt is a shortcut for strconv's ParseInt function. This function will
// not return an error on "" (empty string), it returns 0 for n
func parseInt(str string) (int64, error) {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil && str != "" {
		return n, err
	}

	return n, nil
}
