package lengthof

import "fmt"
import "testing"
import "github.com/nowk/assert"
import "github.com/gostrut/strut"

func TestLengthOfExact(t *testing.T) {
	type Address struct {
		Street string `length_of:"3"`
	}

	type Person struct {
		Addresses []Address `length_of:"2"`
	}

	a := Address{Street: "12"}
	b := Person{
		Addresses: []Address{Address{}},
	}
	blank := Address{}
	empty := Person{}

	val := strut.NewValidator()
	val.Checks("length_of", Validator)

	for _, v := range []struct {
		f string
		n int
		i interface{}
	}{
		{"Street", 3, a},
		{"Street", 3, blank},
		{"Addresses", 2, b},
		{"Addresses", 2, empty},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.False(t, fields.Valid())
		assert.Equal(t, 1, len(fields))
		f := fields[0]
		assert.Equal(t, f.Error(), fmt.Sprintf("%s must have an exact length of %d", v.f, v.n))
	}

	c := Address{Street: "123"}
	d := Person{
		Addresses: []Address{Address{}, Address{}},
	}

	for _, v := range []struct {
		i interface{}
	}{
		{c},
		{d},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.True(t, fields.Valid())
	}
}

func TestLengthOfLess(t *testing.T) {
	type Address struct {
		Street string `length_of:":3"`
	}

	type Person struct {
		Addresses []Address `length_of:":2"`
	}

	a := Address{Street: "1234"}
	b := Person{
		Addresses: []Address{Address{}, Address{}, Address{}},
	}

	val := strut.NewValidator()
	val.Checks("length_of", Validator)

	for _, v := range []struct {
		f string
		n int
		i interface{}
	}{
		{"Street", 3, a},
		{"Addresses", 2, b},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.False(t, fields.Valid())
		assert.Equal(t, 1, len(fields))
		f := fields[0]
		assert.Equal(t, f.Error(), fmt.Sprintf("%s must have a length less than or equal to %d", v.f, v.n))
	}

	c := Address{Street: "123"}
	d := Address{Street: "12"}
	e := Person{
		Addresses: []Address{Address{}, Address{}},
	}
	f := Person{
		Addresses: []Address{Address{}, Address{}},
	}
	blank := Address{}
	empty := Person{}

	for _, v := range []struct {
		i interface{}
	}{
		{c},
		{d},
		{e},
		{f},
		{blank},
		{empty},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.True(t, fields.Valid())
	}
}

func TestLengthOfGreater(t *testing.T) {
	type Address struct {
		Street string `length_of:"3:"`
	}

	type Person struct {
		Addresses []Address `length_of:"2:"`
	}

	a := Address{Street: "12"}
	b := Person{
		Addresses: []Address{Address{}},
	}
	blank := Address{}
	empty := Person{}

	val := strut.NewValidator()
	val.Checks("length_of", Validator)

	for _, v := range []struct {
		f string
		n int
		i interface{}
	}{
		{"Street", 3, a},
		{"Street", 3, blank},
		{"Addresses", 2, b},
		{"Addresses", 2, empty},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.False(t, fields.Valid())
		assert.Equal(t, 1, len(fields))
		f := fields[0]
		assert.Equal(t, f.Error(), fmt.Sprintf("%s must have a length greater than or equal to %d", v.f, v.n))
	}

	c := Address{Street: "123"}
	d := Address{Street: "1234"}
	e := Person{
		Addresses: []Address{Address{}, Address{}},
	}
	f := Person{
		Addresses: []Address{Address{}, Address{}, Address{}},
	}

	for _, v := range []struct {
		i interface{}
	}{
		{c},
		{d},
		{e},
		{f},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.True(t, fields.Valid())
	}
}

func TestLengthOfRange(t *testing.T) {
	type Address struct {
		Street string `length_of:"3:5"`
	}

	type Person struct {
		Addresses []Address `length_of:"2:3"`
	}

	a := Address{Street: "12"}
	b := Address{Street: "123456"}
	c := Person{
		Addresses: []Address{Address{}},
	}
	d := Person{
		Addresses: []Address{Address{}, Address{}, Address{}, Address{}},
	}
	blank := Address{}
	empty := Person{}

	val := strut.NewValidator()
	val.Checks("length_of", Validator)

	for _, v := range []struct {
		f string
		n int
		m int
		i interface{}
	}{
		{"Street", 3, 5, a},
		{"Street", 3, 5, b},
		{"Street", 3, 5, blank},
		{"Addresses", 2, 3, c},
		{"Addresses", 2, 3, d},
		{"Addresses", 2, 3, empty},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.False(t, fields.Valid())
		assert.Equal(t, 1, len(fields))
		f := fields[0]
		assert.Equal(t, f.Error(), fmt.Sprintf("%s must have a length greater than or equal to %d and less than or equal to %d", v.f, v.n, v.m))
	}

	e := Address{Street: "123"}
	f := Address{Street: "1234"}
	g := Address{Street: "12345"}
	h := Person{
		Addresses: []Address{Address{}, Address{}},
	}
	i := Person{
		Addresses: []Address{Address{}, Address{}, Address{}},
	}

	for _, v := range []struct {
		i interface{}
	}{
		{e},
		{f},
		{g},
		{h},
		{i},
	} {
		fields, err := val.Validates(v.i)
		assert.Nil(t, err)
		assert.True(t, fields.Valid())
	}
}

func TestLengthOfError(t *testing.T) {
	type A struct {
		V string `length_of:""` // NOTE this validation is skipped, must have a tag value
	}

	type B struct {
		V string `length_of:"a"`
	}

	type C struct {
		V string `length_of:"1:2:"`
	}

	type D struct {
		V C `length_of:"1"`
	}

	// a := A{}
	b := B{}
	c := C{}
	d := D{V: C{}}

	val := strut.NewValidator()
	val.Checks("length_of", Validator)

	for _, v := range []struct {
		i interface{}
		e string
	}{
		// {a, `strconv.ParseInt: parsing "": invalid syntax`},
		{b, `strconv.ParseInt: parsing "a": invalid syntax`},
		{c, "error: tag: unprocessable value `1:2:`"},
		{d, `reflect: call of reflect.Value.Len on struct Value`},
	} {
		_, err := val.Validates(v.i)
		assert.Equal(t, err.Error(), v.e)
	}
}
