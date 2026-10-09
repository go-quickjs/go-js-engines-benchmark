package main

import (
	"arewefastyet-v8v7/engines"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type V8V7TestScore struct {
	Name  string
	Score int
}

type EngineScore struct {
	Engine    string
	Reference bool
	Duration  time.Duration
	Scores    []V8V7TestScore
}

func benchmarkEngine(engine engines.JSEngine) EngineScore {
	runtime.GC()
	time.Sleep(10 * time.Millisecond)
	engineName := engine.Name()
	start := time.Now()
	output := must(engine.Run("./v8-v7/run.js"))
	duration := time.Since(start)
	score := must(parseScore(engineName, output))
	score.Duration = duration
	fmt.Printf("Duration: %.3fs\n", duration.Seconds())

	return score
}

func main() {
	engs := engines.Engines()
	var allScores []EngineScore

	for _, eng := range engs {
		func(eng engines.JSEngine) {
			fmt.Printf("Running engine: %s\n", eng.Name())
			must(0, eng.Init())
			score := benchmarkEngine(eng)
			_, score.Reference = eng.(*engines.Node)
			allScores = append(allScores, score)
			eng.Close()
			fmt.Println()
		}(eng)
	}

	var primary, reference []EngineScore
	for _, score := range allScores {
		if score.Reference {
			reference = append(reference, score)
		} else {
			primary = append(primary, score)
		}
	}
	fmt.Println("## Go-Hosted Engine Results")
	generateMarkdownTable(primary)
	fmt.Println("\n## Node.js Reference Results")
	generateMarkdownTable(reference)
}

func generateMarkdownTable(scores []EngineScore) {
	if len(scores) == 0 {
		return
	}

	t := table.NewWriter()
	header := table.Row{"Metric"}
	for _, score := range scores {
		header = append(header, score.Engine)
	}
	t.AppendHeader(header)

	metricNames := make([]string, 0)
	for _, test := range scores[0].Scores {
		metricNames = append(metricNames, test.Name)
	}

	for _, metricName := range metricNames {
		row := table.Row{metricName}
		maxScore := 0
		for _, engineScore := range scores {
			for _, test := range engineScore.Scores {
				if test.Name == metricName && test.Score > maxScore {
					maxScore = test.Score
				}
			}
		}

		for _, engineScore := range scores {
			for _, test := range engineScore.Scores {
				if test.Name == metricName {
					if test.Score == maxScore {
						row = append(row, fmt.Sprintf("**%d**", test.Score))
					} else {
						row = append(row, fmt.Sprintf("%d", test.Score))
					}
					break
				}
			}
		}
		t.AppendRow(row)
	}

	row := table.Row{"Duration (seconds)"}
	minDuration := time.Duration(0)
	for _, engineScore := range scores {
		if minDuration == 0 || engineScore.Duration < minDuration {
			minDuration = engineScore.Duration
		}
	}

	for _, engineScore := range scores {
		if engineScore.Duration == minDuration {
			row = append(row, fmt.Sprintf("**%.3fs**", engineScore.Duration.Seconds()))
		} else {
			row = append(row, fmt.Sprintf("%.3fs", engineScore.Duration.Seconds()))
		}
	}

	t.AppendRow(row)
	fmt.Println(t.RenderMarkdown())
}

// Example output:
// Richards: 382
// DeltaBlue: 450
// Crypto: 199
// RayTrace: 442
// EarleyBoyer: 869
// RegExp: 298
// Splay: 1332
// NavierStokes: 357
// ----
// Score (version 7): 456
func parseScore(engineName string, output [][]string) (EngineScore, error) {
	var engineScore EngineScore
	engineScore.Engine = engineName
	for _, line := range output {
		if len(line) == 0 {
			continue
		}
		text := line[0]
		if text == "----" {
			continue
		}

		parts := strings.Split(text, ": ")
		if len(parts) != 2 {
			return EngineScore{}, fmt.Errorf("could not parse line: %s", text)
		}

		var score int
		if _, err := fmt.Sscanf(parts[1], "%d", &score); err != nil {
			return EngineScore{}, fmt.Errorf("could not parse score from line: %s, error: %v", text, err)
		}

		engineScore.Scores = append(engineScore.Scores, V8V7TestScore{
			Name:  parts[0],
			Score: score,
		})
	}

	return engineScore, nil
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
