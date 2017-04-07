

# optional
`import "github.com/leighmcculloch/optional"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package optional provides Optional types that wrap builtin types as an alternative to using pointers or zero values to represent the lack of something. The Optional types require conditional unwrapping to get to the inner value, which ensures by way of the compiler that you do not end up with a nil value.

### Examples
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
		optional.OfInt(2017),
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
  * [func OfNilableBool(ptr *bool) Bool](#OfNilableBool)
  * [func (o Bool) IfPresent(f func(value bool))](#Bool.IfPresent)
  * [func (o Bool) IsPresent() bool](#Bool.IsPresent)
  * [func (o Bool) OrElse(value bool) bool](#Bool.OrElse)
* [type Byte](#Byte)
  * [func EmptyByte() Byte](#EmptyByte)
  * [func OfByte(value byte) Byte](#OfByte)
  * [func OfNilableByte(ptr *byte) Byte](#OfNilableByte)
  * [func (o Byte) IfPresent(f func(value byte))](#Byte.IfPresent)
  * [func (o Byte) IsPresent() bool](#Byte.IsPresent)
  * [func (o Byte) OrElse(value byte) byte](#Byte.OrElse)
* [type Complex128](#Complex128)
  * [func EmptyComplex128() Complex128](#EmptyComplex128)
  * [func OfComplex128(value complex128) Complex128](#OfComplex128)
  * [func OfNilableComplex128(ptr *complex128) Complex128](#OfNilableComplex128)
  * [func (o Complex128) IfPresent(f func(value complex128))](#Complex128.IfPresent)
  * [func (o Complex128) IsPresent() bool](#Complex128.IsPresent)
  * [func (o Complex128) OrElse(value complex128) complex128](#Complex128.OrElse)
* [type Complex64](#Complex64)
  * [func EmptyComplex64() Complex64](#EmptyComplex64)
  * [func OfComplex64(value complex64) Complex64](#OfComplex64)
  * [func OfNilableComplex64(ptr *complex64) Complex64](#OfNilableComplex64)
  * [func (o Complex64) IfPresent(f func(value complex64))](#Complex64.IfPresent)
  * [func (o Complex64) IsPresent() bool](#Complex64.IsPresent)
  * [func (o Complex64) OrElse(value complex64) complex64](#Complex64.OrElse)
* [type Error](#Error)
  * [func EmptyError() Error](#EmptyError)
  * [func OfError(value error) Error](#OfError)
  * [func OfNilableError(ptr *error) Error](#OfNilableError)
  * [func (o Error) IfPresent(f func(value error))](#Error.IfPresent)
  * [func (o Error) IsPresent() bool](#Error.IsPresent)
  * [func (o Error) OrElse(value error) error](#Error.OrElse)
* [type Float32](#Float32)
  * [func EmptyFloat32() Float32](#EmptyFloat32)
  * [func OfFloat32(value float32) Float32](#OfFloat32)
  * [func OfNilableFloat32(ptr *float32) Float32](#OfNilableFloat32)
  * [func (o Float32) IfPresent(f func(value float32))](#Float32.IfPresent)
  * [func (o Float32) IsPresent() bool](#Float32.IsPresent)
  * [func (o Float32) OrElse(value float32) float32](#Float32.OrElse)
* [type Float64](#Float64)
  * [func EmptyFloat64() Float64](#EmptyFloat64)
  * [func OfFloat64(value float64) Float64](#OfFloat64)
  * [func OfNilableFloat64(ptr *float64) Float64](#OfNilableFloat64)
  * [func (o Float64) IfPresent(f func(value float64))](#Float64.IfPresent)
  * [func (o Float64) IsPresent() bool](#Float64.IsPresent)
  * [func (o Float64) OrElse(value float64) float64](#Float64.OrElse)
* [type Int](#Int)
  * [func EmptyInt() Int](#EmptyInt)
  * [func OfInt(value int) Int](#OfInt)
  * [func OfNilableInt(ptr *int) Int](#OfNilableInt)
  * [func (o Int) IfPresent(f func(value int))](#Int.IfPresent)
  * [func (o Int) IsPresent() bool](#Int.IsPresent)
  * [func (o Int) OrElse(value int) int](#Int.OrElse)
* [type Int16](#Int16)
  * [func EmptyInt16() Int16](#EmptyInt16)
  * [func OfInt16(value int16) Int16](#OfInt16)
  * [func OfNilableInt16(ptr *int16) Int16](#OfNilableInt16)
  * [func (o Int16) IfPresent(f func(value int16))](#Int16.IfPresent)
  * [func (o Int16) IsPresent() bool](#Int16.IsPresent)
  * [func (o Int16) OrElse(value int16) int16](#Int16.OrElse)
* [type Int32](#Int32)
  * [func EmptyInt32() Int32](#EmptyInt32)
  * [func OfInt32(value int32) Int32](#OfInt32)
  * [func OfNilableInt32(ptr *int32) Int32](#OfNilableInt32)
  * [func (o Int32) IfPresent(f func(value int32))](#Int32.IfPresent)
  * [func (o Int32) IsPresent() bool](#Int32.IsPresent)
  * [func (o Int32) OrElse(value int32) int32](#Int32.OrElse)
* [type Int64](#Int64)
  * [func EmptyInt64() Int64](#EmptyInt64)
  * [func OfInt64(value int64) Int64](#OfInt64)
  * [func OfNilableInt64(ptr *int64) Int64](#OfNilableInt64)
  * [func (o Int64) IfPresent(f func(value int64))](#Int64.IfPresent)
  * [func (o Int64) IsPresent() bool](#Int64.IsPresent)
  * [func (o Int64) OrElse(value int64) int64](#Int64.OrElse)
* [type Int8](#Int8)
  * [func EmptyInt8() Int8](#EmptyInt8)
  * [func OfInt8(value int8) Int8](#OfInt8)
  * [func OfNilableInt8(ptr *int8) Int8](#OfNilableInt8)
  * [func (o Int8) IfPresent(f func(value int8))](#Int8.IfPresent)
  * [func (o Int8) IsPresent() bool](#Int8.IsPresent)
  * [func (o Int8) OrElse(value int8) int8](#Int8.OrElse)
* [type Rune](#Rune)
  * [func EmptyRune() Rune](#EmptyRune)
  * [func OfNilableRune(ptr *rune) Rune](#OfNilableRune)
  * [func OfRune(value rune) Rune](#OfRune)
  * [func (o Rune) IfPresent(f func(value rune))](#Rune.IfPresent)
  * [func (o Rune) IsPresent() bool](#Rune.IsPresent)
  * [func (o Rune) OrElse(value rune) rune](#Rune.OrElse)
* [type String](#String)
  * [func EmptyString() String](#EmptyString)
  * [func OfNilableString(ptr *string) String](#OfNilableString)
  * [func OfString(value string) String](#OfString)
  * [func (o String) IfPresent(f func(value string))](#String.IfPresent)
  * [func (o String) IsPresent() bool](#String.IsPresent)
  * [func (o String) OrElse(value string) string](#String.OrElse)
* [type Uint](#Uint)
  * [func EmptyUint() Uint](#EmptyUint)
  * [func OfNilableUint(ptr *uint) Uint](#OfNilableUint)
  * [func OfUint(value uint) Uint](#OfUint)
  * [func (o Uint) IfPresent(f func(value uint))](#Uint.IfPresent)
  * [func (o Uint) IsPresent() bool](#Uint.IsPresent)
  * [func (o Uint) OrElse(value uint) uint](#Uint.OrElse)
* [type Uint16](#Uint16)
  * [func EmptyUint16() Uint16](#EmptyUint16)
  * [func OfNilableUint16(ptr *uint16) Uint16](#OfNilableUint16)
  * [func OfUint16(value uint16) Uint16](#OfUint16)
  * [func (o Uint16) IfPresent(f func(value uint16))](#Uint16.IfPresent)
  * [func (o Uint16) IsPresent() bool](#Uint16.IsPresent)
  * [func (o Uint16) OrElse(value uint16) uint16](#Uint16.OrElse)
* [type Uint32](#Uint32)
  * [func EmptyUint32() Uint32](#EmptyUint32)
  * [func OfNilableUint32(ptr *uint32) Uint32](#OfNilableUint32)
  * [func OfUint32(value uint32) Uint32](#OfUint32)
  * [func (o Uint32) IfPresent(f func(value uint32))](#Uint32.IfPresent)
  * [func (o Uint32) IsPresent() bool](#Uint32.IsPresent)
  * [func (o Uint32) OrElse(value uint32) uint32](#Uint32.OrElse)
* [type Uint64](#Uint64)
  * [func EmptyUint64() Uint64](#EmptyUint64)
  * [func OfNilableUint64(ptr *uint64) Uint64](#OfNilableUint64)
  * [func OfUint64(value uint64) Uint64](#OfUint64)
  * [func (o Uint64) IfPresent(f func(value uint64))](#Uint64.IfPresent)
  * [func (o Uint64) IsPresent() bool](#Uint64.IsPresent)
  * [func (o Uint64) OrElse(value uint64) uint64](#Uint64.OrElse)
* [type Uint8](#Uint8)
  * [func EmptyUint8() Uint8](#EmptyUint8)
  * [func OfNilableUint8(ptr *uint8) Uint8](#OfNilableUint8)
  * [func OfUint8(value uint8) Uint8](#OfUint8)
  * [func (o Uint8) IfPresent(f func(value uint8))](#Uint8.IfPresent)
  * [func (o Uint8) IsPresent() bool](#Uint8.IsPresent)
  * [func (o Uint8) OrElse(value uint8) uint8](#Uint8.OrElse)
* [type Uintptr](#Uintptr)
  * [func EmptyUintptr() Uintptr](#EmptyUintptr)
  * [func OfNilableUintptr(ptr *uintptr) Uintptr](#OfNilableUintptr)
  * [func OfUintptr(value uintptr) Uintptr](#OfUintptr)
  * [func (o Uintptr) IfPresent(f func(value uintptr))](#Uintptr.IfPresent)
  * [func (o Uintptr) IsPresent() bool](#Uintptr.IsPresent)
  * [func (o Uintptr) OrElse(value uintptr) uintptr](#Uintptr.OrElse)

#### <a name="pkg-examples">Examples</a>
* [Int.IfPresent (Present)](#example_Int_IfPresent_present)
* [Int.OrElse](#example_Int_OrElse)

#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/leighmcculloch/optional/doc.go) [gotemplate_Bool.go](/src/github.com/leighmcculloch/optional/gotemplate_Bool.go) [gotemplate_Byte.go](/src/github.com/leighmcculloch/optional/gotemplate_Byte.go) [gotemplate_Complex128.go](/src/github.com/leighmcculloch/optional/gotemplate_Complex128.go) [gotemplate_Complex64.go](/src/github.com/leighmcculloch/optional/gotemplate_Complex64.go) [gotemplate_Error.go](/src/github.com/leighmcculloch/optional/gotemplate_Error.go) [gotemplate_Float32.go](/src/github.com/leighmcculloch/optional/gotemplate_Float32.go) [gotemplate_Float64.go](/src/github.com/leighmcculloch/optional/gotemplate_Float64.go) [gotemplate_Int.go](/src/github.com/leighmcculloch/optional/gotemplate_Int.go) [gotemplate_Int16.go](/src/github.com/leighmcculloch/optional/gotemplate_Int16.go) [gotemplate_Int32.go](/src/github.com/leighmcculloch/optional/gotemplate_Int32.go) [gotemplate_Int64.go](/src/github.com/leighmcculloch/optional/gotemplate_Int64.go) [gotemplate_Int8.go](/src/github.com/leighmcculloch/optional/gotemplate_Int8.go) [gotemplate_Rune.go](/src/github.com/leighmcculloch/optional/gotemplate_Rune.go) [gotemplate_String.go](/src/github.com/leighmcculloch/optional/gotemplate_String.go) [gotemplate_Uint.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint.go) [gotemplate_Uint16.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint16.go) [gotemplate_Uint32.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint32.go) [gotemplate_Uint64.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint64.go) [gotemplate_Uint8.go](/src/github.com/leighmcculloch/optional/gotemplate_Uint8.go) [gotemplate_Uintptr.go](/src/github.com/leighmcculloch/optional/gotemplate_Uintptr.go) [types.go](/src/github.com/leighmcculloch/optional/types.go) 






## <a name="Bool">type</a> [Bool](/src/target/gotemplate_Bool.go?s=180:213#L1)
``` go
type Bool struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyBool">func</a> [EmptyBool](/src/target/gotemplate_Bool.go?s=402:423#L11)
``` go
func EmptyBool() Bool
```
Empty returns an empty Optional.


### <a name="OfBool">func</a> [OfBool](/src/target/gotemplate_Bool.go?s=253:281#L2)
``` go
func OfBool(value bool) Bool
```
Of wraps the value in an Optional.


### <a name="OfNilableBool">func</a> [OfNilableBool](/src/target/gotemplate_Bool.go?s=308:342#L6)
``` go
func OfNilableBool(ptr *bool) Bool
```




### <a name="Bool.IfPresent">func</a> (Bool) [IfPresent](/src/target/gotemplate_Bool.go?s=653:696#L21)
``` go
func (o Bool) IfPresent(f func(value bool))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Bool.IsPresent">func</a> (Bool) [IsPresent](/src/target/gotemplate_Bool.go?s=516:546#L16)
``` go
func (o Bool) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Bool.OrElse">func</a> (Bool) [OrElse](/src/target/gotemplate_Bool.go?s=867:904#L29)
``` go
func (o Bool) OrElse(value bool) bool
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Byte">type</a> [Byte](/src/target/gotemplate_Byte.go?s=180:213#L1)
``` go
type Byte struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyByte">func</a> [EmptyByte](/src/target/gotemplate_Byte.go?s=402:423#L11)
``` go
func EmptyByte() Byte
```
Empty returns an empty Optional.


### <a name="OfByte">func</a> [OfByte](/src/target/gotemplate_Byte.go?s=253:281#L2)
``` go
func OfByte(value byte) Byte
```
Of wraps the value in an Optional.


### <a name="OfNilableByte">func</a> [OfNilableByte](/src/target/gotemplate_Byte.go?s=308:342#L6)
``` go
func OfNilableByte(ptr *byte) Byte
```




### <a name="Byte.IfPresent">func</a> (Byte) [IfPresent](/src/target/gotemplate_Byte.go?s=653:696#L21)
``` go
func (o Byte) IfPresent(f func(value byte))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Byte.IsPresent">func</a> (Byte) [IsPresent](/src/target/gotemplate_Byte.go?s=516:546#L16)
``` go
func (o Byte) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Byte.OrElse">func</a> (Byte) [OrElse](/src/target/gotemplate_Byte.go?s=867:904#L29)
``` go
func (o Byte) OrElse(value byte) byte
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Complex128">type</a> [Complex128](/src/target/gotemplate_Complex128.go?s=180:225#L1)
``` go
type Complex128 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyComplex128">func</a> [EmptyComplex128](/src/target/gotemplate_Complex128.go?s=462:495#L11)
``` go
func EmptyComplex128() Complex128
```
Empty returns an empty Optional.


### <a name="OfComplex128">func</a> [OfComplex128](/src/target/gotemplate_Complex128.go?s=265:311#L2)
``` go
func OfComplex128(value complex128) Complex128
```
Of wraps the value in an Optional.


### <a name="OfNilableComplex128">func</a> [OfNilableComplex128](/src/target/gotemplate_Complex128.go?s=344:396#L6)
``` go
func OfNilableComplex128(ptr *complex128) Complex128
```




### <a name="Complex128.IfPresent">func</a> (Complex128) [IfPresent](/src/target/gotemplate_Complex128.go?s=737:792#L21)
``` go
func (o Complex128) IfPresent(f func(value complex128))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Complex128.IsPresent">func</a> (Complex128) [IsPresent](/src/target/gotemplate_Complex128.go?s=594:630#L16)
``` go
func (o Complex128) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Complex128.OrElse">func</a> (Complex128) [OrElse](/src/target/gotemplate_Complex128.go?s=963:1018#L29)
``` go
func (o Complex128) OrElse(value complex128) complex128
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Complex64">type</a> [Complex64](/src/target/gotemplate_Complex64.go?s=180:223#L1)
``` go
type Complex64 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyComplex64">func</a> [EmptyComplex64](/src/target/gotemplate_Complex64.go?s=452:483#L11)
``` go
func EmptyComplex64() Complex64
```
Empty returns an empty Optional.


### <a name="OfComplex64">func</a> [OfComplex64](/src/target/gotemplate_Complex64.go?s=263:306#L2)
``` go
func OfComplex64(value complex64) Complex64
```
Of wraps the value in an Optional.


### <a name="OfNilableComplex64">func</a> [OfNilableComplex64](/src/target/gotemplate_Complex64.go?s=338:387#L6)
``` go
func OfNilableComplex64(ptr *complex64) Complex64
```




### <a name="Complex64.IfPresent">func</a> (Complex64) [IfPresent](/src/target/gotemplate_Complex64.go?s=723:776#L21)
``` go
func (o Complex64) IfPresent(f func(value complex64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Complex64.IsPresent">func</a> (Complex64) [IsPresent](/src/target/gotemplate_Complex64.go?s=581:616#L16)
``` go
func (o Complex64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Complex64.OrElse">func</a> (Complex64) [OrElse](/src/target/gotemplate_Complex64.go?s=947:999#L29)
``` go
func (o Complex64) OrElse(value complex64) complex64
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Error">type</a> [Error](/src/target/gotemplate_Error.go?s=180:215#L1)
``` go
type Error struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyError">func</a> [EmptyError](/src/target/gotemplate_Error.go?s=412:435#L11)
``` go
func EmptyError() Error
```
Empty returns an empty Optional.


### <a name="OfError">func</a> [OfError](/src/target/gotemplate_Error.go?s=255:286#L2)
``` go
func OfError(value error) Error
```
Of wraps the value in an Optional.


### <a name="OfNilableError">func</a> [OfNilableError](/src/target/gotemplate_Error.go?s=314:351#L6)
``` go
func OfNilableError(ptr *error) Error
```




### <a name="Error.IfPresent">func</a> (Error) [IfPresent](/src/target/gotemplate_Error.go?s=667:712#L21)
``` go
func (o Error) IfPresent(f func(value error))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Error.IsPresent">func</a> (Error) [IsPresent](/src/target/gotemplate_Error.go?s=529:560#L16)
``` go
func (o Error) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Error.OrElse">func</a> (Error) [OrElse](/src/target/gotemplate_Error.go?s=883:923#L29)
``` go
func (o Error) OrElse(value error) error
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Float32">type</a> [Float32](/src/target/gotemplate_Float32.go?s=180:219#L1)
``` go
type Float32 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyFloat32">func</a> [EmptyFloat32](/src/target/gotemplate_Float32.go?s=432:459#L11)
``` go
func EmptyFloat32() Float32
```
Empty returns an empty Optional.


### <a name="OfFloat32">func</a> [OfFloat32](/src/target/gotemplate_Float32.go?s=259:296#L2)
``` go
func OfFloat32(value float32) Float32
```
Of wraps the value in an Optional.


### <a name="OfNilableFloat32">func</a> [OfNilableFloat32](/src/target/gotemplate_Float32.go?s=326:369#L6)
``` go
func OfNilableFloat32(ptr *float32) Float32
```




### <a name="Float32.IfPresent">func</a> (Float32) [IfPresent](/src/target/gotemplate_Float32.go?s=695:744#L21)
``` go
func (o Float32) IfPresent(f func(value float32))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Float32.IsPresent">func</a> (Float32) [IsPresent](/src/target/gotemplate_Float32.go?s=555:588#L16)
``` go
func (o Float32) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Float32.OrElse">func</a> (Float32) [OrElse](/src/target/gotemplate_Float32.go?s=915:961#L29)
``` go
func (o Float32) OrElse(value float32) float32
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Float64">type</a> [Float64](/src/target/gotemplate_Float64.go?s=180:219#L1)
``` go
type Float64 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyFloat64">func</a> [EmptyFloat64](/src/target/gotemplate_Float64.go?s=432:459#L11)
``` go
func EmptyFloat64() Float64
```
Empty returns an empty Optional.


### <a name="OfFloat64">func</a> [OfFloat64](/src/target/gotemplate_Float64.go?s=259:296#L2)
``` go
func OfFloat64(value float64) Float64
```
Of wraps the value in an Optional.


### <a name="OfNilableFloat64">func</a> [OfNilableFloat64](/src/target/gotemplate_Float64.go?s=326:369#L6)
``` go
func OfNilableFloat64(ptr *float64) Float64
```




### <a name="Float64.IfPresent">func</a> (Float64) [IfPresent](/src/target/gotemplate_Float64.go?s=695:744#L21)
``` go
func (o Float64) IfPresent(f func(value float64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Float64.IsPresent">func</a> (Float64) [IsPresent](/src/target/gotemplate_Float64.go?s=555:588#L16)
``` go
func (o Float64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Float64.OrElse">func</a> (Float64) [OrElse](/src/target/gotemplate_Float64.go?s=915:961#L29)
``` go
func (o Float64) OrElse(value float64) float64
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Int">type</a> [Int](/src/target/gotemplate_Int.go?s=180:211#L1)
``` go
type Int struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt">func</a> [EmptyInt](/src/target/gotemplate_Int.go?s=392:411#L11)
``` go
func EmptyInt() Int
```
Empty returns an empty Optional.


### <a name="OfInt">func</a> [OfInt](/src/target/gotemplate_Int.go?s=251:276#L2)
``` go
func OfInt(value int) Int
```
Of wraps the value in an Optional.


### <a name="OfNilableInt">func</a> [OfNilableInt](/src/target/gotemplate_Int.go?s=302:333#L6)
``` go
func OfNilableInt(ptr *int) Int
```




### <a name="Int.IfPresent">func</a> (Int) [IfPresent](/src/target/gotemplate_Int.go?s=639:680#L21)
``` go
func (o Int) IfPresent(f func(value int))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int.IsPresent">func</a> (Int) [IsPresent](/src/target/gotemplate_Int.go?s=503:532#L16)
``` go
func (o Int) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int.OrElse">func</a> (Int) [OrElse](/src/target/gotemplate_Int.go?s=851:885#L29)
``` go
func (o Int) OrElse(value int) int
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Int16">type</a> [Int16](/src/target/gotemplate_Int16.go?s=180:215#L1)
``` go
type Int16 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt16">func</a> [EmptyInt16](/src/target/gotemplate_Int16.go?s=412:435#L11)
``` go
func EmptyInt16() Int16
```
Empty returns an empty Optional.


### <a name="OfInt16">func</a> [OfInt16](/src/target/gotemplate_Int16.go?s=255:286#L2)
``` go
func OfInt16(value int16) Int16
```
Of wraps the value in an Optional.


### <a name="OfNilableInt16">func</a> [OfNilableInt16](/src/target/gotemplate_Int16.go?s=314:351#L6)
``` go
func OfNilableInt16(ptr *int16) Int16
```




### <a name="Int16.IfPresent">func</a> (Int16) [IfPresent](/src/target/gotemplate_Int16.go?s=667:712#L21)
``` go
func (o Int16) IfPresent(f func(value int16))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int16.IsPresent">func</a> (Int16) [IsPresent](/src/target/gotemplate_Int16.go?s=529:560#L16)
``` go
func (o Int16) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int16.OrElse">func</a> (Int16) [OrElse](/src/target/gotemplate_Int16.go?s=883:923#L29)
``` go
func (o Int16) OrElse(value int16) int16
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Int32">type</a> [Int32](/src/target/gotemplate_Int32.go?s=180:215#L1)
``` go
type Int32 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt32">func</a> [EmptyInt32](/src/target/gotemplate_Int32.go?s=412:435#L11)
``` go
func EmptyInt32() Int32
```
Empty returns an empty Optional.


### <a name="OfInt32">func</a> [OfInt32](/src/target/gotemplate_Int32.go?s=255:286#L2)
``` go
func OfInt32(value int32) Int32
```
Of wraps the value in an Optional.


### <a name="OfNilableInt32">func</a> [OfNilableInt32](/src/target/gotemplate_Int32.go?s=314:351#L6)
``` go
func OfNilableInt32(ptr *int32) Int32
```




### <a name="Int32.IfPresent">func</a> (Int32) [IfPresent](/src/target/gotemplate_Int32.go?s=667:712#L21)
``` go
func (o Int32) IfPresent(f func(value int32))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int32.IsPresent">func</a> (Int32) [IsPresent](/src/target/gotemplate_Int32.go?s=529:560#L16)
``` go
func (o Int32) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int32.OrElse">func</a> (Int32) [OrElse](/src/target/gotemplate_Int32.go?s=883:923#L29)
``` go
func (o Int32) OrElse(value int32) int32
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Int64">type</a> [Int64](/src/target/gotemplate_Int64.go?s=180:215#L1)
``` go
type Int64 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt64">func</a> [EmptyInt64](/src/target/gotemplate_Int64.go?s=412:435#L11)
``` go
func EmptyInt64() Int64
```
Empty returns an empty Optional.


### <a name="OfInt64">func</a> [OfInt64](/src/target/gotemplate_Int64.go?s=255:286#L2)
``` go
func OfInt64(value int64) Int64
```
Of wraps the value in an Optional.


### <a name="OfNilableInt64">func</a> [OfNilableInt64](/src/target/gotemplate_Int64.go?s=314:351#L6)
``` go
func OfNilableInt64(ptr *int64) Int64
```




### <a name="Int64.IfPresent">func</a> (Int64) [IfPresent](/src/target/gotemplate_Int64.go?s=667:712#L21)
``` go
func (o Int64) IfPresent(f func(value int64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int64.IsPresent">func</a> (Int64) [IsPresent](/src/target/gotemplate_Int64.go?s=529:560#L16)
``` go
func (o Int64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int64.OrElse">func</a> (Int64) [OrElse](/src/target/gotemplate_Int64.go?s=883:923#L29)
``` go
func (o Int64) OrElse(value int64) int64
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Int8">type</a> [Int8](/src/target/gotemplate_Int8.go?s=180:213#L1)
``` go
type Int8 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyInt8">func</a> [EmptyInt8](/src/target/gotemplate_Int8.go?s=402:423#L11)
``` go
func EmptyInt8() Int8
```
Empty returns an empty Optional.


### <a name="OfInt8">func</a> [OfInt8](/src/target/gotemplate_Int8.go?s=253:281#L2)
``` go
func OfInt8(value int8) Int8
```
Of wraps the value in an Optional.


### <a name="OfNilableInt8">func</a> [OfNilableInt8](/src/target/gotemplate_Int8.go?s=308:342#L6)
``` go
func OfNilableInt8(ptr *int8) Int8
```




### <a name="Int8.IfPresent">func</a> (Int8) [IfPresent](/src/target/gotemplate_Int8.go?s=653:696#L21)
``` go
func (o Int8) IfPresent(f func(value int8))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int8.IsPresent">func</a> (Int8) [IsPresent](/src/target/gotemplate_Int8.go?s=516:546#L16)
``` go
func (o Int8) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int8.OrElse">func</a> (Int8) [OrElse](/src/target/gotemplate_Int8.go?s=867:904#L29)
``` go
func (o Int8) OrElse(value int8) int8
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Rune">type</a> [Rune](/src/target/gotemplate_Rune.go?s=180:213#L1)
``` go
type Rune struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyRune">func</a> [EmptyRune](/src/target/gotemplate_Rune.go?s=402:423#L11)
``` go
func EmptyRune() Rune
```
Empty returns an empty Optional.


### <a name="OfNilableRune">func</a> [OfNilableRune](/src/target/gotemplate_Rune.go?s=308:342#L6)
``` go
func OfNilableRune(ptr *rune) Rune
```

### <a name="OfRune">func</a> [OfRune](/src/target/gotemplate_Rune.go?s=253:281#L2)
``` go
func OfRune(value rune) Rune
```
Of wraps the value in an Optional.





### <a name="Rune.IfPresent">func</a> (Rune) [IfPresent](/src/target/gotemplate_Rune.go?s=653:696#L21)
``` go
func (o Rune) IfPresent(f func(value rune))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Rune.IsPresent">func</a> (Rune) [IsPresent](/src/target/gotemplate_Rune.go?s=516:546#L16)
``` go
func (o Rune) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Rune.OrElse">func</a> (Rune) [OrElse](/src/target/gotemplate_Rune.go?s=867:904#L29)
``` go
func (o Rune) OrElse(value rune) rune
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="String">type</a> [String](/src/target/gotemplate_String.go?s=180:217#L1)
``` go
type String struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyString">func</a> [EmptyString](/src/target/gotemplate_String.go?s=422:447#L11)
``` go
func EmptyString() String
```
Empty returns an empty Optional.


### <a name="OfNilableString">func</a> [OfNilableString](/src/target/gotemplate_String.go?s=320:360#L6)
``` go
func OfNilableString(ptr *string) String
```

### <a name="OfString">func</a> [OfString](/src/target/gotemplate_String.go?s=257:291#L2)
``` go
func OfString(value string) String
```
Of wraps the value in an Optional.





### <a name="String.IfPresent">func</a> (String) [IfPresent](/src/target/gotemplate_String.go?s=681:728#L21)
``` go
func (o String) IfPresent(f func(value string))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="String.IsPresent">func</a> (String) [IsPresent](/src/target/gotemplate_String.go?s=542:574#L16)
``` go
func (o String) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="String.OrElse">func</a> (String) [OrElse](/src/target/gotemplate_String.go?s=899:942#L29)
``` go
func (o String) OrElse(value string) string
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Uint">type</a> [Uint](/src/target/gotemplate_Uint.go?s=180:213#L1)
``` go
type Uint struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint">func</a> [EmptyUint](/src/target/gotemplate_Uint.go?s=402:423#L11)
``` go
func EmptyUint() Uint
```
Empty returns an empty Optional.


### <a name="OfNilableUint">func</a> [OfNilableUint](/src/target/gotemplate_Uint.go?s=308:342#L6)
``` go
func OfNilableUint(ptr *uint) Uint
```

### <a name="OfUint">func</a> [OfUint](/src/target/gotemplate_Uint.go?s=253:281#L2)
``` go
func OfUint(value uint) Uint
```
Of wraps the value in an Optional.





### <a name="Uint.IfPresent">func</a> (Uint) [IfPresent](/src/target/gotemplate_Uint.go?s=653:696#L21)
``` go
func (o Uint) IfPresent(f func(value uint))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint.IsPresent">func</a> (Uint) [IsPresent](/src/target/gotemplate_Uint.go?s=516:546#L16)
``` go
func (o Uint) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint.OrElse">func</a> (Uint) [OrElse](/src/target/gotemplate_Uint.go?s=867:904#L29)
``` go
func (o Uint) OrElse(value uint) uint
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Uint16">type</a> [Uint16](/src/target/gotemplate_Uint16.go?s=180:217#L1)
``` go
type Uint16 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint16">func</a> [EmptyUint16](/src/target/gotemplate_Uint16.go?s=422:447#L11)
``` go
func EmptyUint16() Uint16
```
Empty returns an empty Optional.


### <a name="OfNilableUint16">func</a> [OfNilableUint16](/src/target/gotemplate_Uint16.go?s=320:360#L6)
``` go
func OfNilableUint16(ptr *uint16) Uint16
```

### <a name="OfUint16">func</a> [OfUint16](/src/target/gotemplate_Uint16.go?s=257:291#L2)
``` go
func OfUint16(value uint16) Uint16
```
Of wraps the value in an Optional.





### <a name="Uint16.IfPresent">func</a> (Uint16) [IfPresent](/src/target/gotemplate_Uint16.go?s=681:728#L21)
``` go
func (o Uint16) IfPresent(f func(value uint16))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint16.IsPresent">func</a> (Uint16) [IsPresent](/src/target/gotemplate_Uint16.go?s=542:574#L16)
``` go
func (o Uint16) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint16.OrElse">func</a> (Uint16) [OrElse](/src/target/gotemplate_Uint16.go?s=899:942#L29)
``` go
func (o Uint16) OrElse(value uint16) uint16
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Uint32">type</a> [Uint32](/src/target/gotemplate_Uint32.go?s=180:217#L1)
``` go
type Uint32 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint32">func</a> [EmptyUint32](/src/target/gotemplate_Uint32.go?s=422:447#L11)
``` go
func EmptyUint32() Uint32
```
Empty returns an empty Optional.


### <a name="OfNilableUint32">func</a> [OfNilableUint32](/src/target/gotemplate_Uint32.go?s=320:360#L6)
``` go
func OfNilableUint32(ptr *uint32) Uint32
```

### <a name="OfUint32">func</a> [OfUint32](/src/target/gotemplate_Uint32.go?s=257:291#L2)
``` go
func OfUint32(value uint32) Uint32
```
Of wraps the value in an Optional.





### <a name="Uint32.IfPresent">func</a> (Uint32) [IfPresent](/src/target/gotemplate_Uint32.go?s=681:728#L21)
``` go
func (o Uint32) IfPresent(f func(value uint32))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint32.IsPresent">func</a> (Uint32) [IsPresent](/src/target/gotemplate_Uint32.go?s=542:574#L16)
``` go
func (o Uint32) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint32.OrElse">func</a> (Uint32) [OrElse](/src/target/gotemplate_Uint32.go?s=899:942#L29)
``` go
func (o Uint32) OrElse(value uint32) uint32
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Uint64">type</a> [Uint64](/src/target/gotemplate_Uint64.go?s=180:217#L1)
``` go
type Uint64 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint64">func</a> [EmptyUint64](/src/target/gotemplate_Uint64.go?s=422:447#L11)
``` go
func EmptyUint64() Uint64
```
Empty returns an empty Optional.


### <a name="OfNilableUint64">func</a> [OfNilableUint64](/src/target/gotemplate_Uint64.go?s=320:360#L6)
``` go
func OfNilableUint64(ptr *uint64) Uint64
```

### <a name="OfUint64">func</a> [OfUint64](/src/target/gotemplate_Uint64.go?s=257:291#L2)
``` go
func OfUint64(value uint64) Uint64
```
Of wraps the value in an Optional.





### <a name="Uint64.IfPresent">func</a> (Uint64) [IfPresent](/src/target/gotemplate_Uint64.go?s=681:728#L21)
``` go
func (o Uint64) IfPresent(f func(value uint64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint64.IsPresent">func</a> (Uint64) [IsPresent](/src/target/gotemplate_Uint64.go?s=542:574#L16)
``` go
func (o Uint64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint64.OrElse">func</a> (Uint64) [OrElse](/src/target/gotemplate_Uint64.go?s=899:942#L29)
``` go
func (o Uint64) OrElse(value uint64) uint64
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Uint8">type</a> [Uint8](/src/target/gotemplate_Uint8.go?s=180:215#L1)
``` go
type Uint8 struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUint8">func</a> [EmptyUint8](/src/target/gotemplate_Uint8.go?s=412:435#L11)
``` go
func EmptyUint8() Uint8
```
Empty returns an empty Optional.


### <a name="OfNilableUint8">func</a> [OfNilableUint8](/src/target/gotemplate_Uint8.go?s=314:351#L6)
``` go
func OfNilableUint8(ptr *uint8) Uint8
```

### <a name="OfUint8">func</a> [OfUint8](/src/target/gotemplate_Uint8.go?s=255:286#L2)
``` go
func OfUint8(value uint8) Uint8
```
Of wraps the value in an Optional.





### <a name="Uint8.IfPresent">func</a> (Uint8) [IfPresent](/src/target/gotemplate_Uint8.go?s=667:712#L21)
``` go
func (o Uint8) IfPresent(f func(value uint8))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint8.IsPresent">func</a> (Uint8) [IsPresent](/src/target/gotemplate_Uint8.go?s=529:560#L16)
``` go
func (o Uint8) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint8.OrElse">func</a> (Uint8) [OrElse](/src/target/gotemplate_Uint8.go?s=883:923#L29)
``` go
func (o Uint8) OrElse(value uint8) uint8
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.




## <a name="Uintptr">type</a> [Uintptr](/src/target/gotemplate_Uintptr.go?s=180:219#L1)
``` go
type Uintptr struct {
    // contains filtered or unexported fields
}
```
Optional wraps a value that may or may not be nil.
If a value is present, it may be unwrapped to expose the underlying value.







### <a name="EmptyUintptr">func</a> [EmptyUintptr](/src/target/gotemplate_Uintptr.go?s=432:459#L11)
``` go
func EmptyUintptr() Uintptr
```
Empty returns an empty Optional.


### <a name="OfNilableUintptr">func</a> [OfNilableUintptr](/src/target/gotemplate_Uintptr.go?s=326:369#L6)
``` go
func OfNilableUintptr(ptr *uintptr) Uintptr
```

### <a name="OfUintptr">func</a> [OfUintptr](/src/target/gotemplate_Uintptr.go?s=259:296#L2)
``` go
func OfUintptr(value uintptr) Uintptr
```
Of wraps the value in an Optional.





### <a name="Uintptr.IfPresent">func</a> (Uintptr) [IfPresent](/src/target/gotemplate_Uintptr.go?s=695:744#L21)
``` go
func (o Uintptr) IfPresent(f func(value uintptr))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uintptr.IsPresent">func</a> (Uintptr) [IsPresent](/src/target/gotemplate_Uintptr.go?s=555:588#L16)
``` go
func (o Uintptr) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uintptr.OrElse">func</a> (Uintptr) [OrElse](/src/target/gotemplate_Uintptr.go?s=915:961#L29)
``` go
func (o Uintptr) OrElse(value uintptr) uintptr
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
