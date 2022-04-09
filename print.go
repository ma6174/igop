package gossa

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"

	"golang.org/x/tools/go/ssa"
)

func writevalue(buf *bytes.Buffer, v interface{}) {
	switch v := v.(type) {
	case float64:
		writefloat(buf, v)
	case float32:
		writefloat(buf, float64(v))
	case complex128:
		writecomplex(buf, v)
	case complex64:
		writecomplex(buf, complex128(v))
	case nil, bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		string:
		fmt.Fprintf(buf, "%v", v)
	case *ssa.Function, *ssa.Builtin, *closure:
		fmt.Fprintf(buf, "%p", v) // (an address)
	case tuple:
		// Unreachable in well-formed Go programs
		buf.WriteString("(")
		for i, e := range v {
			if i > 0 {
				buf.WriteString(", ")
			}
			writeValue(buf, e)
		}
		buf.WriteString(")")
	default:
		i := reflect.ValueOf(v)
		switch i.Kind() {
		case reflect.Float32, reflect.Float64:
			writefloat(buf, i.Float())
		case reflect.Complex64, reflect.Complex128:
			writecomplex(buf, i.Complex())
		case reflect.Map, reflect.Ptr, reflect.Func, reflect.Chan, reflect.UnsafePointer:
			fmt.Fprintf(buf, "%p", v)
		case reflect.Slice:
			fmt.Fprintf(buf, "[%v/%v]%p", i.Len(), i.Cap(), v)
		case reflect.String:
			fmt.Fprintf(buf, "%v", v)
		case reflect.Interface:
			eface := *(*emptyInterface)(unsafe.Pointer(&i))
			fmt.Fprintf(buf, "(%p,%p)", eface.typ, eface.word)
		case reflect.Struct, reflect.Array:
			panic(fmt.Errorf("illegal types for operand: print %T", v))
		default:
			fmt.Fprintf(buf, "%v", v)
		}
	}
}

func writeinterface(out *bytes.Buffer, i interface{}) {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	fmt.Fprintf(out, "(%p,%p)", eface.typ, eface.word)
}

func writecomplex(out *bytes.Buffer, c complex128) {
	out.WriteByte('(')
	writefloat(out, real(c))
	writefloat(out, imag(c))
	out.WriteString("i)")
}

func writefloat(out *bytes.Buffer, v float64) {
	switch {
	case v != v:
		out.WriteString("NaN")
		return
	case v+v == v && v > 0:
		out.WriteString("+Inf")
		return
	case v+v == v && v < 0:
		out.WriteString("-Inf")
		return
	}

	const n = 7 // digits printed
	var buf [n + 7]byte
	buf[0] = '+'
	e := 0 // exp
	if v == 0 {
		if 1/v < 0 {
			buf[0] = '-'
		}
	} else {
		if v < 0 {
			v = -v
			buf[0] = '-'
		}

		// normalize
		for v >= 10 {
			e++
			v /= 10
		}
		for v < 1 {
			e--
			v *= 10
		}

		// round
		h := 5.0
		for i := 0; i < n; i++ {
			h /= 10
		}
		v += h
		if v >= 10 {
			e++
			v /= 10
		}
	}

	// format +d.dddd+edd
	for i := 0; i < n; i++ {
		s := int(v)
		buf[i+2] = byte(s + '0')
		v -= float64(s)
		v *= 10
	}
	buf[1] = buf[2]
	buf[2] = '.'

	buf[n+2] = 'e'
	buf[n+3] = '+'
	if e < 0 {
		e = -e
		buf[n+3] = '-'
	}

	buf[n+4] = byte(e/100) + '0'
	buf[n+5] = byte(e/10)%10 + '0'
	buf[n+6] = byte(e%10) + '0'
	out.Write(buf[:])
}
