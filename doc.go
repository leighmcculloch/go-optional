/*
Package optional provides types that wrap builtin types to represent the lack of value. The types guarantee safety by requiring the developer to unwrap them to get to the inner value. This ensures ensures a nil value is never operated on. Optionals marshal to text, XML and JSON like their underlying type, and omitempty works just like their wrapped type without the use of pointers.

These types are are an alternative to using pointers, zero values, or similar null wrapper packages. Unlike similar solutions these omitempty without the use of pointers and ensure their value is not used when empty.

The package also contains a template that you can use with go generate to create optional types for your own types. See below for instructions on how to use the template.

Examples

Take a pointer to something and make it an optional to force code you share it with to only use it if it's not nil:

	var i *int = ...

	o := optional.OfIntPtr(i)

	o.If(func(i int) {
		// called if o is not empty
	})

	// returns 100 if o is empty
	_ := o.Else(100)
	_ := o.ElseFunc(func() { return 100; })

Support XML, JSON and other encoding/decoding out of the box, including omitempty without needing to use pointers:

	s := struct {
		F1 optional.Int `xml:"f1,omitempty"`
		F2 optional.Int `xml:"f2,omitempty"`
		F3 optional.Int `xml:"f3,omitempty"`
	}{
		F1: optional.EmptyInt(),
		F2: optional.OfInt(1000),
		F3: optional.OfIntPtr(nil),
	}

	output, _ := xml.Marshal(s)

	// output = <v><o2>1000</o2></v>

Templates

Use the Optional template for your own types by installing gotemplate.

	go get github.com/ncw/gotemplate

Then adding a go generate comment for your type.

	//go:generate gotemplate "github.com/leighmcculloch/optional/template" OptionalMyType(MyType)

*/
package optional
