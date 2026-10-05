# Go JavaScript Engine Benchmarks

Performance benchmarks for four JavaScript engines in Go.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go JavaScript engine
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly
- **[GoQuickJS](https://github.com/go-quickjs/go-quickjs)**: Pure Go reimplementation of QuickJS-NG

Both benchmark modules use the following engine versions:

| Engine | Version |
| --- | --- |
| Goja | `v0.0.0-20261004200024-481fdb442bb4` |
| ModerncQuickJS | `v0.25.0` |
| QJS | `v0.0.6` |
| GoQuickJS | `v0.21.1-0.20261005231136-9b660e93edba` |

The results below are historical measurements from the earlier three-engine setup. They do not represent these updated versions or include GoQuickJS. Run both benchmarks to generate current four-engine comparisons.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS |
| --- | --- | --- | --- |
| 1 | 1.096s | 1.427s | 700.832ms |
| 2 | 1.067s | 1.451s | 689.655ms |
| 3 | 1.113s | 1.598s | 771.041ms |
| 4 | 1.085s | 1.469s | 709.371ms |
| 5 | 1.090s | 1.487s | 715.391ms |
| Average | 1.090s | 1.486s | **717.258ms** |
| Total | 5.450s | 7.432s | **3.586s** |
| Relative Time (lower is better) | 1.52x | 2.07x | 1.00x |

*Benchmarks run on AMD Ryzen 7 7840HS, 32GB RAM, Linux*

## V8v7 Benchmark Results

| Metric | GOJA | ModerncQuickJS | QJS |
| --- | --- | --- | --- |
| Richards | 353 | 221 | **448** |
| DeltaBlue | 448 | 260 | **490** |
| Crypto | 209 | 218 | **414** |
| RayTrace | 419 | 392 | **536** |
| EarleyBoyer | 803 | 625 | **873** |
| RegExp | **343** | 162 | 141 |
| Splay | 1361 | 939 | **1511** |
| NavierStokes | 336 | 417 | **611** |
| Score (version 7) | 452 | 343 | **522** |
| Duration (seconds) | 76.086s | 86.525s | **69.060s** |

*Benchmarks run on AMD Ryzen 7 7840HS, 32GB RAM, Linux*

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
| QJS | WASM linear memory | No |
| ModerncQuickJS | mmap allocations | No |

Goja and GoQuickJS use normal Go memory that shows up in `runtime.MemStats`. QJS and ModerncQuickJS use memory that Go cannot see. This makes comparisons based on Go memory statistics misleading.

Only execution time is compared across all four engines.

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
- QJS: MIT
- ModerncQuickJS: BSD-3-Clause
