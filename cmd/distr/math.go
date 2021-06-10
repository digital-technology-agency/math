package main

import (
	"flag"
	"github.com/digital-technology-agency/math"
	"log"
)

func main() {
	n := *flag.Int64(`n`, 1000, ``)
	k := *flag.Int64(`k`, 2, ``)
	flag.Parse()
	log.Printf(`Params N=%d K=%d`, n, k)
	log.Printf(`Uint64 %d`, math.CnkUint(n, k))
	log.Printf(`String %s`, math.CnkStr(n, k))
}
