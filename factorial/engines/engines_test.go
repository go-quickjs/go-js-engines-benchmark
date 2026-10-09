package engines

import (
	"os/exec"
	"testing"
)

func TestEngines(t *testing.T) {
	for _, engine := range Engines() {
		t.Run(engine.Name(), func(t *testing.T) {
			if _, isNode := engine.(*Node); isNode {
				if _, err := exec.LookPath("node"); err != nil {
					t.Skip("Node.js is not installed")
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
			if err := engine.Run(`
				function factorial(n) {
					return n <= 1 ? 1 : n * factorial(n - 1);
				}
				for (var i = 0; i < 100; i++) {
					if (factorial(10) !== 3628800) throw new Error("incorrect factorial");
				}
			`); err != nil {
				t.Fatal(err)
			}
			if err := engine.Run(`throw new Error("expected failure")`); err == nil {
				t.Fatal("expected JavaScript error")
			}
			if err := engine.Close(); err != nil {
				t.Fatal(err)
			}
			if err := engine.Init(); err != nil {
				t.Fatal(err)
			}
			if err := engine.Run(`if (typeof factorial !== "undefined") throw new Error("stale runtime")`); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestNodeMissingExecutable(t *testing.T) {
	t.Setenv("PATH", t.TempDir())
	engine := &Node{}
	if err := engine.Init(); err == nil {
		t.Fatal("expected error when Node.js is not installed")
	}
	if err := engine.Close(); err != nil {
		t.Fatal(err)
	}
}
