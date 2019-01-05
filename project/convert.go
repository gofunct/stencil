package project

import "strings"

//Convert takes a type and converts it into a different type

func (p Path) ConvertToString() string { return string(p) }

// ConvertStringToComment comments every line of in.
func (p *Project) ConvertStringToComment(in string) string {
	var newlines []string
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "//") {
			newlines = append(newlines, line)
		} else {
			if line == "" {
				newlines = append(newlines, "//")
			} else {
				newlines = append(newlines, "// "+line)
			}
		}
	}
	return strings.Join(newlines, "\n")
}
