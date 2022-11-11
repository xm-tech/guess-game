package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

// 胜利的组合
type winTuple [][]string

// 胜利组合是否包含该输入
func (self *winTuple) has(tuple []string) bool {
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

func from_prompt() (string, error) {
	prompt := promptui.Prompt{
		Label: "请出拳[剪刀/石头/布]",
		Validate: func(s string) error {
			s = strings.TrimSpace(s)
			if !toGuess.contains(s) {
				fmt.Println("invalid input:", s)
				return errors.New(fmt.Sprintf("Invalid input:%v", s))
			}

			return nil
		},
	}
	u, err := prompt.Run()
	return u, err
}

func from_input() string {
	var input string
	fmt.Scanf("%s\n", &input)
	if !toGuess.contains(input) {
		return ""
	}
	return input
}

// 出拳集合
var toGuess set

// 获胜的出拳组合
var wins winTuple

func init() {
	toGuess = set([]string{"剪刀", "石头", "布"})
	wins = winTuple([][]string{{"剪刀", "布"}, {"布", "石头"}, {"石头", "剪刀"}})
}

func main() {

	for {
		// u, _ := from_prompt()
		u := from_input()
		if u == "" {
			fmt.Println("wrong input:", u)
			continue
		}

		fmt.Println("u: ", u)

		time.Sleep(time.Second * 1)

		// 电脑随机
		ai := toGuess[rand.Intn(len(toGuess))]
		fmt.Println("ai: ", ai)
		tuple := []string{ai, u}

		time.Sleep(time.Second * 1)

		if ai == u {
			fmt.Println("draw")
		} else if wins.has(tuple) {
			fmt.Println("ai win")
		} else {
			fmt.Println("u win")
		}

		break
	}
}
