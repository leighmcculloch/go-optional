package optional_test

import (
	"fmt"

	"github.com/leighmcculloch/optional"
)

func ExampleInt_IfPresent_present() {
	i := 1001
	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(1000),
		optional.OfIntPtr(nil),
		optional.OfIntPtr(&i),
	}

	for _, v := range values {
		v.IfPresent(func(i int) {
			fmt.Println(i)
		})
	}

	// Output:
	// 1000
	// 1001
}

func ExampleInt_OrElse() {
	i := 1001
	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(1000),
		optional.OfIntPtr(nil),
		optional.OfIntPtr(&i),
	}

	for _, v := range values {
		fmt.Println(v.OrElse(1))
	}

	// Output:
	// 1
	// 1000
	// 1
	// 1001
}
