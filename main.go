package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 胜利的组合
type winGroups [][]string

// 胜利组合是否包含该输入
func (self *winGroups) has(tuple []string) bool {
	for _, item := range [][]string(*self) {
		if item[0] == tuple[0] && item[1] == tuple[1] {
			return true
		}
	}
	return false
}

type set []string

func (self *set) contains(s string) bool {
	aset := make(map[string]struct{}, len(*self))
	for _, v := range *self {
		aset[v] = struct{}{}
	}
	_, ok := aset[s]
	return ok
}

func main() {
	var u string
	guessList := set([]string{"剪刀", "石头", "布"})
	winGroups := winGroups([][]string{{"剪刀", "布"}, {"布", "石头"}, {"石头", "剪刀"}})

	for {
		// 解析玩家输入
		fmt.Scanf("%s\n", &u)
		if !guessList.contains(u) {
			fmt.Println("invalid input:", u)
			continue
		}

		// 电脑随机
		index := rand.Intn(len(guessList))
		ai := guessList[index]

		fmt.Println("u: ", u)

		time.Sleep(time.Second * 1)

		fmt.Println("ai: ", ai)
		tuple := []string{ai, u}

		time.Sleep(time.Second * 1)

		if ai == u {
			fmt.Println("draw")
		} else if winGroups.has(tuple) {
			fmt.Println("ai win")
		} else {
			fmt.Println("u win")
		}
	}
}
