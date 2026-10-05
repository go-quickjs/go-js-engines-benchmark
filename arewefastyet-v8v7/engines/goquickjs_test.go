package engines

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGoQuickJSRun(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := os.Mkdir("v8-v7", 0755); err != nil {
		t.Fatal(err)
	}
	scripts := map[string]string{
		"loaded.js":  `function factorial(n) { return n <= 1 ? 1 : n * factorial(n - 1); }`,
		"run.js":     `load("loaded.js"); print("Factorial: " + factorial(10)); print("values", 42, true);`,
		"error.js":   `throw new Error("expected failure");`,
		"syntax.js":  `function (`,
		"missing.js": `load("does-not-exist.js");`,
		"badload.js": `load("syntax.js");`,
	}
	for name, script := range scripts {
		if err := os.WriteFile(filepath.Join("v8-v7", name), []byte(script), 0644); err != nil {
			t.Fatal(err)
		}
	}

	engine := &GoQuickJS{}
	if err := engine.Init(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := engine.Close(); err != nil {
			t.Error(err)
		}
	})
	output, err := engine.Run("v8-v7/run.js")
	if err != nil {
		t.Fatal(err)
	}
	want := [][]string{{"Factorial: 3628800"}, {"values", " ", "42", " ", "true"}}
	if !reflect.DeepEqual(output, want) {
		t.Fatalf("output = %v, want %v", output, want)
	}
	for _, name := range []string{"error.js", "syntax.js", "missing.js", "badload.js", "does-not-exist.js"} {
		t.Run(name, func(t *testing.T) {
			if _, err := engine.Run(filepath.Join("v8-v7", name)); err == nil {
				t.Fatal("expected error")
			}
		})
	}
	if err := engine.Close(); err != nil {
		t.Fatal(err)
	}
	if err := engine.Close(); err != nil {
		t.Fatal(err)
	}
	if err := engine.Init(); err != nil {
		t.Fatal(err)
	}
	output, err = engine.Run("v8-v7/run.js")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(output, want) {
		t.Fatalf("output after reinitialization = %v, want %v", output, want)
	}
}
