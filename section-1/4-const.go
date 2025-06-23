package main

import "fmt"

const STATUS_OK = 200
const (
	STATUS_NOT_FOUND    = 404
	STATUS_SERVER_ERROR = 500
)

func main() {
	fmt.Println("OK : ", STATUS_OK)
	fmt.Println("NOT FOUND : ", STATUS_NOT_FOUND)
	fmt.Println("SERVER ERROR : ", STATUS_SERVER_ERROR)
}
