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

func Example_xmlMarshalOmitEmpty() {
	s := struct {
		XMLName xml.Name      `xml:"s"`
		Bool    optional.Bool `xml:"bool,omitempty"`
		Int     optional.Int  `xml:"int,omitempty"`
	}{
		Bool: optional.EmptyBool(),
		Int:  optional.EmptyInt(),
	}

	output, _ := xml.MarshalIndent(s, "", "  ")
	fmt.Println(string(output))

	// Output:
	// <s></s>
}

func Example_xmlMarshalEmpty() {
	s := struct {
		XMLName xml.Name      `xml:"s"`
		Bool    optional.Bool `xml:"bool"`
		Int     optional.Int  `xml:"int"`
	}{
		Bool: optional.EmptyBool(),
		Int:  optional.EmptyInt(),
	}

	output, _ := xml.MarshalIndent(s, "", "  ")
	fmt.Println(string(output))

	// Output:
	// <s>
	//   <bool></bool>
	//   <int></int>
	// </s>
}

func Example_xmlMarshalPresent() {
	s := struct {
		XMLName xml.Name      `xml:"s"`
		Bool    optional.Bool `xml:"bool,omitempty"`
		Int     optional.Int  `xml:"int,omitempty"`
	}{
		Int:  optional.OfInt(0),
		Bool: optional.OfBool(false),
	}

	output, _ := xml.MarshalIndent(s, "", "  ")
	fmt.Println(string(output))

	// Output:
	// <s>
	//   <bool>false</bool>
	//   <int>0</int>
	// </s>
}

func Example_xmlUnmarshalPresent() {
	s := struct {
		XMLName xml.Name      `xml:"s"`
		Bool    optional.Bool `xml:"bool,omitempty"`
		Int     optional.Int  `xml:"int,omitempty"`
	}{}

	x := `<s>
  <bool>false</bool>
  <int>0</int>
</s>`
	xml.Unmarshal([]byte(x), &s)

	fmt.Println("Bool:", s.Bool.IsPresent(), s.Bool)
	fmt.Println("Int:", s.Int.IsPresent(), s.Int)

	// Output:
	// Bool: true false
	// Int: true 0
}

func Example_xmlUnmarshalEmpty() {
	s := struct {
		XMLName xml.Name      `xml:"s"`
		Bool    optional.Bool `xml:"bool,omitempty"`
		Int     optional.Int  `xml:"int,omitempty"`
	}{}

	x := `<s></s>`
	xml.Unmarshal([]byte(x), &s)
	fmt.Println("Bool:", s.Bool.IsPresent())
	fmt.Println("Int:", s.Int.IsPresent())

	// Output:
	// Bool: false
	// Int: false
}
