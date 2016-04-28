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

// IsOrderByParamValid function
func IsOrderByParamValid(orderByQuery string) (string, bool) {
	orderBy := "-numberOfEpisodes"
	if orderByQuery != "" {
		if !(strings.Compare(orderByQuery, "-numberOfEpisodes") == 0 || strings.Compare(orderByQuery, "numberOfEpisodes") == 0) {
			return "", false
		}
		orderBy = orderByQuery
	}
	return orderBy, true
}
