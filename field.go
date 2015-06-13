package lengthof

import (
	"fmt"
)

type field struct {
	name string
}

func (f field) Name() string {
	return f.name
}

func (f field) Validator() string {
	return "LengthOf"
}

func (f field) Error() string {
	panic("field error() should be not be called directly")
}

// LengthOfExactError .
type LengthOfExactError struct {
	*field

	n int64
}

func (e LengthOfExactError) Error() string {
	return fmt.Sprintf("%s must have an exact length of %d", e.name, e.n)
}

// LengthOfLessError .
type LengthOfLessError struct {
	*field

	n int64
}

func (e LengthOfLessError) Error() string {
	return fmt.Sprintf(
		"%s must have a length less than or equal to %d", e.name, e.n)
}

// LengthOfGreaterError .
type LengthOfGreaterError struct {
	*field

	n int64
}

func (e LengthOfGreaterError) Error() string {
	return fmt.Sprintf(
		"%s must have a length greater than or equal to %d", e.name, e.n)
}

// LengthOfRangeError .
type LengthOfRangeError struct {
	*field

	n int64
	m int64
}

func (e LengthOfRangeError) Error() string {
	return fmt.Sprintf(
		"%s must have a length greater than or equal to %d and less than or "+
			"equal to %d", e.name, e.n, e.m)
}
