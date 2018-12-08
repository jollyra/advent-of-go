package main

import (
	"container/heap"
	"fmt"

	"github.com/jollyra/go-advent-util"
)

type job struct {
	Timer int
	Name  rune
}

func (j *job) String() string { return fmt.Sprintf("%c %d", j.Name, j.Timer) }

func (j job) Equal(r rune) bool { return j.Name == r }

func containsJobs(xs []*job, r rune) bool {
	for _, x := range xs {
		if x.Equal(r) {
			return true
		}
	}
	return false
}

type jobHeap []*job

func (h jobHeap) Len() int           { return len(h) }
func (h jobHeap) Less(i, j int) bool { return h[i].Name < h[j].Name }
func (h jobHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *jobHeap) Push(x interface{}) {
	if containsJobs(*h, x.(*job).Name) {
		return
	}

	*h = append(*h, x.(*job))
}

func (h *jobHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func parseEdges(lines []string) (map[rune][]rune, map[rune][]rune, []rune) {
	edges := make(map[rune][]rune)
	for key := range edges {
		edges[key] = make([]rune, 0)
	}

	dependency := make(map[rune][]rune)
	for key := range dependency {
		dependency[key] = make([]rune, 0)
	}

	srcs := make([]rune, 0)
	dsts := make([]rune, 0)
	for _, line := range lines {
		edge := advent.Split(line)
		a := rune(edge[0][0])
		b := rune(edge[1][0])
		edges[a] = append(edges[a], b)
		dependency[b] = append(dependency[b], a)
		srcs = append(srcs, a)
		dsts = append(dsts, b)
	}

	starts := diffStartingNodes(srcs, dsts)

	return edges, dependency, starts
}

func diffStartingNodes(xs, ys []rune) []rune {
	starts := make([]rune, 0)
	for _, x := range xs {
		if !contains(ys, x) {
			starts = append(starts, x)
		}
	}
	return starts
}

func contains(xs []rune, y rune) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func dependenciesMet(dependency []rune, complete []*job) bool {
	if len(dependency) == 0 {
		return true
	}

	for _, d := range dependency {
		if !containsJobs(complete, d) {
			return false
		}
		for i := range complete {
			if complete[i].Name == d && complete[i].Timer > 0 {
				return false
			}
		}
	}
	return true
}

func getCurrentN(horizon *jobHeap, deps map[rune][]rune, complete []*job, n int) []*job {
	curs := make([]*job, 0)
	for i := 0; i < n; i++ {
		var cur *job
		blocked := make([]*job, 0)
		found := false
		for found == false {
			if horizon.Len() == 0 {
				break
			}
			cur = heap.Pop(horizon).(*job)
			if dependenciesMet(deps[cur.Name], complete) {
				for _, r := range blocked {
					heap.Push(horizon, r)
				}
				found = true
				curs = append(curs, cur)
			} else {
				blocked = append(blocked, cur)
			}
		}
	}
	return curs
}

func countIncomplete(jobs []*job) int {
	count := 0
	for i := range jobs {
		if jobs[i].Timer > 0 {
			count++
		}
	}
	return count
}

func bfs(edges, dependency map[rune][]rune, starts []rune) int {
	elves := 5
	path := make([]*job, 0)
	complete := make([]*job, 0)
	horizon := &jobHeap{}
	for _, start := range starts {
		heap.Push(horizon, &job{Name: start, Timer: int(start) - 4})
	}
	heap.Init(horizon)
	complete = append(complete, getCurrentN(horizon, dependency, complete, elves)...)
	timer := 0
	for countIncomplete(complete) > 0 {
		timer++
		for _, cur := range complete {
			fmt.Println(cur)
			if cur.Timer == 0 {
				continue
			}
			cur.Timer--
			if cur.Timer == 0 {
				path = append(path, cur)
				for _, next := range edges[cur.Name] {
					heap.Push(horizon, &job{Name: next, Timer: int(next) - 4})
				}
				newJobs := elves - countIncomplete(complete)
				complete = append(complete,
					getCurrentN(horizon, dependency, complete, newJobs)...)
			}
		}
	}

	return timer
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	edges, dependency, starts := parseEdges(lines)
	fmt.Println(bfs(edges, dependency, starts))
}
