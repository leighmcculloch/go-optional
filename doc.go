/*
Package optional provides Optional types that wrap builtin types as an alternative to using pointers or zero values to represent the lack of something. The Optional types require conditional unwrapping to get to the inner value, which ensures by way of the compiler that you do not end up with a nil value.

Examples

Perform operations only if the optional is not empty:

	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(2017),
	}

	for _, v := range values {
		v.IfPresent(func(i int) {
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
		fmt.Println(v.OrElse(1))
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
