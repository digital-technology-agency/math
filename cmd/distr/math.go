package main

import (
	"flag"
	"github.com/digital-technology-agency/math/pkg/combinatorics"
	"log"
	"strconv"
)

func Int64(input string) int64 {
	result, _ := strconv.ParseInt(input, 10, 64)
	return result
}

func main() {
	typeValue := flag.String(`t`, "", "cnk\nfactorial\npermutations\n")
	n := flag.String(`n`, "5", ``)
	k := flag.String(`k`, "2", ``)
	flag.Parse()
	switch *typeValue {
	default:
		flag.PrintDefaults()
		break
	case `cnk`:
		log.Printf(`Params N=%s K=%s`, *n, *k)
		log.Printf(`Uint64 [%d]`, combinatorics.CnkUint(Int64(*n), Int64(*k)))
		log.Printf(`String [%s]`, combinatorics.CnkStr(Int64(*n), Int64(*k)))
		break
	case `factorial`:
		log.Printf(`Params N=%s `, *n)
		log.Printf(`Int64 [%v]`, combinatorics.FactorialInt(Int64(*n)))
		log.Printf(`String [%s]`, combinatorics.FactorialString(Int64(*n)))
	case `permutations`:
		log.Printf(`Params N=%s `, *n)
		log.Printf(`Int64 [%v]`, combinatorics.PermutationsInt(Int64(*n)))
		log.Printf(`String [%s]`, combinatorics.PermutationStr(Int64(*n)))
	}
}
