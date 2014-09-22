package lengthof

import "fmt"

type iField struct {
	name    string
	message string
}

func (f iField) Name() string {
	return f.name
}

func (f iField) Validator() string {
	return "LengthOf"
}

func (f iField) Error() string {
	panic("iField error() should be not be called directly")
}

// LengthOfExactError .
type LengthOfExactError struct {
	iField
	name string
	n    int64
}

func (e LengthOfExactError) Error() string {
	return fmt.Sprintf("%s must have an exact length of %d", e.name, e.n)
}

// LengthOfLessError .
type LengthOfLessError struct {
	iField
	name string
	n    int64
}

func (e LengthOfLessError) Error() string {
	return fmt.Sprintf("%s must have a length less than or equal to %d", e.name, e.n)
}

// LengthOfGreaterError .
type LengthOfGreaterError struct {
	iField
	name string
	n    int64
}

func (e LengthOfGreaterError) Error() string {
	return fmt.Sprintf("%s must have a length greater than or equal to %d", e.name, e.n)
}

// LengthOfRangeError .
type LengthOfRangeError struct {
	iField
	name string
	n    int64
	m    int64
}

func (e LengthOfRangeError) Error() string {
	return fmt.Sprintf("%s must have a length greater than or equal to %d and less than or equal to %d", e.name, e.n, e.m)
}
