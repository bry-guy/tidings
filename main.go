package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Tiding struct {
	message string
	author  string
	host    string
}

type Options struct {
	id int
}

var idFlag *int
var options *Options

func init() {
	idFlag = flag.Int("id", 0, "specify tiding by id")

	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	if *idFlag > 0 {
		options = &Options{id: *idFlag}
	}

	hostname, _ := os.Hostname()

	tidings := []Tiding{
		{message: "A little bit of duplication beats a little bit of complexity.", author: "Brain", host: hostname},
		{message: "Don't communicate by sharing memory, share memory by communicating.", author: "Go Proverbs", host: hostname},
		{message: "Clear is better than clever.", author: "Go Proverbs", host: hostname},
		{message: "People will forget what you said, people will forget what you did, but people will never forget how you made them feel.", author: "Maya Angelou", host: hostname},
	}

	var tiding Tiding

	if options != nil {
		tiding = tidings[options.id]
	} else {
		tiding = tidings[rand.Intn(len(tidings))]
	}

	fmt.Println(tiding)
}

func usage() {
	fmt.Println("tidings\t<arg>")
}

func (t Tiding) String() string {
	return fmt.Sprintf("%v - %v", t.message, t.author)
}
