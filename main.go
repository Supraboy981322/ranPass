package main

import(
	"os"
	"log"
	"strconv"
	"crypto/rand"
	"math/big"
	"net/http"
)

var (
	port = "2628"
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
)

func main() {
	http.HandleFunc("/gen", genHan)
	log.Printf("listening on port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func genHan(w http.ResponseWriter, r *http.Request) {
	if len(r.Header.Get("len")) >= 18 {
		log.Printf("overflow attempt: direct ip %s\n", r.RemoteAddr)
		w.Write([]byte("Nice try moron.\n"))
		w.Write([]byte("Detected IP:  " + r.RemoteAddr + "\n"))
	} else {
		l, err := strconv.ParseInt(r.Header.Get("len"), 10, 64)
		if l < 0 {
			l = -l
		}
		hanErr(err)
		log.Printf("req: /gen  ;  len: %d", l)
		w.Write([]byte(genStr(l)))
	}
}

func genStr(l int64) string {
	var res string
	var i int64
	for i = 0; i < l; i++ {
		ranDig := genInt(len(chars))
		res += chars[ranDig]
	}
	return res
}

func genInt(l int) int {
	bigInt := big.NewInt(int64(l))
	i, err := rand.Int(rand.Reader, bigInt)
	hanErr(err)
	return int(i.Int64())
}

func hanErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
