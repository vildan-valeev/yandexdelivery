package main

import (
	"fmt"
	"github.com/vildan-valeev/yandexlogistic"
)

func main() {
	cl := yandexlogistic.NewYandexClient("")
	fmt.Println(cl)
}
