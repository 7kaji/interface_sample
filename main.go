package main

import (
	"fmt"

	Crier "github.com/7kaji/interface_sample/interface"
	"github.com/7kaji/interface_sample/model"
)

func main() {
	// インターフェースの型を持つ変数に, インターフェースを満たす構造体は代入できる
	// (pochi と donald どちらの書き方でも可能)
	pochi := model.NewDog("Pochi")
	var donald Crier.Animal = model.NewDuck("Donald")

	taro := model.NewHuman("Taro")
	fmt.Println(taro.Touch(pochi))
	fmt.Println(taro.Touch(donald))
}
