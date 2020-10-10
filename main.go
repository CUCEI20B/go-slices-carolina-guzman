package main

import "fmt"

func main()  {
	var n uint64
	var i uint64
	fmt.Scan(&n)
	var r int
	var numero int
	var s [] int
	for i= 0; i < n; i++ {
		fmt.Scan(&numero)
		s = append(s, numero)
		//fmt.Println(s[i])
		r += s[i]
	}
	fmt.Println(r)


}