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
| Goja | `v0.0.0-20261121162556-79d6ad37b976` |
| ModerncQuickJS | `v0.27.0` |
| QJS | `v0.0.6` |
| GoQuickJS / GoQuickJSJIT | `v0.43.0` |
| Paserati | `v0.10.1` |
| MoeJS | `v0.1.1` |
| Ekuma | `v0.0.0-20261122110136-d62f469797c3` |
| Node / NodeJitless | `v26.13.1` |

Engines use their latest stable releases where available. Goja and Ekuma have no tagged releases and use their latest Go module pseudo-versions.

Results collected on November 22, 2026 using the engine versions listed above and Go 1.27.3. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, MoeJS, and Ekuma are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJSJIT, with a 32 MiB JIT memory limit, has the highest V8v7 score among the eight Go-hosted engine configurations in this run: 19886, compared with 7656 for the GoQuickJS interpreter, 5785 for MoeJS, and 5679 for Ekuma. Its factorial average is 65.839ms, compared with 143.746ms for MoeJS, 170.010ms for Ekuma, and 173.822ms for the GoQuickJS interpreter. Among the Go-hosted interpreter configurations, GoQuickJS has the highest V8v7 score, while MoeJS has the fastest factorial average. Enabling JIT gives GoQuickJS a 2.60x higher V8v7 score and runs factorial 2.64x faster than its interpreter. Compilation and warmup are included. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS | Ekuma |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 713.425ms | 539.434ms | 623.574ms | 180.584ms | 62.303ms | 629.252ms | 140.473ms | 167.966ms |
| 2 | 711.296ms | 532.907ms | 584.510ms | 169.875ms | 62.462ms | 625.508ms | 139.975ms | 172.326ms |
| 3 | 703.569ms | 551.956ms | 578.707ms | 175.095ms | 64.171ms | 625.067ms | 146.300ms | 166.733ms |
| 4 | 683.691ms | 543.867ms | 586.423ms | 173.450ms | 63.150ms | 621.797ms | 151.417ms | 174.701ms |
| 5 | 722.558ms | 535.462ms | 584.820ms | 170.108ms | 77.109ms | 625.713ms | 140.567ms | 168.323ms |
| Average | 706.908ms | 540.725ms | 591.607ms | 173.822ms | **65.839ms** | 625.467ms | 143.746ms | 170.010ms |
| Total | 3.535s | 2.704s | 2.958s | 869.114ms | **329.197ms** | 3.127s | 718.734ms | 850.050ms |
| Relative Time (lower is better) | 10.74x | 8.21x | 8.99x | 2.64x | 1.00x | 9.50x | 2.18x | 2.58x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 26.575ms | 177.158ms |
| 2 | 21.335ms | 164.880ms |
| 3 | 22.010ms | 157.292ms |
| 4 | 21.074ms | 165.372ms |
| 5 | 22.069ms | 156.591ms |
| Average | **22.613ms** | 164.259ms |
| Total | **113.066ms** | 821.296ms |
| Relative Time (lower is better) | 1.00x | 7.26x |

*Node v26.13.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | GoQuickJSJIT | Paserati | MoeJS | Ekuma |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Richards | 533 | 592 | 591 | 3280 | **15600** | 473 | 2951 | 2984 |
| DeltaBlue | 656 | 596 | 651 | 3603 | **14255** | 592 | 3038 | 3537 |
| Crypto | 350 | 598 | 471 | 5392 | **21881** | 662 | 2791 | 3866 |
| RayTrace | 811 | 1242 | 1037 | 10708 | **36149** | 1089 | 7540 | 8206 |
| EarleyBoyer | 1514 | 1547 | 1543 | 11375 | **15733** | 1956 | 10798 | 9125 |
| RegExp | 588 | 550 | 266 | 8848 | **11103** | 1141 | 5983 | 4360 |
| Splay | 2646 | 3199 | 2684 | 16282 | **20886** | 3427 | 15386 | 10594 |
| NavierStokes | 584 | 1551 | 737 | 10558 | **38101** | 1300 | 6693 | 7671 |
| Score (version 7) | 779 | 1013 | 791 | 7656 | **19886** | 1091 | 5785 | 5679 |
| Duration (seconds) | 50.801s | 39.821s | 50.678s | 20.147s | **20.084s** | 37.538s | 21.282s | 20.262s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.3.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **82211** | 3032 |
| DeltaBlue | **235869** | 3395 |
| Crypto | **102134** | 2827 |
| RayTrace | **228065** | 9462 |
| EarleyBoyer | **190266** | 14003 |
| RegExp | **26236** | 7509 |
| Splay | **98669** | 16461 |
| NavierStokes | **78504** | 4448 |
| Score (version 7) | **107219** | 6177 |
| Duration (seconds) | **20.050s** | 22.184s |

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
