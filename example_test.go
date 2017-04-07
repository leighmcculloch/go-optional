package optional_test

import (
	"fmt"

	"github.com/leighmcculloch/optional"
)

func ExampleInt_IfPresent_present() {
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
}

func ExampleInt_OrElse() {
	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(2017),
	}

	for _, v := range values {
		fmt.Println(v.OrElse(1))
	}

	// Output:
	// 1
	// 2017
}
