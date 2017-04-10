

# optional
`import "github.com/leighmcculloch/optional"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package optional provides types that wrap builtin types as a alternative to using pointers or zero values to represent the lack of value. The Optional types require the developer to unwrap them to get to the inner value, which ensures a nil value doesn't get operated on by way of the compiler. Optionals marshal to text, XML and JSON like their underlying type, and omitempty when the omitempty tag option is specified.

The package also contains a template that you can use with go generate to create optional types for your own types. See below for instructions on how to use the template.

### Examples
Take a pointer to something and make it an optional to force code you share it with to only use it if it's not nil:


	var i *int = ...
	
	o := optional.OfIntPtr(i)
	
	o.If(func(i int) {
		// only called if i was not originally nil
		// use i here
	})
	
	// returns 100 if o is empty
	_ := o.Else(100)
	_ := o.ElseFunc(func() { return 100; })

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

### Templates
Use the Optional template for your own types by installing gotemplate.


	go get github.com/ncw/gotemplate

Then adding a go generate comment for your type.


	//go:generate gotemplate "github.com/leighmcculloch/optional/template" OptionalMyType(MyType)




## <a name="pkg-index">Index</a>
* [type Bool](#Bool)
  * [func EmptyBool() Bool](#EmptyBool)
  * [func OfBool(value bool) Bool](#OfBool)
  * [func OfBoolPtr(ptr *bool) Bool](#OfBoolPtr)
  * [func (o Bool) Else(elseValue bool) (value bool)](#Bool.Else)
  * [func (o Bool) ElseFunc(f func() bool) (value bool)](#Bool.ElseFunc)
  * [func (o Bool) If(f func(value bool))](#Bool.If)
  * [func (o Bool) IsEmpty() bool](#Bool.IsEmpty)
  * [func (o Bool) IsPresent() bool](#Bool.IsPresent)
  * [func (o Bool) MarshalText() (text []byte, err error)](#Bool.MarshalText)
  * [func (o Bool) String() string](#Bool.String)
  * [func (o *Bool) UnmarshalText(text []byte) error](#Bool.UnmarshalText)
* [type Byte](#Byte)
  * [func EmptyByte() Byte](#EmptyByte)
  * [func OfByte(value byte) Byte](#OfByte)
  * [func OfBytePtr(ptr *byte) Byte](#OfBytePtr)
  * [func (o Byte) Else(elseValue byte) (value byte)](#Byte.Else)
  * [func (o Byte) ElseFunc(f func() byte) (value byte)](#Byte.ElseFunc)
  * [func (o Byte) If(f func(value byte))](#Byte.If)
  * [func (o Byte) IsEmpty() bool](#Byte.IsEmpty)
  * [func (o Byte) IsPresent() bool](#Byte.IsPresent)
  * [func (o Byte) String() string](#Byte.String)
* [type Complex128](#Complex128)
  * [func EmptyComplex128() Complex128](#EmptyComplex128)
  * [func OfComplex128(value complex128) Complex128](#OfComplex128)
  * [func OfComplex128Ptr(ptr *complex128) Complex128](#OfComplex128Ptr)
  * [func (o Complex128) Else(elseValue complex128) (value complex128)](#Complex128.Else)
  * [func (o Complex128) ElseFunc(f func() complex128) (value complex128)](#Complex128.ElseFunc)
  * [func (o Complex128) If(f func(value complex128))](#Complex128.If)
  * [func (o Complex128) IsEmpty() bool](#Complex128.IsEmpty)
  * [func (o Complex128) IsPresent() bool](#Complex128.IsPresent)
  * [func (o Complex128) String() string](#Complex128.String)
* [type Complex64](#Complex64)
  * [func EmptyComplex64() Complex64](#EmptyComplex64)
  * [func OfComplex64(value complex64) Complex64](#OfComplex64)
  * [func OfComplex64Ptr(ptr *complex64) Complex64](#OfComplex64Ptr)
  * [func (o Complex64) Else(elseValue complex64) (value complex64)](#Complex64.Else)
  * [func (o Complex64) ElseFunc(f func() complex64) (value complex64)](#Complex64.ElseFunc)
  * [func (o Complex64) If(f func(value complex64))](#Complex64.If)
  * [func (o Complex64) IsEmpty() bool](#Complex64.IsEmpty)
  * [func (o Complex64) IsPresent() bool](#Complex64.IsPresent)
  * [func (o Complex64) String() string](#Complex64.String)
* [type Float32](#Float32)
  * [func EmptyFloat32() Float32](#EmptyFloat32)
  * [func OfFloat32(value float32) Float32](#OfFloat32)
  * [func OfFloat32Ptr(ptr *float32) Float32](#OfFloat32Ptr)
  * [func (o Float32) Else(elseValue float32) (value float32)](#Float32.Else)
  * [func (o Float32) ElseFunc(f func() float32) (value float32)](#Float32.ElseFunc)
  * [func (o Float32) If(f func(value float32))](#Float32.If)
  * [func (o Float32) IsEmpty() bool](#Float32.IsEmpty)
  * [func (o Float32) IsPresent() bool](#Float32.IsPresent)
  * [func (o Float32) String() string](#Float32.String)
* [type Float64](#Float64)
  * [func EmptyFloat64() Float64](#EmptyFloat64)
  * [func OfFloat64(value float64) Float64](#OfFloat64)
  * [func OfFloat64Ptr(ptr *float64) Float64](#OfFloat64Ptr)
  * [func (o Float64) Else(elseValue float64) (value float64)](#Float64.Else)
  * [func (o Float64) ElseFunc(f func() float64) (value float64)](#Float64.ElseFunc)
  * [func (o Float64) If(f func(value float64))](#Float64.If)
  * [func (o Float64) IsEmpty() bool](#Float64.IsEmpty)
  * [func (o Float64) IsPresent() bool](#Float64.IsPresent)
  * [func (o Float64) String() string](#Float64.String)
* [type Int](#Int)
  * [func EmptyInt() Int](#EmptyInt)
  * [func OfInt(value int) Int](#OfInt)
  * [func OfIntPtr(ptr *int) Int](#OfIntPtr)
  * [func (o Int) Else(elseValue int) (value int)](#Int.Else)
  * [func (o Int) ElseFunc(f func() int) (value int)](#Int.ElseFunc)
  * [func (o Int) If(f func(value int))](#Int.If)
  * [func (o Int) IsEmpty() bool](#Int.IsEmpty)
  * [func (o Int) IsPresent() bool](#Int.IsPresent)
  * [func (o Int) MarshalText() (text []byte, err error)](#Int.MarshalText)
  * [func (o Int) String() string](#Int.String)
  * [func (o *Int) UnmarshalText(text []byte) error](#Int.UnmarshalText)
* [type Int16](#Int16)
  * [func EmptyInt16() Int16](#EmptyInt16)
  * [func OfInt16(value int16) Int16](#OfInt16)
  * [func OfInt16Ptr(ptr *int16) Int16](#OfInt16Ptr)
  * [func (o Int16) Else(elseValue int16) (value int16)](#Int16.Else)
  * [func (o Int16) ElseFunc(f func() int16) (value int16)](#Int16.ElseFunc)
  * [func (o Int16) If(f func(value int16))](#Int16.If)
  * [func (o Int16) IsEmpty() bool](#Int16.IsEmpty)
  * [func (o Int16) IsPresent() bool](#Int16.IsPresent)
  * [func (o Int16) String() string](#Int16.String)
* [type Int32](#Int32)
  * [func EmptyInt32() Int32](#EmptyInt32)
  * [func OfInt32(value int32) Int32](#OfInt32)
  * [func OfInt32Ptr(ptr *int32) Int32](#OfInt32Ptr)
  * [func (o Int32) Else(elseValue int32) (value int32)](#Int32.Else)
  * [func (o Int32) ElseFunc(f func() int32) (value int32)](#Int32.ElseFunc)
  * [func (o Int32) If(f func(value int32))](#Int32.If)
  * [func (o Int32) IsEmpty() bool](#Int32.IsEmpty)
  * [func (o Int32) IsPresent() bool](#Int32.IsPresent)
  * [func (o Int32) String() string](#Int32.String)
* [type Int64](#Int64)
  * [func EmptyInt64() Int64](#EmptyInt64)
  * [func OfInt64(value int64) Int64](#OfInt64)
  * [func OfInt64Ptr(ptr *int64) Int64](#OfInt64Ptr)
  * [func (o Int64) Else(elseValue int64) (value int64)](#Int64.Else)
  * [func (o Int64) ElseFunc(f func() int64) (value int64)](#Int64.ElseFunc)
  * [func (o Int64) If(f func(value int64))](#Int64.If)
  * [func (o Int64) IsEmpty() bool](#Int64.IsEmpty)
  * [func (o Int64) IsPresent() bool](#Int64.IsPresent)
  * [func (o Int64) String() string](#Int64.String)
* [type Int8](#Int8)
  * [func EmptyInt8() Int8](#EmptyInt8)
  * [func OfInt8(value int8) Int8](#OfInt8)
  * [func OfInt8Ptr(ptr *int8) Int8](#OfInt8Ptr)
  * [func (o Int8) Else(elseValue int8) (value int8)](#Int8.Else)
  * [func (o Int8) ElseFunc(f func() int8) (value int8)](#Int8.ElseFunc)
  * [func (o Int8) If(f func(value int8))](#Int8.If)
  * [func (o Int8) IsEmpty() bool](#Int8.IsEmpty)
  * [func (o Int8) IsPresent() bool](#Int8.IsPresent)
  * [func (o Int8) String() string](#Int8.String)
* [type Rune](#Rune)
  * [func EmptyRune() Rune](#EmptyRune)
  * [func OfRune(value rune) Rune](#OfRune)
  * [func OfRunePtr(ptr *rune) Rune](#OfRunePtr)
  * [func (o Rune) Else(elseValue rune) (value rune)](#Rune.Else)
  * [func (o Rune) ElseFunc(f func() rune) (value rune)](#Rune.ElseFunc)
  * [func (o Rune) If(f func(value rune))](#Rune.If)
  * [func (o Rune) IsEmpty() bool](#Rune.IsEmpty)
  * [func (o Rune) IsPresent() bool](#Rune.IsPresent)
  * [func (o Rune) String() string](#Rune.String)
* [type String](#String)
  * [func EmptyString() String](#EmptyString)
  * [func OfString(value string) String](#OfString)
  * [func OfStringPtr(ptr *string) String](#OfStringPtr)
  * [func (o String) Else(elseValue string) (value string)](#String.Else)
  * [func (o String) ElseFunc(f func() string) (value string)](#String.ElseFunc)
  * [func (o String) If(f func(value string))](#String.If)
  * [func (o String) IsEmpty() bool](#String.IsEmpty)
  * [func (o String) IsPresent() bool](#String.IsPresent)
  * [func (o String) String() string](#String.String)
* [type Uint](#Uint)
  * [func EmptyUint() Uint](#EmptyUint)
  * [func OfUint(value uint) Uint](#OfUint)
  * [func OfUintPtr(ptr *uint) Uint](#OfUintPtr)
  * [func (o Uint) Else(elseValue uint) (value uint)](#Uint.Else)
  * [func (o Uint) ElseFunc(f func() uint) (value uint)](#Uint.ElseFunc)
  * [func (o Uint) If(f func(value uint))](#Uint.If)
  * [func (o Uint) IsEmpty() bool](#Uint.IsEmpty)
  * [func (o Uint) IsPresent() bool](#Uint.IsPresent)
  * [func (o Uint) String() string](#Uint.String)
* [type Uint16](#Uint16)
  * [func EmptyUint16() Uint16](#EmptyUint16)
  * [func OfUint16(value uint16) Uint16](#OfUint16)
  * [func OfUint16Ptr(ptr *uint16) Uint16](#OfUint16Ptr)
  * [func (o Uint16) Else(elseValue uint16) (value uint16)](#Uint16.Else)
  * [func (o Uint16) ElseFunc(f func() uint16) (value uint16)](#Uint16.ElseFunc)
  * [func (o Uint16) If(f func(value uint16))](#Uint16.If)
  * [func (o Uint16) IsEmpty() bool](#Uint16.IsEmpty)
  * [func (o Uint16) IsPresent() bool](#Uint16.IsPresent)
  * [func (o Uint16) String() string](#Uint16.String)
* [type Uint32](#Uint32)
  * [func EmptyUint32() Uint32](#EmptyUint32)
  * [func OfUint32(value uint32) Uint32](#OfUint32)
  * [func OfUint32Ptr(ptr *uint32) Uint32](#OfUint32Ptr)
  * [func (o Uint32) Else(elseValue uint32) (value uint32)](#Uint32.Else)
  * [func (o Uint32) ElseFunc(f func() uint32) (value uint32)](#Uint32.ElseFunc)
  * [func (o Uint32) If(f func(value uint32))](#Uint32.If)
  * [func (o Uint32) IsEmpty() bool](#Uint32.IsEmpty)
  * [func (o Uint32) IsPresent() bool](#Uint32.IsPresent)
  * [func (o Uint32) String() string](#Uint32.String)
* [type Uint64](#Uint64)
  * [func EmptyUint64() Uint64](#EmptyUint64)
  * [func OfUint64(value uint64) Uint64](#OfUint64)
  * [func OfUint64Ptr(ptr *uint64) Uint64](#OfUint64Ptr)
  * [func (o Uint64) Else(elseValue uint64) (value uint64)](#Uint64.Else)
  * [func (o Uint64) ElseFunc(f func() uint64) (value uint64)](#Uint64.ElseFunc)
  * [func (o Uint64) If(f func(value uint64))](#Uint64.If)
  * [func (o Uint64) IsEmpty() bool](#Uint64.IsEmpty)
  * [func (o Uint64) IsPresent() bool](#Uint64.IsPresent)
  * [func (o Uint64) String() string](#Uint64.String)
* [type Uint8](#Uint8)
  * [func EmptyUint8() Uint8](#EmptyUint8)
  * [func OfUint8(value uint8) Uint8](#OfUint8)
  * [func OfUint8Ptr(ptr *uint8) Uint8](#OfUint8Ptr)
  * [func (o Uint8) Else(elseValue uint8) (value uint8)](#Uint8.Else)
  * [func (o Uint8) ElseFunc(f func() uint8) (value uint8)](#Uint8.ElseFunc)
  * [func (o Uint8) If(f func(value uint8))](#Uint8.If)
  * [func (o Uint8) IsEmpty() bool](#Uint8.IsEmpty)
  * [func (o Uint8) IsPresent() bool](#Uint8.IsPresent)
  * [func (o Uint8) String() string](#Uint8.String)
* [type Uintptr](#Uintptr)
  * [func EmptyUintptr() Uintptr](#EmptyUintptr)
  * [func OfUintptr(value uintptr) Uintptr](#OfUintptr)
  * [func OfUintptrPtr(ptr *uintptr) Uintptr](#OfUintptrPtr)
  * [func (o Uintptr) Else(elseValue uintptr) (value uintptr)](#Uintptr.Else)
  * [func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr)](#Uintptr.ElseFunc)
  * [func (o Uintptr) If(f func(value uintptr))](#Uintptr.If)
  * [func (o Uintptr) IsEmpty() bool](#Uintptr.IsEmpty)
  * [func (o Uintptr) IsPresent() bool](#Uintptr.IsPresent)
  * [func (o Uintptr) String() string](#Uintptr.String)

#### <a name="pkg-examples">Examples</a>
* [Int.Else](#example_Int_Else)
* [Int.If (Present)](#example_Int_If_present)
* [Int (XmlMarshal)](#example_Int_xmlMarshal)
* [Int (XmlUnmarshal)](#example_Int_xmlUnmarshal)

#### <a name="pkg-files">Package files</a>
[bool_generated.go](/src/github.com/leighmcculloch/optional/bool_generated.go) [bool_marshal.go](/src/github.com/leighmcculloch/optional/bool_marshal.go) [byte_generated.go](/src/github.com/leighmcculloch/optional/byte_generated.go) [complex128_generated.go](/src/github.com/leighmcculloch/optional/complex128_generated.go) [complex64_generated.go](/src/github.com/leighmcculloch/optional/complex64_generated.go) [doc.go](/src/github.com/leighmcculloch/optional/doc.go) [float32_generated.go](/src/github.com/leighmcculloch/optional/float32_generated.go) [float64_generated.go](/src/github.com/leighmcculloch/optional/float64_generated.go) [int16_generated.go](/src/github.com/leighmcculloch/optional/int16_generated.go) [int32_generated.go](/src/github.com/leighmcculloch/optional/int32_generated.go) [int64_generated.go](/src/github.com/leighmcculloch/optional/int64_generated.go) [int8_generated.go](/src/github.com/leighmcculloch/optional/int8_generated.go) [int_generated.go](/src/github.com/leighmcculloch/optional/int_generated.go) [int_marshal.go](/src/github.com/leighmcculloch/optional/int_marshal.go) [rune_generated.go](/src/github.com/leighmcculloch/optional/rune_generated.go) [string_generated.go](/src/github.com/leighmcculloch/optional/string_generated.go) [types.go](/src/github.com/leighmcculloch/optional/types.go) [uint16_generated.go](/src/github.com/leighmcculloch/optional/uint16_generated.go) [uint32_generated.go](/src/github.com/leighmcculloch/optional/uint32_generated.go) [uint64_generated.go](/src/github.com/leighmcculloch/optional/uint64_generated.go) [uint8_generated.go](/src/github.com/leighmcculloch/optional/uint8_generated.go) [uint_generated.go](/src/github.com/leighmcculloch/optional/uint_generated.go) [uintptr_generated.go](/src/github.com/leighmcculloch/optional/uintptr_generated.go) 






## <a name="Bool">type</a> [Bool](/src/target/bool_generated.go?s=199:221#L1)
``` go
type Bool optionalBool
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyBool">func</a> [EmptyBool](/src/target/bool_generated.go?s=532:553#L23)
``` go
func EmptyBool() Bool
```
Empty returns an empty Optional.


### <a name="OfBool">func</a> [OfBool](/src/target/bool_generated.go?s=319:347#L10)
``` go
func OfBool(value bool) Bool
```
Of wraps the value in an Optional.


### <a name="OfBoolPtr">func</a> [OfBoolPtr](/src/target/bool_generated.go?s=387:417#L14)
``` go
func OfBoolPtr(ptr *bool) Bool
```




### <a name="Bool.Else">func</a> (Bool) [Else](/src/target/bool_generated.go?s=1256:1303#L55)
``` go
func (o Bool) Else(elseValue bool) (value bool)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Bool.ElseFunc">func</a> (Bool) [ElseFunc](/src/target/bool_generated.go?s=985:1035#L44)
``` go
func (o Bool) ElseFunc(f func() bool) (value bool)
```



### <a name="Bool.If">func</a> (Bool) [If](/src/target/bool_generated.go?s=899:935#L38)
``` go
func (o Bool) If(f func(value bool))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Bool.IsEmpty">func</a> (Bool) [IsEmpty](/src/target/bool_generated.go?s=648:676#L28)
``` go
func (o Bool) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Bool.IsPresent">func</a> (Bool) [IsPresent](/src/target/bool_generated.go?s=771:801#L33)
``` go
func (o Bool) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Bool.MarshalText">func</a> (Bool) [MarshalText](/src/target/bool_marshal.go?s=94:146#L1)
``` go
func (o Bool) MarshalText() (text []byte, err error)
```
MarshalText returns text for marshaling this Int.




### <a name="Bool.String">func</a> (Bool) [String](/src/target/bool_generated.go?s=1471:1500#L60)
``` go
func (o Bool) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




### <a name="Bool.UnmarshalText">func</a> (\*Bool) [UnmarshalText](/src/target/bool_marshal.go?s=270:317#L6)
``` go
func (o *Bool) UnmarshalText(text []byte) error
```
UnmarshalText parse the text into this Int




## <a name="Byte">type</a> [Byte](/src/target/byte_generated.go?s=199:221#L1)
``` go
type Byte optionalByte
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyByte">func</a> [EmptyByte](/src/target/byte_generated.go?s=532:553#L23)
``` go
func EmptyByte() Byte
```
Empty returns an empty Optional.


### <a name="OfByte">func</a> [OfByte](/src/target/byte_generated.go?s=319:347#L10)
``` go
func OfByte(value byte) Byte
```
Of wraps the value in an Optional.


### <a name="OfBytePtr">func</a> [OfBytePtr](/src/target/byte_generated.go?s=387:417#L14)
``` go
func OfBytePtr(ptr *byte) Byte
```




### <a name="Byte.Else">func</a> (Byte) [Else](/src/target/byte_generated.go?s=1256:1303#L55)
``` go
func (o Byte) Else(elseValue byte) (value byte)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Byte.ElseFunc">func</a> (Byte) [ElseFunc](/src/target/byte_generated.go?s=985:1035#L44)
``` go
func (o Byte) ElseFunc(f func() byte) (value byte)
```



### <a name="Byte.If">func</a> (Byte) [If](/src/target/byte_generated.go?s=899:935#L38)
``` go
func (o Byte) If(f func(value byte))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Byte.IsEmpty">func</a> (Byte) [IsEmpty](/src/target/byte_generated.go?s=648:676#L28)
``` go
func (o Byte) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Byte.IsPresent">func</a> (Byte) [IsPresent](/src/target/byte_generated.go?s=771:801#L33)
``` go
func (o Byte) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Byte.String">func</a> (Byte) [String](/src/target/byte_generated.go?s=1471:1500#L60)
``` go
func (o Byte) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Complex128">type</a> [Complex128](/src/target/complex128_generated.go?s=199:233#L1)
``` go
type Complex128 optionalComplex128
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyComplex128">func</a> [EmptyComplex128](/src/target/complex128_generated.go?s=622:655#L23)
``` go
func EmptyComplex128() Complex128
```
Empty returns an empty Optional.


### <a name="OfComplex128">func</a> [OfComplex128](/src/target/complex128_generated.go?s=349:395#L10)
``` go
func OfComplex128(value complex128) Complex128
```
Of wraps the value in an Optional.


### <a name="OfComplex128Ptr">func</a> [OfComplex128Ptr](/src/target/complex128_generated.go?s=447:495#L14)
``` go
func OfComplex128Ptr(ptr *complex128) Complex128
```




### <a name="Complex128.Else">func</a> (Complex128) [Else](/src/target/complex128_generated.go?s=1412:1477#L55)
``` go
func (o Complex128) Else(elseValue complex128) (value complex128)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Complex128.ElseFunc">func</a> (Complex128) [ElseFunc](/src/target/complex128_generated.go?s=1117:1185#L44)
``` go
func (o Complex128) ElseFunc(f func() complex128) (value complex128)
```



### <a name="Complex128.If">func</a> (Complex128) [If](/src/target/complex128_generated.go?s=1013:1061#L38)
``` go
func (o Complex128) If(f func(value complex128))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Complex128.IsEmpty">func</a> (Complex128) [IsEmpty](/src/target/complex128_generated.go?s=750:784#L28)
``` go
func (o Complex128) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Complex128.IsPresent">func</a> (Complex128) [IsPresent](/src/target/complex128_generated.go?s=879:915#L33)
``` go
func (o Complex128) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Complex128.String">func</a> (Complex128) [String](/src/target/complex128_generated.go?s=1651:1686#L60)
``` go
func (o Complex128) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Complex64">type</a> [Complex64](/src/target/complex64_generated.go?s=199:231#L1)
``` go
type Complex64 optionalComplex64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyComplex64">func</a> [EmptyComplex64](/src/target/complex64_generated.go?s=607:638#L23)
``` go
func EmptyComplex64() Complex64
```
Empty returns an empty Optional.


### <a name="OfComplex64">func</a> [OfComplex64](/src/target/complex64_generated.go?s=344:387#L10)
``` go
func OfComplex64(value complex64) Complex64
```
Of wraps the value in an Optional.


### <a name="OfComplex64Ptr">func</a> [OfComplex64Ptr](/src/target/complex64_generated.go?s=437:482#L14)
``` go
func OfComplex64Ptr(ptr *complex64) Complex64
```




### <a name="Complex64.Else">func</a> (Complex64) [Else](/src/target/complex64_generated.go?s=1386:1448#L55)
``` go
func (o Complex64) Else(elseValue complex64) (value complex64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Complex64.ElseFunc">func</a> (Complex64) [ElseFunc](/src/target/complex64_generated.go?s=1095:1160#L44)
``` go
func (o Complex64) ElseFunc(f func() complex64) (value complex64)
```



### <a name="Complex64.If">func</a> (Complex64) [If](/src/target/complex64_generated.go?s=994:1040#L38)
``` go
func (o Complex64) If(f func(value complex64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Complex64.IsEmpty">func</a> (Complex64) [IsEmpty](/src/target/complex64_generated.go?s=733:766#L28)
``` go
func (o Complex64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Complex64.IsPresent">func</a> (Complex64) [IsPresent](/src/target/complex64_generated.go?s=861:896#L33)
``` go
func (o Complex64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Complex64.String">func</a> (Complex64) [String](/src/target/complex64_generated.go?s=1621:1655#L60)
``` go
func (o Complex64) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Float32">type</a> [Float32](/src/target/float32_generated.go?s=199:227#L1)
``` go
type Float32 optionalFloat32
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyFloat32">func</a> [EmptyFloat32](/src/target/float32_generated.go?s=577:604#L23)
``` go
func EmptyFloat32() Float32
```
Empty returns an empty Optional.


### <a name="OfFloat32">func</a> [OfFloat32](/src/target/float32_generated.go?s=334:371#L10)
``` go
func OfFloat32(value float32) Float32
```
Of wraps the value in an Optional.


### <a name="OfFloat32Ptr">func</a> [OfFloat32Ptr](/src/target/float32_generated.go?s=417:456#L14)
``` go
func OfFloat32Ptr(ptr *float32) Float32
```




### <a name="Float32.Else">func</a> (Float32) [Else](/src/target/float32_generated.go?s=1334:1390#L55)
``` go
func (o Float32) Else(elseValue float32) (value float32)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Float32.ElseFunc">func</a> (Float32) [ElseFunc](/src/target/float32_generated.go?s=1051:1110#L44)
``` go
func (o Float32) ElseFunc(f func() float32) (value float32)
```



### <a name="Float32.If">func</a> (Float32) [If](/src/target/float32_generated.go?s=956:998#L38)
``` go
func (o Float32) If(f func(value float32))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Float32.IsEmpty">func</a> (Float32) [IsEmpty](/src/target/float32_generated.go?s=699:730#L28)
``` go
func (o Float32) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Float32.IsPresent">func</a> (Float32) [IsPresent](/src/target/float32_generated.go?s=825:858#L33)
``` go
func (o Float32) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Float32.String">func</a> (Float32) [String](/src/target/float32_generated.go?s=1561:1593#L60)
``` go
func (o Float32) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Float64">type</a> [Float64](/src/target/float64_generated.go?s=199:227#L1)
``` go
type Float64 optionalFloat64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyFloat64">func</a> [EmptyFloat64](/src/target/float64_generated.go?s=577:604#L23)
``` go
func EmptyFloat64() Float64
```
Empty returns an empty Optional.


### <a name="OfFloat64">func</a> [OfFloat64](/src/target/float64_generated.go?s=334:371#L10)
``` go
func OfFloat64(value float64) Float64
```
Of wraps the value in an Optional.


### <a name="OfFloat64Ptr">func</a> [OfFloat64Ptr](/src/target/float64_generated.go?s=417:456#L14)
``` go
func OfFloat64Ptr(ptr *float64) Float64
```




### <a name="Float64.Else">func</a> (Float64) [Else](/src/target/float64_generated.go?s=1334:1390#L55)
``` go
func (o Float64) Else(elseValue float64) (value float64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Float64.ElseFunc">func</a> (Float64) [ElseFunc](/src/target/float64_generated.go?s=1051:1110#L44)
``` go
func (o Float64) ElseFunc(f func() float64) (value float64)
```



### <a name="Float64.If">func</a> (Float64) [If](/src/target/float64_generated.go?s=956:998#L38)
``` go
func (o Float64) If(f func(value float64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Float64.IsEmpty">func</a> (Float64) [IsEmpty](/src/target/float64_generated.go?s=699:730#L28)
``` go
func (o Float64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Float64.IsPresent">func</a> (Float64) [IsPresent](/src/target/float64_generated.go?s=825:858#L33)
``` go
func (o Float64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Float64.String">func</a> (Float64) [String](/src/target/float64_generated.go?s=1561:1593#L60)
``` go
func (o Float64) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Int">type</a> [Int](/src/target/int_generated.go?s=199:219#L1)
``` go
type Int optionalInt
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt">func</a> [EmptyInt](/src/target/int_generated.go?s=517:536#L23)
``` go
func EmptyInt() Int
```
Empty returns an empty Optional.


### <a name="OfInt">func</a> [OfInt](/src/target/int_generated.go?s=314:339#L10)
``` go
func OfInt(value int) Int
```
Of wraps the value in an Optional.


### <a name="OfIntPtr">func</a> [OfIntPtr](/src/target/int_generated.go?s=377:404#L14)
``` go
func OfIntPtr(ptr *int) Int
```




### <a name="Int.Else">func</a> (Int) [Else](/src/target/int_generated.go?s=1230:1274#L55)
``` go
func (o Int) Else(elseValue int) (value int)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int.ElseFunc">func</a> (Int) [ElseFunc](/src/target/int_generated.go?s=963:1010#L44)
``` go
func (o Int) ElseFunc(f func() int) (value int)
```



### <a name="Int.If">func</a> (Int) [If](/src/target/int_generated.go?s=880:914#L38)
``` go
func (o Int) If(f func(value int))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int.IsEmpty">func</a> (Int) [IsEmpty](/src/target/int_generated.go?s=631:658#L28)
``` go
func (o Int) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int.IsPresent">func</a> (Int) [IsPresent](/src/target/int_generated.go?s=753:782#L33)
``` go
func (o Int) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int.MarshalText">func</a> (Int) [MarshalText](/src/target/int_marshal.go?s=94:145#L1)
``` go
func (o Int) MarshalText() (text []byte, err error)
```
MarshalText returns text for marshaling this Int.




### <a name="Int.String">func</a> (Int) [String](/src/target/int_generated.go?s=1441:1469#L60)
``` go
func (o Int) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




### <a name="Int.UnmarshalText">func</a> (\*Int) [UnmarshalText](/src/target/int_marshal.go?s=278:324#L6)
``` go
func (o *Int) UnmarshalText(text []byte) error
```
UnmarshalText parse the text into this Int




## <a name="Int16">type</a> [Int16](/src/target/int16_generated.go?s=199:223#L1)
``` go
type Int16 optionalInt16
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt16">func</a> [EmptyInt16](/src/target/int16_generated.go?s=547:570#L23)
``` go
func EmptyInt16() Int16
```
Empty returns an empty Optional.


### <a name="OfInt16">func</a> [OfInt16](/src/target/int16_generated.go?s=324:355#L10)
``` go
func OfInt16(value int16) Int16
```
Of wraps the value in an Optional.


### <a name="OfInt16Ptr">func</a> [OfInt16Ptr](/src/target/int16_generated.go?s=397:430#L14)
``` go
func OfInt16Ptr(ptr *int16) Int16
```




### <a name="Int16.Else">func</a> (Int16) [Else](/src/target/int16_generated.go?s=1282:1332#L55)
``` go
func (o Int16) Else(elseValue int16) (value int16)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int16.ElseFunc">func</a> (Int16) [ElseFunc](/src/target/int16_generated.go?s=1007:1060#L44)
``` go
func (o Int16) ElseFunc(f func() int16) (value int16)
```



### <a name="Int16.If">func</a> (Int16) [If](/src/target/int16_generated.go?s=918:956#L38)
``` go
func (o Int16) If(f func(value int16))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int16.IsEmpty">func</a> (Int16) [IsEmpty](/src/target/int16_generated.go?s=665:694#L28)
``` go
func (o Int16) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int16.IsPresent">func</a> (Int16) [IsPresent](/src/target/int16_generated.go?s=789:820#L33)
``` go
func (o Int16) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int16.String">func</a> (Int16) [String](/src/target/int16_generated.go?s=1501:1531#L60)
``` go
func (o Int16) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Int32">type</a> [Int32](/src/target/int32_generated.go?s=199:223#L1)
``` go
type Int32 optionalInt32
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt32">func</a> [EmptyInt32](/src/target/int32_generated.go?s=547:570#L23)
``` go
func EmptyInt32() Int32
```
Empty returns an empty Optional.


### <a name="OfInt32">func</a> [OfInt32](/src/target/int32_generated.go?s=324:355#L10)
``` go
func OfInt32(value int32) Int32
```
Of wraps the value in an Optional.


### <a name="OfInt32Ptr">func</a> [OfInt32Ptr](/src/target/int32_generated.go?s=397:430#L14)
``` go
func OfInt32Ptr(ptr *int32) Int32
```




### <a name="Int32.Else">func</a> (Int32) [Else](/src/target/int32_generated.go?s=1282:1332#L55)
``` go
func (o Int32) Else(elseValue int32) (value int32)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int32.ElseFunc">func</a> (Int32) [ElseFunc](/src/target/int32_generated.go?s=1007:1060#L44)
``` go
func (o Int32) ElseFunc(f func() int32) (value int32)
```



### <a name="Int32.If">func</a> (Int32) [If](/src/target/int32_generated.go?s=918:956#L38)
``` go
func (o Int32) If(f func(value int32))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int32.IsEmpty">func</a> (Int32) [IsEmpty](/src/target/int32_generated.go?s=665:694#L28)
``` go
func (o Int32) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int32.IsPresent">func</a> (Int32) [IsPresent](/src/target/int32_generated.go?s=789:820#L33)
``` go
func (o Int32) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int32.String">func</a> (Int32) [String](/src/target/int32_generated.go?s=1501:1531#L60)
``` go
func (o Int32) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Int64">type</a> [Int64](/src/target/int64_generated.go?s=199:223#L1)
``` go
type Int64 optionalInt64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt64">func</a> [EmptyInt64](/src/target/int64_generated.go?s=547:570#L23)
``` go
func EmptyInt64() Int64
```
Empty returns an empty Optional.


### <a name="OfInt64">func</a> [OfInt64](/src/target/int64_generated.go?s=324:355#L10)
``` go
func OfInt64(value int64) Int64
```
Of wraps the value in an Optional.


### <a name="OfInt64Ptr">func</a> [OfInt64Ptr](/src/target/int64_generated.go?s=397:430#L14)
``` go
func OfInt64Ptr(ptr *int64) Int64
```




### <a name="Int64.Else">func</a> (Int64) [Else](/src/target/int64_generated.go?s=1282:1332#L55)
``` go
func (o Int64) Else(elseValue int64) (value int64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int64.ElseFunc">func</a> (Int64) [ElseFunc](/src/target/int64_generated.go?s=1007:1060#L44)
``` go
func (o Int64) ElseFunc(f func() int64) (value int64)
```



### <a name="Int64.If">func</a> (Int64) [If](/src/target/int64_generated.go?s=918:956#L38)
``` go
func (o Int64) If(f func(value int64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int64.IsEmpty">func</a> (Int64) [IsEmpty](/src/target/int64_generated.go?s=665:694#L28)
``` go
func (o Int64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int64.IsPresent">func</a> (Int64) [IsPresent](/src/target/int64_generated.go?s=789:820#L33)
``` go
func (o Int64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int64.String">func</a> (Int64) [String](/src/target/int64_generated.go?s=1501:1531#L60)
``` go
func (o Int64) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Int8">type</a> [Int8](/src/target/int8_generated.go?s=199:221#L1)
``` go
type Int8 optionalInt8
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt8">func</a> [EmptyInt8](/src/target/int8_generated.go?s=532:553#L23)
``` go
func EmptyInt8() Int8
```
Empty returns an empty Optional.


### <a name="OfInt8">func</a> [OfInt8](/src/target/int8_generated.go?s=319:347#L10)
``` go
func OfInt8(value int8) Int8
```
Of wraps the value in an Optional.


### <a name="OfInt8Ptr">func</a> [OfInt8Ptr](/src/target/int8_generated.go?s=387:417#L14)
``` go
func OfInt8Ptr(ptr *int8) Int8
```




### <a name="Int8.Else">func</a> (Int8) [Else](/src/target/int8_generated.go?s=1256:1303#L55)
``` go
func (o Int8) Else(elseValue int8) (value int8)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int8.ElseFunc">func</a> (Int8) [ElseFunc](/src/target/int8_generated.go?s=985:1035#L44)
``` go
func (o Int8) ElseFunc(f func() int8) (value int8)
```



### <a name="Int8.If">func</a> (Int8) [If](/src/target/int8_generated.go?s=899:935#L38)
``` go
func (o Int8) If(f func(value int8))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int8.IsEmpty">func</a> (Int8) [IsEmpty](/src/target/int8_generated.go?s=648:676#L28)
``` go
func (o Int8) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int8.IsPresent">func</a> (Int8) [IsPresent](/src/target/int8_generated.go?s=771:801#L33)
``` go
func (o Int8) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int8.String">func</a> (Int8) [String](/src/target/int8_generated.go?s=1471:1500#L60)
``` go
func (o Int8) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Rune">type</a> [Rune](/src/target/rune_generated.go?s=199:221#L1)
``` go
type Rune optionalRune
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyRune">func</a> [EmptyRune](/src/target/rune_generated.go?s=532:553#L23)
``` go
func EmptyRune() Rune
```
Empty returns an empty Optional.


### <a name="OfRune">func</a> [OfRune](/src/target/rune_generated.go?s=319:347#L10)
``` go
func OfRune(value rune) Rune
```
Of wraps the value in an Optional.


### <a name="OfRunePtr">func</a> [OfRunePtr](/src/target/rune_generated.go?s=387:417#L14)
``` go
func OfRunePtr(ptr *rune) Rune
```




### <a name="Rune.Else">func</a> (Rune) [Else](/src/target/rune_generated.go?s=1256:1303#L55)
``` go
func (o Rune) Else(elseValue rune) (value rune)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Rune.ElseFunc">func</a> (Rune) [ElseFunc](/src/target/rune_generated.go?s=985:1035#L44)
``` go
func (o Rune) ElseFunc(f func() rune) (value rune)
```



### <a name="Rune.If">func</a> (Rune) [If](/src/target/rune_generated.go?s=899:935#L38)
``` go
func (o Rune) If(f func(value rune))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Rune.IsEmpty">func</a> (Rune) [IsEmpty](/src/target/rune_generated.go?s=648:676#L28)
``` go
func (o Rune) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Rune.IsPresent">func</a> (Rune) [IsPresent](/src/target/rune_generated.go?s=771:801#L33)
``` go
func (o Rune) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Rune.String">func</a> (Rune) [String](/src/target/rune_generated.go?s=1471:1500#L60)
``` go
func (o Rune) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="String">type</a> [String](/src/target/string_generated.go?s=199:225#L1)
``` go
type String optionalString
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyString">func</a> [EmptyString](/src/target/string_generated.go?s=562:587#L23)
``` go
func EmptyString() String
```
Empty returns an empty Optional.


### <a name="OfString">func</a> [OfString](/src/target/string_generated.go?s=329:363#L10)
``` go
func OfString(value string) String
```
Of wraps the value in an Optional.


### <a name="OfStringPtr">func</a> [OfStringPtr](/src/target/string_generated.go?s=407:443#L14)
``` go
func OfStringPtr(ptr *string) String
```




### <a name="String.Else">func</a> (String) [Else](/src/target/string_generated.go?s=1308:1361#L55)
``` go
func (o String) Else(elseValue string) (value string)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="String.ElseFunc">func</a> (String) [ElseFunc](/src/target/string_generated.go?s=1029:1085#L44)
``` go
func (o String) ElseFunc(f func() string) (value string)
```



### <a name="String.If">func</a> (String) [If](/src/target/string_generated.go?s=937:977#L38)
``` go
func (o String) If(f func(value string))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="String.IsEmpty">func</a> (String) [IsEmpty](/src/target/string_generated.go?s=682:712#L28)
``` go
func (o String) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="String.IsPresent">func</a> (String) [IsPresent](/src/target/string_generated.go?s=807:839#L33)
``` go
func (o String) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="String.String">func</a> (String) [String](/src/target/string_generated.go?s=1531:1562#L60)
``` go
func (o String) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Uint">type</a> [Uint](/src/target/uint_generated.go?s=199:221#L1)
``` go
type Uint optionalUint
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint">func</a> [EmptyUint](/src/target/uint_generated.go?s=532:553#L23)
``` go
func EmptyUint() Uint
```
Empty returns an empty Optional.


### <a name="OfUint">func</a> [OfUint](/src/target/uint_generated.go?s=319:347#L10)
``` go
func OfUint(value uint) Uint
```
Of wraps the value in an Optional.


### <a name="OfUintPtr">func</a> [OfUintPtr](/src/target/uint_generated.go?s=387:417#L14)
``` go
func OfUintPtr(ptr *uint) Uint
```




### <a name="Uint.Else">func</a> (Uint) [Else](/src/target/uint_generated.go?s=1256:1303#L55)
``` go
func (o Uint) Else(elseValue uint) (value uint)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint.ElseFunc">func</a> (Uint) [ElseFunc](/src/target/uint_generated.go?s=985:1035#L44)
``` go
func (o Uint) ElseFunc(f func() uint) (value uint)
```



### <a name="Uint.If">func</a> (Uint) [If](/src/target/uint_generated.go?s=899:935#L38)
``` go
func (o Uint) If(f func(value uint))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint.IsEmpty">func</a> (Uint) [IsEmpty](/src/target/uint_generated.go?s=648:676#L28)
``` go
func (o Uint) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint.IsPresent">func</a> (Uint) [IsPresent](/src/target/uint_generated.go?s=771:801#L33)
``` go
func (o Uint) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint.String">func</a> (Uint) [String](/src/target/uint_generated.go?s=1471:1500#L60)
``` go
func (o Uint) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Uint16">type</a> [Uint16](/src/target/uint16_generated.go?s=199:225#L1)
``` go
type Uint16 optionalUint16
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint16">func</a> [EmptyUint16](/src/target/uint16_generated.go?s=562:587#L23)
``` go
func EmptyUint16() Uint16
```
Empty returns an empty Optional.


### <a name="OfUint16">func</a> [OfUint16](/src/target/uint16_generated.go?s=329:363#L10)
``` go
func OfUint16(value uint16) Uint16
```
Of wraps the value in an Optional.


### <a name="OfUint16Ptr">func</a> [OfUint16Ptr](/src/target/uint16_generated.go?s=407:443#L14)
``` go
func OfUint16Ptr(ptr *uint16) Uint16
```




### <a name="Uint16.Else">func</a> (Uint16) [Else](/src/target/uint16_generated.go?s=1308:1361#L55)
``` go
func (o Uint16) Else(elseValue uint16) (value uint16)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint16.ElseFunc">func</a> (Uint16) [ElseFunc](/src/target/uint16_generated.go?s=1029:1085#L44)
``` go
func (o Uint16) ElseFunc(f func() uint16) (value uint16)
```



### <a name="Uint16.If">func</a> (Uint16) [If](/src/target/uint16_generated.go?s=937:977#L38)
``` go
func (o Uint16) If(f func(value uint16))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint16.IsEmpty">func</a> (Uint16) [IsEmpty](/src/target/uint16_generated.go?s=682:712#L28)
``` go
func (o Uint16) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint16.IsPresent">func</a> (Uint16) [IsPresent](/src/target/uint16_generated.go?s=807:839#L33)
``` go
func (o Uint16) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint16.String">func</a> (Uint16) [String](/src/target/uint16_generated.go?s=1531:1562#L60)
``` go
func (o Uint16) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Uint32">type</a> [Uint32](/src/target/uint32_generated.go?s=199:225#L1)
``` go
type Uint32 optionalUint32
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint32">func</a> [EmptyUint32](/src/target/uint32_generated.go?s=562:587#L23)
``` go
func EmptyUint32() Uint32
```
Empty returns an empty Optional.


### <a name="OfUint32">func</a> [OfUint32](/src/target/uint32_generated.go?s=329:363#L10)
``` go
func OfUint32(value uint32) Uint32
```
Of wraps the value in an Optional.


### <a name="OfUint32Ptr">func</a> [OfUint32Ptr](/src/target/uint32_generated.go?s=407:443#L14)
``` go
func OfUint32Ptr(ptr *uint32) Uint32
```




### <a name="Uint32.Else">func</a> (Uint32) [Else](/src/target/uint32_generated.go?s=1308:1361#L55)
``` go
func (o Uint32) Else(elseValue uint32) (value uint32)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint32.ElseFunc">func</a> (Uint32) [ElseFunc](/src/target/uint32_generated.go?s=1029:1085#L44)
``` go
func (o Uint32) ElseFunc(f func() uint32) (value uint32)
```



### <a name="Uint32.If">func</a> (Uint32) [If](/src/target/uint32_generated.go?s=937:977#L38)
``` go
func (o Uint32) If(f func(value uint32))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint32.IsEmpty">func</a> (Uint32) [IsEmpty](/src/target/uint32_generated.go?s=682:712#L28)
``` go
func (o Uint32) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint32.IsPresent">func</a> (Uint32) [IsPresent](/src/target/uint32_generated.go?s=807:839#L33)
``` go
func (o Uint32) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint32.String">func</a> (Uint32) [String](/src/target/uint32_generated.go?s=1531:1562#L60)
``` go
func (o Uint32) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Uint64">type</a> [Uint64](/src/target/uint64_generated.go?s=199:225#L1)
``` go
type Uint64 optionalUint64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint64">func</a> [EmptyUint64](/src/target/uint64_generated.go?s=562:587#L23)
``` go
func EmptyUint64() Uint64
```
Empty returns an empty Optional.


### <a name="OfUint64">func</a> [OfUint64](/src/target/uint64_generated.go?s=329:363#L10)
``` go
func OfUint64(value uint64) Uint64
```
Of wraps the value in an Optional.


### <a name="OfUint64Ptr">func</a> [OfUint64Ptr](/src/target/uint64_generated.go?s=407:443#L14)
``` go
func OfUint64Ptr(ptr *uint64) Uint64
```




### <a name="Uint64.Else">func</a> (Uint64) [Else](/src/target/uint64_generated.go?s=1308:1361#L55)
``` go
func (o Uint64) Else(elseValue uint64) (value uint64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint64.ElseFunc">func</a> (Uint64) [ElseFunc](/src/target/uint64_generated.go?s=1029:1085#L44)
``` go
func (o Uint64) ElseFunc(f func() uint64) (value uint64)
```



### <a name="Uint64.If">func</a> (Uint64) [If](/src/target/uint64_generated.go?s=937:977#L38)
``` go
func (o Uint64) If(f func(value uint64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint64.IsEmpty">func</a> (Uint64) [IsEmpty](/src/target/uint64_generated.go?s=682:712#L28)
``` go
func (o Uint64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint64.IsPresent">func</a> (Uint64) [IsPresent](/src/target/uint64_generated.go?s=807:839#L33)
``` go
func (o Uint64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint64.String">func</a> (Uint64) [String](/src/target/uint64_generated.go?s=1531:1562#L60)
``` go
func (o Uint64) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Uint8">type</a> [Uint8](/src/target/uint8_generated.go?s=199:223#L1)
``` go
type Uint8 optionalUint8
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint8">func</a> [EmptyUint8](/src/target/uint8_generated.go?s=547:570#L23)
``` go
func EmptyUint8() Uint8
```
Empty returns an empty Optional.


### <a name="OfUint8">func</a> [OfUint8](/src/target/uint8_generated.go?s=324:355#L10)
``` go
func OfUint8(value uint8) Uint8
```
Of wraps the value in an Optional.


### <a name="OfUint8Ptr">func</a> [OfUint8Ptr](/src/target/uint8_generated.go?s=397:430#L14)
``` go
func OfUint8Ptr(ptr *uint8) Uint8
```




### <a name="Uint8.Else">func</a> (Uint8) [Else](/src/target/uint8_generated.go?s=1282:1332#L55)
``` go
func (o Uint8) Else(elseValue uint8) (value uint8)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint8.ElseFunc">func</a> (Uint8) [ElseFunc](/src/target/uint8_generated.go?s=1007:1060#L44)
``` go
func (o Uint8) ElseFunc(f func() uint8) (value uint8)
```



### <a name="Uint8.If">func</a> (Uint8) [If](/src/target/uint8_generated.go?s=918:956#L38)
``` go
func (o Uint8) If(f func(value uint8))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint8.IsEmpty">func</a> (Uint8) [IsEmpty](/src/target/uint8_generated.go?s=665:694#L28)
``` go
func (o Uint8) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint8.IsPresent">func</a> (Uint8) [IsPresent](/src/target/uint8_generated.go?s=789:820#L33)
``` go
func (o Uint8) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint8.String">func</a> (Uint8) [String](/src/target/uint8_generated.go?s=1501:1531#L60)
``` go
func (o Uint8) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.




## <a name="Uintptr">type</a> [Uintptr](/src/target/uintptr_generated.go?s=199:227#L1)
``` go
type Uintptr optionalUintptr
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUintptr">func</a> [EmptyUintptr](/src/target/uintptr_generated.go?s=577:604#L23)
``` go
func EmptyUintptr() Uintptr
```
Empty returns an empty Optional.


### <a name="OfUintptr">func</a> [OfUintptr](/src/target/uintptr_generated.go?s=334:371#L10)
``` go
func OfUintptr(value uintptr) Uintptr
```
Of wraps the value in an Optional.


### <a name="OfUintptrPtr">func</a> [OfUintptrPtr](/src/target/uintptr_generated.go?s=417:456#L14)
``` go
func OfUintptrPtr(ptr *uintptr) Uintptr
```




### <a name="Uintptr.Else">func</a> (Uintptr) [Else](/src/target/uintptr_generated.go?s=1334:1390#L55)
``` go
func (o Uintptr) Else(elseValue uintptr) (value uintptr)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uintptr.ElseFunc">func</a> (Uintptr) [ElseFunc](/src/target/uintptr_generated.go?s=1051:1110#L44)
``` go
func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr)
```



### <a name="Uintptr.If">func</a> (Uintptr) [If](/src/target/uintptr_generated.go?s=956:998#L38)
``` go
func (o Uintptr) If(f func(value uintptr))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uintptr.IsEmpty">func</a> (Uintptr) [IsEmpty](/src/target/uintptr_generated.go?s=699:730#L28)
``` go
func (o Uintptr) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uintptr.IsPresent">func</a> (Uintptr) [IsPresent](/src/target/uintptr_generated.go?s=825:858#L33)
``` go
func (o Uintptr) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uintptr.String">func</a> (Uintptr) [String](/src/target/uintptr_generated.go?s=1561:1593#L60)
``` go
func (o Uintptr) String() string
```
String returns a string representation of the wrapped value if one is present, otherwise an empty string.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
