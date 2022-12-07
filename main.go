package main

import (
	"ChsSeltApi/models"
	"ChsSeltApi/routers"
	"fmt"
)

func main() {
	fmt.Println(models.Test1("aaaaa"))
	routers.Init()
}
