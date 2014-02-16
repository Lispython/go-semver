package semver

import (

"testing"
)

var TEST_VERSION *Version = &Version{10, 5, 6}

func TestToString(t *testing.T){
	if v := TEST_VERSION.ToString(); v != "10.5.6" {
		t.Fatalf("Invalid ToString: %s != 10.5.6", v)
	}
}

func TestAsSlist(t *testing.T) {

	if v := TEST_VERSION.AsSList(); v[0] != "10" || v[1] != "5" || v[2] != "6" {
		t.Fatalf("Invalid AsSlist")
	}
}