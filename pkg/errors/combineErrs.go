package errs

import (
	"fmt"
	"strings"
)

func Combine(errsArr []error) string {
	if len(errsArr) == 0 {
		return ""
	}

	setOfErrs := make(map[string]int)

	for _, err := range errsArr {
		if _, exists := setOfErrs[err.Error()]; !exists {
			setOfErrs[err.Error()] = 0
		}
		setOfErrs[err.Error()]++
	}

	var result strings.Builder
	result.WriteString("errors(count): ")
	for errMsg, qnt := range setOfErrs {
		result.WriteString(fmt.Sprintf("%s(%d);", errMsg, qnt))
	}

	finalMsg := result.String()
	if len(finalMsg) > 0 {
		finalMsg = finalMsg[:len(finalMsg)-1]
	}

	return finalMsg
}
