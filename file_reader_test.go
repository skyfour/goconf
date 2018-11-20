package goconf

import "testing"

const testFileName  = "readme.md"

// read config file
func TestReadConfigFile(t *testing.T) {
	d := ReadConfigFile(testFileName)
	if d == nil {
		t.FailNow()
	}
}
