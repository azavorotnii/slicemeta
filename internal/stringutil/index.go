package stringutil

// Generated by slicemeta-0.0.1a (2018-01-05T20:27:51-08:00)

func Index(in []string, value string) int {
	for i, v := range in {
		if v == value {
			return i
		}
	}
	return -1
}

func IndexAny(in []string, values []string) int {
	for i, v := range in {
		for _, value := range values {
			if v == value {
				return i
			}
		}
	}
	return -1
}

func IndexFunc(in []string, f func(string) bool) int {
	for i, v := range in {
		if f(v) {
			return i
		}
	}
	return -1
}