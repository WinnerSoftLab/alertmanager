package slackV2

import (
	"github.com/prometheus/alertmanager/template"
	"unicode/utf8"
)

const SummaryMessageDiffThreshold = 3

func UniqStr(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

func getMapValue(data template.KV, key string) string {
	if value, ok := data[key]; ok {
		return value
	} else {
		return ""
	}
}

func levenshteinDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return utf8.RuneCountInString(s2)
	} else if len(s2) == 0 {
		return utf8.RuneCountInString(s1)
	} else if s1 == s2 {
		return 0
	}

	min := func(values ...int) int {
		m := values[0]
		for _, v := range values {
			if v < m {
				m = v
			}
		}
		return m
	}
	r1, r2 := []rune(s1), []rune(s2)
	n, m := len(r1), len(r2)
	if n > m {
		r1, r2 = r2, r1
		n, m = m, n
	}
	currentRow := make([]int, n+1)
	previousRow := make([]int, n+1)
	for i := range currentRow {
		currentRow[i] = i
	}
	for i := 1; i <= m; i++ {
		for j := range currentRow {
			previousRow[j] = currentRow[j]
			if j == 0 {
				currentRow[j] = i
				continue
			} else {
				currentRow[j] = 0
			}
			add, del, change := previousRow[j]+1, currentRow[j-1]+1, previousRow[j-1]
			if r1[j-1] != r2[i-1] {
				change++
			}
			currentRow[j] = min(add, del, change)
		}
	}
	return currentRow[n]
}

func mergeSameMessages(arr []string) []string {
	result := make([]string, 0)
	if len(arr) > 0 {
		result = append(result, arr[0])
	}

	for _, val := range arr {
		differs := 0
		for _, res := range result {
			if levenshteinDistance(val, res) > SummaryMessageDiffThreshold {
				differs++
			}
		}
		if differs == len(result) {
			result = append(result, val)
		}
	}

	result = UniqStr(result)
	return result
}
