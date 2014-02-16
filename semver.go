package semver

import (
	"strconv"
	"strings"
)

const SEPARATOR string = "."

type Version struct {
	Major int64
	Minor int64
	Patch int64
}

// Return Version as string in format MAJOR.MINOR.PATCH
func (v Version) ToString() string {
	return strings.Join(v.AsSList(), SEPARATOR)
}

// Return Version as array of strings
func (v Version) AsSList() []string {
	return []string{strconv.FormatInt(v.Major, 10), strconv.FormatInt(v.Minor, 10), strconv.FormatInt(v.Patch, 10)}
}

// Current version of go-semver
var VERSION *Version = &Version{1, 0, 0}
