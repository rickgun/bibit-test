package testdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

var basepath string

type FuncCaller struct {
	IsCalled bool
	Input    []interface{}
	Output   []interface{}
}

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	fmt.Printf(currentFile)
	basepath = filepath.Dir(currentFile)
}

func path(relPath string) string {
	if filepath.IsAbs(relPath) {
		return relPath
	}

	return filepath.Join(basepath, relPath)
}

func GetGolden(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := ioutil.ReadFile(path(filename + ".golden"))
	if err != nil {
		t.Fatal(t)
	}

	return b
}

func GoldenJSONUnmarshal(t *testing.T, filename string, input interface{}) {
	_bytes := GetGolden(t, filename)

	err := json.Unmarshal(_bytes, &input)
	if err != nil {
		t.Fatal(t)
	}
}
