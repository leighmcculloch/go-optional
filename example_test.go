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

func Example_xmlMarshal() {
	s := struct {
		XMLName     xml.Name      `xml:"s"`
		IntEmpty    optional.Int  `xml:"int_empty,omitempty"`
		IntPresent  optional.Int  `xml:"int_present,omitempty"`
		BoolEmpty   optional.Bool `xml:"bool_empty,omitempty"`
		BoolPresent optional.Bool `xml:"bool_present,omitempty"`
	}{
		IntEmpty:    optional.EmptyInt(),
		IntPresent:  optional.OfInt(0),
		BoolEmpty:   optional.EmptyBool(),
		BoolPresent: optional.OfBool(false),
	}

	output, _ := xml.MarshalIndent(s, "", "  ")
	fmt.Println(string(output))

	// Output:
	// <s>
	//   <int_present>0</int_present>
	//   <bool_present>false</bool_present>
	// </s>
}

func Example_xmlUnmarshal() {
	s := struct {
		XMLName     xml.Name      `xml:"s"`
		IntEmpty    optional.Int  `xml:"int_empty,omitempty"`
		IntPresent  optional.Int  `xml:"int_present,omitempty"`
		BoolEmpty   optional.Bool `xml:"bool_empty,omitempty"`
		BoolPresent optional.Bool `xml:"bool_present,omitempty"`
	}{}

	x := `<s>
  <int_present>0</int_present>
  <bool_present>false</bool_present>
</s>`
	xml.Unmarshal([]byte(x), &s)

	fmt.Println("IntEmpty:", s.IntEmpty.IsPresent())
	fmt.Println("IntPresent:", s.IntPresent.IsPresent(), s.IntPresent)
	fmt.Println("BoolEmpty:", s.BoolEmpty.IsPresent())
	fmt.Println("BoolPresent:", s.BoolPresent.IsPresent(), s.BoolPresent)

	// Output:
	// IntEmpty: false
	// IntPresent: true 0
	// BoolEmpty: false
	// BoolPresent: true false
}
