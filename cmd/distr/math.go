package main

import (
	"flag"
	"github.com/digital-technology-agency/math"
	"log"
	"strconv"
)

func Int64(input string) int64 {
	result, _ := strconv.ParseInt(input, 10, 64)
	return result
}

func main() {
	n := flag.String(`n`, "1000", ``)
	k := flag.String(`k`, "2", ``)
	flag.Parse()
	log.Printf(`Params N=%s K=%s`, *n, *k)
	log.Printf(`Uint64 %d`, math.CnkUint(Int64(*n), Int64(*k)))
	log.Printf(`String %s`, math.CnkStr(Int64(*n), Int64(*k)))
}
