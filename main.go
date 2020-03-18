package main

import (
	"log"

	"github.com/lzyyauto/auto4play/router"
)

func main() {
	err := router.Init()
	// defer mysql
	if err != nil {
		log.Panic("启动失败")
	}
}
