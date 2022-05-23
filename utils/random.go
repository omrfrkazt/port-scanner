package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//init random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Create every file with name like "1234567890.txt"
func CreateFileName() string {
	var number string
	for i := 0; i < 10; i++ {
		number += strconv.Itoa(rand.Intn(10))
	}
	fmt.Println("qwe", number+".txt")
	return number + ".txt"
}
