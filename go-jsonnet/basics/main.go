package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/pkg/errors"
)

func main() {
	vm := jsonnet.MakeVM()
	snippet := `{
		person1: {
		    name: "Alice",
		    welcome: "Hello " + self.name + "!",
		},
		person2: self.person1 { name: "Bob" },
	}`

	// anonymous snippet
	jsonStr, err := vm.EvaluateAnonymousSnippet("example1.jsonnet", snippet)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonStr)

	// native fn
	var x = []interface{}{"abc123"}
	out, err := SHA256().Func(x)
	fmt.Println(out)
}

func SHA256() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "sha256",
		Params: ast.Identifiers{"value"},
		Func: func(dataString []interface{}) (res interface{}, err error) {
			if len(dataString) != 1 {
				return nil, fmt.Errorf("bad arguments to sha256: needs %d argument", 1)
			}
			data, ok := dataString[0].(string)
			if !ok {
				return nil, errors.New("sha256 failed to read input")
			}
			h := sha256.New()
			_, err = h.Write([]byte(data))
			return fmt.Sprintf("%x", h.Sum(nil)), err
		},
	}
}
