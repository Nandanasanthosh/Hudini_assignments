package kata

func GetMiddle(s string) string {
	str := ""
	n := len(s)
	middle := len(s) / 2
	if n%2 == 0 {
		str += string(s[middle-1]) + string(s[middle])
	} else {
		str += string(s[middle])
	}
	return str
}
