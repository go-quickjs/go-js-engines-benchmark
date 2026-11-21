# Go JavaScript Engine Benchmarks

Performance benchmarks for six Go-hosted JavaScript engines, including GoQuickJS with and without JIT, with normal and JITless Node.js runs as separate references.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go JavaScript engine
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly
- **[GoQuickJS](https://github.com/go-quickjs/go-quickjs)**: Pure Go reimplementation of QuickJS-NG
- **GoQuickJS JIT**: The same engine with its experimental native JIT enabled (`GoQuickJSJIT`), using a 32 MiB JIT memory limit
- **[Paserati](https://github.com/nooga/paserati)**: Pure Go TypeScript/JavaScript runtime with a bytecode VM; benchmarks run with type checking disabled
- **[MoeJS](https://github.com/MoeForge/moejs)**: Pure Go JavaScript runtime with a register bytecode interpreter
- **[Node.js](https://nodejs.org/)**: V8 with JIT compilation enabled (`Node`), shown as a reference
- **Node.js JITless**: The same Node.js executable with `--jitless` (`NodeJitless`), shown as a reference

Both benchmark modules use the following engine versions:

| Engine | Version |
| --- | --- |
| Goja | `v0.0.0-20261118204651-8dfece3d1e1d` |
| ModerncQuickJS | `v0.27.0` |
| QJS | `v0.0.6` |
| GoQuickJS / GoQuickJSJIT | `v0.42.0` |
| Paserati | `v0.10.1` |
| MoeJS | `v0.1.1` |
| Node / NodeJitless | `v26.13.1` |

Engines use their latest stable releases; Goja has no tagged releases and uses the latest Go module pseudo-version.

Results collected on November 20, 2026 using the engine versions listed above and Go 1.27.3. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, and MoeJS are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJSJIT, with a 32 MiB JIT memory limit, has the highest V8v7 score among the seven Go-hosted engine configurations in this run: 19082, compared with 7327 for the GoQuickJS interpreter and 5695 for MoeJS. Its factorial average is 61.119ms, compared with 153.012ms for MoeJS and 172.574ms for the GoQuickJS interpreter. Among the Go-hosted interpreter configurations, GoQuickJS has the highest V8v7 score, while MoeJS has the fastest factorial average. Enabling JIT gives GoQuickJS a 2.60x higher V8v7 score and runs factorial 2.82x faster than its interpreter. Compilation and warmup are included. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 696.683ms | 543.726ms | 623.641ms | 169.577ms | 64.297ms | 634.480ms | 143.102ms |
| 2 | 707.390ms | 572.865ms | 646.872ms | 172.030ms | 65.169ms | 632.673ms | 145.965ms |
| 3 | 714.161ms | 550.993ms | 606.293ms | 178.235ms | 61.946ms | 664.199ms | 138.094ms |
| 4 | 697.074ms | 549.126ms | 617.827ms | 174.892ms | 62.290ms | 646.417ms | 140.718ms |
| 5 | 707.963ms | 575.069ms | 607.545ms | 168.134ms | 51.894ms | 672.869ms | 197.184ms |
| Average | 704.654ms | 558.356ms | 620.436ms | 172.574ms | **61.119ms** | 650.127ms | 153.012ms |
| Total | 3.523s | 2.792s | 3.102s | 862.870ms | **305.598ms** | 3.251s | 765.064ms |
| Relative Time (lower is better) | 11.53x | 9.14x | 10.15x | 2.82x | 1.00x | 10.64x | 2.50x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 26.729ms | 162.561ms |
| 2 | 21.147ms | 160.964ms |
| 3 | 21.404ms | 164.801ms |
| 4 | 21.187ms | 172.311ms |
| 5 | 18.081ms | 159.685ms |
| Average | **21.710ms** | 164.064ms |
| Total | **108.550ms** | 820.323ms |
| Relative Time (lower is better) | 1.00x | 7.56x |

*Node v26.13.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Richards | 501 | 565 | 571 | 3319 | **15272** | 494 | 2835 |
| DeltaBlue | 623 | 572 | 628 | 3587 | **13891** | 584 | 2992 |
| Crypto | 337 | 578 | 445 | 5513 | **20870** | 649 | 2569 |
| RayTrace | 777 | 1161 | 1007 | 9305 | **32634** | 1054 | 7495 |
| EarleyBoyer | 1434 | 1416 | 1525 | 9707 | **12852** | 1877 | 10928 |
| RegExp | 572 | 530 | 257 | 8027 | **10344** | 1135 | 5733 |
| Splay | 2518 | 3056 | 2532 | 17056 | **23771** | 3403 | 15540 |
| NavierStokes | 562 | 1544 | 726 | 10240 | **38507** | 1259 | 6961 |
| Score (version 7) | 745 | 968 | 764 | 7327 | **19082** | 1076 | 5695 |
| Duration (seconds) | 50.791s | 41.030s | 52.273s | 20.165s | **20.085s** | 37.900s | 21.222s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **84008** | 2958 |
| DeltaBlue | **226368** | 3349 |
| Crypto | **102046** | 2893 |
| RayTrace | **229767** | 9415 |
| EarleyBoyer | **189990** | 14047 |
| RegExp | **26392** | 7182 |
| Splay | **94073** | 16274 |
| NavierStokes | **85033** | 4319 |
| Score (version 7) | **107538** | 6099 |
| Duration (seconds) | **20.049s** | 22.207s |

*Node v26.13.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

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
Use Node `v26.13.1` to reproduce the recorded reference configuration.

The commands below include GoQuickJSJIT using `-tags quickjs_jit`. Its native tier supports `linux/amd64`, `linux/arm64`, `windows/amd64`, and `darwin/arm64`. Builds without that tag, or on unsupported platforms, omit the GoQuickJSJIT configuration. `GoQuickJS` always runs with JIT disabled.

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
