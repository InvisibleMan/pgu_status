package ppot

import (
	"io/ioutil"
	// "log"
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
		er     bool
		ummsID string
	}{
		{"xml/response_success.xml", false, "161015734"},
		{"xml/response_fail.xml", true, "161015735"},
	}

	for _, table := range tables {
		msg, _ := NewResultParser().Parse([]byte(ReadFixture(table.path)))
		if msg.IsError() != table.er {
			t.Errorf("Result was incorrect, got: %t, want: %t.", msg.IsError(), table.er)
		}

		if msg.UmmsID() != table.ummsID {
			t.Errorf("UmmsID was incorrect, got: %s, want: %s.", msg.UmmsID(), table.ummsID)
		}
	}
}

func init() {
	wd, _ := os.Getwd()
	ROOT = wd
}
