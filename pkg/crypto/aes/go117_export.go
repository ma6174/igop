// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package aes

import (
	q "crypto/aes"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "aes",
		Path: "crypto/aes",
		Deps: map[string]string{
			"crypto/cipher":          "cipher",
			"crypto/internal/subtle": "subtle",
			"crypto/subtle":          "subtle",
			"encoding/binary":        "binary",
			"errors":                 "errors",
			"internal/cpu":           "cpu",
			"strconv":                "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"KeySizeError": {reflect.TypeOf((*q.KeySizeError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
