package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建空map， key是sting类型， value是int类型
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	scanResult := input.Scan()
	fmt.Println(scanResult)
	for scanResult {
		fmt.Println("+++++++++")
		counts[input.Text()]++
		fmt.Printf("%s---", input.Text())
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
