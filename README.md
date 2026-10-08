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
| Goja | `v0.0.0-20261007200356-e2ea74d3d210` |
| ModerncQuickJS | `v0.25.0` |
| QJS | `v0.0.6` |
| GoQuickJS | `v0.24.3` |
| Paserati | `v0.9.13` |

Results collected on October 8, 2026 using the engine versions listed above. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati |
| --- | --- | --- | --- | --- | --- |
| 1 | 757.357ms | 699.730ms | 616.800ms | 215.146ms | 532.040ms |
| 2 | 725.276ms | 706.937ms | 615.922ms | 214.989ms | 530.160ms |
| 3 | 740.876ms | 695.378ms | 617.475ms | 227.441ms | 531.224ms |
| 4 | 750.306ms | 698.285ms | 627.344ms | 215.264ms | 538.193ms |
| 5 | 727.457ms | 694.401ms | 617.112ms | 215.013ms | 534.047ms |
| Average | 740.254ms | 698.946ms | 618.930ms | **217.571ms** | 533.133ms |
| Total | 3.701s | 3.495s | 3.095s | **1.088s** | 2.666s |
| Relative Time (lower is better) | 3.40x | 3.21x | 2.84x | 1.00x | 2.45x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.0.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati |
| --- | --- | --- | --- | --- | --- |
| Richards | 485 | 471 | 565 | **1959** | 473 |
| DeltaBlue | 592 | 495 | 608 | **2371** | 614 |
| Crypto | 319 | 413 | 442 | **2956** | 623 |
| RayTrace | 739 | 1033 | 974 | **4826** | 1080 |
| EarleyBoyer | 1315 | 1244 | 1467 | **5486** | 1924 |
| RegExp | 547 | 332 | 248 | **3965** | 1055 |
| Splay | 2327 | 2608 | 2481 | **7644** | 3366 |
| NavierStokes | 531 | 1028 | 700 | **6078** | 1214 |
| Score (version 7) | 704 | 759 | 743 | **4011** | 1062 |
| Duration (seconds) | 53.940s | 50.370s | 53.124s | **22.350s** | 38.207s |

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
