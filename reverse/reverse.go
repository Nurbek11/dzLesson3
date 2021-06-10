package dz3Lesson

func Reverse(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	return result
}
