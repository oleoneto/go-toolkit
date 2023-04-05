package cmd

import "fmt"

type Version struct {
	Major string `json:"major" yaml:"major" toml:"major" xml:"major"`
	Minor string `json:"minor" yaml:"minor" toml:"minor" xml:"minor"`
	Patch string `json:"patch" yaml:"patch" toml:"patch" xml:"patch"`
}

func (V *Version) String() string {
	return fmt.Sprintf(`go-toolkit %v.%v.%v`, V.Major, V.Minor, V.Patch)
}

var version = Version{
	Major: "0",
	Minor: "1",
	Patch: "0",
}
