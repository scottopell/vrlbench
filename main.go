package main

import (
	"fmt"
	"time"

	flog "github.com/mingrammer/flog"
)

func setupFFIRunner() {
	// no-op
}

func setupWasmRunner() {
	// no-op
}

type Processor interface {
	Process(string) string
	Describe() string
}

func performTest(logCh chan string, testDuration time.Duration, processor Processor) string {
	recorder := BlackholeRecorder[string]{}
	start := time.Now()

	for log := range logCh {
		_ = processor.Process(log)
		recorder.Consume(log)
		if time.Since(start) > testDuration {
			break
		}
	}
	return recorder.AvgThroughput()
}

func main() {
	logCh := flog.GenerateInfiniteLogs()
	testDuration := time.Second * 5

	processor := NewRegexProcessor(`\b\w{4}\b`)
	res := performTest(logCh, testDuration, processor)
	fmt.Println(processor.Describe(), "Avg Throughput:", res)

	processor = NewRegexProcessor(`\b\w{4}`)
	res = performTest(logCh, testDuration, processor)
	fmt.Println(processor.Describe(), "Avg Throughput:", res)

	processor = NewRegexProcessor(`\w{4}`)
	res = performTest(logCh, testDuration, processor)
	fmt.Println(processor.Describe(), "Avg Throughput:", res)

	noopProcessor := NewNoopProcessor()
	res = performTest(logCh, testDuration, noopProcessor)
	fmt.Println(noopProcessor.Describe(), "Avg Throughput:", res)

}
