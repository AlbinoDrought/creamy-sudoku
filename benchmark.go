package main

import (
	"fmt"
	"reflect"
	"time"
)

type runResult struct {
	Start time.Time
	End   time.Time
}

func (result *runResult) Duration() time.Duration {
	return result.End.Sub(result.Start)
}

type benchmarkResult struct {
	Target     interface{}
	Benchmark  time.Duration
	Runs       int
	RunResults []*runResult
}

func (result *benchmarkResult) LogRun(run *runResult) {
	result.RunResults = append(result.RunResults, run)
	result.Runs = len(result.RunResults)
	result.Benchmark += run.Duration()
}

func (result *benchmarkResult) Bench(callback func()) *runResult {
	run := &runResult{
		Start: time.Now(),
	}
	callback()
	run.End = time.Now()
	result.LogRun(run)
	return run
}

func (result *benchmarkResult) Print() {
	fmt.Printf(
		"%v: %vms (%vms/puzzle)\n",
		reflect.TypeOf(result.Target).String(),
		result.Benchmark.Milliseconds(),
		result.Benchmark.Milliseconds()/int64(result.Runs),
	)
}
