package helper

import "strings"

func CheckMandatoryHeaders(source []string, target map[string][]string) bool {
	for _, value := range source {
		found := false
		for key, _ := range target {
			if strings.ToLower(value) == strings.ToLower(key) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
