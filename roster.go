package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var candidates = flag.String("c", "", "List of candidates")
var prioritized = flag.String("p", "", "Prioritized candidates")

func main(){
	flag.Parse()
	c := strings.Split(*candidates, ",")
	p := strings.Split(*prioritized, ",")

	fmt.Println("*Hockey PROs selected (Candidates randomized, Prioritized prioritized):*")

	prioLength := 0 //to push starting ID of the candidates behind the list of prioritized
	if *prioritized != "" {
  	for n,i := range(p){
			fmt.Println(n+1,i)
		}
		prioLength = len(p)
	}

	if *candidates != "" {
		// Randomize supplied candidates
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
		//fmt.Println("Random order:",c)

		for n,i := range(c){
			fmt.Println(prioLength+n+1,i)
		}
	}
}
