package runtime

import (
	"io"
	"io/ioutil"
	"text/template"
)

// writes templates to the writer
func execTemplate(file string, wr io.Writer, data interface{}) {
	dat, err := ioutil.ReadFile(file)
	tmpl, err := template.New("test").Funcs(funcMap()).Parse(string(dat))
	checkErr(err)
	err = tmpl.Execute(wr, data)
	checkErr(err)
}
