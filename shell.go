package sortilege

// ShellSorter implements Shell Sort.
// An optimisation of Insertion Sort that exchanges far-apart elements first.
// Uses a gap sequence starting at n/2, halved each iteration.
// Time: O(n²)  Space: O(1)
type ShellSorter struct{}

func (ShellSorter) Name() string { return "Shell Sort" }

func (ShellSorter) Sort(input []int) []int {
	arr := copySlice(input)
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
	return arr
}
