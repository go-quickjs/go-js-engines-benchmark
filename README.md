# Go JavaScript Engine Benchmarks

Performance benchmarks for six Go-hosted JavaScript engines, with normal and JITless Node.js runs as separate references.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go JavaScript engine
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly
- **[GoQuickJS](https://github.com/go-quickjs/go-quickjs)**: Pure Go reimplementation of QuickJS-NG
- **[Paserati](https://github.com/nooga/paserati)**: Pure Go TypeScript/JavaScript runtime with a bytecode VM; benchmarks run with type checking disabled
- **[MoeJS](https://github.com/Calcium-Ion/moejs)**: Pure Go JavaScript runtime with a register bytecode interpreter; tested using its latest alpha release
- **[Node.js](https://nodejs.org/)**: V8 with JIT compilation enabled (`Node`), shown as a reference
- **Node.js JITless**: The same Node.js executable with `--jitless` (`NodeJitless`), shown as a reference

Both benchmark modules use the following engine versions:

| Engine | Version |
| --- | --- |
| Goja | `v0.0.0-20261024144015-2a9269b6b6de` |
| ModerncQuickJS | `v0.25.0` |
| QJS | `v0.0.6` |
| GoQuickJS | `v0.29.0` |
| Paserati | `v0.9.13` |
| MoeJS | `v0.1.0-alpha.6` |
| Node / NodeJitless | `v26.8.1` |

Results collected on October 25, 2026 using the engine versions listed above. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware. This run uses Go 1.27.3; the previous October 16 run used Go 1.27.0, so differences between those runs cannot be attributed solely to engine updates.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, and MoeJS are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJS v0.29.0 has the highest V8v7 score among the six Go-hosted engines in this run: 4379, compared with 1720 for MoeJS, 1046 for Paserati, and 722 for Goja. GoQuickJS and MoeJS have nearly identical factorial averages at 217.311ms and 217.401ms, respectively; the difference is less than 0.1% and should be read in light of run-to-run variation. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 717.626ms | 674.612ms | 632.688ms | 214.237ms | 533.056ms | 208.598ms |
| 2 | 722.347ms | 689.109ms | 611.138ms | 217.819ms | 539.622ms | 218.493ms |
| 3 | 742.273ms | 677.979ms | 622.914ms | 223.828ms | 538.164ms | 214.182ms |
| 4 | 710.065ms | 692.521ms | 609.894ms | 216.056ms | 525.546ms | 233.010ms |
| 5 | 738.422ms | 719.459ms | 608.984ms | 214.613ms | 524.415ms | 212.724ms |
| Average | 726.147ms | 690.736ms | 617.124ms | **217.311ms** | 532.161ms | 217.401ms |
| Total | 3.631s | 3.454s | 3.086s | **1.087s** | 2.661s | 1.087s |
| Relative Time (lower is better) | 3.34x | 3.18x | 2.84x | 1.00x | 2.45x | 1.00x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 21.740ms | 163.929ms |
| 2 | 20.810ms | 163.617ms |
| 3 | 21.491ms | 159.917ms |
| 4 | 21.125ms | 164.387ms |
| 5 | 20.852ms | 161.793ms |
| Average | **21.203ms** | 162.729ms |
| Total | **106.019ms** | 813.645ms |
| Relative Time (lower is better) | 1.00x | 7.67x |

*Node v26.8.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- |
| Richards | 484 | 477 | 569 | **2115** | 463 | 774 |
| DeltaBlue | 591 | 492 | 625 | **2510** | 588 | 1129 |
| Crypto | 325 | 415 | 448 | **2943** | 615 | 1430 |
| RayTrace | 781 | 1040 | 957 | **5517** | 1037 | 1923 |
| EarleyBoyer | 1360 | 1278 | 1483 | **6486** | 1878 | 2905 |
| RegExp | 581 | 340 | 254 | **4743** | 1078 | 1028 |
| Splay | 2420 | 2703 | 2456 | **8858** | 3390 | 4376 |
| NavierStokes | 534 | 1027 | 697 | **5759** | 1199 | 2435 |
| Score (version 7) | 722 | 769 | 748 | **4379** | 1046 | 1720 |
| Duration (seconds) | 52.451s | 49.612s | 52.669s | **22.314s** | 38.458s | 29.893s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **80979** | 2754 |
| DeltaBlue | **216722** | 3101 |
| Crypto | **97203** | 2786 |
| RayTrace | **228065** | 8844 |
| EarleyBoyer | **176202** | 13167 |
| RegExp | **23887** | 6513 |
| Splay | **41991** | 14652 |
| NavierStokes | **79166** | 4188 |
| Score (version 7) | **92680** | 5695 |
| Duration (seconds) | **20.055s** | 22.191s |

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
| MoeJS | Go heap and stack | Yes |
| QJS | WASM linear memory | No |
| ModerncQuickJS | mmap allocations | No |
| Node / NodeJitless | V8 heap in a separate process | No |

Goja, GoQuickJS, Paserati, and MoeJS use normal Go memory that shows up in `runtime.MemStats`. QJS and ModerncQuickJS use memory that Go cannot see, and Node runs in a separate process. This makes comparisons based on Go memory statistics misleading.

Only execution time is measured for the six Go-hosted engines and the two Node reference configurations.

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

### MoeJS

Parses classic JavaScript scripts and compiles them to register bytecode in pure Go. Each benchmark creates a runtime before timing starts; parsing, compilation, and execution are included in the measured interval. Loaded V8 scripts share one global environment, with native `load` and `print` callbacks. V8v7 enables `MutableIntrinsics` because the suite replaces `Math.random` with a deterministic generator; factorial uses the default options. Memory is managed by Go's garbage collector.

These results use the normal Go build settings, without MoeJS's optional plugin-workload PGO profile, matching the build approach used for the other engines.

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
- MoeJS: Apache-2.0
- QJS: MIT
- ModerncQuickJS: BSD-3-Clause
- Node.js: MIT, with separately licensed bundled dependencies
