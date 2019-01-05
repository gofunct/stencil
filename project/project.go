package project

import (
	"bytes"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

// Name returns the ProjName of project, e.g. "github.com/spf13/cobra"
func (p Project) Name() string {
	return p.ProjName
}

// CmdPath returns absolute path to directory, where all commands are located.
func (p *Project) CmdPath(v *viper.Viper) string {
	if p.AbsolutePath == "" {
		return ""
	}
	if p.CommandPath == "" {
		p.CommandPath = filepath.Join(p.AbsolutePath, p.FindCmdDir(p.AbsolutePath, v))
	}
	return p.CommandPath
}

// AbsPath returns absolute path of project.
func (p Project) AbsPath() string {
	return p.AbsolutePath
}

// SrcPath returns absolute path to $GOPATH/src where project is located.
func (p *Project) SrcPath(v *viper.Viper) string {
	srcPaths := p.FindSourcePath(v)
	if p.SourcePath != "" {
		return p.SourcePath
	}
	if p.AbsolutePath == "" {
		p.SourcePath = srcPaths[0]
		return p.SourcePath
	}

	for _, SourcePath := range srcPaths {
		if p.CheckFilepathHasPrefix(p.AbsolutePath, SourcePath) {
			p.SourcePath = SourcePath
			break
		}
	}

	return p.SourcePath
}

// License returns the License object of project.
func (p *Project) License() License {
	if p.LicenseType.Text == "" && p.LicenseType.Name != "None" {
		p.LicenseType = p.GetLicense()
	}
	return p.LicenseType
}

func (p *Project) CopyrightLine(v *viper.Viper) string {
	author := v.GetString("author")

	year := v.GetString("year") // For tests.
	if year == "" {
		year = time.Now().Format("2006")
	}

	return "Copyright © " + year + " " + author
}

func (p *Project) CreateLicenseFile(path string, v *viper.Viper) {
	data := make(map[string]interface{})
	data["copyright"] = p.CopyrightLine(v)
	LicenseType := Licenses["mit"]
	// Generate LicenseType template from text and data.
	text, err := p.ExecuteTemplate(LicenseType.Text, data)
	if err != nil {
		p.Exit(err)
	}

	// Write LicenseType text to LICENSE file.
	err = p.WriteStringToFile(filepath.Join(path, "LICENSE"), text)
	if err != nil {
		p.Exit(err)
	}
}

func (p *Project) GetLicense() License {
	// If user didn't set any LicenseType, use Apache 2.0 by default.
	return Licenses["mit"]
}

func (p *Project) Exit(msg interface{}) {
	p.Notify.Error(msg.(string))
	os.Exit(1)
}

func (p *Project) BuildSpec() string {
	buf := bytes.NewBufferString("")
	var constraint string
	switch {
	case p.Revision != "":
		constraint = p.Revision
	case p.Branch != "":
		constraint = p.Branch
	case p.HEAD:
		constraint = "master"
	case p.Version != "":
		constraint = p.Version
	}
	if constraint != "" {
		buf.WriteString("@" + constraint)
	}
	return buf.String()
}
