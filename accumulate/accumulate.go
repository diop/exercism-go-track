package accumulate

// Accumulate applies an operation on each element of the collection
func Accumulate(collection []string, operation func(string) string) (results []string) {
	for _, elem := range collection {
		op := operation(elem)
		results = append(results, op)
	}
	return results
}
