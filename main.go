package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		log.Fatalln("No start time defined. Define as hh:mm")
	}
	s := strings.Split(args[1], ":")
	if len(s) != 2 {
		log.Fatalf("%s is not formatted correctly.", args[1])
	}

	sh, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatalf("%s is not a valid hour", s[0])
	}
	sm, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatalf("%s is not a valid minute", s[1])
	}

	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), sh, sm, 0, 0, now.Location())

	diff := time.Since(start)
	dh := int(diff.Minutes() / 60)
	dm := int(diff.Minutes()) - dh*60
	fmt.Printf("Time passed: %02d:%02d\n", dh, dm)
}
