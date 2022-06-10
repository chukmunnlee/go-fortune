package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	FORTUNE              = "FORTUNE"
	PORT                 = "PORT"
	DEFAULT_FORTUNE_FILE = "./fortune.txt"
	DEFAULT_PORT         = 3000
)

func loadFortunes(path string) []string {
	buff, err := ioutil.ReadFile(path)
	if nil != err {
		log.Fatalf("Error reading %s: %v\n", path, err)
	}
	lines := strings.Split(string(buff), "\n")
	return lines[:len(lines)-1]
}

func defaultFortune() string {
	value, present := os.LookupEnv(FORTUNE)
	if present {
		return value
	}
	return DEFAULT_FORTUNE_FILE
}

func defaultPort() (int, error) {
	value, present := os.LookupEnv(PORT)
	if present {
		return strconv.Atoi(value)
	}
	return DEFAULT_PORT, nil
}

func getFortunes(fortune []string, count int) []string {
	idx := rand.Perm(len(fortune))[:count]
	f := make([]string, count)
	for i := 0; i < count; i++ {
		f[i] = fortune[idx[i]]
	}
	return f
}

func main() {

	var fortuneFile string
	var port int
	defPort, err := defaultPort()

	if nil != err {
		log.Fatalf("Error: %v", err)
	}

	flag.StringVar(&fortuneFile, "fortune", defaultFortune(), "Fortune file")
	flag.IntVar(&port, "port", defPort, "port")
	flag.Parse()

	log.Printf("fortune file: %s, port: %d", fortuneFile, port)

	fortunes := loadFortunes(fortuneFile)
	log.Printf("Loaded %s fortunes file\n", fortuneFile)

	rand.Seed(time.Now().UnixNano())

	getFortunes(fortunes, 3)

	//r := gin.Default()

}
