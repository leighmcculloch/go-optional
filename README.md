

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
Take a pointer to something and make it an optional to force users to only use it if it's not nil:


	var i *int = ...
	
	o := optional.OfIntPtr(v)
	
	o.IfPresent(func(i int) {
		// only called if i was not originally nil
		// use i here
	})

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
  * [func (o Bool) UnmarshalText(text []byte) error](#Bool.UnmarshalText)
* [type Byte](#Byte)
  * [func EmptyByte() Byte](#EmptyByte)
  * [func OfByte(value byte) Byte](#OfByte)
  * [func OfBytePtr(ptr *byte) Byte](#OfBytePtr)
  * [func (o Byte) Else(elseValue byte) (value byte)](#Byte.Else)
  * [func (o Byte) ElseFunc(f func() byte) (value byte)](#Byte.ElseFunc)
  * [func (o Byte) If(f func(value byte))](#Byte.If)
  * [func (o Byte) IsEmpty() bool](#Byte.IsEmpty)
  * [func (o Byte) IsPresent() bool](#Byte.IsPresent)
  * [func (o Byte) MarshalText() (text []byte, err error)](#Byte.MarshalText)
  * [func (o Byte) UnmarshalText(text []byte) error](#Byte.UnmarshalText)
* [type Complex128](#Complex128)
  * [func EmptyComplex128() Complex128](#EmptyComplex128)
  * [func OfComplex128(value complex128) Complex128](#OfComplex128)
  * [func OfComplex128Ptr(ptr *complex128) Complex128](#OfComplex128Ptr)
  * [func (o Complex128) Else(elseValue complex128) (value complex128)](#Complex128.Else)
  * [func (o Complex128) ElseFunc(f func() complex128) (value complex128)](#Complex128.ElseFunc)
  * [func (o Complex128) If(f func(value complex128))](#Complex128.If)
  * [func (o Complex128) IsEmpty() bool](#Complex128.IsEmpty)
  * [func (o Complex128) IsPresent() bool](#Complex128.IsPresent)
  * [func (o Complex128) MarshalText() (text []byte, err error)](#Complex128.MarshalText)
  * [func (o Complex128) UnmarshalText(text []byte) error](#Complex128.UnmarshalText)
* [type Complex64](#Complex64)
  * [func EmptyComplex64() Complex64](#EmptyComplex64)
  * [func OfComplex64(value complex64) Complex64](#OfComplex64)
  * [func OfComplex64Ptr(ptr *complex64) Complex64](#OfComplex64Ptr)
  * [func (o Complex64) Else(elseValue complex64) (value complex64)](#Complex64.Else)
  * [func (o Complex64) ElseFunc(f func() complex64) (value complex64)](#Complex64.ElseFunc)
  * [func (o Complex64) If(f func(value complex64))](#Complex64.If)
  * [func (o Complex64) IsEmpty() bool](#Complex64.IsEmpty)
  * [func (o Complex64) IsPresent() bool](#Complex64.IsPresent)
  * [func (o Complex64) MarshalText() (text []byte, err error)](#Complex64.MarshalText)
  * [func (o Complex64) UnmarshalText(text []byte) error](#Complex64.UnmarshalText)
* [type Error](#Error)
  * [func EmptyError() Error](#EmptyError)
  * [func OfError(value error) Error](#OfError)
  * [func OfErrorPtr(ptr *error) Error](#OfErrorPtr)
  * [func (o Error) Else(elseValue error) (value error)](#Error.Else)
  * [func (o Error) ElseFunc(f func() error) (value error)](#Error.ElseFunc)
  * [func (o Error) If(f func(value error))](#Error.If)
  * [func (o Error) IsEmpty() bool](#Error.IsEmpty)
  * [func (o Error) IsPresent() bool](#Error.IsPresent)
  * [func (o Error) MarshalText() (text []byte, err error)](#Error.MarshalText)
  * [func (o Error) UnmarshalText(text []byte) error](#Error.UnmarshalText)
* [type Float32](#Float32)
  * [func EmptyFloat32() Float32](#EmptyFloat32)
  * [func OfFloat32(value float32) Float32](#OfFloat32)
  * [func OfFloat32Ptr(ptr *float32) Float32](#OfFloat32Ptr)
  * [func (o Float32) Else(elseValue float32) (value float32)](#Float32.Else)
  * [func (o Float32) ElseFunc(f func() float32) (value float32)](#Float32.ElseFunc)
  * [func (o Float32) If(f func(value float32))](#Float32.If)
  * [func (o Float32) IsEmpty() bool](#Float32.IsEmpty)
  * [func (o Float32) IsPresent() bool](#Float32.IsPresent)
  * [func (o Float32) MarshalText() (text []byte, err error)](#Float32.MarshalText)
  * [func (o Float32) UnmarshalText(text []byte) error](#Float32.UnmarshalText)
* [type Float64](#Float64)
  * [func EmptyFloat64() Float64](#EmptyFloat64)
  * [func OfFloat64(value float64) Float64](#OfFloat64)
  * [func OfFloat64Ptr(ptr *float64) Float64](#OfFloat64Ptr)
  * [func (o Float64) Else(elseValue float64) (value float64)](#Float64.Else)
  * [func (o Float64) ElseFunc(f func() float64) (value float64)](#Float64.ElseFunc)
  * [func (o Float64) If(f func(value float64))](#Float64.If)
  * [func (o Float64) IsEmpty() bool](#Float64.IsEmpty)
  * [func (o Float64) IsPresent() bool](#Float64.IsPresent)
  * [func (o Float64) MarshalText() (text []byte, err error)](#Float64.MarshalText)
  * [func (o Float64) UnmarshalText(text []byte) error](#Float64.UnmarshalText)
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
  * [func (o Int) UnmarshalText(text []byte) error](#Int.UnmarshalText)
* [type Int16](#Int16)
  * [func EmptyInt16() Int16](#EmptyInt16)
  * [func OfInt16(value int16) Int16](#OfInt16)
  * [func OfInt16Ptr(ptr *int16) Int16](#OfInt16Ptr)
  * [func (o Int16) Else(elseValue int16) (value int16)](#Int16.Else)
  * [func (o Int16) ElseFunc(f func() int16) (value int16)](#Int16.ElseFunc)
  * [func (o Int16) If(f func(value int16))](#Int16.If)
  * [func (o Int16) IsEmpty() bool](#Int16.IsEmpty)
  * [func (o Int16) IsPresent() bool](#Int16.IsPresent)
  * [func (o Int16) MarshalText() (text []byte, err error)](#Int16.MarshalText)
  * [func (o Int16) UnmarshalText(text []byte) error](#Int16.UnmarshalText)
* [type Int32](#Int32)
  * [func EmptyInt32() Int32](#EmptyInt32)
  * [func OfInt32(value int32) Int32](#OfInt32)
  * [func OfInt32Ptr(ptr *int32) Int32](#OfInt32Ptr)
  * [func (o Int32) Else(elseValue int32) (value int32)](#Int32.Else)
  * [func (o Int32) ElseFunc(f func() int32) (value int32)](#Int32.ElseFunc)
  * [func (o Int32) If(f func(value int32))](#Int32.If)
  * [func (o Int32) IsEmpty() bool](#Int32.IsEmpty)
  * [func (o Int32) IsPresent() bool](#Int32.IsPresent)
  * [func (o Int32) MarshalText() (text []byte, err error)](#Int32.MarshalText)
  * [func (o Int32) UnmarshalText(text []byte) error](#Int32.UnmarshalText)
* [type Int64](#Int64)
  * [func EmptyInt64() Int64](#EmptyInt64)
  * [func OfInt64(value int64) Int64](#OfInt64)
  * [func OfInt64Ptr(ptr *int64) Int64](#OfInt64Ptr)
  * [func (o Int64) Else(elseValue int64) (value int64)](#Int64.Else)
  * [func (o Int64) ElseFunc(f func() int64) (value int64)](#Int64.ElseFunc)
  * [func (o Int64) If(f func(value int64))](#Int64.If)
  * [func (o Int64) IsEmpty() bool](#Int64.IsEmpty)
  * [func (o Int64) IsPresent() bool](#Int64.IsPresent)
  * [func (o Int64) MarshalText() (text []byte, err error)](#Int64.MarshalText)
  * [func (o Int64) UnmarshalText(text []byte) error](#Int64.UnmarshalText)
* [type Int8](#Int8)
  * [func EmptyInt8() Int8](#EmptyInt8)
  * [func OfInt8(value int8) Int8](#OfInt8)
  * [func OfInt8Ptr(ptr *int8) Int8](#OfInt8Ptr)
  * [func (o Int8) Else(elseValue int8) (value int8)](#Int8.Else)
  * [func (o Int8) ElseFunc(f func() int8) (value int8)](#Int8.ElseFunc)
  * [func (o Int8) If(f func(value int8))](#Int8.If)
  * [func (o Int8) IsEmpty() bool](#Int8.IsEmpty)
  * [func (o Int8) IsPresent() bool](#Int8.IsPresent)
  * [func (o Int8) MarshalText() (text []byte, err error)](#Int8.MarshalText)
  * [func (o Int8) UnmarshalText(text []byte) error](#Int8.UnmarshalText)
* [type Rune](#Rune)
  * [func EmptyRune() Rune](#EmptyRune)
  * [func OfRune(value rune) Rune](#OfRune)
  * [func OfRunePtr(ptr *rune) Rune](#OfRunePtr)
  * [func (o Rune) Else(elseValue rune) (value rune)](#Rune.Else)
  * [func (o Rune) ElseFunc(f func() rune) (value rune)](#Rune.ElseFunc)
  * [func (o Rune) If(f func(value rune))](#Rune.If)
  * [func (o Rune) IsEmpty() bool](#Rune.IsEmpty)
  * [func (o Rune) IsPresent() bool](#Rune.IsPresent)
  * [func (o Rune) MarshalText() (text []byte, err error)](#Rune.MarshalText)
  * [func (o Rune) UnmarshalText(text []byte) error](#Rune.UnmarshalText)
* [type String](#String)
  * [func EmptyString() String](#EmptyString)
  * [func OfString(value string) String](#OfString)
  * [func OfStringPtr(ptr *string) String](#OfStringPtr)
  * [func (o String) Else(elseValue string) (value string)](#String.Else)
  * [func (o String) ElseFunc(f func() string) (value string)](#String.ElseFunc)
  * [func (o String) If(f func(value string))](#String.If)
  * [func (o String) IsEmpty() bool](#String.IsEmpty)
  * [func (o String) IsPresent() bool](#String.IsPresent)
  * [func (o String) MarshalText() (text []byte, err error)](#String.MarshalText)
  * [func (o String) UnmarshalText(text []byte) error](#String.UnmarshalText)
* [type Uint](#Uint)
  * [func EmptyUint() Uint](#EmptyUint)
  * [func OfUint(value uint) Uint](#OfUint)
  * [func OfUintPtr(ptr *uint) Uint](#OfUintPtr)
  * [func (o Uint) Else(elseValue uint) (value uint)](#Uint.Else)
  * [func (o Uint) ElseFunc(f func() uint) (value uint)](#Uint.ElseFunc)
  * [func (o Uint) If(f func(value uint))](#Uint.If)
  * [func (o Uint) IsEmpty() bool](#Uint.IsEmpty)
  * [func (o Uint) IsPresent() bool](#Uint.IsPresent)
  * [func (o Uint) MarshalText() (text []byte, err error)](#Uint.MarshalText)
  * [func (o Uint) UnmarshalText(text []byte) error](#Uint.UnmarshalText)
* [type Uint16](#Uint16)
  * [func EmptyUint16() Uint16](#EmptyUint16)
  * [func OfUint16(value uint16) Uint16](#OfUint16)
  * [func OfUint16Ptr(ptr *uint16) Uint16](#OfUint16Ptr)
  * [func (o Uint16) Else(elseValue uint16) (value uint16)](#Uint16.Else)
  * [func (o Uint16) ElseFunc(f func() uint16) (value uint16)](#Uint16.ElseFunc)
  * [func (o Uint16) If(f func(value uint16))](#Uint16.If)
  * [func (o Uint16) IsEmpty() bool](#Uint16.IsEmpty)
  * [func (o Uint16) IsPresent() bool](#Uint16.IsPresent)
  * [func (o Uint16) MarshalText() (text []byte, err error)](#Uint16.MarshalText)
  * [func (o Uint16) UnmarshalText(text []byte) error](#Uint16.UnmarshalText)
* [type Uint32](#Uint32)
  * [func EmptyUint32() Uint32](#EmptyUint32)
  * [func OfUint32(value uint32) Uint32](#OfUint32)
  * [func OfUint32Ptr(ptr *uint32) Uint32](#OfUint32Ptr)
  * [func (o Uint32) Else(elseValue uint32) (value uint32)](#Uint32.Else)
  * [func (o Uint32) ElseFunc(f func() uint32) (value uint32)](#Uint32.ElseFunc)
  * [func (o Uint32) If(f func(value uint32))](#Uint32.If)
  * [func (o Uint32) IsEmpty() bool](#Uint32.IsEmpty)
  * [func (o Uint32) IsPresent() bool](#Uint32.IsPresent)
  * [func (o Uint32) MarshalText() (text []byte, err error)](#Uint32.MarshalText)
  * [func (o Uint32) UnmarshalText(text []byte) error](#Uint32.UnmarshalText)
* [type Uint64](#Uint64)
  * [func EmptyUint64() Uint64](#EmptyUint64)
  * [func OfUint64(value uint64) Uint64](#OfUint64)
  * [func OfUint64Ptr(ptr *uint64) Uint64](#OfUint64Ptr)
  * [func (o Uint64) Else(elseValue uint64) (value uint64)](#Uint64.Else)
  * [func (o Uint64) ElseFunc(f func() uint64) (value uint64)](#Uint64.ElseFunc)
  * [func (o Uint64) If(f func(value uint64))](#Uint64.If)
  * [func (o Uint64) IsEmpty() bool](#Uint64.IsEmpty)
  * [func (o Uint64) IsPresent() bool](#Uint64.IsPresent)
  * [func (o Uint64) MarshalText() (text []byte, err error)](#Uint64.MarshalText)
  * [func (o Uint64) UnmarshalText(text []byte) error](#Uint64.UnmarshalText)
* [type Uint8](#Uint8)
  * [func EmptyUint8() Uint8](#EmptyUint8)
  * [func OfUint8(value uint8) Uint8](#OfUint8)
  * [func OfUint8Ptr(ptr *uint8) Uint8](#OfUint8Ptr)
  * [func (o Uint8) Else(elseValue uint8) (value uint8)](#Uint8.Else)
  * [func (o Uint8) ElseFunc(f func() uint8) (value uint8)](#Uint8.ElseFunc)
  * [func (o Uint8) If(f func(value uint8))](#Uint8.If)
  * [func (o Uint8) IsEmpty() bool](#Uint8.IsEmpty)
  * [func (o Uint8) IsPresent() bool](#Uint8.IsPresent)
  * [func (o Uint8) MarshalText() (text []byte, err error)](#Uint8.MarshalText)
  * [func (o Uint8) UnmarshalText(text []byte) error](#Uint8.UnmarshalText)
* [type Uintptr](#Uintptr)
  * [func EmptyUintptr() Uintptr](#EmptyUintptr)
  * [func OfUintptr(value uintptr) Uintptr](#OfUintptr)
  * [func OfUintptrPtr(ptr *uintptr) Uintptr](#OfUintptrPtr)
  * [func (o Uintptr) Else(elseValue uintptr) (value uintptr)](#Uintptr.Else)
  * [func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr)](#Uintptr.ElseFunc)
  * [func (o Uintptr) If(f func(value uintptr))](#Uintptr.If)
  * [func (o Uintptr) IsEmpty() bool](#Uintptr.IsEmpty)
  * [func (o Uintptr) IsPresent() bool](#Uintptr.IsPresent)
  * [func (o Uintptr) MarshalText() (text []byte, err error)](#Uintptr.MarshalText)
  * [func (o Uintptr) UnmarshalText(text []byte) error](#Uintptr.UnmarshalText)

#### <a name="pkg-examples">Examples</a>
* [Int.Else](#example_Int_Else)
* [Int.If (Present)](#example_Int_If_present)

#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/leighmcculloch/optional/doc.go) [gotemplate_Bool.go](/src/github.com/leighmcculloch/optional/gotemplate_Bool.go) [gotemplate_Byte.go](/src/github.com/leighmcculloch/optional/gotemplate_Byte.go) [gotemplate_Complex128.go](/src/github.com/leighmcculloch/optional/gotemplate_Complex128.go) [gotemplate_Complex64.go](/src/github.com/leighmcculloch/optional/gotemplate_Complex64.go) [gotemplate_Error.go](/src/github.com/leighmcculloch/optional/gotemplate_Error.go) [gotemplate_Float32.go](/src/github.com/leighmcculloch/optional/gotemplate_Float32.go) [gotemplate_Float64.go](/src/github.com/leighmcculloch/optional/gotemplate_Float64.go) [gotemplate_Int.go](/src/github.com/leighmcculloch/optional/gotemplate_Int.go) [gotemplate_Int16.go](/src/github.com/leighmcculloch/optional/gotemplate_Int16.go) [gotemplate_Int32.go](/src/github.com/leighmcculloch/optional/gotemplate_Int32.go) [gotemplate_Int64.go](/src/github.com/leighmcculloch/optional/gotemplate_Int64.go) [gotemplate_Int8.go](/src/github.com/leighmcculloch/optional/gotemplate_Int8.go) [gotemplate_Rune.go](/src/github.com/leighmcculloch/optional/gotemplate_Rune.go) [gotemplate_String.go](/src/github.com/leighmcculloch/optional/gotemplate_String.go) [gotemplate_Uint.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint.go) [gotemplate_Uint16.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint16.go) [gotemplate_Uint32.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint32.go) [gotemplate_Uint64.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint64.go) [gotemplate_Uint8.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint8.go) [gotemplate_Uintptr.go](/src/github.com/leighmcculloch/optional/gotemplate_Uintptr.go) [types.go](/src/github.com/leighmcculloch/optional/types.go) 






## <a name="Bool">type</a> [Bool](/src/target/gotemplate_Bool.go?s=214:241#L2)
``` go
type Bool map[keyBool]*bool
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyBool">func</a> [EmptyBool](/src/target/gotemplate_Bool.go?s=553:574#L24)
``` go
func EmptyBool() Bool
```
Empty returns an empty Optional.


### <a name="OfBool">func</a> [OfBool](/src/target/gotemplate_Bool.go?s=339:367#L11)
``` go
func OfBool(value bool) Bool
```
Of wraps the value in an Optional.


### <a name="OfBoolPtr">func</a> [OfBoolPtr](/src/target/gotemplate_Bool.go?s=408:438#L15)
``` go
func OfBoolPtr(ptr *bool) Bool
```




### <a name="Bool.Else">func</a> (Bool) [Else](/src/target/gotemplate_Bool.go?s=1276:1323#L56)
``` go
func (o Bool) Else(elseValue bool) (value bool)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Bool.ElseFunc">func</a> (Bool) [ElseFunc](/src/target/gotemplate_Bool.go?s=1007:1057#L45)
``` go
func (o Bool) ElseFunc(f func() bool) (value bool)
```



### <a name="Bool.If">func</a> (Bool) [If](/src/target/gotemplate_Bool.go?s=920:956#L39)
``` go
func (o Bool) If(f func(value bool))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Bool.IsEmpty">func</a> (Bool) [IsEmpty](/src/target/gotemplate_Bool.go?s=669:697#L29)
``` go
func (o Bool) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Bool.IsPresent">func</a> (Bool) [IsPresent](/src/target/gotemplate_Bool.go?s=792:822#L34)
``` go
func (o Bool) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Bool.MarshalText">func</a> (Bool) [MarshalText](/src/target/gotemplate_Bool.go?s=1382:1434#L60)
``` go
func (o Bool) MarshalText() (text []byte, err error)
```



### <a name="Bool.UnmarshalText">func</a> (Bool) [UnmarshalText](/src/target/gotemplate_Bool.go?s=1631:1677#L74)
``` go
func (o Bool) UnmarshalText(text []byte) error
```



## <a name="Byte">type</a> [Byte](/src/target/gotemplate_Byte.go?s=214:241#L2)
``` go
type Byte map[keyByte]*byte
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyByte">func</a> [EmptyByte](/src/target/gotemplate_Byte.go?s=553:574#L24)
``` go
func EmptyByte() Byte
```
Empty returns an empty Optional.


### <a name="OfByte">func</a> [OfByte](/src/target/gotemplate_Byte.go?s=339:367#L11)
``` go
func OfByte(value byte) Byte
```
Of wraps the value in an Optional.


### <a name="OfBytePtr">func</a> [OfBytePtr](/src/target/gotemplate_Byte.go?s=408:438#L15)
``` go
func OfBytePtr(ptr *byte) Byte
```




### <a name="Byte.Else">func</a> (Byte) [Else](/src/target/gotemplate_Byte.go?s=1276:1323#L56)
``` go
func (o Byte) Else(elseValue byte) (value byte)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Byte.ElseFunc">func</a> (Byte) [ElseFunc](/src/target/gotemplate_Byte.go?s=1007:1057#L45)
``` go
func (o Byte) ElseFunc(f func() byte) (value byte)
```



### <a name="Byte.If">func</a> (Byte) [If](/src/target/gotemplate_Byte.go?s=920:956#L39)
``` go
func (o Byte) If(f func(value byte))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Byte.IsEmpty">func</a> (Byte) [IsEmpty](/src/target/gotemplate_Byte.go?s=669:697#L29)
``` go
func (o Byte) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Byte.IsPresent">func</a> (Byte) [IsPresent](/src/target/gotemplate_Byte.go?s=792:822#L34)
``` go
func (o Byte) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Byte.MarshalText">func</a> (Byte) [MarshalText](/src/target/gotemplate_Byte.go?s=1382:1434#L60)
``` go
func (o Byte) MarshalText() (text []byte, err error)
```



### <a name="Byte.UnmarshalText">func</a> (Byte) [UnmarshalText](/src/target/gotemplate_Byte.go?s=1631:1677#L74)
``` go
func (o Byte) UnmarshalText(text []byte) error
```



## <a name="Complex128">type</a> [Complex128](/src/target/gotemplate_Complex128.go?s=214:259#L2)
``` go
type Complex128 map[keyComplex128]*complex128
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyComplex128">func</a> [EmptyComplex128](/src/target/gotemplate_Complex128.go?s=649:682#L24)
``` go
func EmptyComplex128() Complex128
```
Empty returns an empty Optional.


### <a name="OfComplex128">func</a> [OfComplex128](/src/target/gotemplate_Complex128.go?s=375:421#L11)
``` go
func OfComplex128(value complex128) Complex128
```
Of wraps the value in an Optional.


### <a name="OfComplex128Ptr">func</a> [OfComplex128Ptr](/src/target/gotemplate_Complex128.go?s=474:522#L15)
``` go
func OfComplex128Ptr(ptr *complex128) Complex128
```




### <a name="Complex128.Else">func</a> (Complex128) [Else](/src/target/gotemplate_Complex128.go?s=1438:1503#L56)
``` go
func (o Complex128) Else(elseValue complex128) (value complex128)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Complex128.ElseFunc">func</a> (Complex128) [ElseFunc](/src/target/gotemplate_Complex128.go?s=1145:1213#L45)
``` go
func (o Complex128) ElseFunc(f func() complex128) (value complex128)
```



### <a name="Complex128.If">func</a> (Complex128) [If](/src/target/gotemplate_Complex128.go?s=1040:1088#L39)
``` go
func (o Complex128) If(f func(value complex128))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Complex128.IsEmpty">func</a> (Complex128) [IsEmpty](/src/target/gotemplate_Complex128.go?s=777:811#L29)
``` go
func (o Complex128) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Complex128.IsPresent">func</a> (Complex128) [IsPresent](/src/target/gotemplate_Complex128.go?s=906:942#L34)
``` go
func (o Complex128) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Complex128.MarshalText">func</a> (Complex128) [MarshalText](/src/target/gotemplate_Complex128.go?s=1568:1626#L60)
``` go
func (o Complex128) MarshalText() (text []byte, err error)
```



### <a name="Complex128.UnmarshalText">func</a> (Complex128) [UnmarshalText](/src/target/gotemplate_Complex128.go?s=1829:1881#L74)
``` go
func (o Complex128) UnmarshalText(text []byte) error
```



## <a name="Complex64">type</a> [Complex64](/src/target/gotemplate_Complex64.go?s=214:256#L2)
``` go
type Complex64 map[keyComplex64]*complex64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyComplex64">func</a> [EmptyComplex64](/src/target/gotemplate_Complex64.go?s=633:664#L24)
``` go
func EmptyComplex64() Complex64
```
Empty returns an empty Optional.


### <a name="OfComplex64">func</a> [OfComplex64](/src/target/gotemplate_Complex64.go?s=369:412#L11)
``` go
func OfComplex64(value complex64) Complex64
```
Of wraps the value in an Optional.


### <a name="OfComplex64Ptr">func</a> [OfComplex64Ptr](/src/target/gotemplate_Complex64.go?s=463:508#L15)
``` go
func OfComplex64Ptr(ptr *complex64) Complex64
```




### <a name="Complex64.Else">func</a> (Complex64) [Else](/src/target/gotemplate_Complex64.go?s=1411:1473#L56)
``` go
func (o Complex64) Else(elseValue complex64) (value complex64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Complex64.ElseFunc">func</a> (Complex64) [ElseFunc](/src/target/gotemplate_Complex64.go?s=1122:1187#L45)
``` go
func (o Complex64) ElseFunc(f func() complex64) (value complex64)
```



### <a name="Complex64.If">func</a> (Complex64) [If](/src/target/gotemplate_Complex64.go?s=1020:1066#L39)
``` go
func (o Complex64) If(f func(value complex64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Complex64.IsEmpty">func</a> (Complex64) [IsEmpty](/src/target/gotemplate_Complex64.go?s=759:792#L29)
``` go
func (o Complex64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Complex64.IsPresent">func</a> (Complex64) [IsPresent](/src/target/gotemplate_Complex64.go?s=887:922#L34)
``` go
func (o Complex64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Complex64.MarshalText">func</a> (Complex64) [MarshalText](/src/target/gotemplate_Complex64.go?s=1537:1594#L60)
``` go
func (o Complex64) MarshalText() (text []byte, err error)
```



### <a name="Complex64.UnmarshalText">func</a> (Complex64) [UnmarshalText](/src/target/gotemplate_Complex64.go?s=1796:1847#L74)
``` go
func (o Complex64) UnmarshalText(text []byte) error
```



## <a name="Error">type</a> [Error](/src/target/gotemplate_Error.go?s=214:244#L2)
``` go
type Error map[keyError]*error
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyError">func</a> [EmptyError](/src/target/gotemplate_Error.go?s=569:592#L24)
``` go
func EmptyError() Error
```
Empty returns an empty Optional.


### <a name="OfError">func</a> [OfError](/src/target/gotemplate_Error.go?s=345:376#L11)
``` go
func OfError(value error) Error
```
Of wraps the value in an Optional.


### <a name="OfErrorPtr">func</a> [OfErrorPtr](/src/target/gotemplate_Error.go?s=419:452#L15)
``` go
func OfErrorPtr(ptr *error) Error
```




### <a name="Error.Else">func</a> (Error) [Else](/src/target/gotemplate_Error.go?s=1303:1353#L56)
``` go
func (o Error) Else(elseValue error) (value error)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Error.ElseFunc">func</a> (Error) [ElseFunc](/src/target/gotemplate_Error.go?s=1030:1083#L45)
``` go
func (o Error) ElseFunc(f func() error) (value error)
```



### <a name="Error.If">func</a> (Error) [If](/src/target/gotemplate_Error.go?s=940:978#L39)
``` go
func (o Error) If(f func(value error))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Error.IsEmpty">func</a> (Error) [IsEmpty](/src/target/gotemplate_Error.go?s=687:716#L29)
``` go
func (o Error) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Error.IsPresent">func</a> (Error) [IsPresent](/src/target/gotemplate_Error.go?s=811:842#L34)
``` go
func (o Error) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Error.MarshalText">func</a> (Error) [MarshalText](/src/target/gotemplate_Error.go?s=1413:1466#L60)
``` go
func (o Error) MarshalText() (text []byte, err error)
```



### <a name="Error.UnmarshalText">func</a> (Error) [UnmarshalText](/src/target/gotemplate_Error.go?s=1664:1711#L74)
``` go
func (o Error) UnmarshalText(text []byte) error
```



## <a name="Float32">type</a> [Float32](/src/target/gotemplate_Float32.go?s=214:250#L2)
``` go
type Float32 map[keyFloat32]*float32
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyFloat32">func</a> [EmptyFloat32](/src/target/gotemplate_Float32.go?s=601:628#L24)
``` go
func EmptyFloat32() Float32
```
Empty returns an empty Optional.


### <a name="OfFloat32">func</a> [OfFloat32](/src/target/gotemplate_Float32.go?s=357:394#L11)
``` go
func OfFloat32(value float32) Float32
```
Of wraps the value in an Optional.


### <a name="OfFloat32Ptr">func</a> [OfFloat32Ptr](/src/target/gotemplate_Float32.go?s=441:480#L15)
``` go
func OfFloat32Ptr(ptr *float32) Float32
```




### <a name="Float32.Else">func</a> (Float32) [Else](/src/target/gotemplate_Float32.go?s=1357:1413#L56)
``` go
func (o Float32) Else(elseValue float32) (value float32)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Float32.ElseFunc">func</a> (Float32) [ElseFunc](/src/target/gotemplate_Float32.go?s=1076:1135#L45)
``` go
func (o Float32) ElseFunc(f func() float32) (value float32)
```



### <a name="Float32.If">func</a> (Float32) [If](/src/target/gotemplate_Float32.go?s=980:1022#L39)
``` go
func (o Float32) If(f func(value float32))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Float32.IsEmpty">func</a> (Float32) [IsEmpty](/src/target/gotemplate_Float32.go?s=723:754#L29)
``` go
func (o Float32) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Float32.IsPresent">func</a> (Float32) [IsPresent](/src/target/gotemplate_Float32.go?s=849:882#L34)
``` go
func (o Float32) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Float32.MarshalText">func</a> (Float32) [MarshalText](/src/target/gotemplate_Float32.go?s=1475:1530#L60)
``` go
func (o Float32) MarshalText() (text []byte, err error)
```



### <a name="Float32.UnmarshalText">func</a> (Float32) [UnmarshalText](/src/target/gotemplate_Float32.go?s=1730:1779#L74)
``` go
func (o Float32) UnmarshalText(text []byte) error
```



## <a name="Float64">type</a> [Float64](/src/target/gotemplate_Float64.go?s=214:250#L2)
``` go
type Float64 map[keyFloat64]*float64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyFloat64">func</a> [EmptyFloat64](/src/target/gotemplate_Float64.go?s=601:628#L24)
``` go
func EmptyFloat64() Float64
```
Empty returns an empty Optional.


### <a name="OfFloat64">func</a> [OfFloat64](/src/target/gotemplate_Float64.go?s=357:394#L11)
``` go
func OfFloat64(value float64) Float64
```
Of wraps the value in an Optional.


### <a name="OfFloat64Ptr">func</a> [OfFloat64Ptr](/src/target/gotemplate_Float64.go?s=441:480#L15)
``` go
func OfFloat64Ptr(ptr *float64) Float64
```




### <a name="Float64.Else">func</a> (Float64) [Else](/src/target/gotemplate_Float64.go?s=1357:1413#L56)
``` go
func (o Float64) Else(elseValue float64) (value float64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Float64.ElseFunc">func</a> (Float64) [ElseFunc](/src/target/gotemplate_Float64.go?s=1076:1135#L45)
``` go
func (o Float64) ElseFunc(f func() float64) (value float64)
```



### <a name="Float64.If">func</a> (Float64) [If](/src/target/gotemplate_Float64.go?s=980:1022#L39)
``` go
func (o Float64) If(f func(value float64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Float64.IsEmpty">func</a> (Float64) [IsEmpty](/src/target/gotemplate_Float64.go?s=723:754#L29)
``` go
func (o Float64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Float64.IsPresent">func</a> (Float64) [IsPresent](/src/target/gotemplate_Float64.go?s=849:882#L34)
``` go
func (o Float64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Float64.MarshalText">func</a> (Float64) [MarshalText](/src/target/gotemplate_Float64.go?s=1475:1530#L60)
``` go
func (o Float64) MarshalText() (text []byte, err error)
```



### <a name="Float64.UnmarshalText">func</a> (Float64) [UnmarshalText](/src/target/gotemplate_Float64.go?s=1730:1779#L74)
``` go
func (o Float64) UnmarshalText(text []byte) error
```



## <a name="Int">type</a> [Int](/src/target/gotemplate_Int.go?s=214:238#L2)
``` go
type Int map[keyInt]*int
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt">func</a> [EmptyInt](/src/target/gotemplate_Int.go?s=537:556#L24)
``` go
func EmptyInt() Int
```
Empty returns an empty Optional.


### <a name="OfInt">func</a> [OfInt](/src/target/gotemplate_Int.go?s=333:358#L11)
``` go
func OfInt(value int) Int
```
Of wraps the value in an Optional.


### <a name="OfIntPtr">func</a> [OfIntPtr](/src/target/gotemplate_Int.go?s=397:424#L15)
``` go
func OfIntPtr(ptr *int) Int
```




### <a name="Int.Else">func</a> (Int) [Else](/src/target/gotemplate_Int.go?s=1249:1293#L56)
``` go
func (o Int) Else(elseValue int) (value int)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int.ElseFunc">func</a> (Int) [ElseFunc](/src/target/gotemplate_Int.go?s=984:1031#L45)
``` go
func (o Int) ElseFunc(f func() int) (value int)
```



### <a name="Int.If">func</a> (Int) [If](/src/target/gotemplate_Int.go?s=900:934#L39)
``` go
func (o Int) If(f func(value int))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int.IsEmpty">func</a> (Int) [IsEmpty](/src/target/gotemplate_Int.go?s=651:678#L29)
``` go
func (o Int) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int.IsPresent">func</a> (Int) [IsPresent](/src/target/gotemplate_Int.go?s=773:802#L34)
``` go
func (o Int) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int.MarshalText">func</a> (Int) [MarshalText](/src/target/gotemplate_Int.go?s=1351:1402#L60)
``` go
func (o Int) MarshalText() (text []byte, err error)
```



### <a name="Int.UnmarshalText">func</a> (Int) [UnmarshalText](/src/target/gotemplate_Int.go?s=1598:1643#L74)
``` go
func (o Int) UnmarshalText(text []byte) error
```



## <a name="Int16">type</a> [Int16](/src/target/gotemplate_Int16.go?s=214:244#L2)
``` go
type Int16 map[keyInt16]*int16
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt16">func</a> [EmptyInt16](/src/target/gotemplate_Int16.go?s=569:592#L24)
``` go
func EmptyInt16() Int16
```
Empty returns an empty Optional.


### <a name="OfInt16">func</a> [OfInt16](/src/target/gotemplate_Int16.go?s=345:376#L11)
``` go
func OfInt16(value int16) Int16
```
Of wraps the value in an Optional.


### <a name="OfInt16Ptr">func</a> [OfInt16Ptr](/src/target/gotemplate_Int16.go?s=419:452#L15)
``` go
func OfInt16Ptr(ptr *int16) Int16
```




### <a name="Int16.Else">func</a> (Int16) [Else](/src/target/gotemplate_Int16.go?s=1303:1353#L56)
``` go
func (o Int16) Else(elseValue int16) (value int16)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int16.ElseFunc">func</a> (Int16) [ElseFunc](/src/target/gotemplate_Int16.go?s=1030:1083#L45)
``` go
func (o Int16) ElseFunc(f func() int16) (value int16)
```



### <a name="Int16.If">func</a> (Int16) [If](/src/target/gotemplate_Int16.go?s=940:978#L39)
``` go
func (o Int16) If(f func(value int16))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int16.IsEmpty">func</a> (Int16) [IsEmpty](/src/target/gotemplate_Int16.go?s=687:716#L29)
``` go
func (o Int16) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int16.IsPresent">func</a> (Int16) [IsPresent](/src/target/gotemplate_Int16.go?s=811:842#L34)
``` go
func (o Int16) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int16.MarshalText">func</a> (Int16) [MarshalText](/src/target/gotemplate_Int16.go?s=1413:1466#L60)
``` go
func (o Int16) MarshalText() (text []byte, err error)
```



### <a name="Int16.UnmarshalText">func</a> (Int16) [UnmarshalText](/src/target/gotemplate_Int16.go?s=1664:1711#L74)
``` go
func (o Int16) UnmarshalText(text []byte) error
```



## <a name="Int32">type</a> [Int32](/src/target/gotemplate_Int32.go?s=214:244#L2)
``` go
type Int32 map[keyInt32]*int32
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt32">func</a> [EmptyInt32](/src/target/gotemplate_Int32.go?s=569:592#L24)
``` go
func EmptyInt32() Int32
```
Empty returns an empty Optional.


### <a name="OfInt32">func</a> [OfInt32](/src/target/gotemplate_Int32.go?s=345:376#L11)
``` go
func OfInt32(value int32) Int32
```
Of wraps the value in an Optional.


### <a name="OfInt32Ptr">func</a> [OfInt32Ptr](/src/target/gotemplate_Int32.go?s=419:452#L15)
``` go
func OfInt32Ptr(ptr *int32) Int32
```




### <a name="Int32.Else">func</a> (Int32) [Else](/src/target/gotemplate_Int32.go?s=1303:1353#L56)
``` go
func (o Int32) Else(elseValue int32) (value int32)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int32.ElseFunc">func</a> (Int32) [ElseFunc](/src/target/gotemplate_Int32.go?s=1030:1083#L45)
``` go
func (o Int32) ElseFunc(f func() int32) (value int32)
```



### <a name="Int32.If">func</a> (Int32) [If](/src/target/gotemplate_Int32.go?s=940:978#L39)
``` go
func (o Int32) If(f func(value int32))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int32.IsEmpty">func</a> (Int32) [IsEmpty](/src/target/gotemplate_Int32.go?s=687:716#L29)
``` go
func (o Int32) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int32.IsPresent">func</a> (Int32) [IsPresent](/src/target/gotemplate_Int32.go?s=811:842#L34)
``` go
func (o Int32) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int32.MarshalText">func</a> (Int32) [MarshalText](/src/target/gotemplate_Int32.go?s=1413:1466#L60)
``` go
func (o Int32) MarshalText() (text []byte, err error)
```



### <a name="Int32.UnmarshalText">func</a> (Int32) [UnmarshalText](/src/target/gotemplate_Int32.go?s=1664:1711#L74)
``` go
func (o Int32) UnmarshalText(text []byte) error
```



## <a name="Int64">type</a> [Int64](/src/target/gotemplate_Int64.go?s=214:244#L2)
``` go
type Int64 map[keyInt64]*int64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt64">func</a> [EmptyInt64](/src/target/gotemplate_Int64.go?s=569:592#L24)
``` go
func EmptyInt64() Int64
```
Empty returns an empty Optional.


### <a name="OfInt64">func</a> [OfInt64](/src/target/gotemplate_Int64.go?s=345:376#L11)
``` go
func OfInt64(value int64) Int64
```
Of wraps the value in an Optional.


### <a name="OfInt64Ptr">func</a> [OfInt64Ptr](/src/target/gotemplate_Int64.go?s=419:452#L15)
``` go
func OfInt64Ptr(ptr *int64) Int64
```




### <a name="Int64.Else">func</a> (Int64) [Else](/src/target/gotemplate_Int64.go?s=1303:1353#L56)
``` go
func (o Int64) Else(elseValue int64) (value int64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int64.ElseFunc">func</a> (Int64) [ElseFunc](/src/target/gotemplate_Int64.go?s=1030:1083#L45)
``` go
func (o Int64) ElseFunc(f func() int64) (value int64)
```



### <a name="Int64.If">func</a> (Int64) [If](/src/target/gotemplate_Int64.go?s=940:978#L39)
``` go
func (o Int64) If(f func(value int64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int64.IsEmpty">func</a> (Int64) [IsEmpty](/src/target/gotemplate_Int64.go?s=687:716#L29)
``` go
func (o Int64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int64.IsPresent">func</a> (Int64) [IsPresent](/src/target/gotemplate_Int64.go?s=811:842#L34)
``` go
func (o Int64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int64.MarshalText">func</a> (Int64) [MarshalText](/src/target/gotemplate_Int64.go?s=1413:1466#L60)
``` go
func (o Int64) MarshalText() (text []byte, err error)
```



### <a name="Int64.UnmarshalText">func</a> (Int64) [UnmarshalText](/src/target/gotemplate_Int64.go?s=1664:1711#L74)
``` go
func (o Int64) UnmarshalText(text []byte) error
```



## <a name="Int8">type</a> [Int8](/src/target/gotemplate_Int8.go?s=214:241#L2)
``` go
type Int8 map[keyInt8]*int8
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt8">func</a> [EmptyInt8](/src/target/gotemplate_Int8.go?s=553:574#L24)
``` go
func EmptyInt8() Int8
```
Empty returns an empty Optional.


### <a name="OfInt8">func</a> [OfInt8](/src/target/gotemplate_Int8.go?s=339:367#L11)
``` go
func OfInt8(value int8) Int8
```
Of wraps the value in an Optional.


### <a name="OfInt8Ptr">func</a> [OfInt8Ptr](/src/target/gotemplate_Int8.go?s=408:438#L15)
``` go
func OfInt8Ptr(ptr *int8) Int8
```




### <a name="Int8.Else">func</a> (Int8) [Else](/src/target/gotemplate_Int8.go?s=1276:1323#L56)
``` go
func (o Int8) Else(elseValue int8) (value int8)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Int8.ElseFunc">func</a> (Int8) [ElseFunc](/src/target/gotemplate_Int8.go?s=1007:1057#L45)
``` go
func (o Int8) ElseFunc(f func() int8) (value int8)
```



### <a name="Int8.If">func</a> (Int8) [If](/src/target/gotemplate_Int8.go?s=920:956#L39)
``` go
func (o Int8) If(f func(value int8))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Int8.IsEmpty">func</a> (Int8) [IsEmpty](/src/target/gotemplate_Int8.go?s=669:697#L29)
``` go
func (o Int8) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Int8.IsPresent">func</a> (Int8) [IsPresent](/src/target/gotemplate_Int8.go?s=792:822#L34)
``` go
func (o Int8) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Int8.MarshalText">func</a> (Int8) [MarshalText](/src/target/gotemplate_Int8.go?s=1382:1434#L60)
``` go
func (o Int8) MarshalText() (text []byte, err error)
```



### <a name="Int8.UnmarshalText">func</a> (Int8) [UnmarshalText](/src/target/gotemplate_Int8.go?s=1631:1677#L74)
``` go
func (o Int8) UnmarshalText(text []byte) error
```



## <a name="Rune">type</a> [Rune](/src/target/gotemplate_Rune.go?s=214:241#L2)
``` go
type Rune map[keyRune]*rune
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyRune">func</a> [EmptyRune](/src/target/gotemplate_Rune.go?s=553:574#L24)
``` go
func EmptyRune() Rune
```
Empty returns an empty Optional.


### <a name="OfRune">func</a> [OfRune](/src/target/gotemplate_Rune.go?s=339:367#L11)
``` go
func OfRune(value rune) Rune
```
Of wraps the value in an Optional.


### <a name="OfRunePtr">func</a> [OfRunePtr](/src/target/gotemplate_Rune.go?s=408:438#L15)
``` go
func OfRunePtr(ptr *rune) Rune
```




### <a name="Rune.Else">func</a> (Rune) [Else](/src/target/gotemplate_Rune.go?s=1276:1323#L56)
``` go
func (o Rune) Else(elseValue rune) (value rune)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Rune.ElseFunc">func</a> (Rune) [ElseFunc](/src/target/gotemplate_Rune.go?s=1007:1057#L45)
``` go
func (o Rune) ElseFunc(f func() rune) (value rune)
```



### <a name="Rune.If">func</a> (Rune) [If](/src/target/gotemplate_Rune.go?s=920:956#L39)
``` go
func (o Rune) If(f func(value rune))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Rune.IsEmpty">func</a> (Rune) [IsEmpty](/src/target/gotemplate_Rune.go?s=669:697#L29)
``` go
func (o Rune) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Rune.IsPresent">func</a> (Rune) [IsPresent](/src/target/gotemplate_Rune.go?s=792:822#L34)
``` go
func (o Rune) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Rune.MarshalText">func</a> (Rune) [MarshalText](/src/target/gotemplate_Rune.go?s=1382:1434#L60)
``` go
func (o Rune) MarshalText() (text []byte, err error)
```



### <a name="Rune.UnmarshalText">func</a> (Rune) [UnmarshalText](/src/target/gotemplate_Rune.go?s=1631:1677#L74)
``` go
func (o Rune) UnmarshalText(text []byte) error
```



## <a name="String">type</a> [String](/src/target/gotemplate_String.go?s=214:247#L2)
``` go
type String map[keyString]*string
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyString">func</a> [EmptyString](/src/target/gotemplate_String.go?s=585:610#L24)
``` go
func EmptyString() String
```
Empty returns an empty Optional.


### <a name="OfString">func</a> [OfString](/src/target/gotemplate_String.go?s=351:385#L11)
``` go
func OfString(value string) String
```
Of wraps the value in an Optional.


### <a name="OfStringPtr">func</a> [OfStringPtr](/src/target/gotemplate_String.go?s=430:466#L15)
``` go
func OfStringPtr(ptr *string) String
```




### <a name="String.Else">func</a> (String) [Else](/src/target/gotemplate_String.go?s=1330:1383#L56)
``` go
func (o String) Else(elseValue string) (value string)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="String.ElseFunc">func</a> (String) [ElseFunc](/src/target/gotemplate_String.go?s=1053:1109#L45)
``` go
func (o String) ElseFunc(f func() string) (value string)
```



### <a name="String.If">func</a> (String) [If](/src/target/gotemplate_String.go?s=960:1000#L39)
``` go
func (o String) If(f func(value string))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="String.IsEmpty">func</a> (String) [IsEmpty](/src/target/gotemplate_String.go?s=705:735#L29)
``` go
func (o String) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="String.IsPresent">func</a> (String) [IsPresent](/src/target/gotemplate_String.go?s=830:862#L34)
``` go
func (o String) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="String.MarshalText">func</a> (String) [MarshalText](/src/target/gotemplate_String.go?s=1444:1498#L60)
``` go
func (o String) MarshalText() (text []byte, err error)
```



### <a name="String.UnmarshalText">func</a> (String) [UnmarshalText](/src/target/gotemplate_String.go?s=1697:1745#L74)
``` go
func (o String) UnmarshalText(text []byte) error
```



## <a name="Uint">type</a> [Uint](/src/target/gotemplate_Uint.go?s=214:241#L2)
``` go
type Uint map[keyUint]*uint
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint">func</a> [EmptyUint](/src/target/gotemplate_Uint.go?s=553:574#L24)
``` go
func EmptyUint() Uint
```
Empty returns an empty Optional.


### <a name="OfUint">func</a> [OfUint](/src/target/gotemplate_Uint.go?s=339:367#L11)
``` go
func OfUint(value uint) Uint
```
Of wraps the value in an Optional.


### <a name="OfUintPtr">func</a> [OfUintPtr](/src/target/gotemplate_Uint.go?s=408:438#L15)
``` go
func OfUintPtr(ptr *uint) Uint
```




### <a name="Uint.Else">func</a> (Uint) [Else](/src/target/gotemplate_Uint.go?s=1276:1323#L56)
``` go
func (o Uint) Else(elseValue uint) (value uint)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint.ElseFunc">func</a> (Uint) [ElseFunc](/src/target/gotemplate_Uint.go?s=1007:1057#L45)
``` go
func (o Uint) ElseFunc(f func() uint) (value uint)
```



### <a name="Uint.If">func</a> (Uint) [If](/src/target/gotemplate_Uint.go?s=920:956#L39)
``` go
func (o Uint) If(f func(value uint))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint.IsEmpty">func</a> (Uint) [IsEmpty](/src/target/gotemplate_Uint.go?s=669:697#L29)
``` go
func (o Uint) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint.IsPresent">func</a> (Uint) [IsPresent](/src/target/gotemplate_Uint.go?s=792:822#L34)
``` go
func (o Uint) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint.MarshalText">func</a> (Uint) [MarshalText](/src/target/gotemplate_Uint.go?s=1382:1434#L60)
``` go
func (o Uint) MarshalText() (text []byte, err error)
```



### <a name="Uint.UnmarshalText">func</a> (Uint) [UnmarshalText](/src/target/gotemplate_Uint.go?s=1631:1677#L74)
``` go
func (o Uint) UnmarshalText(text []byte) error
```



## <a name="Uint16">type</a> [Uint16](/src/target/gotemplate_Uint16.go?s=214:247#L2)
``` go
type Uint16 map[keyUint16]*uint16
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint16">func</a> [EmptyUint16](/src/target/gotemplate_Uint16.go?s=585:610#L24)
``` go
func EmptyUint16() Uint16
```
Empty returns an empty Optional.


### <a name="OfUint16">func</a> [OfUint16](/src/target/gotemplate_Uint16.go?s=351:385#L11)
``` go
func OfUint16(value uint16) Uint16
```
Of wraps the value in an Optional.


### <a name="OfUint16Ptr">func</a> [OfUint16Ptr](/src/target/gotemplate_Uint16.go?s=430:466#L15)
``` go
func OfUint16Ptr(ptr *uint16) Uint16
```




### <a name="Uint16.Else">func</a> (Uint16) [Else](/src/target/gotemplate_Uint16.go?s=1330:1383#L56)
``` go
func (o Uint16) Else(elseValue uint16) (value uint16)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint16.ElseFunc">func</a> (Uint16) [ElseFunc](/src/target/gotemplate_Uint16.go?s=1053:1109#L45)
``` go
func (o Uint16) ElseFunc(f func() uint16) (value uint16)
```



### <a name="Uint16.If">func</a> (Uint16) [If](/src/target/gotemplate_Uint16.go?s=960:1000#L39)
``` go
func (o Uint16) If(f func(value uint16))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint16.IsEmpty">func</a> (Uint16) [IsEmpty](/src/target/gotemplate_Uint16.go?s=705:735#L29)
``` go
func (o Uint16) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint16.IsPresent">func</a> (Uint16) [IsPresent](/src/target/gotemplate_Uint16.go?s=830:862#L34)
``` go
func (o Uint16) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint16.MarshalText">func</a> (Uint16) [MarshalText](/src/target/gotemplate_Uint16.go?s=1444:1498#L60)
``` go
func (o Uint16) MarshalText() (text []byte, err error)
```



### <a name="Uint16.UnmarshalText">func</a> (Uint16) [UnmarshalText](/src/target/gotemplate_Uint16.go?s=1697:1745#L74)
``` go
func (o Uint16) UnmarshalText(text []byte) error
```



## <a name="Uint32">type</a> [Uint32](/src/target/gotemplate_Uint32.go?s=214:247#L2)
``` go
type Uint32 map[keyUint32]*uint32
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint32">func</a> [EmptyUint32](/src/target/gotemplate_Uint32.go?s=585:610#L24)
``` go
func EmptyUint32() Uint32
```
Empty returns an empty Optional.


### <a name="OfUint32">func</a> [OfUint32](/src/target/gotemplate_Uint32.go?s=351:385#L11)
``` go
func OfUint32(value uint32) Uint32
```
Of wraps the value in an Optional.


### <a name="OfUint32Ptr">func</a> [OfUint32Ptr](/src/target/gotemplate_Uint32.go?s=430:466#L15)
``` go
func OfUint32Ptr(ptr *uint32) Uint32
```




### <a name="Uint32.Else">func</a> (Uint32) [Else](/src/target/gotemplate_Uint32.go?s=1330:1383#L56)
``` go
func (o Uint32) Else(elseValue uint32) (value uint32)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint32.ElseFunc">func</a> (Uint32) [ElseFunc](/src/target/gotemplate_Uint32.go?s=1053:1109#L45)
``` go
func (o Uint32) ElseFunc(f func() uint32) (value uint32)
```



### <a name="Uint32.If">func</a> (Uint32) [If](/src/target/gotemplate_Uint32.go?s=960:1000#L39)
``` go
func (o Uint32) If(f func(value uint32))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint32.IsEmpty">func</a> (Uint32) [IsEmpty](/src/target/gotemplate_Uint32.go?s=705:735#L29)
``` go
func (o Uint32) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint32.IsPresent">func</a> (Uint32) [IsPresent](/src/target/gotemplate_Uint32.go?s=830:862#L34)
``` go
func (o Uint32) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint32.MarshalText">func</a> (Uint32) [MarshalText](/src/target/gotemplate_Uint32.go?s=1444:1498#L60)
``` go
func (o Uint32) MarshalText() (text []byte, err error)
```



### <a name="Uint32.UnmarshalText">func</a> (Uint32) [UnmarshalText](/src/target/gotemplate_Uint32.go?s=1697:1745#L74)
``` go
func (o Uint32) UnmarshalText(text []byte) error
```



## <a name="Uint64">type</a> [Uint64](/src/target/gotemplate_Uint64.go?s=214:247#L2)
``` go
type Uint64 map[keyUint64]*uint64
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint64">func</a> [EmptyUint64](/src/target/gotemplate_Uint64.go?s=585:610#L24)
``` go
func EmptyUint64() Uint64
```
Empty returns an empty Optional.


### <a name="OfUint64">func</a> [OfUint64](/src/target/gotemplate_Uint64.go?s=351:385#L11)
``` go
func OfUint64(value uint64) Uint64
```
Of wraps the value in an Optional.


### <a name="OfUint64Ptr">func</a> [OfUint64Ptr](/src/target/gotemplate_Uint64.go?s=430:466#L15)
``` go
func OfUint64Ptr(ptr *uint64) Uint64
```




### <a name="Uint64.Else">func</a> (Uint64) [Else](/src/target/gotemplate_Uint64.go?s=1330:1383#L56)
``` go
func (o Uint64) Else(elseValue uint64) (value uint64)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint64.ElseFunc">func</a> (Uint64) [ElseFunc](/src/target/gotemplate_Uint64.go?s=1053:1109#L45)
``` go
func (o Uint64) ElseFunc(f func() uint64) (value uint64)
```



### <a name="Uint64.If">func</a> (Uint64) [If](/src/target/gotemplate_Uint64.go?s=960:1000#L39)
``` go
func (o Uint64) If(f func(value uint64))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint64.IsEmpty">func</a> (Uint64) [IsEmpty](/src/target/gotemplate_Uint64.go?s=705:735#L29)
``` go
func (o Uint64) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint64.IsPresent">func</a> (Uint64) [IsPresent](/src/target/gotemplate_Uint64.go?s=830:862#L34)
``` go
func (o Uint64) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint64.MarshalText">func</a> (Uint64) [MarshalText](/src/target/gotemplate_Uint64.go?s=1444:1498#L60)
``` go
func (o Uint64) MarshalText() (text []byte, err error)
```



### <a name="Uint64.UnmarshalText">func</a> (Uint64) [UnmarshalText](/src/target/gotemplate_Uint64.go?s=1697:1745#L74)
``` go
func (o Uint64) UnmarshalText(text []byte) error
```



## <a name="Uint8">type</a> [Uint8](/src/target/gotemplate_Uint8.go?s=214:244#L2)
``` go
type Uint8 map[keyUint8]*uint8
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint8">func</a> [EmptyUint8](/src/target/gotemplate_Uint8.go?s=569:592#L24)
``` go
func EmptyUint8() Uint8
```
Empty returns an empty Optional.


### <a name="OfUint8">func</a> [OfUint8](/src/target/gotemplate_Uint8.go?s=345:376#L11)
``` go
func OfUint8(value uint8) Uint8
```
Of wraps the value in an Optional.


### <a name="OfUint8Ptr">func</a> [OfUint8Ptr](/src/target/gotemplate_Uint8.go?s=419:452#L15)
``` go
func OfUint8Ptr(ptr *uint8) Uint8
```




### <a name="Uint8.Else">func</a> (Uint8) [Else](/src/target/gotemplate_Uint8.go?s=1303:1353#L56)
``` go
func (o Uint8) Else(elseValue uint8) (value uint8)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uint8.ElseFunc">func</a> (Uint8) [ElseFunc](/src/target/gotemplate_Uint8.go?s=1030:1083#L45)
``` go
func (o Uint8) ElseFunc(f func() uint8) (value uint8)
```



### <a name="Uint8.If">func</a> (Uint8) [If](/src/target/gotemplate_Uint8.go?s=940:978#L39)
``` go
func (o Uint8) If(f func(value uint8))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uint8.IsEmpty">func</a> (Uint8) [IsEmpty](/src/target/gotemplate_Uint8.go?s=687:716#L29)
``` go
func (o Uint8) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uint8.IsPresent">func</a> (Uint8) [IsPresent](/src/target/gotemplate_Uint8.go?s=811:842#L34)
``` go
func (o Uint8) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uint8.MarshalText">func</a> (Uint8) [MarshalText](/src/target/gotemplate_Uint8.go?s=1413:1466#L60)
``` go
func (o Uint8) MarshalText() (text []byte, err error)
```



### <a name="Uint8.UnmarshalText">func</a> (Uint8) [UnmarshalText](/src/target/gotemplate_Uint8.go?s=1664:1711#L74)
``` go
func (o Uint8) UnmarshalText(text []byte) error
```



## <a name="Uintptr">type</a> [Uintptr](/src/target/gotemplate_Uintptr.go?s=214:250#L2)
``` go
type Uintptr map[keyUintptr]*uintptr
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUintptr">func</a> [EmptyUintptr](/src/target/gotemplate_Uintptr.go?s=601:628#L24)
``` go
func EmptyUintptr() Uintptr
```
Empty returns an empty Optional.


### <a name="OfUintptr">func</a> [OfUintptr](/src/target/gotemplate_Uintptr.go?s=357:394#L11)
``` go
func OfUintptr(value uintptr) Uintptr
```
Of wraps the value in an Optional.


### <a name="OfUintptrPtr">func</a> [OfUintptrPtr](/src/target/gotemplate_Uintptr.go?s=441:480#L15)
``` go
func OfUintptrPtr(ptr *uintptr) Uintptr
```




### <a name="Uintptr.Else">func</a> (Uintptr) [Else](/src/target/gotemplate_Uintptr.go?s=1357:1413#L56)
``` go
func (o Uintptr) Else(elseValue uintptr) (value uintptr)
```
Else returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




### <a name="Uintptr.ElseFunc">func</a> (Uintptr) [ElseFunc](/src/target/gotemplate_Uintptr.go?s=1076:1135#L45)
``` go
func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr)
```



### <a name="Uintptr.If">func</a> (Uintptr) [If](/src/target/gotemplate_Uintptr.go?s=980:1022#L39)
``` go
func (o Uintptr) If(f func(value uintptr))
```
If calls the function if there is a value wrapped by this Optional.




### <a name="Uintptr.IsEmpty">func</a> (Uintptr) [IsEmpty](/src/target/gotemplate_Uintptr.go?s=723:754#L29)
``` go
func (o Uintptr) IsEmpty() bool
```
IsEmpty returns true if there there is no value wrapped by this Optional.




### <a name="Uintptr.IsPresent">func</a> (Uintptr) [IsPresent](/src/target/gotemplate_Uintptr.go?s=849:882#L34)
``` go
func (o Uintptr) IsPresent() bool
```
IsPresent returns true if there is a value wrapped by this Optional.




### <a name="Uintptr.MarshalText">func</a> (Uintptr) [MarshalText](/src/target/gotemplate_Uintptr.go?s=1475:1530#L60)
``` go
func (o Uintptr) MarshalText() (text []byte, err error)
```



### <a name="Uintptr.UnmarshalText">func</a> (Uintptr) [UnmarshalText](/src/target/gotemplate_Uintptr.go?s=1730:1779#L74)
``` go
func (o Uintptr) UnmarshalText(text []byte) error
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
