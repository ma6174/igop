// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package errors

import (
	q "errors"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "errors",
		Path: "errors",
		Deps: map[string]string{
			"internal/reflectlite": "reflectlite",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"As":     reflect.ValueOf(q.As),
			"Is":     reflect.ValueOf(q.Is),
			"New":    reflect.ValueOf(q.New),
			"Unwrap": reflect.ValueOf(q.Unwrap),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
