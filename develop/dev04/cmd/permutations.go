package cmd

// getPermutations возвращает все перестановки для а
func getPermutations(a string) []string {
	result := make([]string, 0)

	if a == "" {
		return result // если строка пустая, у неё нет перестановок
	}

	permute([]rune(a), 0, &result)

	return result
}

func permute(a []rune, i int, result *[]string) {
	if i > len(a) {
		*result = append(*result, string(a))
		return
	}

	permute(a, i+1, result)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permute(a, i+1, result)
		a[i], a[j] = a[j], a[i]
	}
}
