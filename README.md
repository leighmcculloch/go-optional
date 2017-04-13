# optional

[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/optional)](https://goreportcard.com/report/github.com/leighmcculloch/optional)
[![Codecov](https://img.shields.io/codecov/c/github/leighmcculloch/optional.svg)](https://codecov.io/gh/leighmcculloch/optional)
[![Build Status](https://img.shields.io/travis/leighmcculloch/optional.svg)](https://travis-ci.org/leighmcculloch/optional)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/leighmcculloch/optional)

```
go get github.com/leighmcculloch/optional
```

```
import github.com/leighmcculloch/optional
```

Package optional provides types that wrap builtin types to represent the lack of
value. The types guarantee safety by requiring the developer to unwrap them to
get to the inner value. This prevents a nil value being operated on. Optionals
marshal to XML and JSON like their underlying type, and omitempty works just
like their wrapped type would with a pointer, but without the use of pointers.

These types are an alternative to using pointers, zero values, or similar null
wrapper packages. Unlike similar solutions these will omit correctly from XML
and JSON without the use of pointers and the compiler will ensure their value is
not used when empty.

The package also contains a template that you can use with go generate to create
optional types for your own types. See below for instructions on how to use the
template.


### Examples

Take a pointer to something and make it an optional to force code you share it
with to only use it if it's not nil:

    var i *int = ...

    o := optional.OfIntPtr(i)

    o.If(func(i int) { ... }) // called if o is not empty

    _, ok := o.Get() { ... }) // ok is true if o is not empty

    _ := o.Else(100) // returns 100 if o is empty

Support XML and JSON out of the box, including omitempty without needing to use
pointers:

    s := struct {
    	Int1 optional.Int `xml:"int1,omitempty"`
    	Int2 optional.Int `xml:"int2,omitempty"`
    	Int3 optional.Int `xml:"int3,omitempty"`
    }{
    	Int1: optional.EmptyInt(),
    	Int2: optional.OfInt(1000),
    	Int3: optional.OfIntPtr(nil),
    }

    output, _ := xml.Marshal(s)

    // output = <s><int2>1000</int2></s>


### Templates

Use the Optional template for your own types by installing gotemplate.

    go get github.com/ncw/gotemplate

Then adding a go generate comment for your type.

    //go:generate gotemplate "github.com/leighmcculloch/optional/template" OptionalMyType(MyType)


### Examples

See the examples for more approaches to use.

### Documentation

See the [godoc](https://godoc.org/github.com/leighmcculloch/optional).
