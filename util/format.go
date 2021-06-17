package util

import "strings"

// EndCryptTel utils
func EndCryptTel(tel string) string {
	var enTel string
	for i, v := range tel {
		if i < 3 {
			enTel += "*"
		} else if i == len(tel)-2 || i == len(tel)-1 {
			enTel += "*"
		} else {
			enTel += string(v)
		}
	}
	return enTel
}

// Remove zero first
func RemoveZeroFirst(tel string) string {
	arr := strings.Split(tel, "")
	if arr[0] == "0" {
		copy(arr[0:], arr[0+1:])
		arr[len(arr)-1] = ""
		arr = arr[:len(arr)-1]
		tel = strings.Join(arr, "")
	}
	return tel
}
