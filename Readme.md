# lengthof 

[![Build Status](https://travis-ci.org/gostrut/lengthof.svg?branch=master)](https://travis-ci.org/gostrut/lengthof)

Validate length of

## Example

    type Person struct {
      Name string `length_of:"3"`
    }

    val := NewValidator()
    val.Checks("length_of", lengthof.Validator)

    p := Person{Name: "Foo"}
    fields, err := val.Validates(p)

**Exact**

    Field type `length_of:"n"`

**Less than or equal**

    Field type `length_of:":n"`

**Greater than or equal**

    Field type `length_of:"n:"`

**Range**

    Field type `length_of:"n:m"`

*Range validations include `n` and `m`.*

## License

MIT