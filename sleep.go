package sortilege

import (
	"sync"
	"time"
)

// SleepSorter implements Sleep Sort.
// Launches a goroutine per element; each goroutine sleeps for a duration
// proportional to its value, then appends to the result. Elements with smaller
// values wake earlier and appear first.
//
// ScaleFactor controls how long each unit of value sleeps. Defaults to 1ms.
// Only works correctly with non-negative integer inputs.
// Time: O(max(input))  Space: O(n)
type SleepSorter struct {
	ScaleFactor time.Duration
}

func (s SleepSorter) Name() string { return "Sleep Sort" }

func (s SleepSorter) Sort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	scale := s.ScaleFactor
	if scale == 0 {
		scale = time.Millisecond
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	result := make([]int, 0, len(input))
	for _, v := range input {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			time.Sleep(time.Duration(val) * scale)
			mu.Lock()
			result = append(result, val)
			mu.Unlock()
		}(v)
	}
	wg.Wait()
	return result
}
