package main

import (
	"log"
	"os"
)

func main() {
	/*
		_, err := os.Open("nofile.txt")
		if err != nil {
			// print to stdout
			// fmt.Println("err happened", err)

			// log includes date and timestamp for log events
			// log.Println("error happened", err)
			log.Fatalln(err) // adds in a non zero exit and exits immediately
		}
	*/

	lf, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln("can't open log file: ", err)
	}
	defer lf.Close()
	log.SetOutput(lf)

	f2, err := os.Open("nofile.txt")
	if err != nil {
		log.Println("err happened", err) // now log will write to a file
	}
	defer f2.Close()
}
