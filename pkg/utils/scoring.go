package utils

import (
	"math"
	"strings"
	"unicode"
)

// CalculateScore computes similarity score between answer and correct text.
// Returns score in [0,1] and isCorrect based on threshold from config.
func CalculateScore(answer, correct string, threshold float64) (score float64, isCorrect bool) {
	// 1. Edge case: empty answer
	if strings.TrimSpace(answer) == "" {
		return 0, false
	}
	// 2. Normalize: lowercase, remove punctuation (keep apostrophes)
	a := normalizeText(answer)
	c := normalizeText(correct)

	// 3. Levenshtein distance → score [0, 1]
	dist := levenshteinDistance(a, c)
	ar := []rune(a)
	cr := []rune(c)
	maxLen := math.Max(float64(len(ar)), float64(len(cr)))
	if maxLen == 0 {
		return 1, true
	}
	score = 1 - float64(dist)/maxLen
	if score < 0 {
		score = 0
	}
	// 4. Threshold read from config
	isCorrect = score >= threshold
	return
}

// normalizeText lowercases and removes punctuation (keeps apostrophe for contractions)
func normalizeText(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'' || r == ' ' {
			b.WriteRune(r)
		}
	}
	return strings.TrimSpace(b.String())
}

// levenshteinDistance computes the edit distance between two strings (rune-aware)
func levenshteinDistance(a, b string) int {
	ar := []rune(a)
	br := []rune(b)
	la, lb := len(ar), len(br)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}
	prev := make([]int, lb+1)
	curr := make([]int, lb+1)
	for j := 0; j <= lb; j++ {
		prev[j] = j
	}
	for i := 1; i <= la; i++ {
		curr[0] = i
		for j := 1; j <= lb; j++ {
			cost := 1
			if ar[i-1] == br[j-1] {
				cost = 0
			}
			curr[j] = minOf3(curr[j-1]+1, prev[j]+1, prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}
	return prev[lb]
}

func minOf3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
