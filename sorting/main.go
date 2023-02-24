package main

type IntSlice []int

func (a IntSlice) Len() int           { return len(a) }
func (a IntSlice) Less(i, j int) bool { return a[i] < a[j] }
func (a IntSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	// nums := []float64{3.4, 67.4, 2.3, 65.3, 90.0, 54.3, 12.3}
	// sort.Float64s(nums) // * sorts in increasing order
	// fmt.Println(sort.Float64sAreSorted(nums)) // * checks wheter it is sorted or not
	// nums := []int{3, 2, 1, 5, 2, 6, 32, 53, 32}
	// sort.Ints(nums) // * sorts slice in incresing order
	// fmt.Println(nums)
}
