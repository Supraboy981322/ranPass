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
	//port (might make toml)
	port = "2628"

	//cli args
	args = os.Args[1:]

	//char pool
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

	log.Printf(
		"listening on port: %s\n", port)

	log.Fatal(
		http.ListenAndServe(":"+port, nil))
}

func genHan(w http.ResponseWriter, r *http.Request) {
	//if overflow attempt deny req
	if len(r.Header.Get("len")) >= 18 {
		log.Printf(
			"overflow attempt: ip %s\n",
			r.RemoteAddr)

		//felt it was mildly cringe, so I wrote it
		w.Write([]byte("Nice try moron.\n"))
		w.Write([]byte("Detected IP:  " + 
			r.RemoteAddr + "\n"))
		w.Write([]byte("Event logged.\n"))
	} else {
		//get val len from header 
		l, err := strconv.ParseInt(
			r.Header.Get("len"), 10, 64)
		hanErr(err)
		//make sure they're not being
		// a moron
		if l < 0 {
			w.Write([]byte(
				"Trade offer\n"+
				"  You give me:\n"+
				"    negative\n"+
				"  I give you\n"+
				"    positive\n"))
			w.Write([]byte("Too bad I don't "+
				"know your answer; positive it "+
				"is\n"))
			l = -l
		}
		if l > 56527 {
			w.Write([]byte("The hell do you need a random string longer than 56,527 characters for?\n"))
		} else {
			log.Printf("req: /gen  ;  len: %d", l)
			w.Write([]byte(genStr(l)))
		}
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
