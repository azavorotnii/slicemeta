package stringutil

// Generated by slicemeta-0.0.1a (2018-01-05T20:27:51-08:00)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}