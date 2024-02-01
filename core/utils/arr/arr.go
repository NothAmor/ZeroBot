package arr

// ContainsString 判断字符串切片中是否包含某个字符串
func ContainsString(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
