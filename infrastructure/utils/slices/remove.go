package slices

func Remove[T comparable](slice []T, elements ...T) []T {
	newSlice := make([]T, 0)

	for sliceIdx := range slice {
		if !Contains(elements, slice[sliceIdx]) {
			newSlice = append(newSlice, slice[sliceIdx])
		}
	}

	return newSlice
}
