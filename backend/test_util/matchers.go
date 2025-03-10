package test_util

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/stretchr/testify/mock"
)

// SQLMatcher creates a custom matcher that normalizes SQL queries
func SQLMatcher(expectedSQL string) interface{} {
	return mock.MatchedBy(func(actualSQL string) bool {
		normalized1 := normalizeSQL(expectedSQL)
		normalized2 := normalizeSQL(actualSQL)

		if normalized1 != normalized2 {
			fmt.Printf("\nSQL Comparison Failed\nExpected: %s\nActual: %s\nNormalized Expected: %s\nNormalized Actual: %s\n",
				expectedSQL, actualSQL, normalized1, normalized2)
			return false
		}
		return true
	})
}

func normalizeSQL(sql string) string {
	// Replace newlines, tabs, and multiple spaces with single spaces
	spaceRegex := regexp.MustCompile(`[\n\t\r ]+`)
	normalized := spaceRegex.ReplaceAllString(sql, " ")
	// Trim leading/trailing spaces
	return strings.TrimSpace(normalized)
}
