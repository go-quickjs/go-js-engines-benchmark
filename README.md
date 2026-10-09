# Go JavaScript Engine Benchmarks

Performance benchmarks for five Go-hosted JavaScript engines, with normal and JITless Node.js runs as separate references.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go JavaScript engine
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly
- **[GoQuickJS](https://github.com/go-quickjs/go-quickjs)**: Pure Go reimplementation of QuickJS-NG
- **[Paserati](https://github.com/nooga/paserati)**: Pure Go TypeScript/JavaScript runtime with a bytecode VM; benchmarks run with type checking disabled
- **[Node.js](https://nodejs.org/)**: V8 with JIT compilation enabled (`Node`), shown as a reference
- **Node.js JITless**: The same Node.js executable with `--jitless` (`NodeJitless`), shown as a reference

Both benchmark modules use the following engine versions:

| Engine | Version |
| --- | --- |
| Goja | `v0.0.0-20261007200356-e2ea74d3d210` |
| ModerncQuickJS | `v0.25.0` |
| QJS | `v0.0.6` |
| GoQuickJS | `v0.24.3` |
| Paserati | `v0.9.13` |
| Node / NodeJitless | `v26.8.1` |

Results collected on October 8, 2026 using the engine versions listed above. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, and Paserati are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati |
| --- | --- | --- | --- | --- | --- |
| 1 | 723.468ms | 676.680ms | 618.591ms | 228.784ms | 533.396ms |
| 2 | 698.477ms | 673.048ms | 615.720ms | 210.347ms | 531.844ms |
| 3 | 730.755ms | 672.215ms | 604.424ms | 210.459ms | 532.983ms |
| 4 | 775.860ms | 694.209ms | 613.146ms | 220.651ms | 530.048ms |
| 5 | 748.697ms | 688.465ms | 611.407ms | 230.680ms | 524.646ms |
| Average | 735.452ms | 680.923ms | 612.657ms | **220.184ms** | 530.583ms |
| Total | 3.677s | 3.405s | 3.063s | **1.101s** | 2.653s |
| Relative Time (lower is better) | 3.34x | 3.09x | 2.78x | 1.00x | 2.41x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.0.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 21.239ms | 169.119ms |
| 2 | 20.880ms | 168.414ms |
| 3 | 20.977ms | 167.212ms |
| 4 | 24.341ms | 167.054ms |
| 5 | 18.744ms | 163.946ms |
| Average | **21.236ms** | 167.149ms |
| Total | **106.183ms** | 835.747ms |
| Relative Time (lower is better) | 1.00x | 7.87x |

*Node v26.8.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati |
| --- | --- | --- | --- | --- | --- |
| Richards | 482 | 474 | 573 | **1980** | 481 |
| DeltaBlue | 568 | 484 | 632 | **2395** | 626 |
| Crypto | 317 | 407 | 431 | **2958** | 628 |
| RayTrace | 785 | 1019 | 957 | **4758** | 1065 |
| EarleyBoyer | 1434 | 1226 | 1437 | **5777** | 1948 |
| RegExp | 573 | 348 | 251 | **4051** | 1054 |
| Splay | 2212 | 2741 | 2600 | **7505** | 3333 |
| NavierStokes | 531 | 1047 | 716 | **6060** | 1232 |
| Score (version 7) | 712 | 765 | 750 | **4041** | 1068 |
| Duration (seconds) | 52.740s | 49.056s | 52.372s | **22.364s** | 38.329s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.0.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **80436** | 2849 |
| DeltaBlue | **217191** | 3098 |
| Crypto | **99938** | 2737 |
| RayTrace | **229249** | 8745 |
| EarleyBoyer | **180628** | 13091 |
| RegExp | **25366** | 6449 |
| Splay | **43255** | 14619 |
| NavierStokes | **80284** | 4197 |
| Score (version 7) | **94516** | 5687 |
| Duration (seconds) | **20.044s** | 22.253s |

*Node v26.8.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

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
| Node / NodeJitless | V8 heap in a separate process | No |

Goja, GoQuickJS, and Paserati use normal Go memory that shows up in `runtime.MemStats`. QJS and ModerncQuickJS use memory that Go cannot see, and Node runs in a separate process. This makes comparisons based on Go memory statistics misleading.

Only execution time is measured for the five Go-hosted engines and the two Node reference configurations.

## How to Run

Requires Go 1.26 or newer, as required by ModerncQuickJS `v0.25.0`.
Also requires `node` on `PATH`, with support for `--jitless` and `--no-jitless`.
Each command runs both Node configurations and prints their results separately.

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

### Node.js References

Starts a Node.js worker in `Init` and waits until it is ready before timing begins. `Node` explicitly enables JIT compilation with `--no-jitless`; `NodeJitless` uses `--jitless`. Scripts run with `vm.runInThisContext` in Node's native global realm, avoiding the proxy overhead of a sandboxed context. The `load` and `print` callbacks share that realm, and worker variables stay inside a closure to avoid collisions with benchmark globals.

The measured interval includes script compilation, execution, and communication with the worker over JSON pipes. Process startup and shutdown are excluded. These external-process measurements provide context and do not affect the Go-hosted engine rankings.

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
5. Excludes runtime startup (WASM compilation and the Node readiness handshake happen before timing starts)
6. Reports Node configurations separately; their timings include communication overhead

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
- Node.js: MIT, with separately licensed bundled dependencies
