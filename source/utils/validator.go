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
func IsOrderByParamValid(orderByQueries []string, validOrderBys []string) []string {
	var orderBy []string
	for _, validOrderBy := range validOrderBys {
		for _, orderByQuery := range orderByQueries {
			if strings.Compare(validOrderBy, orderByQuery) == 0 {
				orderBy = append(orderBy, orderByQuery)
			}
		}
	}
	if len(orderBy) == 0 {
		orderBy = validOrderBys[0:1]
	}
	return orderBy
}
