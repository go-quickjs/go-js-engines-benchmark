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
| Goja | `v0.0.0-20261007200356-e2ea74d3d210` |
| ModerncQuickJS | `v0.25.0` |
| QJS | `v0.0.6` |
| GoQuickJS | `v0.27.0` |
| Paserati | `v0.9.13` |
| MoeJS | `v0.1.0-alpha.6` |
| Node / NodeJitless | `v26.8.1` |

Results collected on October 16, 2026 using the engine versions listed above. The factorial benchmark averages five runs per engine; V8v7 reports one full suite run per engine. Results vary with system load and hardware.

Node reference results are presented separately and excluded from the Go-hosted engine rankings and relative times. Goja, GoQuickJS, Paserati, and MoeJS are pure-Go engines; ModerncQuickJS, QJS, and Node use different runtime architectures. Node provides a heavily optimized, JIT-enabled runtime reference, while NodeJitless provides an interpreter reference.

The factorial workload discards its return values. An optimizing engine can eliminate redundant work, so these timings measure optimization as well as execution.

GoQuickJS v0.27.0 has the highest V8v7 score among the six Go-hosted engines in this run: 4257, compared with 1767 for MoeJS, 1077 for Paserati, and 725 for Goja. MoeJS has the lowest factorial average at 217.442ms, followed closely by GoQuickJS at 220.546ms (about 1.4% more time). The small factorial gap should be read in light of run-to-run variation. These workloads measure different aspects of execution and do not establish a ranking for every application.

## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 726.823ms | 689.747ms | 604.245ms | 208.063ms | 537.690ms | 221.727ms |
| 2 | 722.220ms | 694.812ms | 617.076ms | 215.207ms | 535.019ms | 221.158ms |
| 3 | 732.469ms | 692.784ms | 616.240ms | 244.844ms | 549.086ms | 207.329ms |
| 4 | 723.920ms | 701.936ms | 623.265ms | 212.263ms | 528.038ms | 223.355ms |
| 5 | 742.624ms | 681.168ms | 619.002ms | 222.353ms | 536.816ms | 213.639ms |
| Average | 729.611ms | 692.089ms | 615.965ms | 220.546ms | 537.330ms | **217.442ms** |
| Total | 3.648s | 3.460s | 3.080s | 1.103s | 2.687s | **1.087s** |
| Relative Time (lower is better) | 3.36x | 3.18x | 2.83x | 1.01x | 2.47x | 1.00x |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.0.*

### Node.js Reference Results

| Iteration | Node | NodeJitless |
| --- | --- | --- |
| 1 | 21.143ms | 164.316ms |
| 2 | 21.244ms | 172.361ms |
| 3 | 21.241ms | 162.373ms |
| 4 | 21.780ms | 160.557ms |
| 5 | 19.691ms | 165.884ms |
| Average | **21.020ms** | 165.098ms |
| Total | **105.102ms** | 825.493ms |
| Relative Time (lower is better) | 1.00x | 7.85x |

*Node v26.8.1, using the same hardware and OS as above. Reference rankings compare only the two Node configurations.*

## V8v7 Benchmark Results

Higher scores are better; lower duration is better.

| Metric | GOJA | ModerncQuickJS | QJS | GoQuickJS | Paserati | MoeJS |
| --- | --- | --- | --- | --- | --- | --- |
| Richards | 506 | 470 | 549 | **2053** | 475 | 795 |
| DeltaBlue | 615 | 493 | 609 | **2600** | 607 | 1144 |
| Crypto | 329 | 408 | 443 | **2881** | 618 | 1454 |
| RayTrace | 748 | 1037 | 976 | **5681** | 1142 | 1944 |
| EarleyBoyer | 1393 | 1242 | 1442 | **6211** | 1923 | 2965 |
| RegExp | 579 | 313 | 256 | **4019** | 1088 | 1035 |
| Splay | 2304 | 2640 | 2456 | **8687** | 3463 | 4824 |
| NavierStokes | 535 | 1026 | 704 | **5696** | 1225 | 2498 |
| Score (version 7) | 725 | 753 | 742 | **4257** | 1077 | 1767 |
| Duration (seconds) | 53.497s | 50.939s | 52.380s | **22.321s** | 37.321s | 29.747s |

*Benchmarks run on Apple M5 Max, 128 GiB RAM, macOS ARM64, Go 1.27.0.*

### Node.js Reference Results

| Metric | Node | NodeJitless |
| --- | --- | --- |
| Richards | **77703** | 2800 |
| DeltaBlue | **207690** | 3061 |
| Crypto | **97809** | 2786 |
| RayTrace | **226585** | 8965 |
| EarleyBoyer | **177814** | 13330 |
| RegExp | **23251** | 6571 |
| Splay | **88043** | 15247 |
| NavierStokes | **78277** | 4211 |
| Score (version 7) | **100237** | 5756 |
| Duration (seconds) | **20.046s** | 22.236s |

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
