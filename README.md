

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
		optional.ForInt(2017),
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
		optional.ForInt(2017),
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
  * [func ForBool(value bool) Bool](#ForBool)
  * [func ForBoolPtr(ptr *bool) Bool](#ForBoolPtr)
  * [func (o Bool) IfPresent(f func(value bool))](#Bool.IfPresent)
  * [func (o Bool) IsPresent() bool](#Bool.IsPresent)
  * [func (o Bool) OrElse(value bool) bool](#Bool.OrElse)
* [type Byte](#Byte)
  * [func EmptyByte() Byte](#EmptyByte)
  * [func ForByte(value byte) Byte](#ForByte)
  * [func ForBytePtr(ptr *byte) Byte](#ForBytePtr)
  * [func (o Byte) IfPresent(f func(value byte))](#Byte.IfPresent)
  * [func (o Byte) IsPresent() bool](#Byte.IsPresent)
  * [func (o Byte) OrElse(value byte) byte](#Byte.OrElse)
* [type Complex128](#Complex128)
  * [func EmptyComplex128() Complex128](#EmptyComplex128)
  * [func ForComplex128(value complex128) Complex128](#ForComplex128)
  * [func ForComplex128Ptr(ptr *complex128) Complex128](#ForComplex128Ptr)
  * [func (o Complex128) IfPresent(f func(value complex128))](#Complex128.IfPresent)
  * [func (o Complex128) IsPresent() bool](#Complex128.IsPresent)
  * [func (o Complex128) OrElse(value complex128) complex128](#Complex128.OrElse)
* [type Complex64](#Complex64)
  * [func EmptyComplex64() Complex64](#EmptyComplex64)
  * [func ForComplex64(value complex64) Complex64](#ForComplex64)
  * [func ForComplex64Ptr(ptr *complex64) Complex64](#ForComplex64Ptr)
  * [func (o Complex64) IfPresent(f func(value complex64))](#Complex64.IfPresent)
  * [func (o Complex64) IsPresent() bool](#Complex64.IsPresent)
  * [func (o Complex64) OrElse(value complex64) complex64](#Complex64.OrElse)
* [type Error](#Error)
  * [func EmptyError() Error](#EmptyError)
  * [func ForError(value error) Error](#ForError)
  * [func ForErrorPtr(ptr *error) Error](#ForErrorPtr)
  * [func (o Error) IfPresent(f func(value error))](#Error.IfPresent)
  * [func (o Error) IsPresent() bool](#Error.IsPresent)
  * [func (o Error) OrElse(value error) error](#Error.OrElse)
* [type Float32](#Float32)
  * [func EmptyFloat32() Float32](#EmptyFloat32)
  * [func ForFloat32(value float32) Float32](#ForFloat32)
  * [func ForFloat32Ptr(ptr *float32) Float32](#ForFloat32Ptr)
  * [func (o Float32) IfPresent(f func(value float32))](#Float32.IfPresent)
  * [func (o Float32) IsPresent() bool](#Float32.IsPresent)
  * [func (o Float32) OrElse(value float32) float32](#Float32.OrElse)
* [type Float64](#Float64)
  * [func EmptyFloat64() Float64](#EmptyFloat64)
  * [func ForFloat64(value float64) Float64](#ForFloat64)
  * [func ForFloat64Ptr(ptr *float64) Float64](#ForFloat64Ptr)
  * [func (o Float64) IfPresent(f func(value float64))](#Float64.IfPresent)
  * [func (o Float64) IsPresent() bool](#Float64.IsPresent)
  * [func (o Float64) OrElse(value float64) float64](#Float64.OrElse)
* [type Int](#Int)
  * [func EmptyInt() Int](#EmptyInt)
  * [func ForInt(value int) Int](#ForInt)
  * [func ForIntPtr(ptr *int) Int](#ForIntPtr)
  * [func (o Int) IfPresent(f func(value int))](#Int.IfPresent)
  * [func (o Int) IsPresent() bool](#Int.IsPresent)
  * [func (o Int) OrElse(value int) int](#Int.OrElse)
* [type Int16](#Int16)
  * [func EmptyInt16() Int16](#EmptyInt16)
  * [func ForInt16(value int16) Int16](#ForInt16)
  * [func ForInt16Ptr(ptr *int16) Int16](#ForInt16Ptr)
  * [func (o Int16) IfPresent(f func(value int16))](#Int16.IfPresent)
  * [func (o Int16) IsPresent() bool](#Int16.IsPresent)
  * [func (o Int16) OrElse(value int16) int16](#Int16.OrElse)
* [type Int32](#Int32)
  * [func EmptyInt32() Int32](#EmptyInt32)
  * [func ForInt32(value int32) Int32](#ForInt32)
  * [func ForInt32Ptr(ptr *int32) Int32](#ForInt32Ptr)
  * [func (o Int32) IfPresent(f func(value int32))](#Int32.IfPresent)
  * [func (o Int32) IsPresent() bool](#Int32.IsPresent)
  * [func (o Int32) OrElse(value int32) int32](#Int32.OrElse)
* [type Int64](#Int64)
  * [func EmptyInt64() Int64](#EmptyInt64)
  * [func ForInt64(value int64) Int64](#ForInt64)
  * [func ForInt64Ptr(ptr *int64) Int64](#ForInt64Ptr)
  * [func (o Int64) IfPresent(f func(value int64))](#Int64.IfPresent)
  * [func (o Int64) IsPresent() bool](#Int64.IsPresent)
  * [func (o Int64) OrElse(value int64) int64](#Int64.OrElse)
* [type Int8](#Int8)
  * [func EmptyInt8() Int8](#EmptyInt8)
  * [func ForInt8(value int8) Int8](#ForInt8)
  * [func ForInt8Ptr(ptr *int8) Int8](#ForInt8Ptr)
  * [func (o Int8) IfPresent(f func(value int8))](#Int8.IfPresent)
  * [func (o Int8) IsPresent() bool](#Int8.IsPresent)
  * [func (o Int8) OrElse(value int8) int8](#Int8.OrElse)
* [type Rune](#Rune)
  * [func EmptyRune() Rune](#EmptyRune)
  * [func ForRune(value rune) Rune](#ForRune)
  * [func ForRunePtr(ptr *rune) Rune](#ForRunePtr)
  * [func (o Rune) IfPresent(f func(value rune))](#Rune.IfPresent)
  * [func (o Rune) IsPresent() bool](#Rune.IsPresent)
  * [func (o Rune) OrElse(value rune) rune](#Rune.OrElse)
* [type String](#String)
  * [func EmptyString() String](#EmptyString)
  * [func ForString(value string) String](#ForString)
  * [func ForStringPtr(ptr *string) String](#ForStringPtr)
  * [func (o String) IfPresent(f func(value string))](#String.IfPresent)
  * [func (o String) IsPresent() bool](#String.IsPresent)
  * [func (o String) OrElse(value string) string](#String.OrElse)
* [type Uint](#Uint)
  * [func EmptyUint() Uint](#EmptyUint)
  * [func ForUint(value uint) Uint](#ForUint)
  * [func ForUintPtr(ptr *uint) Uint](#ForUintPtr)
  * [func (o Uint) IfPresent(f func(value uint))](#Uint.IfPresent)
  * [func (o Uint) IsPresent() bool](#Uint.IsPresent)
  * [func (o Uint) OrElse(value uint) uint](#Uint.OrElse)
* [type Uint16](#Uint16)
  * [func EmptyUint16() Uint16](#EmptyUint16)
  * [func ForUint16(value uint16) Uint16](#ForUint16)
  * [func ForUint16Ptr(ptr *uint16) Uint16](#ForUint16Ptr)
  * [func (o Uint16) IfPresent(f func(value uint16))](#Uint16.IfPresent)
  * [func (o Uint16) IsPresent() bool](#Uint16.IsPresent)
  * [func (o Uint16) OrElse(value uint16) uint16](#Uint16.OrElse)
* [type Uint32](#Uint32)
  * [func EmptyUint32() Uint32](#EmptyUint32)
  * [func ForUint32(value uint32) Uint32](#ForUint32)
  * [func ForUint32Ptr(ptr *uint32) Uint32](#ForUint32Ptr)
  * [func (o Uint32) IfPresent(f func(value uint32))](#Uint32.IfPresent)
  * [func (o Uint32) IsPresent() bool](#Uint32.IsPresent)
  * [func (o Uint32) OrElse(value uint32) uint32](#Uint32.OrElse)
* [type Uint64](#Uint64)
  * [func EmptyUint64() Uint64](#EmptyUint64)
  * [func ForUint64(value uint64) Uint64](#ForUint64)
  * [func ForUint64Ptr(ptr *uint64) Uint64](#ForUint64Ptr)
  * [func (o Uint64) IfPresent(f func(value uint64))](#Uint64.IfPresent)
  * [func (o Uint64) IsPresent() bool](#Uint64.IsPresent)
  * [func (o Uint64) OrElse(value uint64) uint64](#Uint64.OrElse)
* [type Uint8](#Uint8)
  * [func EmptyUint8() Uint8](#EmptyUint8)
  * [func ForUint8(value uint8) Uint8](#ForUint8)
  * [func ForUint8Ptr(ptr *uint8) Uint8](#ForUint8Ptr)
  * [func (o Uint8) IfPresent(f func(value uint8))](#Uint8.IfPresent)
  * [func (o Uint8) IsPresent() bool](#Uint8.IsPresent)
  * [func (o Uint8) OrElse(value uint8) uint8](#Uint8.OrElse)
* [type Uintptr](#Uintptr)
  * [func EmptyUintptr() Uintptr](#EmptyUintptr)
  * [func ForUintptr(value uintptr) Uintptr](#ForUintptr)
  * [func ForUintptrPtr(ptr *uintptr) Uintptr](#ForUintptrPtr)
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







### <a name="EmptyBool">func</a> [EmptyBool](/src/target/gotemplate_Bool.go?s=401:422#L11)
``` go
func EmptyBool() Bool
```
Empty returns an empty Optional.


### <a name="ForBool">func</a> [ForBool](/src/target/gotemplate_Bool.go?s=254:283#L2)
``` go
func ForBool(value bool) Bool
```
For wraps the value in an Optional.


### <a name="ForBoolPtr">func</a> [ForBoolPtr](/src/target/gotemplate_Bool.go?s=310:341#L6)
``` go
func ForBoolPtr(ptr *bool) Bool
```




### <a name="Bool.IfPresent">func</a> (Bool) [IfPresent](/src/target/gotemplate_Bool.go?s=652:695#L21)
``` go
func (o Bool) IfPresent(f func(value bool))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Bool.IsPresent">func</a> (Bool) [IsPresent](/src/target/gotemplate_Bool.go?s=515:545#L16)
``` go
func (o Bool) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Bool.OrElse">func</a> (Bool) [OrElse](/src/target/gotemplate_Bool.go?s=866:903#L29)
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







### <a name="EmptyByte">func</a> [EmptyByte](/src/target/gotemplate_Byte.go?s=401:422#L11)
``` go
func EmptyByte() Byte
```
Empty returns an empty Optional.


### <a name="ForByte">func</a> [ForByte](/src/target/gotemplate_Byte.go?s=254:283#L2)
``` go
func ForByte(value byte) Byte
```
For wraps the value in an Optional.


### <a name="ForBytePtr">func</a> [ForBytePtr](/src/target/gotemplate_Byte.go?s=310:341#L6)
``` go
func ForBytePtr(ptr *byte) Byte
```




### <a name="Byte.IfPresent">func</a> (Byte) [IfPresent](/src/target/gotemplate_Byte.go?s=652:695#L21)
``` go
func (o Byte) IfPresent(f func(value byte))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Byte.IsPresent">func</a> (Byte) [IsPresent](/src/target/gotemplate_Byte.go?s=515:545#L16)
``` go
func (o Byte) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Byte.OrElse">func</a> (Byte) [OrElse](/src/target/gotemplate_Byte.go?s=866:903#L29)
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







### <a name="EmptyComplex128">func</a> [EmptyComplex128](/src/target/gotemplate_Complex128.go?s=461:494#L11)
``` go
func EmptyComplex128() Complex128
```
Empty returns an empty Optional.


### <a name="ForComplex128">func</a> [ForComplex128](/src/target/gotemplate_Complex128.go?s=266:313#L2)
``` go
func ForComplex128(value complex128) Complex128
```
For wraps the value in an Optional.


### <a name="ForComplex128Ptr">func</a> [ForComplex128Ptr](/src/target/gotemplate_Complex128.go?s=346:395#L6)
``` go
func ForComplex128Ptr(ptr *complex128) Complex128
```




### <a name="Complex128.IfPresent">func</a> (Complex128) [IfPresent](/src/target/gotemplate_Complex128.go?s=736:791#L21)
``` go
func (o Complex128) IfPresent(f func(value complex128))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Complex128.IsPresent">func</a> (Complex128) [IsPresent](/src/target/gotemplate_Complex128.go?s=593:629#L16)
``` go
func (o Complex128) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Complex128.OrElse">func</a> (Complex128) [OrElse](/src/target/gotemplate_Complex128.go?s=962:1017#L29)
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







### <a name="EmptyComplex64">func</a> [EmptyComplex64](/src/target/gotemplate_Complex64.go?s=451:482#L11)
``` go
func EmptyComplex64() Complex64
```
Empty returns an empty Optional.


### <a name="ForComplex64">func</a> [ForComplex64](/src/target/gotemplate_Complex64.go?s=264:308#L2)
``` go
func ForComplex64(value complex64) Complex64
```
For wraps the value in an Optional.


### <a name="ForComplex64Ptr">func</a> [ForComplex64Ptr](/src/target/gotemplate_Complex64.go?s=340:386#L6)
``` go
func ForComplex64Ptr(ptr *complex64) Complex64
```




### <a name="Complex64.IfPresent">func</a> (Complex64) [IfPresent](/src/target/gotemplate_Complex64.go?s=722:775#L21)
``` go
func (o Complex64) IfPresent(f func(value complex64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Complex64.IsPresent">func</a> (Complex64) [IsPresent](/src/target/gotemplate_Complex64.go?s=580:615#L16)
``` go
func (o Complex64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Complex64.OrElse">func</a> (Complex64) [OrElse](/src/target/gotemplate_Complex64.go?s=946:998#L29)
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







### <a name="EmptyError">func</a> [EmptyError](/src/target/gotemplate_Error.go?s=411:434#L11)
``` go
func EmptyError() Error
```
Empty returns an empty Optional.


### <a name="ForError">func</a> [ForError](/src/target/gotemplate_Error.go?s=256:288#L2)
``` go
func ForError(value error) Error
```
For wraps the value in an Optional.


### <a name="ForErrorPtr">func</a> [ForErrorPtr](/src/target/gotemplate_Error.go?s=316:350#L6)
``` go
func ForErrorPtr(ptr *error) Error
```




### <a name="Error.IfPresent">func</a> (Error) [IfPresent](/src/target/gotemplate_Error.go?s=666:711#L21)
``` go
func (o Error) IfPresent(f func(value error))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Error.IsPresent">func</a> (Error) [IsPresent](/src/target/gotemplate_Error.go?s=528:559#L16)
``` go
func (o Error) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Error.OrElse">func</a> (Error) [OrElse](/src/target/gotemplate_Error.go?s=882:922#L29)
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







### <a name="EmptyFloat32">func</a> [EmptyFloat32](/src/target/gotemplate_Float32.go?s=431:458#L11)
``` go
func EmptyFloat32() Float32
```
Empty returns an empty Optional.


### <a name="ForFloat32">func</a> [ForFloat32](/src/target/gotemplate_Float32.go?s=260:298#L2)
``` go
func ForFloat32(value float32) Float32
```
For wraps the value in an Optional.


### <a name="ForFloat32Ptr">func</a> [ForFloat32Ptr](/src/target/gotemplate_Float32.go?s=328:368#L6)
``` go
func ForFloat32Ptr(ptr *float32) Float32
```




### <a name="Float32.IfPresent">func</a> (Float32) [IfPresent](/src/target/gotemplate_Float32.go?s=694:743#L21)
``` go
func (o Float32) IfPresent(f func(value float32))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Float32.IsPresent">func</a> (Float32) [IsPresent](/src/target/gotemplate_Float32.go?s=554:587#L16)
``` go
func (o Float32) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Float32.OrElse">func</a> (Float32) [OrElse](/src/target/gotemplate_Float32.go?s=914:960#L29)
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







### <a name="EmptyFloat64">func</a> [EmptyFloat64](/src/target/gotemplate_Float64.go?s=431:458#L11)
``` go
func EmptyFloat64() Float64
```
Empty returns an empty Optional.


### <a name="ForFloat64">func</a> [ForFloat64](/src/target/gotemplate_Float64.go?s=260:298#L2)
``` go
func ForFloat64(value float64) Float64
```
For wraps the value in an Optional.


### <a name="ForFloat64Ptr">func</a> [ForFloat64Ptr](/src/target/gotemplate_Float64.go?s=328:368#L6)
``` go
func ForFloat64Ptr(ptr *float64) Float64
```




### <a name="Float64.IfPresent">func</a> (Float64) [IfPresent](/src/target/gotemplate_Float64.go?s=694:743#L21)
``` go
func (o Float64) IfPresent(f func(value float64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Float64.IsPresent">func</a> (Float64) [IsPresent](/src/target/gotemplate_Float64.go?s=554:587#L16)
``` go
func (o Float64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Float64.OrElse">func</a> (Float64) [OrElse](/src/target/gotemplate_Float64.go?s=914:960#L29)
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







### <a name="EmptyInt">func</a> [EmptyInt](/src/target/gotemplate_Int.go?s=391:410#L11)
``` go
func EmptyInt() Int
```
Empty returns an empty Optional.


### <a name="ForInt">func</a> [ForInt](/src/target/gotemplate_Int.go?s=252:278#L2)
``` go
func ForInt(value int) Int
```
For wraps the value in an Optional.


### <a name="ForIntPtr">func</a> [ForIntPtr](/src/target/gotemplate_Int.go?s=304:332#L6)
``` go
func ForIntPtr(ptr *int) Int
```




### <a name="Int.IfPresent">func</a> (Int) [IfPresent](/src/target/gotemplate_Int.go?s=638:679#L21)
``` go
func (o Int) IfPresent(f func(value int))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int.IsPresent">func</a> (Int) [IsPresent](/src/target/gotemplate_Int.go?s=502:531#L16)
``` go
func (o Int) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int.OrElse">func</a> (Int) [OrElse](/src/target/gotemplate_Int.go?s=850:884#L29)
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







### <a name="EmptyInt16">func</a> [EmptyInt16](/src/target/gotemplate_Int16.go?s=411:434#L11)
``` go
func EmptyInt16() Int16
```
Empty returns an empty Optional.


### <a name="ForInt16">func</a> [ForInt16](/src/target/gotemplate_Int16.go?s=256:288#L2)
``` go
func ForInt16(value int16) Int16
```
For wraps the value in an Optional.


### <a name="ForInt16Ptr">func</a> [ForInt16Ptr](/src/target/gotemplate_Int16.go?s=316:350#L6)
``` go
func ForInt16Ptr(ptr *int16) Int16
```




### <a name="Int16.IfPresent">func</a> (Int16) [IfPresent](/src/target/gotemplate_Int16.go?s=666:711#L21)
``` go
func (o Int16) IfPresent(f func(value int16))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int16.IsPresent">func</a> (Int16) [IsPresent](/src/target/gotemplate_Int16.go?s=528:559#L16)
``` go
func (o Int16) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int16.OrElse">func</a> (Int16) [OrElse](/src/target/gotemplate_Int16.go?s=882:922#L29)
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







### <a name="EmptyInt32">func</a> [EmptyInt32](/src/target/gotemplate_Int32.go?s=411:434#L11)
``` go
func EmptyInt32() Int32
```
Empty returns an empty Optional.


### <a name="ForInt32">func</a> [ForInt32](/src/target/gotemplate_Int32.go?s=256:288#L2)
``` go
func ForInt32(value int32) Int32
```
For wraps the value in an Optional.


### <a name="ForInt32Ptr">func</a> [ForInt32Ptr](/src/target/gotemplate_Int32.go?s=316:350#L6)
``` go
func ForInt32Ptr(ptr *int32) Int32
```




### <a name="Int32.IfPresent">func</a> (Int32) [IfPresent](/src/target/gotemplate_Int32.go?s=666:711#L21)
``` go
func (o Int32) IfPresent(f func(value int32))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int32.IsPresent">func</a> (Int32) [IsPresent](/src/target/gotemplate_Int32.go?s=528:559#L16)
``` go
func (o Int32) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int32.OrElse">func</a> (Int32) [OrElse](/src/target/gotemplate_Int32.go?s=882:922#L29)
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







### <a name="EmptyInt64">func</a> [EmptyInt64](/src/target/gotemplate_Int64.go?s=411:434#L11)
``` go
func EmptyInt64() Int64
```
Empty returns an empty Optional.


### <a name="ForInt64">func</a> [ForInt64](/src/target/gotemplate_Int64.go?s=256:288#L2)
``` go
func ForInt64(value int64) Int64
```
For wraps the value in an Optional.


### <a name="ForInt64Ptr">func</a> [ForInt64Ptr](/src/target/gotemplate_Int64.go?s=316:350#L6)
``` go
func ForInt64Ptr(ptr *int64) Int64
```




### <a name="Int64.IfPresent">func</a> (Int64) [IfPresent](/src/target/gotemplate_Int64.go?s=666:711#L21)
``` go
func (o Int64) IfPresent(f func(value int64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int64.IsPresent">func</a> (Int64) [IsPresent](/src/target/gotemplate_Int64.go?s=528:559#L16)
``` go
func (o Int64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int64.OrElse">func</a> (Int64) [OrElse](/src/target/gotemplate_Int64.go?s=882:922#L29)
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







### <a name="EmptyInt8">func</a> [EmptyInt8](/src/target/gotemplate_Int8.go?s=401:422#L11)
``` go
func EmptyInt8() Int8
```
Empty returns an empty Optional.


### <a name="ForInt8">func</a> [ForInt8](/src/target/gotemplate_Int8.go?s=254:283#L2)
``` go
func ForInt8(value int8) Int8
```
For wraps the value in an Optional.


### <a name="ForInt8Ptr">func</a> [ForInt8Ptr](/src/target/gotemplate_Int8.go?s=310:341#L6)
``` go
func ForInt8Ptr(ptr *int8) Int8
```




### <a name="Int8.IfPresent">func</a> (Int8) [IfPresent](/src/target/gotemplate_Int8.go?s=652:695#L21)
``` go
func (o Int8) IfPresent(f func(value int8))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Int8.IsPresent">func</a> (Int8) [IsPresent](/src/target/gotemplate_Int8.go?s=515:545#L16)
``` go
func (o Int8) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Int8.OrElse">func</a> (Int8) [OrElse](/src/target/gotemplate_Int8.go?s=866:903#L29)
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







### <a name="EmptyRune">func</a> [EmptyRune](/src/target/gotemplate_Rune.go?s=401:422#L11)
``` go
func EmptyRune() Rune
```
Empty returns an empty Optional.


### <a name="ForRune">func</a> [ForRune](/src/target/gotemplate_Rune.go?s=254:283#L2)
``` go
func ForRune(value rune) Rune
```
For wraps the value in an Optional.


### <a name="ForRunePtr">func</a> [ForRunePtr](/src/target/gotemplate_Rune.go?s=310:341#L6)
``` go
func ForRunePtr(ptr *rune) Rune
```




### <a name="Rune.IfPresent">func</a> (Rune) [IfPresent](/src/target/gotemplate_Rune.go?s=652:695#L21)
``` go
func (o Rune) IfPresent(f func(value rune))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Rune.IsPresent">func</a> (Rune) [IsPresent](/src/target/gotemplate_Rune.go?s=515:545#L16)
``` go
func (o Rune) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Rune.OrElse">func</a> (Rune) [OrElse](/src/target/gotemplate_Rune.go?s=866:903#L29)
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







### <a name="EmptyString">func</a> [EmptyString](/src/target/gotemplate_String.go?s=421:446#L11)
``` go
func EmptyString() String
```
Empty returns an empty Optional.


### <a name="ForString">func</a> [ForString](/src/target/gotemplate_String.go?s=258:293#L2)
``` go
func ForString(value string) String
```
For wraps the value in an Optional.


### <a name="ForStringPtr">func</a> [ForStringPtr](/src/target/gotemplate_String.go?s=322:359#L6)
``` go
func ForStringPtr(ptr *string) String
```




### <a name="String.IfPresent">func</a> (String) [IfPresent](/src/target/gotemplate_String.go?s=680:727#L21)
``` go
func (o String) IfPresent(f func(value string))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="String.IsPresent">func</a> (String) [IsPresent](/src/target/gotemplate_String.go?s=541:573#L16)
``` go
func (o String) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="String.OrElse">func</a> (String) [OrElse](/src/target/gotemplate_String.go?s=898:941#L29)
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







### <a name="EmptyUint">func</a> [EmptyUint](/src/target/gotemplate_Uint.go?s=401:422#L11)
``` go
func EmptyUint() Uint
```
Empty returns an empty Optional.


### <a name="ForUint">func</a> [ForUint](/src/target/gotemplate_Uint.go?s=254:283#L2)
``` go
func ForUint(value uint) Uint
```
For wraps the value in an Optional.


### <a name="ForUintPtr">func</a> [ForUintPtr](/src/target/gotemplate_Uint.go?s=310:341#L6)
``` go
func ForUintPtr(ptr *uint) Uint
```




### <a name="Uint.IfPresent">func</a> (Uint) [IfPresent](/src/target/gotemplate_Uint.go?s=652:695#L21)
``` go
func (o Uint) IfPresent(f func(value uint))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint.IsPresent">func</a> (Uint) [IsPresent](/src/target/gotemplate_Uint.go?s=515:545#L16)
``` go
func (o Uint) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint.OrElse">func</a> (Uint) [OrElse](/src/target/gotemplate_Uint.go?s=866:903#L29)
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







### <a name="EmptyUint16">func</a> [EmptyUint16](/src/target/gotemplate_Uint16.go?s=421:446#L11)
``` go
func EmptyUint16() Uint16
```
Empty returns an empty Optional.


### <a name="ForUint16">func</a> [ForUint16](/src/target/gotemplate_Uint16.go?s=258:293#L2)
``` go
func ForUint16(value uint16) Uint16
```
For wraps the value in an Optional.


### <a name="ForUint16Ptr">func</a> [ForUint16Ptr](/src/target/gotemplate_Uint16.go?s=322:359#L6)
``` go
func ForUint16Ptr(ptr *uint16) Uint16
```




### <a name="Uint16.IfPresent">func</a> (Uint16) [IfPresent](/src/target/gotemplate_Uint16.go?s=680:727#L21)
``` go
func (o Uint16) IfPresent(f func(value uint16))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint16.IsPresent">func</a> (Uint16) [IsPresent](/src/target/gotemplate_Uint16.go?s=541:573#L16)
``` go
func (o Uint16) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint16.OrElse">func</a> (Uint16) [OrElse](/src/target/gotemplate_Uint16.go?s=898:941#L29)
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







### <a name="EmptyUint32">func</a> [EmptyUint32](/src/target/gotemplate_Uint32.go?s=421:446#L11)
``` go
func EmptyUint32() Uint32
```
Empty returns an empty Optional.


### <a name="ForUint32">func</a> [ForUint32](/src/target/gotemplate_Uint32.go?s=258:293#L2)
``` go
func ForUint32(value uint32) Uint32
```
For wraps the value in an Optional.


### <a name="ForUint32Ptr">func</a> [ForUint32Ptr](/src/target/gotemplate_Uint32.go?s=322:359#L6)
``` go
func ForUint32Ptr(ptr *uint32) Uint32
```




### <a name="Uint32.IfPresent">func</a> (Uint32) [IfPresent](/src/target/gotemplate_Uint32.go?s=680:727#L21)
``` go
func (o Uint32) IfPresent(f func(value uint32))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint32.IsPresent">func</a> (Uint32) [IsPresent](/src/target/gotemplate_Uint32.go?s=541:573#L16)
``` go
func (o Uint32) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint32.OrElse">func</a> (Uint32) [OrElse](/src/target/gotemplate_Uint32.go?s=898:941#L29)
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







### <a name="EmptyUint64">func</a> [EmptyUint64](/src/target/gotemplate_Uint64.go?s=421:446#L11)
``` go
func EmptyUint64() Uint64
```
Empty returns an empty Optional.


### <a name="ForUint64">func</a> [ForUint64](/src/target/gotemplate_Uint64.go?s=258:293#L2)
``` go
func ForUint64(value uint64) Uint64
```
For wraps the value in an Optional.


### <a name="ForUint64Ptr">func</a> [ForUint64Ptr](/src/target/gotemplate_Uint64.go?s=322:359#L6)
``` go
func ForUint64Ptr(ptr *uint64) Uint64
```




### <a name="Uint64.IfPresent">func</a> (Uint64) [IfPresent](/src/target/gotemplate_Uint64.go?s=680:727#L21)
``` go
func (o Uint64) IfPresent(f func(value uint64))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint64.IsPresent">func</a> (Uint64) [IsPresent](/src/target/gotemplate_Uint64.go?s=541:573#L16)
``` go
func (o Uint64) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint64.OrElse">func</a> (Uint64) [OrElse](/src/target/gotemplate_Uint64.go?s=898:941#L29)
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







### <a name="EmptyUint8">func</a> [EmptyUint8](/src/target/gotemplate_Uint8.go?s=411:434#L11)
``` go
func EmptyUint8() Uint8
```
Empty returns an empty Optional.


### <a name="ForUint8">func</a> [ForUint8](/src/target/gotemplate_Uint8.go?s=256:288#L2)
``` go
func ForUint8(value uint8) Uint8
```
For wraps the value in an Optional.


### <a name="ForUint8Ptr">func</a> [ForUint8Ptr](/src/target/gotemplate_Uint8.go?s=316:350#L6)
``` go
func ForUint8Ptr(ptr *uint8) Uint8
```




### <a name="Uint8.IfPresent">func</a> (Uint8) [IfPresent](/src/target/gotemplate_Uint8.go?s=666:711#L21)
``` go
func (o Uint8) IfPresent(f func(value uint8))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uint8.IsPresent">func</a> (Uint8) [IsPresent](/src/target/gotemplate_Uint8.go?s=528:559#L16)
``` go
func (o Uint8) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uint8.OrElse">func</a> (Uint8) [OrElse](/src/target/gotemplate_Uint8.go?s=882:922#L29)
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







### <a name="EmptyUintptr">func</a> [EmptyUintptr](/src/target/gotemplate_Uintptr.go?s=431:458#L11)
``` go
func EmptyUintptr() Uintptr
```
Empty returns an empty Optional.


### <a name="ForUintptr">func</a> [ForUintptr](/src/target/gotemplate_Uintptr.go?s=260:298#L2)
``` go
func ForUintptr(value uintptr) Uintptr
```
For wraps the value in an Optional.


### <a name="ForUintptrPtr">func</a> [ForUintptrPtr](/src/target/gotemplate_Uintptr.go?s=328:368#L6)
``` go
func ForUintptrPtr(ptr *uintptr) Uintptr
```




### <a name="Uintptr.IfPresent">func</a> (Uintptr) [IfPresent](/src/target/gotemplate_Uintptr.go?s=694:743#L21)
``` go
func (o Uintptr) IfPresent(f func(value uintptr))
```
IfPresent calls the function if there is a value wrapped by this Optional.




### <a name="Uintptr.IsPresent">func</a> (Uintptr) [IsPresent](/src/target/gotemplate_Uintptr.go?s=554:587#L16)
``` go
func (o Uintptr) IsPresent() bool
```
IsPresent returns whether there is a value wrapped by this Optional.




### <a name="Uintptr.OrElse">func</a> (Uintptr) [OrElse](/src/target/gotemplate_Uintptr.go?s=914:960#L29)
``` go
func (o Uintptr) OrElse(value uintptr) uintptr
```
OrElse returns the value wrapped by this Optional, or the value passed in if
there is no value wrapped by this Optional.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
