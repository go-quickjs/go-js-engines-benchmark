package main

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

const JS = `function factorial(n) {
	return n <= 1 ? 1 : n * factorial(n - 1);
}

var i = 0;

while (i++ < 1e6) {
	factorial(10);
}`

type BenchmarkResult struct {
	Duration time.Duration
}

type EngineMetrics struct {
	Name      string
	Reference bool
	Results   []BenchmarkResult
	Total     BenchmarkResult
	Average   BenchmarkResult
}

func printResultsTable(metrics []EngineMetrics, iterations int) {
	var primary, reference []EngineMetrics
	for _, metric := range metrics {
		if metric.Reference {
			reference = append(reference, metric)
		} else {
			primary = append(primary, metric)
		}
	}
	fmt.Println("\n## Go-Hosted Engine Results")
	printMetricsTable(primary, iterations)
	fmt.Println("\n## Node.js Reference Results")
	printMetricsTable(reference, iterations)
}

func printMetricsTable(metrics []EngineMetrics, iterations int) {
	if len(metrics) == 0 {
		return
	}
	t := table.NewWriter()
	header := table.Row{"Iteration"}
	for _, m := range metrics {
		header = append(header, m.Name)
	}

	t.AppendHeader(header)
	for i := range iterations {
		row := table.Row{fmt.Sprintf("%d", i+1)}
		for _, m := range metrics {
			row = append(row, formatDuration(m.Results[i].Duration))
		}
		t.AppendRow(row)
	}

	t.AppendSeparator()

	fastestAvgIdx := 0
	for i := 1; i < len(metrics); i++ {
		if metrics[i].Average.Duration < metrics[fastestAvgIdx].Average.Duration {
			fastestAvgIdx = i
		}
	}

	avgRow := table.Row{"Average"}
	for i, m := range metrics {
		if i == fastestAvgIdx {
			avgRow = append(avgRow, fmt.Sprintf("**%s**", formatDuration(m.Average.Duration)))
		} else {
			avgRow = append(avgRow, formatDuration(m.Average.Duration))
		}
	}

	t.AppendRow(avgRow)

	fastestTotalIdx := 0
	for i := 1; i < len(metrics); i++ {
		if metrics[i].Total.Duration < metrics[fastestTotalIdx].Total.Duration {
			fastestTotalIdx = i
		}
	}

	totalRow := table.Row{"Total"}
	for i, m := range metrics {
		if i == fastestTotalIdx {
			totalRow = append(totalRow, fmt.Sprintf("**%s**", formatDuration(m.Total.Duration)))
		} else {
			totalRow = append(totalRow, formatDuration(m.Total.Duration))
		}
	}

	t.AppendRow(totalRow)

	fastest := metrics[fastestAvgIdx]
	speedRow := table.Row{"Relative Time (lower is better)"}
	for _, m := range metrics {
		ratio := float64(m.Average.Duration) / float64(fastest.Average.Duration)
		speedRow = append(speedRow, fmt.Sprintf("%.2fx", ratio))
	}
	t.AppendRow(speedRow)

	fmt.Println(t.RenderMarkdown())
}

func formatDuration(d time.Duration) string {
	ms := float64(d.Microseconds()) / 1000.0
	if ms < 1000 {
		return fmt.Sprintf("%.3fms", ms)
	}
	s := ms / 1000.0
	if s < 60 {
		return fmt.Sprintf("%.3fs", s)
	}
	m := s / 60.0
	return fmt.Sprintf("%.3fm", m)
}
