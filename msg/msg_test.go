package msg

import (
	"io/ioutil"
	"os"
	"testing"
)

var ROOT string

func ReadFile(path string) (xml string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ReadFixture(fixtureName string) (xml string) {
	return ReadFile(ROOT + "/../fixuters/" + fixtureName)
}

func TestParse(t *testing.T) {
	tables := []struct {
		path   string
		res    bool
		ummsID string
	}{
		{"xml/response_success.xml", true, "161015734"},
		// {"xml/response_fail.xml", false, "161015734"},
	}

	for _, table := range tables {
		msg, _ := Parse(ReadFixture(table.path))
		if msg.IsError != table.res {
			t.Errorf("Result was incorrect, got: %t, want: %t.", msg.IsError, table.res)
		}
	}
}

func init() {
	wd, _ := os.Getwd()
	ROOT = wd
}
