package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genRandom(n int) (ans string) {
	var str = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+=")
	var tempPass = make([]rune, n)

	for i := range tempPass {
		tempPass[i] = str[rand.Intn(len(str))]
	}

	ans = string(tempPass)

	return ans
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var tempPass = "%s-%s-%s-%s"
	var pass = fmt.Sprintf(tempPass, genRandom(4), genRandom(4), genRandom(4), genRandom(4))

	fmt.Printf("new password: %s \n", pass)
	time.Sleep(60 * time.Second)
}
