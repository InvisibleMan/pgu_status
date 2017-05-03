package sx

import (
	"io/ioutil"
	// "log"
	"os"
	"testing"
)

var ROOT string

func ReadFile(path string) (xml []byte) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func ReadFixture(fixtureName string) (xml []byte) {
	return ReadFile(ROOT + "/../fixuters/" + fixtureName)
}

func TestParse(t *testing.T) {
	tables := []struct {
		path  string
		IsNil bool
	}{
		{"xml/sx_success.xml", true},
		{"xml/sx_fail.xml", false},
	}

	for _, table := range tables {
		xml := ReadFixture(table.path)
		// log.Println(xml)
		err := Parse(xml)
		if (err == nil) != table.IsNil {
			t.Errorf("Result was incorrect: %v", err)
		}
	}
}

func init() {
	wd, _ := os.Getwd()
	ROOT = wd
}
