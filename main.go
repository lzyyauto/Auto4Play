package main

import (
	"log"

	"github.com/lzyyauto/Auto4Play/router"
)

func main() {
	err := router.Init()
	if err != nil {
		log.Panic("启动失败")
	}
}
