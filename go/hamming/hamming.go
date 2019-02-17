package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := errors.New("strands are not of equal length")
		return -1, err
	}
	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
