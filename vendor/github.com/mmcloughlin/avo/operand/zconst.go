// Code generated by make_const.go. DO NOT EDIT.

package operand

import "fmt"

// I8 is a 8-bit signed integer constant.
type I8 int8

func (i I8) Asm() string { return fmt.Sprintf("$%+d", i) }
func (i I8) Bytes() int  { return 1 }
func (i I8) constant()   {}

// U8 is a 8-bit unsigned integer constant.
type U8 uint8

func (u U8) Asm() string { return fmt.Sprintf("$%#02x", u) }
func (u U8) Bytes() int  { return 1 }
func (u U8) constant()   {}

// I16 is a 16-bit signed integer constant.
type I16 int16

func (i I16) Asm() string { return fmt.Sprintf("$%+d", i) }
func (i I16) Bytes() int  { return 2 }
func (i I16) constant()   {}

// U16 is a 16-bit unsigned integer constant.
type U16 uint16

func (u U16) Asm() string { return fmt.Sprintf("$%#04x", u) }
func (u U16) Bytes() int  { return 2 }
func (u U16) constant()   {}

// F32 is a 32-bit floating point constant.
type F32 float32

func (f F32) Asm() string { return fmt.Sprintf("$(%s)", f) }
func (f F32) Bytes() int  { return 4 }
func (f F32) constant()   {}

// I32 is a 32-bit signed integer constant.
type I32 int32

func (i I32) Asm() string { return fmt.Sprintf("$%+d", i) }
func (i I32) Bytes() int  { return 4 }
func (i I32) constant()   {}

// U32 is a 32-bit unsigned integer constant.
type U32 uint32

func (u U32) Asm() string { return fmt.Sprintf("$%#08x", u) }
func (u U32) Bytes() int  { return 4 }
func (u U32) constant()   {}

// F64 is a 64-bit floating point constant.
type F64 float64

func (f F64) Asm() string { return fmt.Sprintf("$(%s)", f) }
func (f F64) Bytes() int  { return 8 }
func (f F64) constant()   {}

// I64 is a 64-bit signed integer constant.
type I64 int64

func (i I64) Asm() string { return fmt.Sprintf("$%+d", i) }
func (i I64) Bytes() int  { return 8 }
func (i I64) constant()   {}

// U64 is a 64-bit unsigned integer constant.
type U64 uint64

func (u U64) Asm() string { return fmt.Sprintf("$%#016x", u) }
func (u U64) Bytes() int  { return 8 }
func (u U64) constant()   {}