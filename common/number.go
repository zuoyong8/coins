package common 

import (
	"strconv"
	"strings"
	"math"
)


//十六转进制十进制
func HexDec(h string) (n int64) {
	if h[:2]=="0x"{
		h = h[2:] 
	}
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
	   c := s[i]
	   if v, ok := hex[c]; ok {
		  c = v
	   }
	   f, err := strconv.ParseFloat(c, 10)
	   if err != nil {
		  return -1
	   }
	   d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
 }
