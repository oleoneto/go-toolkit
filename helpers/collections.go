package helpers

// Contains Checks if the element is contained within the given collection.
//
// Example:
//
//	Contains([]string{"hello", "world", "!"}, "world") // -> true
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// ContainsDuplicate return true if any value appears at least twice in the given collection.
//
// Example:
//
//	ContainsDuplicate([]int{1, 2, 3}) // -> false
func ContainsDuplicate[T comparable](collection []T) bool {
	visited := make(map[T]int)

	for idx, x := range collection {
		if _, found := visited[x]; found {
			return true
		}

		visited[x] = idx
	}

	return len(visited) != len(collection)
}

// Map Applies a `transformer` function to every element in a collection.
//
// Usage:
//
//	Map([]int{3, 4}, func(index, n int) int { return n * n }) // -> [9, 16]
func Map[A any, B any](collection []A, transformFunc func(int, A) B) []B {
	result := make([]B, len(collection))

	for index, item := range collection {
		result[index] = transformFunc(index, item)
	}

	return result
}

// Filter Filters a collection and returns only the elements
// that match the provided `inclusionTest`.
//
// Example:
//
// Filter and return all even numbers:
//
//	Filter([]int{16, 9, 25}, func(i, n int) bool { return n%2 == 0 }) // -> [16]
func Filter[T any](collection []T, inclusionTest func(int, T) bool) []T {
	result := make([]T, 0)

	for index, item := range collection {
		if inclusionTest(index, item) {
			result = append(result, item)
		}
	}

	return result
}

// Reduce applies a counter function to every element of a collection and returns a total sum.
//
// Example:
//
//	Reduce([]Payments{{Amount: 25.99}, {Amount: 4.01}}, func(i int, p Payment) float64 { return p.Amount }, 0) // -> 30
func Reduce[T any](c []T, counter func(int, T) float64, initialCount float64) float64 {
	sum := initialCount
	for index, item := range c {
		sum += counter(index, item)
	}

	return sum
}

// Performs sorting of a collection using the quicksort algorithm.
//
// Example:
//
// QuickSort([]int{1, 9, 4, 2, 9, 3, 5, 4}, func(idx int, item, pivot T) bool { return item < pivot })
// -> [1, 2, 3, 4, 4, 5, 9, 9]
func QuickSort[T comparable](collection []T, lhsFunc func(idx int, a T, b T) bool) []T {
	switch len(collection) {
	case 0, 1:
		return collection
	case 2:
		if !lhsFunc(0, collection[0] /*item*/, collection[1] /*pivot*/) {
			temp := collection[0]
			collection[0] = collection[1]
			collection[1] = temp
		}

		return collection
	}

	pivot := collection[len(collection)-1]
	var lhs, rhs []T
	var pivotOccurrences int

	for idx, el := range collection {
		// In case the pivot appears more than once,
		// keep track of how many times it appears
		if el == pivot {
			pivotOccurrences += 1
			continue
		}

		if lhsFunc(idx, el, pivot) {
			lhs = append(lhs, el)
		} else {
			rhs = append(rhs, el)
		}
	}

	lhs = QuickSort(lhs, lhsFunc)
	rhs = QuickSort(rhs, lhsFunc)

	for range pivotOccurrences {
		// this is equivalent to prepending the pivot to the rhs
		lhs = append(lhs, pivot)
	}

	return append(lhs, rhs...)
}
