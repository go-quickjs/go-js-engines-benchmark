# Go JavaScript Engine Benchmarks

Performance benchmarks for five JavaScript engines in Go.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go JavaScript engine
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly
- **[GoQuickJS](https://github.com/go-quickjs/go-quickjs)**: Pure Go reimplementation of QuickJS-NG
- **[Paserati](https://github.com/nooga/paserati)**: Pure Go TypeScript/JavaScript runtime with a bytecode VM; benchmarks run with type checking disabled

Both benchmark modules use the following engine versions:

| Engine | Version |
| --- | --- |
| Goja | `v0.0.0-20261004200024-481fdb442bb4` |
| ModerncQuickJS | `v0.25.0` |
| QJS | `v0.0.6` |
| GoQuickJS | `v0.21.1-0.20261005231136-9b660e93edba` |
| Paserati | `v0.9.13` |

Results collected on October 5, 2026 using the engine versions listed above. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati |
| --- | --- | --- | --- | --- | --- |
| 1 | 695.428ms | 658.038ms | 576.685ms | 206.547ms | 519.791ms |
| 2 | 688.229ms | 644.812ms | 603.838ms | 217.724ms | 521.393ms |
| 3 | 713.952ms | 646.720ms | 600.329ms | 210.486ms | 507.021ms |
| 4 | 710.495ms | 659.025ms | 593.381ms | 225.509ms | 518.671ms |
| 5 | 709.531ms | 656.934ms | 585.931ms | 206.696ms | 512.245ms |
| Average | 703.527ms | 653.106ms | 592.033ms | **213.392ms** | 515.824ms |
| Total | 3.518s | 3.266s | 2.960s | **1.067s** | 2.579s |
| Relative Time (lower is better) | 3.30x | 3.06x | 2.77x | 1.00x | 2.42x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.0.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati |
| --- | --- | --- | --- | --- | --- |
| Richards | 503 | 484 | 563 | **1774** | 492 |
| DeltaBlue | 596 | 512 | 622 | **2246** | 649 |
| Crypto | 330 | 416 | 448 | **2902** | 649 |
| RayTrace | 820 | 1062 | 994 | **4525** | 1187 |
| EarleyBoyer | 1459 | 1304 | 1496 | **5557** | 2020 |
| RegExp | 605 | 351 | 255 | **4321** | 1125 |
| Splay | 2605 | 2760 | 2635 | **8231** | 3602 |
| NavierStokes | 545 | 1067 | 719 | **5311** | 1273 |
| Score (version 7) | 751 | 787 | 761 | **3912** | 1125 |
| Duration (seconds) | 51.188s | 48.447s | 52.086s | **22.380s** | 36.567s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.0.*

## What Gets Tested

### Factorial Benchmark

Calculates `factorial(10)` one million times using recursion. This tests how fast each engine handles computation and function calls. Each engine runs 5 times with alternating order to reduce bias.

### V8v7 Benchmark Suite

A standard JavaScript benchmark from the V8 project. Includes these tests:

- **Richards**: OS kernel simulation
- **DeltaBlue**: Constraint solving
- **Crypto**: Cryptographic operations
- **RayTrace**: 3D rendering
- **EarleyBoyer**: Parser and logic
- **RegExp**: Regular expressions
- **Splay**: Tree operations
- **NavierStokes**: Fluid dynamics

Each engine gets a score for each test and an overall score.

## Why Only Time Is Measured

Memory usage cannot be compared fairly between these engines. Here's why:

| Engine | Memory Type | Visible to Go |
|--------|-------------|---------------|
| Goja | Go heap and stack | Yes |
| GoQuickJS | Go heap and stack | Yes |
| Paserati | Go heap and stack | Yes |
| QJS | WASM linear memory | No |
| ModerncQuickJS | mmap allocations | No |

Goja, GoQuickJS, and Paserati use normal Go memory that shows up in `runtime.MemStats`. QJS and ModerncQuickJS use memory that Go cannot see. This makes comparisons based on Go memory statistics misleading.

Only execution time is compared across all five engines.

## How to Run

Requires Go 1.26 or newer, as required by ModerncQuickJS `v0.25.0`.

Clone the repository:

```bash
git clone http://github.com/ngocphuongnb/go-js-engines-benchmark.git
cd go-js-engines-benchmark
```

Run the factorial benchmark:

```bash
cd factorial
go run .
```

Run the V8v7 benchmark:

```bash
cd arewefastyet-v8v7
go run .
```

## How Each Engine Works

### Goja
```
┌─────────────────────────┐
│   Goja JS Engine        │
│  ┌──────────────────┐   │
│  │  Go Memory       │   │  <-  Go can see this
│  │  (Heap, Stack)   │   │
│  └──────────────────┘   │
└─────────────────────────┘
```

Written entirely in Go. Memory is managed by Go's garbage collector.

### GoQuickJS

Reimplements QuickJS-NG in pure Go. Memory is managed by Go's garbage collector, and it needs neither cgo nor WebAssembly. Both benchmarks create its runtime before measuring execution time.

### Paserati

Parses JavaScript/TypeScript and compiles it to bytecode for a register VM, implemented in pure Go. The benchmarks disable TypeScript type checking and create the runtime before measuring execution time. The V8 suite's `load` function compiles scripts in a shared global context. This integration adapts [PR #1](https://github.com/ngocphuongnb/go-js-engines-benchmark/pull/1) to the current Paserati release.

### ModerncQuickJS
```
┌─────────────────────────┐
│  QuickJS (ccgo)         │  <-  C code translated to Go
│  ┌──────────────────┐   │
│  │  mmap Memory     │   │  <-  Go cannot see this
│  │  (modernc.org/   │   │
│  │   memory)        │   │
│  └──────────────────┘   │
└─────────────────────────┘
```

Uses ccgo to translate C code to Go. Memory is allocated via mmap, which bypasses Go's memory tracking.

### QJS
```
┌─────────────────────────┐
│   QJS Wrapper (Go)      │
└─────────────────────────┘
           |
┌─────────────────────────┐
│   Wazero Runtime        │
│  ┌──────────────────┐   │
│  │ WASM Memory      │   │  <-  Go cannot see this
│  └──────────────────┘   │
└─────────────────────────┘
```

Runs QuickJS compiled to WebAssembly. Uses Wazero as the WebAssembly runtime. Memory is inside the WASM module.

## Keeping Tests Fair

The factorial benchmark uses several techniques to ensure fair comparison:

1. Runs garbage collection before each test
2. Waits 10ms before measuring, 100ms between tests
3. Alternates engine order every other iteration
4. Runs each engine 5 times and averages results
5. Measures only execution time, not startup (WASM compilation happens before timing starts)

## Contributing

Contributions are welcome. When reporting issues, include:

- Go version (`go version`)
- OS and CPU (`uname -a`)
- Full benchmark output
- System load and background processes

To add a benchmark:

1. Implement the `JSEngine` interface
2. Add it to `engines.Engines()`
3. Follow the existing structure

## License

Provided as-is for testing purposes. Engine licenses:
- Goja: MIT
- GoQuickJS: MIT
- Paserati: MIT
- QJS: MIT
- ModerncQuickJS: BSD-3-Clause
