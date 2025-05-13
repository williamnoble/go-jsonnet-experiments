package lib

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-jsonnet"
	"log"
	"text/template"
)

// F is a function that wraps build and output
func F(file string, function string, testParams string) any {
	build, err := build(file, function, testParams)
	if err != nil {
		log.Fatal(err)
	}
	return output(build)
}

func output(s string) any {
	type Output struct {
		Output any `json:"output"`
	}

	var output Output
	err := json.Unmarshal([]byte(s), &output)
	if err != nil {
		panic(err)
	}
	return output.Output
}

func build(file string, function string, testParams string) (json string, err error) {
	vm := jsonnet.MakeVM()
	snippet := snippetTemplate(file, function, testParams).String()
	return vm.EvaluateAnonymousSnippet("utils.libsonnet", snippet)
}

func snippetTemplate(file string, function string, testParams string) *bytes.Buffer {
	b := new(bytes.Buffer)
	snippet := `
	local utils = import '{{ .ImportFile }}.libsonnet';
	{ output: utils.{{ .Function }}( {{ .TestParams }} ) }
	`
	t := template.Must(template.New("index").Parse(snippet))
	p := struct {
		ImportFile string
		Function   string
		TestParams any // TODO: updated from string
	}{
		ImportFile: file,
		Function:   function,
		TestParams: testParams,
	}
	_ = t.Execute(b, p)
	return b
}
