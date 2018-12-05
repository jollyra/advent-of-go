package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
	"sort"
	"strconv"
	"strings"
	"time"
)

type event struct {
	time time.Time
	raw  string
}

func (e event) String() string { return fmt.Sprintf("event{time=%q, raw=%s}", e.time, e.raw) }

type byTime []event

func (t byTime) Len() int           { return len(t) }
func (t byTime) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byTime) Less(i, j int) bool { return t[i].time.Before(t[j].time) }

type counter map[int][60]int

func (c counter) RangeAdd(key, a, b int) {
	aggregator := c[key]
	for i := a; i < b; i++ {
		aggregator[i]++
	}
	c[key] = aggregator
}

func parse(lines []string) []event {
	const layout = "2006-01-02 15:04"
	events := make([]event, 0)
	for _, line := range lines {
		words := strings.Split(line, ", ")

		t, err := time.Parse(layout, words[0])
		if err != nil {
			panic(err)
		}

		e := event{t, words[1]}
		events = append(events, e)
	}
	return events
}

func newCounter() counter {
	return make(map[int][60]int)
}

func countTimeAsleep(events []event) counter {
	counter := newCounter()
	words := strings.Split(events[0].raw, " ")
	id, _ := strconv.Atoi(strings.TrimSpace(words[1]))
	currentGuard := id
	prev := events[0].time
	for _, event := range events[1:] {
		words := strings.Split(event.raw, " ")
		if words[0] == "Guard" {
			id, _ := strconv.Atoi(strings.TrimSpace(words[1]))
			currentGuard = id
		} else if words[0] == "falls" {
			prev = event.time
		} else if words[0] == "wakes" {
			counter.RangeAdd(currentGuard, prev.Minute(), event.time.Minute())
		} else {
			panic(fmt.Sprintf("Unrecognized event %q", event))
		}
	}
	return counter
}

func max(c counter) (int, int) {
	maxSum := -1
	maxKey := -1
	for key, buckets := range c {
		sum := 0
		for i := range buckets {
			sum += buckets[i]
		}
		if sum > maxSum {
			maxSum = sum
			maxKey = key
		}
	}
	return maxKey, maxSum
}

func maxBucketByKey(buckets [60]int) int {
	maxIndex := 0
	maxVal := 0
	for i, bucket := range buckets {
		if bucket > maxVal {
			maxVal = bucket
			maxIndex = i
		}
	}
	return maxIndex
}

func maxGuardMinute(c counter) (int, int) {
	var maxKey, maxSum, maxBucket int
	for key, buckets := range c {
		for i := range buckets {
			if buckets[i] > maxSum {
				maxSum = buckets[i]
				maxKey = key
				maxBucket = i
			}
		}
	}
	return maxKey, maxBucket
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	events := parse(lines)
	sort.Sort(byTime(events))
	counter := countTimeAsleep(events)
	sleepiestGuard, _ := max(counter)
	minute := maxBucketByKey(counter[sleepiestGuard])
	fmt.Println("Part 1:", sleepiestGuard*minute)

	guard, sleepiestMinute := maxGuardMinute(counter)
	fmt.Println("Part 2:", guard*sleepiestMinute)
}
