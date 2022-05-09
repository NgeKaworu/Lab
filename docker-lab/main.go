package main

import (
	"flag"
	"log"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		addr = flag.String("l", ":8020", "绑定Host地址")
		m    = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
	)

	flag.Parse()
	log.Println(*addr, *m)

}
