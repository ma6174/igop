// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package ring

import (
	q "container/ring"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "ring",
		Path:       "container/ring",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Ring": {reflect.TypeOf((*q.Ring)(nil)).Elem(), "", "Do,Len,Link,Move,Next,Prev,Unlink,init"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
