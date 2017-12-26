package main

import (
	"log"
)

func init() {
	/*Log*/
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)


}
