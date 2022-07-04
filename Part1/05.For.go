package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	} else {
		//条件不需要括号
		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
		return result
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//只有结束条件，省略分号，相当于while，golang没有while。
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	//没有任何条件，死循环。
	for {
		fmt.Println(233)
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(10),
		convertToBin(233333),
		convertToBin(0))
	printFile("test.txt")
	//forever()
}
