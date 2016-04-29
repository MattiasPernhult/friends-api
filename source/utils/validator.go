package utils

import (
	"strconv"
	"strings"
)

// IsLimitParamValid function
func IsLimitParamValid(limitQuery string) (int, bool) {
	limit := 25
	var convertErr error
	if limitQuery != "" {
		limit, convertErr = strconv.Atoi(limitQuery)
		if convertErr != nil || limit < 1 {
			return 0, false
		}
	}
	return limit, true
}

// IsIncludeParamValid function
func IsIncludeParamValid(includeQueries []string, validIncludes []string, include map[string]interface{}) map[string]interface{} {
	for _, validInclude := range validIncludes {
		for _, includeQuery := range includeQueries {
			if strings.Compare(validInclude, includeQuery) == 0 {
				include[includeQuery] = 1
			}
		}
	}
	return include
}

// IsOrderByParamValid function
func IsOrderByParamValid(orderByQuery string, validOrderBys []string) (string, bool) {
	if orderByQuery != "" {
		for _, validOrderBy := range validOrderBys {
			if strings.Compare(orderByQuery, validOrderBy) == 0 {
				return validOrderBy, true
			}
		}
		return "", false
	}
	return validOrderBys[0], true
}
