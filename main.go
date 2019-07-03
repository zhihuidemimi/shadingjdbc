// helloworld project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter your password:")
	var password string
	//fmt.Scanf("%s", &password)
	fmt.scan("*", &password)
	fmt.Printf("your password is: " + password)
	os.Exit(0)
}
