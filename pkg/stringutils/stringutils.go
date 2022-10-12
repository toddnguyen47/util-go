package stringutils

import "unicode"

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
