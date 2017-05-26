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
		path           string
		er             bool
		ExternalCaseID string
	}{
		{"xml/ppot_response/success_01.xml", false, "169568750"},
		{"xml/ppot_response/success_02.xml", false, "161015734"},

		{"xml/ppot_response/fail_01.xml", true, "169568750"},
		{"xml/ppot_response/fail_05.xml", true, "161015734"},
	}

	for _, table := range tables {
		msg, _ := NewResultParser().Parse([]byte(ReadFixture(table.path)))
		if msg.IsError() != table.er {
			t.Errorf("Result was incorrect. File: '%s', got: %t, want: %t.", table.path, msg.IsError(), table.er)
		}

		if msg.ExternalCaseID() != table.ExternalCaseID {
			t.Errorf("UmmsID was incorrect, got: %s, want: %s.", msg.ExternalCaseID(), table.ExternalCaseID)
		}
	}
}

func init() {
	wd, _ := os.Getwd()
	ROOT = wd
}
