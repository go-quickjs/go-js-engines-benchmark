# Go JavaScript Engine Benchmarks

Performance benchmarks for six Go-hosted JavaScript engines, including GoQuickJS with and without JIT, with normal and JITless Node.js runs as separate references.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go JavaScript engine
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly
- **[GoQuickJS](https://github.com/go-quickjs/go-quickjs)**: Pure Go reimplementation of QuickJS-NG
- **GoQuickJS JIT**: The same engine with its experimental native JIT enabled (`GoQuickJSJIT`), using a 32 MiB JIT memory limit
- **[Paserati](https://github.com/nooga/paserati)**: Pure Go TypeScript/JavaScript runtime with a bytecode VM; benchmarks run with type checking disabled
- **[MoeJS](https://github.com/MoeForge/moejs)**: Pure Go JavaScript runtime with a register bytecode interpreter; tested using its latest alpha release
- **[Node.js](https://nodejs.org/)**: V8 with JIT compilation enabled (`Node`), shown as a reference
- **Node.js JITless**: The same Node.js executable with `--jitless` (`NodeJitless`), shown as a reference

Both benchmark modules use the following engine versions:

| Engine | Version |
| --- | --- |
| Goja | `v0.0.0-20261107161106-d6eaa50db8bc` |
| ModerncQuickJS | `v0.27.0` |
| QJS | `v0.0.6` |
| GoQuickJS / GoQuickJSJIT | `v0.37.1` |
| Paserati | `v0.10.1` |
| MoeJS | `v0.1.0-alpha.11` |
| Node / NodeJitless | `v26.13.0` |

Results collected on November 8, 2026 using the engine versions listed above and Go 1.27.3. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, and MoeJS are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJSJIT, with a 32 MiB JIT memory limit, has the highest V8v7 score among the seven Go-hosted engine configurations in this run: 17870, compared with 5236 for the GoQuickJS interpreter and 2266 for MoeJS. Its factorial average is 55.230ms, compared with 217.095ms for the interpreter and 177.350ms for MoeJS. Enabling JIT gives a 3.41x higher V8v7 score and runs factorial 3.93x faster than the interpreter. Compilation and warmup are included. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 734.005ms | 570.669ms | 605.002ms | 217.652ms | 53.186ms | 661.490ms | 175.985ms |
| 2 | 707.670ms | 562.451ms | 615.388ms | 234.463ms | 54.814ms | 658.378ms | 176.476ms |
| 3 | 711.285ms | 562.678ms | 608.309ms | 208.080ms | 57.428ms | 657.964ms | 180.910ms |
| 4 | 720.817ms | 571.855ms | 616.494ms | 212.714ms | 59.149ms | 650.282ms | 177.918ms |
| 5 | 717.714ms | 566.231ms | 609.079ms | 212.567ms | 51.571ms | 662.434ms | 175.459ms |
| Average | 718.298ms | 566.777ms | 610.854ms | 217.095ms | **55.230ms** | 658.109ms | 177.350ms |
| Total | 3.591s | 2.834s | 3.054s | 1.085s | **276.150ms** | 3.291s | 886.750ms |
| Relative Time (lower is better) | 13.01x | 10.26x | 11.06x | 3.93x | 1.00x | 11.92x | 3.21x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 19.895ms | 157.701ms |
| 2 | 21.582ms | 166.363ms |
| 3 | 20.831ms | 164.393ms |
| 4 | 20.608ms | 165.968ms |
| 5 | 20.228ms | 170.021ms |
| Average | **20.629ms** | 164.889ms |
| Total | **103.146ms** | 824.448ms |
| Relative Time (lower is better) | 1.00x | 7.99x |

*Node v26.13.0, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Richards | 519 | 547 | 577 | 2549 | **13831** | 483 | 1206 |
| DeltaBlue | 641 | 554 | 636 | 2645 | **13521** | 583 | 1521 |
| Crypto | 331 | 566 | 458 | 3244 | **20060** | 646 | 1641 |
| RayTrace | 814 | 1217 | 1025 | 6117 | **31608** | 996 | 2418 |
| EarleyBoyer | 1414 | 1490 | 1540 | 7308 | **12766** | 1900 | 5297 |
| RegExp | 554 | 531 | 267 | 8478 | **9155** | 1127 | 1115 |
| Splay | 2420 | 3113 | 2697 | 11808 | **19791** | 3452 | 5990 |
| NavierStokes | 540 | 1524 | 730 | 5776 | **37915** | 1277 | 2696 |
| Score (version 7) | 742 | 970 | 781 | 5236 | **17870** | 1069 | 2266 |
| Duration (seconds) | 52.330s | 41.059s | 51.259s | 22.258s | **20.112s** | 38.113s | 26.922s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **80323** | 2821 |
| DeltaBlue | **223155** | 3134 |
| Crypto | **100049** | 2831 |
| RayTrace | **223107** | 9084 |
| EarleyBoyer | **178356** | 13645 |
| RegExp | **24754** | 7448 |
| Splay | **39559** | 14571 |
| NavierStokes | **80056** | 4149 |
| Score (version 7) | **92997** | 5863 |
| Duration (seconds) | **20.049s** | 22.232s |

*Node v26.13.0, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

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
| GoQuickJSJIT | Go heap plus executable memory | Partly |
| Paserati | Go heap and stack | Yes |
| MoeJS | Go heap and stack | Yes |
| QJS | WASM linear memory | No |
| ModerncQuickJS | mmap allocations | No |
| Node / NodeJitless | V8 heap in a separate process | No |

Goja, GoQuickJS, Paserati, and MoeJS use normal Go memory that shows up in `runtime.MemStats`. GoQuickJSJIT also allocates executable memory outside the Go heap. QJS and ModerncQuickJS use memory that Go cannot see, and Node runs in a separate process. This makes comparisons based on Go memory statistics misleading.

Only execution time is measured for the seven Go-hosted engine configurations and the two Node reference configurations.

## How to Run

Requires Go 1.27.1 or newer, as required by Paserati `v0.10.1`.
Also requires `node` on `PATH`, with support for `--jitless` and `--no-jitless`.
Each command runs both Node configurations and prints their results separately.
Use Node `v26.13.0` to reproduce the recorded reference configuration.

The commands below include GoQuickJSJIT using `-tags quickjs_jit`. Its native tier supports `linux/amd64`, `windows/amd64`, and `darwin/arm64`. Builds without that tag, or on unsupported platforms, omit the GoQuickJSJIT configuration. `GoQuickJS` always runs with JIT disabled.

Clone the repository:

```bash
git clone http://github.com/ngocphuongnb/go-js-engines-benchmark.git
cd go-js-engines-benchmark
```

Run the factorial benchmark:

```bash
cd factorial
go run -tags quickjs_jit .
```

Run the V8v7 benchmark:

```bash
cd arewefastyet-v8v7
go run -tags quickjs_jit .
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

Reimplements QuickJS-NG in pure Go and needs neither cgo nor WebAssembly. Both benchmarks create its runtime before measuring execution time. `GoQuickJS` uses the interpreter with default runtime options. `GoQuickJSJIT` uses `WithJIT()` and `WithJITMemoryLimit(32 * 1024 * 1024)` in a build with `-tags quickjs_jit`. JIT compilation runs synchronously with the default promotion thresholds; compilation and warmup are included in the measured interval. Unsupported functions or denied executable-memory allocations fall back to the interpreter. The 32 MiB limit covers memory held by the JIT, not the JavaScript heap or compiler working memory.

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
