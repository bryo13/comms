package main

import (
	_ "embed"
	"log"
	"t_t/pkg"
)

func init() {
	if err := pkg.WriteIP(); err != nil {
		log.Fatalln(err)
	}

}

// for sharing purposes, we include our assests
// to people we trust

// //go:embed .assets/access.txt
// // var ipToGive []byte

func main() {
	pkg.InitUser()
	pkg.HandleConn()
}
