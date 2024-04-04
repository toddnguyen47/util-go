package stringutils

import (
	"strings"
	"unicode"
)

// IsBlank - Check if string is empty / consists of only whitespace
// Inspired by Apache's Java StringUtils.IsBlank:
// https://github.com/apache/commons-lang/blob/master/src/main/java/org/apache/commons/lang3/StringUtils.java
// IsBlank(nil)      = true
// IsBlank("")        = true
// IsBlank(" ")       = true
// IsBlank("bob")     = false
// IsBlank("  bob  ") = false
func IsBlank(strInput *string) bool {

	if strInput == nil || len(*strInput) == 0 {
		return true
	}

	for _, c1 := range *strInput {
		if !unicode.IsSpace(c1) {
			return false
		}
	}

	return true
}

func IsNotBlank(strInput *string) bool {
	return !IsBlank(strInput)
}

func MakePtr(strInput string) *string {
	return &strInput
}

func JoinExcludeEmpty(elems []string, sep string) string {
	sb := strings.Builder{}
	count := 0
	for _, elem := range elems {
		elemTrimmed := strings.TrimSpace(elem)
		if IsNotBlank(&elemTrimmed) {
			if count > 0 {
				sb.WriteString(sep)
			}
			sb.WriteString(elem)
			count += 1
		}
	}
	return sb.String()
}

// GetSubstring - safely get substring starting from zero index.
// The substring is [0, strIndex) where strIndex is excluded.
func GetSubstring(strInput string, strIndex int) string {
	if strIndex < 0 {
		return strInput
	}
	maxInt := len(strInput)
	if strIndex < maxInt {
		maxInt = strIndex
	}
	var sb strings.Builder
	for i := 0; i < maxInt; i++ {
		sb.WriteByte(strInput[i])
	}
	return sb.String()
}

// WriteToSbWithSep - write to strings.Builder with a separator.
func WriteToSbWithSep(sb *strings.Builder, str string, sep string) {
	if sb.Len() > 0 {
		sb.WriteString(sep)
	}
	sb.WriteString(str)
}

// GetUniqueListCaseSensitive - get unique list by comparing strings in a case-sensitive manner
func GetUniqueListCaseSensitive(inputList []string) []string {
	set1 := make(map[string]struct{})
	newList := make([]string, 0)

	for _, elem := range inputList {
		if _, ok := set1[elem]; !ok {
			set1[elem] = struct{}{}
			newList = append(newList, elem)
		}
	}

	return newList
}
