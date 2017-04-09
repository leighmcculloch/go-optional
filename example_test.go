package optional_test

import (
	"encoding/xml"
	"fmt"

	"github.com/leighmcculloch/optional"
)

func ExampleInt_If_present() {
	i := 1001
	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(1000),
		optional.OfIntPtr(nil),
		optional.OfIntPtr(&i),
	}

	for _, v := range values {
		v.If(func(i int) {
			fmt.Println(i)
		})
	}

	// Output:
	// 1000
	// 1001
}

func ExampleInt_Else() {
	i := 1001
	values := []optional.Int{
		optional.EmptyInt(),
		optional.OfInt(1000),
		optional.OfIntPtr(nil),
		optional.OfIntPtr(&i),
	}

	for _, v := range values {
		fmt.Println(v.Else(1))
	}

	// Output:
	// 1
	// 1000
	// 1
	// 1001
}

func ExampleInt_xml_marshal() {
	i := 1001
	v := struct {
		XMLName xml.Name     `xml:"v"`
		O1      optional.Int `xml:"o1,omitempty"`
		O2      optional.Int `xml:"o2,omitempty"`
		O3      optional.Int `xml:"o3,omitempty"`
		O4      optional.Int `xml:"o4,omitempty"`
	}{
		O1: optional.EmptyInt(),
		O2: optional.OfInt(1000),
		O3: optional.OfIntPtr(nil),
		O4: optional.OfIntPtr(&i),
	}

	output, _ := xml.MarshalIndent(v, "", "  ")
	fmt.Println(string(output))

	// Output:
	// <v>
	//   <o2>1000</o2>
	//   <o4>1001</o4>
	// </v>
}
