package util

func GetLastItem(params ...interface{}) interface{} {
	s := make([]string, 0)
	for _, item := range params {
		s = append(s, item.(string))
	}
	return s[:len(s)-1]
}
