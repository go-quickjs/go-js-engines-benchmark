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
| Goja | `v0.0.0-20261110204830-f493a2b600d6` |
| ModerncQuickJS | `v0.27.0` |
| QJS | `v0.0.6` |
| GoQuickJS / GoQuickJSJIT | `v0.38.0` |
| Paserati | `v0.10.1` |
| MoeJS | `v0.1.0` |
| Node / NodeJitless | `v26.13.1` |

Results collected on November 15, 2026 using the engine versions listed above and Go 1.27.3. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, and MoeJS are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJSJIT, with a 32 MiB JIT memory limit, has the highest V8v7 score among the seven Go-hosted engine configurations in this run: 19288, compared with 5614 for MoeJS and 5356 for the GoQuickJS interpreter. Its factorial average is 62.760ms, compared with 143.276ms for MoeJS and 213.610ms for the GoQuickJS interpreter. MoeJS has the best results among the Go-hosted interpreter configurations on both workloads. Enabling JIT gives GoQuickJS a 3.60x higher V8v7 score and runs factorial 3.40x faster than its interpreter. Compilation and warmup are included. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 689.990ms | 551.888ms | 573.946ms | 211.644ms | 63.101ms | 635.908ms | 141.902ms |
| 2 | 680.233ms | 536.974ms | 597.673ms | 215.188ms | 62.741ms | 632.867ms | 146.115ms |
| 3 | 695.291ms | 541.609ms | 591.816ms | 210.922ms | 62.849ms | 637.684ms | 139.979ms |
| 4 | 693.879ms | 547.058ms | 583.248ms | 216.780ms | 62.938ms | 629.117ms | 147.058ms |
| 5 | 700.737ms | 538.855ms | 587.758ms | 213.516ms | 62.171ms | 633.905ms | 141.327ms |
| Average | 692.026ms | 543.277ms | 586.888ms | 213.610ms | **62.760ms** | 633.896ms | 143.276ms |
| Total | 3.460s | 2.716s | 2.934s | 1.068s | **313.802ms** | 3.169s | 716.383ms |
| Relative Time (lower is better) | 11.03x | 8.66x | 9.35x | 3.40x | 1.00x | 10.10x | 2.28x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 19.972ms | 152.687ms |
| 2 | 21.166ms | 171.860ms |
| 3 | 20.659ms | 161.470ms |
| 4 | 20.637ms | 162.646ms |
| 5 | 21.508ms | 160.241ms |
| Average | **20.788ms** | 161.781ms |
| Total | **103.944ms** | 808.905ms |
| Relative Time (lower is better) | 1.00x | 7.78x |

*Node v26.13.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Richards | 528 | 584 | 595 | 2472 | **14350** | 510 | 2831 |
| DeltaBlue | 653 | 581 | 661 | 2732 | **13872** | 595 | 2876 |
| Crypto | 341 | 597 | 444 | 3247 | **20722** | 665 | 2506 |
| RayTrace | 823 | 1255 | 1006 | 5522 | **33562** | 1049 | 7149 |
| EarleyBoyer | 1491 | 1549 | 1553 | 7642 | **14084** | 1930 | 10821 |
| RegExp | 581 | 551 | 275 | 8965 | **10404** | 1148 | 5880 |
| Splay | 2611 | 3178 | 2556 | 13731 | **23526** | 3472 | 15157 |
| NavierStokes | 568 | 1556 | 747 | 5942 | **40136** | 1293 | 7011 |
| Score (version 7) | 770 | 1009 | 784 | 5356 | **19288** | 1097 | 5614 |
| Duration (seconds) | 51.231s | 40.311s | 52.025s | 22.289s | **20.104s** | 37.739s | 21.236s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **82275** | 3050 |
| DeltaBlue | **222639** | 3365 |
| Crypto | **100595** | 2926 |
| RayTrace | **237684** | 9508 |
| EarleyBoyer | **193315** | 13844 |
| RegExp | **26820** | 7720 |
| Splay | **44380** | 16469 |
| NavierStokes | **85627** | 4159 |
| Score (version 7) | **98174** | 6166 |
| Duration (seconds) | **20.040s** | 22.256s |

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
