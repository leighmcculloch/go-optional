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
	s := struct {
		XMLName xml.Name     `xml:"s"`
		F1      optional.Int `xml:"f1,omitempty"`
		F2      optional.Int `xml:"f2,omitempty"`
		F3      optional.Int `xml:"f3,omitempty"`
	}{
		F1: optional.EmptyInt(),
		F2: optional.OfInt(1000),
		F3: optional.OfIntPtr(nil),
	}

	output, _ := xml.MarshalIndent(s, "", "  ")
	fmt.Println(string(output))

	// Output:
	// <s>
	//   <f2>1000</f2>
	// </s>
}

func ExampleInt_xml_unmarshal() {
	s := struct {
		XMLName xml.Name     `xml:"s"`
		F1      optional.Int `xml:"f1,omitempty"`
		F2      optional.Int `xml:"f2,omitempty"`
		F3      optional.Int `xml:"f3,omitempty"`
	}{}

	xml.Unmarshal([]byte(`<s><f2>1000</f2></s>`), &s)
	fmt.Printf("F1: %s, F2: %s, F3: %s", s.F1, s.F2, s.F3)

	// Output:
	// F1: , F2: 1000, F3:
}
