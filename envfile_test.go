package envfile

import (
	"errors"
	"reflect"
	"testing"
)

func TestBadFile(t *testing.T) {

	envMap := make(map[string]string)

	expectedErr := errors.New("Unable to parse env var on line 4 - Cannot locate token =")

	err := ReadEnvFile("./fixtures/bad.env", envMap)

	if err == nil {
		t.Fatal("Should error reading file")
	}

	if !reflect.DeepEqual(expectedErr, err) {
		t.Fatalf("Expected %s actual %s", expectedErr, err)
	}
}

func TestGoodFile(t *testing.T) {

	envMap := make(map[string]string)

	err := ReadEnvFile("./fixtures/good.env", envMap)

	if err != nil {
		t.Fatalf("Error reading good file %s", err)
	}

	envMapExpected := map[string]string{
		"keyone": "1234",
		"keytwo": "5678",
	}

	if !reflect.DeepEqual(envMapExpected, envMap) {
		t.Fatalf("Expected %s actual %s", envMapExpected, envMap)
	}

}

func TestParse(t *testing.T) {
	var tests = []struct {
		line  string
		key   string
		value string
	}{
		{"seen=bbb ", "seen", "bbb"},
		{"seen=     bbb", "seen", "bbb"},
		{"seen=", "seen", ""},
		{"keyone=1234", "keyone", "1234"},
	}

	for _, entry := range tests {
		_, key, value := parseVariable(entry.line)
		if key != entry.key || value != entry.value {
			t.Fatalf("Doesn't match (%s - %s) (%s - %s)", key, value, entry.key, entry.value)
		}
	}
}

func TestTrim(t *testing.T) {

	var tests = []struct {
		line   string
		result string
	}{
		{"# Some Comment", ""},
		{"   # Some Comment", "   "},
		{"blah = key # Some Comment", "blah = key "},
		{"blah = key # Some Comment # more comment", "blah = key "},
	}

	for _, x := range tests {
		res := trimComment(x.line)
		if res != x.result {
			t.Fatalf("test doesn't match - %s %s", res, x.result)
		}
	}
}
