package main

import(
	"os"
	"strconv"
	"strings"
	"math/big"
	"net/http"
	"crypto/rand" //what, did you expect math.Rand? I'm not stupid

	"github.com/charmbracelet/log"
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
	http.HandleFunc("/", httpHandler)

	log.Infof(
		"listening on port: %s", port)

	log.Fatal(
		http.ListenAndServe(":"+port, nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "mtd not allowed", http.StatusMethodNotAllowed)
		log.Warn(w, "req: %s  ; mtd not allowed!", r.URL.Path)
		return
	}

	lenStr := r.Header.Get("len")

	//if overflow attempt deny req
	if len(lenStr) >= 18 {
		//log it
		log.Warnf(
			"overflow attempt: ip %s\n",
			r.RemoteAddr)

		//felt it was mildly cringe, so I wrote it
		w.Write([]byte("Nice try moron.\n"))
		w.Write([]byte("Detected IP:  " + 
			r.RemoteAddr + "\n"))
		w.Write([]byte("Event logged.\n"))
		return
	}

	//if no len, default to 16
	if lenStr == "" {
		lenStr = "16"
	}

	switch (r.URL.Path) {
	case "/gen":
		//log req
		log.Infof("req: /gen  ;  len: %s", lenStr)
		gen(lenStr, w)
	case "/bld":
		charsRaw := r.Header.Get("chars")
		chars = strings.Split(charsRaw, "")

		//log req
		log.Infof("req: /bld  ;  len: %s  ;  chars: %s", lenStr, charsRaw)
		
		//gen
		gen(lenStr, w)
	}
}

func gen(lStr string, w http.ResponseWriter) {
	//get val len from header as Int64
	l, err := strconv.ParseInt(lStr, 10, 64)
	hanErr(err, w, "invalid number")
	
	//make sure they're not being a moron
	if l < 0 {
		w.Write([]byte(
			"Trade offer\n"+
			"  You give me:\n"+
			"    negative\n"+
			"  I give you:\n"+
			"    positive\n"))

		w.Write([]byte("Too bad I don't "+
			"know your answer; positive it "+
			"is\n"))

		//make positive
		l = -l
	}

	//prevent weirdos from trying to
	//  gen strings longer than
	//  56,527 chars
	if l > 56527 {
		//let client know
		w.Write([]byte("The hell do you "+
			"need a random string longer "+
			"than 56,527 characters for?\n"))

		//log it
		log.Warnf("%s%s%d\n",
			"uhhh, requested length longer ",
			"than 56,527 characters:  ", l)
		return
	}

	//gen and resp
	w.Write([]byte(genStr(l, w)))
}

func genStr(l int64, w http.ResponseWriter) string {
	var res string
	var i int64
	
	//gen ran nums and get val from
	//  char arr, then add to str
	for i = 0; i < l; i++ {
		ranDig := genInt(len(chars), w)
		res += chars[ranDig]
	}

	return res
}

func genInt(m int, w http.ResponseWriter) int {
	//conv max int to bit.Int
	bigInt := big.NewInt(int64(m))
	
	//actually gen num
	i, err := rand.Int(rand.Reader, bigInt)
	hanErr(err, w, "generating num")

	//return as prim int
	return int(i.Int64())
}


//if err != nil {...} solved
func hanErr(err error, w http.ResponseWriter, str string) {
	if err != nil {
		w.Write([]byte("server err, sorry\n"))
		w.Write([]byte("  reason: " + str + "\n"))
		log.Error(err)
	}
}
