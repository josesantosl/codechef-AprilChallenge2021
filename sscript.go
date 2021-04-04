package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k, t int
	var s, c string
	fmt.Scan(&t)
	for tt := 0; tt < t; tt++ {
		fmt.Scan(&n)
		fmt.Scan(&k)
		fmt.Scanln(&s)
		c = strings.Repeat("*", k)
		if strings.Contains(s, c) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
