package engines

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGoQuickJSRun(t *testing.T) {
	testScriptEngine(t, &GoQuickJS{})
}

func TestPaseratiRun(t *testing.T) {
	testScriptEngine(t, &Paserati{})
}

func TestNodeRun(t *testing.T) {
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("Node.js is not installed")
	}
	for _, engine := range []*Node{{}, {Jitless: true}} {
		t.Run(engine.Name(), func(t *testing.T) {
			testScriptEngine(t, engine)
		})
	}
}

func testScriptEngine(t *testing.T, engine JSEngine) {
	t.Helper()
	t.Chdir(t.TempDir())
	if err := os.Mkdir("v8-v7", 0755); err != nil {
		t.Fatal(err)
	}
	scripts := map[string]string{
		"loaded.js":    `function factorial(n) { return n <= 1 ? 1 : n * factorial(n - 1); }`,
		"entry.js":     `load("loaded.js"); print("Factorial: " + factorial(10)); print("values", 42, true);`,
		"error.js":     `throw new Error("expected failure");`,
		"syntax.js":    `function (`,
		"missing.js":   `load("does-not-exist.js");`,
		"badload.js":   `load("syntax.js");`,
		"throwload.js": `load("error.js");`,
		"noarg.js":     `load();`,
		"caught.js":    `try { load("does-not-exist.js"); } catch (e) { print("caught load"); }`,
	}
	for name, script := range scripts {
		if err := os.WriteFile(filepath.Join("v8-v7", name), []byte(script), 0644); err != nil {
			t.Fatal(err)
		}
	}

	if err := engine.Init(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := engine.Close(); err != nil {
			t.Error(err)
		}
	})
	output, err := engine.Run("v8-v7/entry.js")
	if err != nil {
		t.Fatal(err)
	}
	want := [][]string{{"Factorial: 3628800"}, {"values", " ", "42", " ", "true"}}
	if !reflect.DeepEqual(output, want) {
		t.Fatalf("output = %v, want %v", output, want)
	}
	for _, name := range []string{"error.js", "syntax.js", "missing.js", "badload.js", "throwload.js", "noarg.js", "does-not-exist.js"} {
		t.Run(name, func(t *testing.T) {
			if err := engine.Close(); err != nil {
				t.Fatal(err)
			}
			if err := engine.Init(); err != nil {
				t.Fatal(err)
			}
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
	output, err = engine.Run("v8-v7/entry.js")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(output, want) {
		t.Fatalf("output after reinitialization = %v, want %v", output, want)
	}
	output, err = engine.Run("v8-v7/caught.js")
	if err != nil {
		t.Fatal(err)
	}
	want = append(want, []string{"caught load"})
	if !reflect.DeepEqual(output, want) {
		t.Fatalf("output after caught load error = %v, want %v", output, want)
	}
}
