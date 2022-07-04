package main

import "fmt"

//map使用哈希表，必须可以比较相等。
//除了slice，map，function的内建类型都可以作为key。
//Struct类型不包含上述字段，也可作为key。

//寻找最长不含有重复字符的子串(leetcode)
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 >= maxlength {
			maxlength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxlength
}

func main() {
	//map无序
	m1 := map[string]string{
		"name":    "lkidszzzz",
		"courses": "golang",
		"site":    "github",
		"quality": "notbad",
	}
	fmt.Println(m1)
	//创建空map
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2, m3)

	fmt.Println("Traversing map")
	//使用range遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name, ok := m1["name"]
	fmt.Println(name, ok)
	//如果打错了k，则为nil,可以使用ok看是否存在。
	wrong, ok := m1["wrong"]
	fmt.Println(wrong, ok)
	if site, ok := m1["sites"]; ok {
		fmt.Println(site)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	quality, ok := m1["quality"]
	fmt.Println(quality, ok)
	delete(m1, "quality")
	fmt.Println(m1)

	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
}
