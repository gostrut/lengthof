package lengthof

import "fmt"
import "reflect"
import "strconv"
import "strings"
import "github.com/gostrut/invalid"
import . "github.com/nowk/go-calm"

// Validator of lengthof validates the length of a field
func Validator(name, tagStr string, vo *reflect.Value) (invalid.Field, error) {
	var z int64
	if err := Calm(func() {
		z = int64(vo.Len())
	}); err != nil {
		return nil, err
	}

	var split []string
	if tagStr != ":" {
		split = strings.Split(tagStr, ":")
	}

	switch len(split) {
	case 1:
		n, err := parseInt(split[0])
		if err != nil {
			return nil, err
		}

		if z != n {
			return LengthOfExactError{name: name, n: n}, nil
		}

	case 2:
		i, err := parseInt(split[0])
		if err != nil {
			return nil, err
		}

		n, err := parseInt(split[1])
		if err != nil {
			return nil, err
		}

		if i == 0 {
			if !(z <= n) {
				return LengthOfLessError{name: name, n: n}, nil
			}
			break
		}

		if n == 0 {
			if !(z >= i) {
				return LengthOfGreaterError{name: name, n: i}, nil
			}
			break
		}

		if !(z >= i && z <= n) {
			return LengthOfRangeError{name: name, n: i, m: n}, nil
		}

	default:
		return nil, fmt.Errorf("error: tag: unprocessable value `%s`", tagStr)
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
