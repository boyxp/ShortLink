package common

import "strings"

var index string = "ub04m3tGl8WSvfwj9sN5OEegdDQoLAJKYRaFBcPMh7I12HkzZi6VxyrnXUqTCp"
type Short struct {}

func (S Short) Hash(num int64) string {
	var res []byte
	num = num*12345
	for num>0 {
		res = append(res, index[num%62])
		num = num/62
	}

	return string(res)
}

func (S Short) Convert(hash string) int {
	var n int
	for i:=len(hash)-1;i>=0;i-- {
		p := strings.Index(index, string(hash[i]))
		if p==-1 {
			return 0
		}

		n = n*62 + p
	}

	return n/12345
}
