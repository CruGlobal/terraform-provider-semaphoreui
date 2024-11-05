package provider

import (
	"fmt"
	"regexp"
	"strconv"
)

func parseImportFields(input string, requiredFields []string) (map[string]int64, error) {
	result := make(map[string]int64)
	re := regexp.MustCompile(`(\w+)/(\d+)/?`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) != 3 {
			continue
		}
		value, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			return nil, err
		}
		result[match[1]] = value
	}

	for _, field := range requiredFields {
		if _, ok := result[field]; !ok {
			return nil, fmt.Errorf("missing required import field %s", field)
		}
	}

	return result, nil
}
