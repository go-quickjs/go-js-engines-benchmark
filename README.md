# Go JavaScript Engine Benchmarks

Performance benchmarks for seven Go-hosted JavaScript engines, including GoQuickJS with and without JIT, with normal and JITless Node.js runs as separate references.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go JavaScript engine
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly
- **[GoQuickJS](https://github.com/go-quickjs/go-quickjs)**: Pure Go reimplementation of QuickJS-NG
- **GoQuickJS JIT**: The same engine with its experimental native JIT enabled (`GoQuickJSJIT`), using a 32 MiB JIT memory limit
- **[Paserati](https://github.com/nooga/paserati)**: Pure Go TypeScript/JavaScript runtime with a bytecode VM; benchmarks run with type checking disabled
- **[MoeJS](https://github.com/MoeForge/moejs)**: Pure Go JavaScript runtime with a register bytecode interpreter
- **[Ekuma](https://github.com/hajimehoshi/ekuma)**: Pure Go JavaScript engine with a register bytecode interpreter; an early-stage project with an unstable API
- **[Node.js](https://nodejs.org/)**: V8 with JIT compilation enabled (`Node`), shown as a reference
- **Node.js JITless**: The same Node.js executable with `--jitless` (`NodeJitless`), shown as a reference

Both benchmark modules use the following engine versions:

| Engine | Version |
| --- | --- |
| Goja | `v0.0.0-20261125210501-5e92a7e24d0c` |
| ModerncQuickJS | `v0.27.0` |
| QJS | `v0.0.6` |
| GoQuickJS / GoQuickJSJIT | `v0.45.1` |
| Paserati | `v0.10.1` |
| MoeJS | `v0.1.1` |
| Ekuma | `v0.0.0-20261125190603-3adcefe737df` |
| Node / NodeJitless | `v26.13.1` |

Engines use their latest stable releases where available. Goja and Ekuma have no tagged releases and use their latest Go module pseudo-versions.

Results collected on November 25, 2026 using the engine versions listed above and Go 1.27.3. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, MoeJS, and Ekuma are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJSJIT, with a 32 MiB JIT memory limit, has the highest V8v7 score among the eight Go-hosted engine configurations in this run: 20395, compared with 8103 for the GoQuickJS interpreter, 5944 for Ekuma, and 5872 for MoeJS. Its factorial average is 63.302ms, compared with 137.437ms for MoeJS, 169.869ms for the GoQuickJS interpreter, and 172.614ms for Ekuma. Among the Go-hosted interpreter configurations, GoQuickJS has the highest V8v7 score, while MoeJS has the fastest factorial average. Enabling JIT gives GoQuickJS a 2.52x higher V8v7 score and runs factorial 2.68x faster than its interpreter. Compilation and warmup are included. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS | Ekuma |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 684.647ms | 542.991ms | 578.605ms | 170.089ms | 65.810ms | 627.943ms | 135.950ms | 169.162ms |
| 2 | 680.513ms | 545.009ms | 579.161ms | 172.666ms | 63.658ms | 625.686ms | 132.076ms | 174.992ms |
| 3 | 695.435ms | 564.076ms | 585.658ms | 166.348ms | 64.199ms | 632.576ms | 141.142ms | 171.599ms |
| 4 | 695.562ms | 538.893ms | 572.553ms | 174.781ms | 60.884ms | 624.508ms | 138.288ms | 176.613ms |
| 5 | 703.841ms | 546.992ms | 585.333ms | 165.459ms | 61.960ms | 633.362ms | 139.728ms | 170.706ms |
| Average | 692.000ms | 547.592ms | 580.262ms | 169.869ms | **63.302ms** | 628.815ms | 137.437ms | 172.614ms |
| Total | 3.460s | 2.738s | 2.901s | 849.345ms | **316.513ms** | 3.144s | 687.186ms | 863.073ms |
| Relative Time (lower is better) | 10.93x | 8.65x | 9.17x | 2.68x | 1.00x | 9.93x | 2.17x | 2.73x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 24.917ms | 167.077ms |
| 2 | 19.902ms | 165.489ms |
| 3 | 21.182ms | 160.410ms |
| 4 | 20.790ms | 161.148ms |
| 5 | 20.607ms | 166.428ms |
| Average | **21.480ms** | 164.110ms |
| Total | **107.400ms** | 820.554ms |
| Relative Time (lower is better) | 1.00x | 7.64x |

*Node v26.13.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS | Ekuma |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Richards | 526 | 589 | 599 | 3424 | **16101** | 476 | 2910 | 3371 |
| DeltaBlue | 658 | 592 | 666 | 3765 | **15954** | 569 | 3081 | 3670 |
| Crypto | 353 | 609 | 477 | 5707 | **21692** | 653 | 2615 | 3982 |
| RayTrace | 852 | 1265 | 1079 | 11214 | **35371** | 1020 | 7569 | 8411 |
| EarleyBoyer | 1537 | 1558 | 1596 | 12038 | **15081** | 1908 | 11279 | 9382 |
| RegExp | 613 | 537 | 274 | 9383 | **10856** | 1182 | 6182 | 4541 |
| Splay | 2809 | 3265 | 2784 | 18295 | **23282** | 3590 | 15972 | 11392 |
| NavierStokes | 587 | 1581 | 755 | 10905 | **39840** | 1326 | 7157 | 7749 |
| Score (version 7) | 796 | 1018 | 812 | 8103 | **20395** | 1086 | 5872 | 5944 |
| Duration (seconds) | 49.839s | 40.468s | 51.437s | 20.134s | **20.095s** | 37.660s | 21.211s | 20.217s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **85569** | 3061 |
| DeltaBlue | **241086** | 3402 |
| Crypto | **106357** | 2970 |
| RayTrace | **236056** | 9684 |
| EarleyBoyer | **190804** | 14433 |
| RegExp | **27147** | 7834 |
| Splay | **100641** | 16787 |
| NavierStokes | **85627** | 4501 |
| Score (version 7) | **111044** | 6324 |
| Duration (seconds) | **20.044s** | 22.193s |

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
| Ekuma | Go heap and stack | Yes |
| QJS | WASM linear memory | No |
| ModerncQuickJS | mmap allocations | No |
| Node / NodeJitless | V8 heap in a separate process | No |

Goja, GoQuickJS, Paserati, MoeJS, and Ekuma use normal Go memory that shows up in `runtime.MemStats`. GoQuickJSJIT also allocates executable memory outside the Go heap. QJS and ModerncQuickJS use memory that Go cannot see, and Node runs in a separate process. This makes comparisons based on Go memory statistics misleading.

Only execution time is measured for the eight Go-hosted engine configurations and the two Node reference configurations.

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

### Ekuma

Parses JavaScript and compiles it to register bytecode in pure Go, without a JIT. Each benchmark initializes the VM's realm and built-in objects before timing starts; parsing, compilation, and execution are included in the measured interval. V8v7 scripts share one VM and global environment, with native `load` and `print` callbacks that propagate JavaScript exceptions. The optional `intl` package is not imported because these workloads do not use it. Memory is managed by Go's garbage collector.

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
- Ekuma: 0BSD
- QJS: MIT
- ModerncQuickJS: BSD-3-Clause
- Node.js: MIT, with separately licensed bundled dependencies
