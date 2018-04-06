package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var (
	tmpl = template.Must(template.New("protoc").Parse(`protoc -I
	{{- range $index, $include := .Includes -}}
		{{if $index}}` + string(filepath.ListSeparator) + `{{end -}}
			{{.}}
	{{- end -}}
	{{- if .Descriptors}} --include_imports --descriptor_set_out={{.Descriptors}}{{- end }} --
	{{- .Name -}}_out={{if .Plugins}}plugins={{- range $index, $plugin := .Plugins -}}
		{{- if $index}}+{{end}}
		{{- $plugin}}
	{{- end -}}
	,{{- end -}}import_path={{.ImportPath}}
	{{- range $proto, $gopkg := .PackageMap -}},M
		{{- $proto}}={{$gopkg -}}
	{{- end -}}
	:{{- .OutputDir }}
	{{- range .Files}} {{.}}{{end -}}
`))

	tmplV2 = template.Must(template.New("protoc").Parse(`protoc -I
		{{- range $index, $include := .Includes -}}
		{{if $index}}` + string(filepath.ListSeparator) + `{{end -}}
			{{.}}
	{{- end -}}{{$importPath := .ImportPath}}{{$packageMap := .PackageMap}}{{$outputDir := .OutputDir}}
	{{- if .Descriptors}} --include_imports --descriptor_set_out={{.Descriptors}}{{- end}}{{- range $gen := .Generator }} --
	{{- $gen.Name -}}_out={{if $gen.Options}}{{- range $key, $val := $gen.Options -}}
		{{$key}}={{$val}},{{end}}{{end}}{{- if $gen.Plugin}}plugins={{- range $index, $plugin := $gen.Plugin -}}
		{{- if $index}}+{{- end -}}
		{{- $plugin.Name }}{{- if $plugin.Options -}}={{- range $key, $val := $plugin.Options -}}{{- $key -}}={{- $val -}}{{- if $key}},{{end}}{{end}}
	{{end -}}{{- end -}}{{- end -}}
	{{- if $gen.Plugin -}},{{- end -}}import_path={{$importPath}}
	{{- range $proto, $gopkg := $packageMap -}},M
		{{- $proto}}={{$gopkg -}}
	{{- end -}}:{{- $outputDir }}{{- end -}}
	{{- range .Files}} {{.}}{{end -}}
`))
)

// protocParams defines inputs to a protoc command string. V2
type protocCmdV2 struct {
	Includes    []string
	Generator   []Generator
	Descriptors string
	ImportPath  string
	PackageMap  map[string]string
	Files       []string
	OutputDir   string
}

// protocParams defines inputs to a protoc command string.
type protocCmd struct {
	Name        string // backend name
	Includes    []string
	Plugins     []string
	Descriptors string
	ImportPath  string
	PackageMap  map[string]string
	Files       []string
	OutputDir   string
}

func (p *protocCmdV2) mkcmd() (string, error) {
	var buf bytes.Buffer
	if err := tmplV2.Execute(&buf, p); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (p *protocCmd) mkcmd() (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, p); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (p *protocCmdV2) run() error {
	arg, err := p.mkcmd()
	if err != nil {
		return err
	}

	// pass to sh -c so we don't need to re-split here.
	args := []string{shArg, arg}
	cmd := exec.Command(shCmd, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (p *protocCmd) run() error {
	arg, err := p.mkcmd()
	if err != nil {
		return err
	}

	// pass to sh -c so we don't need to re-split here.
	args := []string{shArg, arg}
	cmd := exec.Command(shCmd, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
