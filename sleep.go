package ordinex

import (
	"sync"
	"time"
)

// SleepSorter implements Sleep Sort. It launches one goroutine per element; each
// goroutine sleeps for a duration proportional to its value and then appends
// itself to the result, so smaller values wake earlier and appear first.
//
// Sleep Sort works correctly only with non-negative integer inputs. It launches
// one goroutine per element, so it is unsuited to large inputs.
//
// Time: O(max(input)). Space: O(n).
type SleepSorter struct {
	// ScaleFactor controls how long each unit of value sleeps. If zero, it
	// defaults to one millisecond.
	ScaleFactor time.Duration
}

// Name returns the algorithm's name, "Sleep Sort".
func (s SleepSorter) Name() string { return "Sleep Sort" }

// Sort returns a sorted copy of input using Sleep Sort. The input is not
// modified. Results are reliable only for non-negative values; negative values
// sleep for a non-positive duration and may appear out of order.
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
