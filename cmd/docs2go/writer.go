package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

const (
	prefix = "zzz_"
)

var (
	goTmpl = template.Must(template.New("source").Funcs(funcMap).Parse(source))

	funcMap = map[string]interface{}{
		"formatParents":              formatParents,
		"fromJSObject":               fromJSObject,
		"constructorParameterGoList": constructorParameterGoList,
		"constructorParameterJSList": constructorParameterJSList,
		"methodParameterGoList":      methodParameterGoList,
		"methodParameterJSList":      methodParameterJSList,
		"propertyParameterGoList":    propertyParameterGoList,
		"propertyParameterJSList":    propertyParameterJSList,
		"receiver":                   receiver,
		"toLower":                    toLower,
	}
)

func (c *classes) writeGo(destDir string) error {
	for k, v := range c.m {
		if strings.Contains(k, "&") {
			logf("Skipping %v", k)
			continue
		}

		var buf bytes.Buffer
		if err := goTmpl.Execute(&buf, v); err != nil {
			return fmt.Errorf("%v: template failed: %v", k, err)
		}

		b := buf.Bytes()
		for _, v := range regexpOverrides[k] {
			b = v.re.ReplaceAll(b, []byte(v.repl))
		}

		filename := filepath.Join(destDir, prefix+k+".go")

		clean, err := format.Source(b)
		if err != nil {
			return fmt.Errorf("%v: unable to format: %v\n%v", k, err, string(b))
		}

		logf("Writing %v...", filename)
		if err := ioutil.WriteFile(filename, clean, 0644); err != nil {
			return err
		}
	}
	return nil
}

func formatParents(parents []string) string {
	var lines []string
	for _, v := range parents {
		lines = append(lines, "*"+v)
	}
	if len(lines) == 0 {
		lines = append(lines, "p js.Value")
	}
	lines = append(lines, "ctx js.Value")
	return strings.Join(lines, "\n")
}

func fromJSObject(parents []string) string {
	if len(parents) == 0 {
		return "p: p, ctx: ctx"
	}
	var wraps []string
	for _, p := range parents {
		wraps = append(wraps, fmt.Sprintf("%v: %vFromJSObject(p, ctx)", p, p))
	}
	wraps = append(wraps, "ctx: ctx")
	return strings.Join(wraps, ", ")
}

func receiver(name string) string {
	return strings.ToLower(name[0:1])
}

func toLower(s string) string {
	return strings.ToLower(s)
}

func constructorParameterGoList(key string, s *Signature) string {
	var params []string
	for i, name := range s.GoParamsName {
		params = append(params, fmt.Sprintf("%v %v", name, s.GoParamsType[i]))
	}

	if s.HasOpts {
		params = append(params, fmt.Sprintf("opts *%vOpts", key))
	}

	return strings.Join(params, ", ")
}

func methodParameterGoList(name, key string, s *Signature) string {
	var params []string
	for i, name := range s.GoParamsName {
		params = append(params, fmt.Sprintf("%v %v", name, s.GoParamsType[i]))
	}

	if s.HasOpts {
		params = append(params, fmt.Sprintf("opts *%v%vOpts", name, key))
	}

	return strings.Join(params, ", ")
}

func propertyParameterGoList(key string, s *Signature) string {
	var params []string
	for i, name := range s.GoParamsName {
		params = append(params, fmt.Sprintf("%v %v", name, s.GoParamsType[i]))
	}

	if s.HasOpts {
		params = append(params, fmt.Sprintf("opts *%vOpts", key))
	}

	return strings.Join(params, ", ")
}

func constructorParameterJSList(key string, s *Signature) []string {
	return s.JSParams
}

func methodParameterJSList(key string, s *Signature) string {
	logf("constructorParameterJSList: method: key=%v", key)
	return strings.Join(s.JSParams, ", ")
}

func propertyParameterJSList(key string, s *Signature) string {
	logf("constructorParameterJSList: property: key=%v", key)
	return strings.Join(s.JSParams, ", ")
}

type reOverride struct {
	re   *regexp.Regexp
	repl string
}

var regexpOverrides = map[string][]reOverride{}

const source = `// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

{{$root := .}}
{{$name := .Name}}

// {{$name}} represents a babylon.js {{$name}}.
// {{.Summary}}
{{- with .Description}}
//
// {{.}}
{{- end -}}{{- with .SeeURL}}
//
// See: {{.}}
{{- end}}
type {{$name}} struct{ {{.Parents | formatParents}} }

// JSObject returns the underlying js.Value.
func ({{$name | receiver}} *{{$name}}) JSObject() js.Value { return {{$name | receiver}}.p }

// {{$name}} returns a {{$name}} JavaScript class.
func ({{.ConstructorNamespaceReceiverName}} {{.ConstructorNamespaceReceiverType}}) {{$name}}() *{{$name}} {
	p := {{.ConstructorNamespaceReceiverName}}.ctx.Get("{{$name}}")
	return {{$name}}FromJSObject(p, {{.ConstructorNamespaceReceiverName}}.ctx)
}

// {{$name}}FromJSObject returns a wrapped {{$name}} JavaScript class.
func {{$name}}FromJSObject(p js.Value, ctx js.Value) *{{$name}} {
	return &{{$name}}{ {{.Parents | fromJSObject}} }
}

// {{$name}}ArrayToJSArray returns a JavaScript Array for the wrapped array.
func {{$name}}ArrayToJSArray(array []*{{$name}}) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

{{range $key, $value := .ConstructorNames}}{{if $value.HasOpts}}
// {{$key}}Opts contains optional parameters for {{$key}}.
type {{$key}}Opts struct {
{{range $index, $optType := $value.GoOptsType -}}
  {{index $value.GoOptsName $index}} {{$optType}}
{{end}}
}
{{end}}

// {{$key}} returns a new {{$name}} object.
//
// https://doc.babylonjs.com/api/classes/babylon.{{$name | toLower}}
func ({{$value.ConstructorNamespaceReceiverName}} {{$value.ConstructorNamespaceReceiverType}}) {{$key}}({{constructorParameterGoList $key $value}}) *{{$name}} { {{if $value.HasOpts}}
if opts == nil {
  opts = &{{$key}}Opts{}
}
{{end}}

  args := make([]interface{}, 0, {{$value.JSParams | len}} + {{$value.JSOpts | len}})

  {{range $index, $element := $value.JSParams -}}
  args = append(args, {{$element}})
  {{end}}

  {{range $index, $element := $value.JSOpts -}}{{if index $value.GoOptsType $index | eq "js.Value"}}args = append(args, opts.{{index $value.GoOptsName $index}})
  {{else -}}
  if opts.{{index $value.GoOptsName $index}} == nil {
    args = append(args, js.Undefined())
  } else {
    args = append(args, {{$element}})
  }
  {{end}}{{end}}

	p := {{$value.ConstructorNamespaceReceiverName}}.ctx.Get("{{$name}}").New(args...)
	return {{$name}}FromJSObject(p, {{$value.ConstructorNamespaceReceiverName}}.ctx)
}
{{end}}

{{range $key, $value := .MethodNames}}{{if $value.HasOpts}}
// {{$name}}{{$key}}Opts contains optional parameters for {{$name}}.{{$key}}.
type {{$name}}{{$key}}Opts struct {
{{range $index, $optType := $value.GoOptsType -}}
  {{index $value.GoOptsName $index}} {{$optType}}
{{end}}
}
{{end}}

// {{$key}} calls the {{$key}} method on the {{$name}} object.
//
// https://doc.babylonjs.com/api/classes/babylon.{{$name | toLower}}#{{$key | toLower}}
func ({{$name | receiver}} *{{$name}}) {{$key}}({{methodParameterGoList $name $key $value}}) {{$value.GoReturnType}} { {{if $value.HasOpts}}
if opts == nil {
opts = &{{$name}}{{$key}}Opts{}
}
{{end}}

  {{if or $value.JSParams $value.JSOpts}}args := make([]interface{}, 0, {{$value.JSParams | len}} + {{$value.JSOpts | len}}){{end}}

  {{range $index, $element := $value.JSParams -}}
  args = append(args, {{if index $value.NeedsArrayHelper $index}}{{index $value.NeedsArrayHelper $index}}{{else}}{{$element}}{{end}})
  {{end}}

  {{range $index, $element := $value.JSOpts -}}{{if index $value.GoOptsType $index | eq "js.Value"}}args = append(args, opts.{{index $value.GoOptsName $index}})
  {{else -}}
  if opts.{{index $value.GoOptsName $index}} == nil {
    args = append(args, js.Undefined())
  } else {
    args = append(args, {{$element}})
  }
  {{end}}{{end}}

	{{if $value.GoReturnType}}retVal := {{end}}{{$name | receiver}}.p.Call("{{$value.JSName}}"{{if or $value.JSParams $value.JSOpts}}, args...{{end}}){{if $value.GoReturnType}}
  {{$value.GoReturnStatement}}{{end}}
}
{{end}}

{{range $key, $value := .PropertyNames}}
// {{$key}} returns the {{$key}} property of class {{$name}}.
//
// https://doc.babylonjs.com/api/classes/babylon.{{$name | toLower}}#{{$key | toLower}}
func ({{$name | receiver}} *{{$name}}) {{$key}}() {{$value.GoReturnType}} {
	retVal := {{$name | receiver}}.p.Get("{{$value.JSName}}")
	{{$value.GoReturnStatement}}
}

{{if $value.WriteSetter}}// Set{{$key}} sets the {{$key}} property of class {{$name}}.
//
// https://doc.babylonjs.com/api/classes/babylon.{{$name | toLower}}#{{$key | toLower}}
func ({{$name | receiver}} *{{$name}}) Set{{$key}}({{propertyParameterGoList $key $value}}) *{{$name}} {
	{{$name | receiver}}.p.Set("{{$value.JSName}}", {{index $value.JSParams 0}})
	return {{$name | receiver}}
}
{{end}}{{end}}
`
