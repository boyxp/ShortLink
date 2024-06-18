package common

import "strings"

var index string = "ub04m3tGl8WSvfwj9sN5OEegdDQoLAJKYRaFBcPMh7I12HkzZi6VxyrnXUqTCp"
type Short struct {}

func (S Short) Hash(num int64) string {
	var res []byte
	var total int64
	num = num*12345
	for num>0 {
		mod := num%62
		res  = append(res, index[mod])
		num  = num/62
		total = total+mod
	}

	res = append(res, index[total%62])

	return string(res)
}

func (S Short) Convert(hash string) int {
	if len(hash)<4 {
		return -1
	}

	var n int
	var total int
	for i:=len(hash)-2;i>=0;i-- {
		p := strings.Index(index, string(hash[i]))
		if p==-1 {
			return -1
		}

		n = n*62 + p
		total = total+p
	}

	if index[total%62]!=hash[len(hash)-1] {
		return -1
	}

	return n/12345
}
