package slices

func All[T comparable](predicate predicateFunc[T], slice []T) bool {
	for _, element := range slice {
		if !predicate(element) {
			return false
		}
	}

	return true
}
