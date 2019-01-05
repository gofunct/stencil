package project

import (
	"bytes"
	"github.com/gofunct/common/errors"
	"github.com/gofunct/stencil/funcmap"
	"text/template"
)

func MustCreateTemplate(name, tmpl string) *template.Template {
	return template.Must(template.New(name).Funcs(funcmap.TxtFuncMap()).Parse(tmpl))
}

// TemplateString is a compilable string with text/template package
type TemplateString string

// Compile generates textual output applied a parsed template to the specified values
func (s TemplateString) Compile(v interface{}) (string, error) {
	tmpl, err := template.New("").Funcs(funcmap.TxtFuncMap()).Parse(string(s))
	if err != nil {
		return string(s), errors.Wrapf(err, "failed to parse a template: %q", string(s))
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, v)
	if err != nil {
		return string(s), errors.Wrapf(err, "failed to execute a template: %q", string(s))
	}
	return string(buf.Bytes()), nil
}

func (p *Project) ExecuteTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("").Funcs(funcmap.TxtFuncMap()).Parse(tmplStr)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	return buf.String(), err
}
