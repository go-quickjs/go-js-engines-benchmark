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
| Goja | `v0.0.0-20261116234614-2eded47d7035` |
| ModerncQuickJS | `v0.27.0` |
| QJS | `v0.0.6` |
| GoQuickJS / GoQuickJSJIT | `v0.38.1-0.20261116232310-ac8f5406c806` |
| Paserati | `v0.10.1` |
| MoeJS | `v0.1.1` |
| Node / NodeJitless | `v26.13.1` |

GoQuickJS is pinned to [`main` commit `ac8f5406c806`](https://github.com/go-quickjs/go-quickjs/commit/ac8f5406c80619fd97f42c9f9015f52b2a55253b). Other engines use their latest stable releases; Goja has no tagged releases and uses the latest Go module pseudo-version.

Results collected on November 16, 2026 using the engine versions listed above and Go 1.27.3. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, and MoeJS are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJSJIT, with a 32 MiB JIT memory limit, has the highest V8v7 score among the seven Go-hosted engine configurations in this run: 19950, compared with 6665 for the GoQuickJS interpreter and 5772 for MoeJS. Its factorial average is 62.077ms, compared with 137.985ms for MoeJS and 172.167ms for the GoQuickJS interpreter. Among the Go-hosted interpreter configurations, GoQuickJS has the highest V8v7 score, while MoeJS has the fastest factorial average. Enabling JIT gives GoQuickJS a 2.99x higher V8v7 score and runs factorial 2.77x faster than its interpreter. Compilation and warmup are included. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 724.063ms | 541.128ms | 584.195ms | 177.104ms | 62.846ms | 691.564ms | 134.291ms |
| 2 | 696.346ms | 562.443ms | 585.275ms | 162.783ms | 62.004ms | 644.795ms | 145.542ms |
| 3 | 710.633ms | 549.430ms | 595.752ms | 170.980ms | 59.536ms | 645.623ms | 126.867ms |
| 4 | 674.523ms | 543.044ms | 585.945ms | 173.478ms | 62.530ms | 631.868ms | 149.512ms |
| 5 | 688.230ms | 528.819ms | 584.646ms | 176.492ms | 63.467ms | 635.203ms | 133.712ms |
| Average | 698.759ms | 544.973ms | 587.163ms | 172.167ms | **62.077ms** | 649.811ms | 137.985ms |
| Total | 3.494s | 2.725s | 2.936s | 860.838ms | **310.385ms** | 3.249s | 689.927ms |
| Relative Time (lower is better) | 11.26x | 8.78x | 9.46x | 2.77x | 1.00x | 10.47x | 2.22x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 28.454ms | 157.604ms |
| 2 | 21.065ms | 159.296ms |
| 3 | 21.695ms | 158.243ms |
| 4 | 20.800ms | 172.411ms |
| 5 | 20.867ms | 159.890ms |
| Average | **22.576ms** | 161.489ms |
| Total | **112.883ms** | 807.446ms |
| Relative Time (lower is better) | 1.00x | 7.15x |

*Node v26.13.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Richards | 527 | 574 | 596 | 3237 | **15808** | 513 | 2860 |
| DeltaBlue | 588 | 560 | 649 | 2913 | **14182** | 597 | 3019 |
| Crypto | 337 | 595 | 474 | 5370 | **21275** | 679 | 2548 |
| RayTrace | 810 | 1244 | 1024 | 6949 | **36519** | 1122 | 7474 |
| EarleyBoyer | 1469 | 1538 | 1567 | 8211 | **13845** | 1997 | 11060 |
| RegExp | 594 | 542 | 259 | 8784 | **10932** | 1173 | 6109 |
| Splay | 2676 | 3235 | 2689 | 15272 | **24064** | 3447 | 15573 |
| NavierStokes | 579 | 1556 | 729 | 10051 | **39544** | 1313 | 7123 |
| Score (version 7) | 763 | 1000 | 789 | 6665 | **19950** | 1120 | 5772 |
| Duration (seconds) | 51.229s | 40.264s | 52.254s | 20.171s | **20.082s** | 37.353s | 21.239s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **83599** | 3001 |
| DeltaBlue | **243532** | 3402 |
| Crypto | **104791** | 2931 |
| RayTrace | **239756** | 9546 |
| EarleyBoyer | **195912** | 13837 |
| RegExp | **26236** | 7637 |
| Splay | **43663** | 16339 |
| NavierStokes | **85775** | 4367 |
| Score (version 7) | **99807** | 6189 |
| Duration (seconds) | **20.050s** | 22.220s |

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
