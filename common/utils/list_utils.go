package utils

import "github.com/lib/pq"

func ConvertToPQArray(arr []string) pq.StringArray {
	r := make(pq.StringArray, len(arr))
	for _, t := range arr {
		r = append(r, t)
	}

	return r
}
