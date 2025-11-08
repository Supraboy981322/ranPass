package main

import(
	"os"
	"log"
	"fmt"
	"strconv"
	"crypto/rand"
	"math/big"
)

var (
	args = os.Args[1:]
	chars = []string{
		"a", "b",	"c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z", "A", "B",
		"C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z", "0", "9", "8", "7",
		"6", "5", "4", "3", "2", "1", "!",
		"@", "#", "$", "%", "^", "&", "*",
		"(", ")", "-", "_", "=", "+", "[",
		"]", "{", "}", "|", "\\", ";", ":",
		"'", "\"", "<", ">", "/", "?", ".",
		",",
	}
	res string
)

func main() {
	fmt.Println(len(chars))
	l, err := strconv.Atoi(args[0])
	hanErr(err)
	for i := 0; i < l; i++ {
		ranDig := genInt(len(chars))
		res += chars[ranDig]
	}
	fmt.Println(res)
}

func genInt(len int) int {
	bigInt := big.NewInt(int64(len))
	i, err := rand.Int(rand.Reader, bigInt)
	hanErr(err)
	return int(i.Int64())
}

func hanErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
