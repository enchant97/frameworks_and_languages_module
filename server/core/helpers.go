package core

func containsAll(values []string, toFind []string) bool {
	valuesMap := make(map[interface{}]*interface{})
	for _, v := range values {
		valuesMap[v] = nil
	}

	for _, v := range toFind {
		if _, exists := valuesMap[v]; !exists {
			return false
		}
	}
	return true
}
