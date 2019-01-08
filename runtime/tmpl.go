package runtime

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"go.uber.org/zap"
	"text/template"
)

func (u *UI) ExecuteTemplate(tmplStr string) string {
	tmpl, err := template.New("").Funcs(sprig.GenericFuncMap()).Parse(tmplStr)
	if err != nil {
		u.Z.Fatal("failed to execute template", zap.Error(err))
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, u.Config)
	if err != nil {
		u.Z.Fatal("failed to execute template", zap.Error(err))
	}
	return buf.String()
}
