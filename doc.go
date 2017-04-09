/*
Package optional provides types that wrap builtin types as a alternative to using pointers or zero values to represent the lack of value. The Optional types require the developer to unwrap them to get to the inner value, which ensures a nil value doesn't get operated on by way of the compiler. Optionals marshal to text, XML and JSON like their underlying type, and omitempty when the omitempty tag option is specified.

The package also contains a template that you can use with go generate to create optional types for your own types. See below for instructions on how to use the template.

Examples

Take a pointer to something and make it an optional to force code you share it with to only use it if it's not nil:

	var i *int = ...

	o := optional.OfIntPtr(i)

	o.If(func(i int) {
		// only called if i was not originally nil
		// use i here
	})

Support XML, JSON and other encoding packages out of the box, including omitempty without needing to use pointers:

	s := struct {
		F1      optional.Int `xml:"f1,omitempty"`
		F2      optional.Int `xml:"f2,omitempty"`
		F3      optional.Int `xml:"f3,omitempty"`
	}{
		F1: optional.EmptyInt(),
		F2: optional.OfInt(1000),
		F3: optional.OfIntPtr(nil),
	}

	output, _ := xml.Marshal(s)

	// output = <v><o2>1000</o2></v>

Perform operations only if the optional is not empty:

	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(2017),
	}

	for _, v := range values {
		v.If(func(i int) {
			fmt.Println(i)
		})
	}

	// Output:
	// 2017

Perform operations using an optional with a default:

	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(2016),
	}

	for _, v := range values {
		fmt.Println(v.Else(1))
	}

	// Output:
	// 1
	// 2017

Templates

Use the Optional template for your own types by installing gotemplate.

	go get github.com/ncw/gotemplate

Then adding a go generate comment for your type.

	//go:generate gotemplate "github.com/leighmcculloch/optional/template" OptionalMyType(MyType)

*/
package optional
