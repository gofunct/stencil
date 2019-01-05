package project

import (
	"testing"
)

func (p *Project) TestInitConfig_BuildSpec(t *testing.T) {
	cases := []struct {
		test string
		cfg  Project
		out  string
	}{
		{
			test: "empty",
		},
		{
			test: "HEAD",
			cfg:  Project{HEAD: true},
			out:  "@master",
		},
		{
			test: "branch",
			cfg:  Project{Branch: "foo/bar"},
			out:  "@foo/bar",
		},
		{
			test: "version",
			cfg:  Project{Version: "^0.3.0"},
			out:  "@^0.3.0",
		},
		{
			test: "revision",
			cfg:  Project{Revision: "a2489d2"},
			out:  "@a2489d2",
		},
	}

	for _, tc := range cases {
		t.Run(tc.test, func(t *testing.T) {
			if got, want := tc.cfg.BuildSpec(), tc.out; got != want {
				t.Errorf("BuildSpec() returned %q, want %q", got, want)
			}
		})
	}
}

func TestPath_String(t *testing.T) {
	pathStr := "/go/src/awesomeapp"
	path := Path(pathStr)

	if got, want := path.ConvertToString(), pathStr; got != want {
		t.Errorf("String() returned %q, want %q", got, want)
	}
}

func TestPath_Join(t *testing.T) {
	path := Path("/go/src/awesomeapp")

	if got, want := path.Join("cmd", "server"), Path("/go/src/awesomeapp/cmd/server"); got != want {
		t.Errorf("Join() returned %q, want %q", got, want)
	}
}
