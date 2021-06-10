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
	typeValue := flag.String(`t`, "", "cnk\nfactorial\n")
	n := flag.String(`n`, "5", ``)
	k := flag.String(`k`, "2", ``)
	flag.Parse()
	switch *typeValue {
	default:
		flag.PrintDefaults()
		break
	case `cnk`:
		log.Printf(`Params N=%s K=%s`, *n, *k)
		log.Printf(`Uint64 [%d]`, math.CnkUint(Int64(*n), Int64(*k)))
		log.Printf(`String [%s]`, math.CnkStr(Int64(*n), Int64(*k)))
		break
	case `factorial`:
		log.Printf(`Params N=%s `, *n)
		log.Printf(`Int64 [%v]`, math.FactorialInt(Int64(*n)))
		log.Printf(`String [%s]`, math.FactorialString(Int64(*n)))
	}
}
