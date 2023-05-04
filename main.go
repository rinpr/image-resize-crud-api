package main

import (
	"fmt"
	"time"
	"github.com/rinpr/crud-api-golang/routes"
)

func main() {
	routes.ListenAndServe()
	fmt.Println(time.Now().Local())
}