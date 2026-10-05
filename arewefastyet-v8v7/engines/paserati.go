package engines

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nooga/paserati/pkg/builtins"
	"github.com/nooga/paserati/pkg/driver"
	"github.com/nooga/paserati/pkg/lexer"
	"github.com/nooga/paserati/pkg/parser"
	"github.com/nooga/paserati/pkg/source"
	"github.com/nooga/paserati/pkg/types"
	"github.com/nooga/paserati/pkg/vm"
)

type Paserati struct {
	p       *driver.Paserati
	output  [][]string
	baseDir string
}

func (p *Paserati) Name() string {
	return "Paserati"
}

func (p *Paserati) Init() error {
	baseDir, err := filepath.Abs("v8-v7")
	if err != nil {
		return fmt.Errorf("could not resolve v8-v7 directory: %w", err)
	}
	p.baseDir = baseDir
	p.output = nil
	initializers := append(builtins.GetStandardInitializers(), &paseratiBenchInitializer{engine: p})
	p.p = driver.NewPaseratiWithInitializersAndBaseDir(initializers, baseDir)
	p.p.SetIgnoreTypeErrors(true)
	p.p.SetSkipTypeCheck(true)
	return nil
}

func (p *Paserati) Run(inputFile string) ([][]string, error) {
	err := p.runScript(inputFile)
	return p.output, err
}

func (p *Paserati) runScript(inputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("could not read %s: %w", inputFile, err)
	}
	lx := lexer.NewLexerWithSource(source.FromFile(inputFile, string(content)))
	program, parseErrs := parser.NewParser(lx).ParseProgram()
	if len(parseErrs) > 0 {
		return fmt.Errorf("could not parse %s: %w", inputFile, parseErrs[0])
	}

	// Compile every load as a script in the same session so globals are shared.
	chunk, compileErrs := p.p.CompileProgramAsScript(program)
	if len(compileErrs) > 0 {
		return fmt.Errorf("could not compile %s: %w", inputFile, compileErrs[0])
	}
	if chunk == nil {
		return fmt.Errorf("could not compile %s: nil chunk", inputFile)
	}
	result, runtimeErrs := p.p.InterpretChunk(chunk)
	if len(runtimeErrs) > 0 {
		return fmt.Errorf("could not run %s: %w", inputFile, runtimeErrs[0])
	}
	// Paserati can return a thrown value without adding it to runtimeErrs.
	if p.p.GetVM().IsUnwinding() {
		return p.p.GetVM().NewExceptionError(result)
	}
	return nil
}

func (p *Paserati) Close() error {
	if p.p != nil {
		p.p.Cleanup()
		p.p = nil
	}
	p.output = nil
	p.baseDir = ""
	return nil
}

var _ JSEngine = (*Paserati)(nil)

type paseratiBenchInitializer struct {
	engine *Paserati
}

func (*paseratiBenchInitializer) Name() string {
	return "V8Bench"
}

func (*paseratiBenchInitializer) Priority() int {
	return 1000
}

func (*paseratiBenchInitializer) InitTypes(ctx *builtins.TypeContext) error {
	if err := ctx.DefineGlobal("load", types.NewSimpleFunction([]types.Type{types.String}, types.Void)); err != nil {
		return err
	}
	return ctx.DefineGlobal("print", types.NewVariadicFunction(nil, types.Void, types.Any))
}

func (v *paseratiBenchInitializer) InitRuntime(ctx *builtins.RuntimeContext) error {
	load := vm.NewNativeFunction(1, false, "load", func(args []vm.Value) (vm.Value, error) {
		if len(args) == 0 {
			return vm.Undefined, fmt.Errorf("load: missing filename argument")
		}
		return vm.Undefined, v.engine.runScript(filepath.Join(v.engine.baseDir, args[0].ToString()))
	})
	if err := ctx.DefineGlobal("load", load); err != nil {
		return err
	}
	print := vm.NewNativeFunction(0, true, "print", func(args []vm.Value) (vm.Value, error) {
		var line []string
		for i, arg := range args {
			if i > 0 {
				fmt.Print(" ")
				line = append(line, " ")
			}
			text := arg.ToString()
			fmt.Print(text)
			line = append(line, text)
		}
		fmt.Println()
		v.engine.output = append(v.engine.output, line)
		return vm.Undefined, nil
	})
	return ctx.DefineGlobal("print", print)
}
