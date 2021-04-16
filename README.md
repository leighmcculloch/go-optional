# optional

[![Go Report Card](https://goreportcard.com/badge/4d63.com/optional)](https://goreportcard.com/report/4d63.com/optional)
[![Build Status](https://img.shields.io/travis/leighmcculloch/go-optional.svg)](https://travis-ci.org/leighmcculloch/go-optional)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/4d63.com/optional)

```
go get 4d63.com/optional
```

```
import 4d63.com/optional
```

Package optional exports types that wrap the builtin types (int, bool, etc) to
represent the lack of value. The types guarantee safety by requiring the
developer to unwrap them to get to the inner value. This prevents a nil value
being operated on. Optionals marshal to XML and JSON like their underlying type,
and omitempty works just like their wrapped type would with a pointer, but
without the use of pointers.

These types are an alternative to using pointers, zero values, or similar null
wrapper packages. Unlike similar solutions these will omit correctly from XML
and JSON without the use of pointers and the compiler will ensure their value is
not used when empty.

The package also contains a template that you can use with go generate to create
optional types for your own types. See below for instructions on how to use the
template.


### Examples

Wrap a pointer in an optional:

    var i *int = ...
    o := optional.OfIntPtr(i)

Unwrap it safely:

    o.If(func(i int) {
    	// called if o is not empty
    })

    if _, ok := o.Get(); ok {
    	// ok is true if o is not empty
    }

Or get it's value with a fallback to a default:

    _ := o.ElseZero() // returns the zero value if empty

    _ := o.Else(100) // returns 100 if o is empty

    _ := o.ElseFunc(func() {
    	// called if o is empty
    	return 100
    })

XML and JSON are supported out of the box. Use `omitempty` to omit the field
when the optional is empty:

    s := struct {
    	Int1 optional.Int `json:"int1,omitempty"`
    	Int2 optional.Int `json:"int2,omitempty"`
    	Int3 optional.Int `json:"int3,omitempty"`
    }{
    	Int1: optional.EmptyInt(),
    	Int2: optional.OfInt(1000),
    	Int3: optional.OfIntPtr(nil),
    }

    output, _ := json.Marshal(s)

    // output = {"int2":1000}


### Templates

Use the Optional template for your own types by installing gotemplate.

    go get github.com/ncw/gotemplate

Then add a `go generate` comment for your type to any `.go` file in your
package.

    //go:generate gotemplate "4d63.com/optional/template" OptionalMyType(MyType)


### Examples

See the examples for more approaches to use.

### Documentation

See the [godoc](https://godoc.org/4d63.com/optional).
