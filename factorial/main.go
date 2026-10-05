package main

import (
	"factorial/engines"
	"fmt"
	"runtime"
	"time"
)

const iterations = 5

func benchmarkEngine(engine engines.JSEngine) (BenchmarkResult, error) {
	runtime.GC()
	time.Sleep(10 * time.Millisecond)
	start := time.Now()
	err := engine.Run(JS)
	duration := time.Since(start)

	return BenchmarkResult{Duration: duration}, err
}

func createMetrics() []EngineMetrics {
	engs := engines.Engines()
	metrics := make([]EngineMetrics, len(engs))
	for i, engine := range engs {
		metrics[i] = EngineMetrics{
			Name:    engine.Name(),
			Results: make([]BenchmarkResult, iterations),
		}
	}

	return metrics
}

func main() {
	fmt.Println("Factorial Benchmark")
	fmt.Println("===================")
	fmt.Println("Computing factorial(10) 1,000,000 times")
	fmt.Println("Measuring execution time only")
	fmt.Println()
	fmt.Println("Memory is not measured because:")
	fmt.Println("  - Goja uses Go memory (visible)")
	fmt.Println("  - GoQuickJS uses Go memory (visible)")
	fmt.Println("  - Paserati uses Go memory (visible)")
	fmt.Println("  - QJS uses WASM memory (invisible to Go)")
	fmt.Println("  - ModerncQuickJS uses mmap memory (invisible to Go)")

	engs := engines.Engines()
	metrics := createMetrics()

	for ite := range iterations {
		fmt.Printf("Iteration %d/%d\n", ite+1, iterations)
		indices := make([]int, len(engs))
		for j := range indices {
			indices[j] = j
		}

		// Alternate execution order to reduce bias
		if ite%2 != 0 {
			for j := 0; j < len(indices)/2; j++ {
				indices[j], indices[len(indices)-1-j] = indices[len(indices)-1-j], indices[j]
			}
		}

		engs := engines.Engines()
		for _, i := range indices {
			engine := engs[i]
			if err := engine.Init(); err != nil {
				panic(fmt.Sprintf("failed to init engine %s: %v", engine.Name(), err))
			}

			result, err := benchmarkEngine(engine)
			if err != nil {
				panic(fmt.Sprintf("%s error: %v", metrics[i].Name, err))
			}

			fmt.Printf("  %-15s: %v\n", metrics[i].Name, result.Duration)
			metrics[i].Results[ite] = result
			if err := engine.Close(); err != nil {
				panic(fmt.Sprintf("failed to close engine %s: %v", engine.Name(), err))
			}
			runtime.GC()
			time.Sleep(100 * time.Millisecond)
		}

		runtime.GC()
		time.Sleep(100 * time.Millisecond)
	}

	for i := range metrics {
		for _, result := range metrics[i].Results {
			metrics[i].Total.Duration += result.Duration
		}

		metrics[i].Average = BenchmarkResult{
			Duration: metrics[i].Total.Duration / time.Duration(iterations),
		}
	}

	printResultsTable(metrics, iterations)
}
