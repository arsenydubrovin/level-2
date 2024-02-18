package cmd

import (
	"sort"
	"strings"
)

// GetAnagrams возвращает мапу множеств анаграмм для каждого слова из words
func GetAnagrams(words []string, dict map[string]struct{}) map[string][]string {
	m := make(map[string][]string, len(words))

	words = filterUnique(words)
	words = filterByDict(words, dict)

	for _, word := range words {
		word = strings.ToLower(word)
		perms := getPermutations(word)
		anagrams := filterByDict(perms, dict)
		anagrams = filterUnique(anagrams)
		sort.Strings(perms)

		// множества из одного элемента не должны попасть в результат
		if len(anagrams) <= 1 {
			continue
		}

		m[anagrams[0]] = anagrams
	}

	return m
}

// filterByDict оставлятет в slice только не слова, которые есть в dict
func filterByDict(slice []string, dict map[string]struct{}) []string {
	filteredSlice := make([]string, 0)

	for _, el := range slice {
		if _, exists := dict[el]; exists {
			filteredSlice = append(filteredSlice, el)
		}
	}

	return filteredSlice
}

// filterUnique оставляет в словаре только уникальные строки
func filterUnique(slice []string) []string {
	set := make(map[string]struct{})
	uniqueSlice := make([]string, 0)

	for _, el := range slice {
		if _, ok := set[el]; !ok {
			set[el] = struct{}{}
			uniqueSlice = append(uniqueSlice, el)
		}
	}

	return uniqueSlice
}
