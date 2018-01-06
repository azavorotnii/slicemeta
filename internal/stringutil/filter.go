package stringutil

// Generated by slicemeta-0.0.1a (2018-01-05T20:27:51-08:00)

func Filter(in []string, filter func(string) bool) []string {
	var result []string
	for _, v := range in {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}