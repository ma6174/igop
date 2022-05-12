package gossa

import (
	"reflect"

	"github.com/goplus/gossa/internal/basic"
	"golang.org/x/tools/go/ssa"
)

func makeTypeChangeInstr(pfn *function, instr *ssa.ChangeType) func(fr *frame) {
	typ := pfn.Interp.preToType(instr.Type())
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	if kx.isStatic() {
		var v interface{}
		if vx == nil {
			v = reflect.New(typ).Elem().Interface()
		} else {
			v = reflect.ValueOf(vx).Convert(typ).Interface()
		}
		return func(fr *frame) {
			fr.setReg(ir, v)
		}
	}
	kind := typ.Kind()
	switch kind {
	case reflect.Ptr, reflect.Chan, reflect.Map, reflect.Func, reflect.Slice:
		t := basic.TypeOfType(typ)
		return func(fr *frame) {
			x := fr.reg(ix)
			fr.setReg(ir, basic.ConvertPtr(t, x))
		}
	case reflect.Struct, reflect.Array:
		t := basic.TypeOfType(typ)
		return func(fr *frame) {
			x := fr.reg(ix)
			fr.setReg(ir, basic.ConvertDirect(t, x))
		}
	case reflect.Interface:
		return func(fr *frame) {
			x := fr.reg(ix)
			if x == nil {
				fr.setReg(ir, reflect.New(typ).Elem().Interface())
			} else {
				fr.setReg(ir, reflect.ValueOf(x).Convert(typ).Interface())
			}
		}
	case reflect.Bool:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Bool(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertBool(t, x))
			}
		}
	case reflect.Int:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt(t, x))
			}
		}
	case reflect.Int8:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int8(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt8(t, x))
			}
		}
	case reflect.Int16:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int16(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt16(t, x))
			}
		}
	case reflect.Int32:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int32(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt32(t, x))
			}
		}
	case reflect.Int64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt64(t, x))
			}
		}
	case reflect.Uint:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint(t, x))
			}
		}
	case reflect.Uint8:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint8(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint8(t, x))
			}
		}
	case reflect.Uint16:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint16(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint16(t, x))
			}
		}
	case reflect.Uint32:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint32(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint32(t, x))
			}
		}
	case reflect.Uint64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint64(t, x))
			}
		}
	case reflect.Uintptr:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uintptr(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUintptr(t, x))
			}
		}
	case reflect.Float32:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Float32(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertFloat32(t, x))
			}
		}
	case reflect.Float64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Float64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertFloat64(t, x))
			}
		}
	case reflect.Complex64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Complex64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertComplex64(t, x))
			}
		}
	case reflect.Complex128:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Complex128(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertComplex128(t, x))
			}
		}
	case reflect.String:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.String(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertString(t, x))
			}
		}
	}
	panic("unreachable")
}

func makeConvertInstr(pfn *function, interp *Interp, instr *ssa.Convert) func(fr *frame) {
	typ := interp.preToType(instr.Type())
	xtyp := interp.preToType(instr.X.Type())
	kind := typ.Kind()
	xkind := xtyp.Kind()
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	switch kind {
	case reflect.UnsafePointer:
		if xkind == reflect.Uintptr {
			return func(fr *frame) {
				v := fr.uintptr(ix)
				fr.setReg(ir, toUnsafePointer(v))
			}
		} else if xkind == reflect.Ptr {
			return func(fr *frame) {
				v := fr.pointer(ix)
				fr.setReg(ir, v)
			}
		}
	case reflect.Uintptr:
		if xkind == reflect.UnsafePointer {
			return func(fr *frame) {
				v := fr.pointer(ix)
				fr.setReg(ir, uintptr(v))
			}
		}
	case reflect.Ptr:
		if xkind == reflect.UnsafePointer {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				v := fr.reg(ix)
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	case reflect.Slice:
		if xkind == reflect.String {
			t := basic.TypeOfType(typ)
			elem := typ.Elem()
			switch elem.Kind() {
			case reflect.Uint8:
				return func(fr *frame) {
					v := fr.string(ix)
					fr.setReg(ir, basic.Make(t, []byte(v)))
				}
			case reflect.Int32:
				return func(fr *frame) {
					v := fr.string(ix)
					fr.setReg(ir, basic.Make(t, []rune(v)))
				}
			}
		}
	case reflect.String:
		if xkind == reflect.Slice {
			t := basic.TypeOfType(typ)
			elem := xtyp.Elem()
			switch elem.Kind() {
			case reflect.Uint8:
				return func(fr *frame) {
					v := fr.bytes(ix)
					fr.setReg(ir, basic.Make(t, string(v)))
				}
			case reflect.Int32:
				return func(fr *frame) {
					v := fr.runes(ix)
					fr.setReg(ir, basic.Make(t, string(v)))
				}
			}
		}
	}
	if kx.isStatic() {
		v := reflect.ValueOf(vx)
		return func(fr *frame) {
			fr.setReg(ir, v.Convert(typ).Interface())
		}
	}
	switch kind {
	case reflect.Int:
		return cvtInt(ir, ix, xkind, xtyp, typ)
	case reflect.Int8:
		return cvtInt8(ir, ix, xkind, xtyp, typ)
	case reflect.Int16:
		return cvtInt16(ir, ix, xkind, xtyp, typ)
	case reflect.Int32:
		return cvtInt32(ir, ix, xkind, xtyp, typ)
	case reflect.Int64:
		return cvtInt64(ir, ix, xkind, xtyp, typ)
	case reflect.Uint:
		return cvtUint(ir, ix, xkind, xtyp, typ)
	case reflect.Uint8:
		return cvtUint8(ir, ix, xkind, xtyp, typ)
	case reflect.Uint16:
		return cvtUint16(ir, ix, xkind, xtyp, typ)
	case reflect.Uint32:
		return cvtUint32(ir, ix, xkind, xtyp, typ)
	case reflect.Uint64:
		return cvtUint64(ir, ix, xkind, xtyp, typ)
	case reflect.Uintptr:
		return cvtUintptr(ir, ix, xkind, xtyp, typ)
	case reflect.Float32:
		return cvtFloat32(ir, ix, xkind, xtyp, typ)
	case reflect.Float64:
		return cvtFloat64(ir, ix, xkind, xtyp, typ)
	case reflect.Complex64:
		return cvtComplex64(ir, ix, xkind, xtyp, typ)
	case reflect.Complex128:
		return cvtComplex128(ir, ix, xkind, xtyp, typ)
	}
	return func(fr *frame) {
		v := reflect.ValueOf(fr.reg(ix))
		fr.setReg(ir, v.Convert(typ).Interface())
	}
}