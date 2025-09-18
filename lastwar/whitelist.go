package lastwar

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"os"
	"strings"
)

const WHITELIST_FITLER_ALL = "all"
const WHITELIST_FITLER_WHITELIST = "whitelist"
const WHITELIST_FITLER_BLACKLIST = "blacklist"

func GetArrayOfAlliances() []string {
	delimiter := ","
	array := strings.Split(os.Getenv("ALLIANCES"), delimiter)

	// Trim any leading/trailing whitespace from each element (optional)
	for i, item := range array {
		array[i] = strings.TrimSpace(item)
	}

	return array
}

func IsAllianceOfInterest(ocrOutput string) bool {
	// Extract potential tag from OCR output
	tag := extractPotentialTag(ocrOutput)
	if tag == "" {
		return false
	}

	fmt.Printf("tag %s \n", tag)

	// Compare with whitelisted alliances using fuzzy matching
	for _, alliance := range GetArrayOfAlliances() {
		alliance = strings.ToLower(alliance)
		targetTag := "[" + alliance + "]"

		// Exact match
		if tag == targetTag {
			return true
		}

		// Fuzzy match with Levenshtein distance
		distance := levenshtein.ComputeDistance(tag, targetTag)
		similarity := 1 - float64(distance)/float64(max(len(tag), len(targetTag)))

		// Adjust threshold based on your OCR accuracy
		if similarity >= 0.8 { // 80% similarity
			return true
		}
	}

	return false
}

func extractPotentialTag(input string) string {
	if strings.HasPrefix(input, "I") {
		input = "[" + input[1:]
	}

	if strings.HasPrefix(input, "wopx") {
		input = "[" + input
	}

	input = strings.ToLower(input)

	// Look for bracket patterns
	start := strings.Index(input, "[")
	end := strings.Index(input, "]")

	if start == -1 || end == -1 || end <= start {
		return ""
	}

	tag := input[start : end+1]

	if len(tag) < 3 { // Minimum: [A]
		return ""
	}

	return tag
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
