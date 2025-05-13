package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-jsonnet"
	"text/template"
)

type output struct {
	Output any `json:"output"`
}

func main() {
	vm := jsonnet.MakeVM()

	// from literal
	s, _ := vm.EvaluateAnonymousSnippet("utils.libsonnet", `
	local utils = import 'utils.libsonnet';
	{ output: utils.foo("great","bear") }
	`)

	fmt.Println("Literal:", s)
	fmt.Println(extractOutput(s))

	// from templated
	s, _ = vm.EvaluateAnonymousSnippet("utils.libsonnet", snippetFromTemplate("utils", "foo", `"great", "bear"`).String())

	fmt.Println("Templated:", s)
	fmt.Println(extractOutput(s))
}

// f is a wrapper of mustBuild and extractOutput
func f(file string, function string, testParams string) any { // TODO: Updated
	build := mustBuild(file, function, testParams)
	out := extractOutput(build)
	return out
}

func extractOutput(s string) any { // TODO: Updated
	var output output
	err := json.Unmarshal([]byte(s), &output)
	if err != nil {
		panic(err)
	}
	return output.Output
}

func mustBuild(file string, function string, testParams string) (json string) {
	vm := jsonnet.MakeVM()
	s, err := vm.EvaluateAnonymousSnippet("utils.libsonnet", snippetFromTemplate(file, function, testParams).String())
	if err != nil {
		panic(err)
	}
	return s
}

// snippetFromTemplate builds an anonymous snippet using Go templates
func snippetFromTemplate(file string, function string, testParams string) *bytes.Buffer {
	b := new(bytes.Buffer)
	snippet := `
	local utils = import '{{ .ImportFile }}.libsonnet';
	{ output: utils.{{ .Function }}( {{ .TestParams }} ) }
	`
	t := template.Must(template.New("index").Parse(snippet))
	p := struct {
		ImportFile string
		Function   string
		TestParams any
	}{
		ImportFile: file,
		Function:   function,
		TestParams: testParams,
	}
	err := t.Execute(b, p)
	if err != nil {
		panic(err)
	}
	return b
}
