package main

import (
	"log"

	"./router"
)

func main() {
	err := router.Init()
	// defer mysql
	if err != nil {
		log.Panic("启动失败")
	}
}
