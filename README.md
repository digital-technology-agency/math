# Mathematical and algorithms experiments 

<p><img src="https://rawcdn.githack.com/devicons/devicon/9c6bfdb9783cdfe1018666ed76adcfd3eab6fad6/icons/go/go-original.svg" alt="go" width="100" height="100"/></p>

# Intro
[![GoDoc](https://godoc.org/github.com/digital-technology-agency/math?status.svg)](https://godoc.org/github.com/digital-technology-agency/math)
[![Go](https://github.com/digital-technology-agency/math/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/digital-technology-agency/math/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/digital-technology-agency/math)](https://goreportcard.com/report/github.com/digital-technology-agency/math)
[![License](http://img.shields.io/badge/Licence-MIT-brightgreen.svg)](LICENSE)
[![Website dta.agency](https://img.shields.io/website-up-down-green-red/http/shields.io.svg)](https://dta.agency)
[![GitHub release](https://img.shields.io/github/v/release/digital-technology-agency/math)](https://github.com/digital-technology-agency/math/releases/latest)
<img alt="open collective badge" src="https://opencollective.com/dta-agency/tiers/backers/badge.svg?label=backer&color=brightgreen" />

<p>&nbsp;<img align="center" src="https://github-readme-stats.vercel.app/api?username=digitaltechnologyagency&show_icons=true&count_private=true" alt="Агентство цифровых технологий" width="80%"/></p>


## Menu
`Sorted`
* [Sort insertion](https://github.com/digital-technology-agency/math/wiki/Sort-insertion)
* [Merge sort](https://github.com/digital-technology-agency/math/wiki/Merge-sort)

`Mathematical`

* [C from n to k](https://github.com/digital-technology-agency/math/wiki/C-from-n-to-k)
* [Factorial](https://github.com/digital-technology-agency/math/wiki/Factorial)
* [Permutation](https://github.com/digital-technology-agency/math/wiki/Permutation)
* [Placements](https://github.com/digital-technology-agency/math/wiki/Placements)


## Examples

### C from n to k

![Card information](./pic/c-from-n-to-k.png)

Let there be **n** different objects. To find the number of combinations of **n** objects by **k**, we will choose combinations of **m** objects in all possible ways, while paying attention to the different composition of the combinations, but not the order (it is not important here, unlike the placements).

### Factorial 

![Card information](./pic/factorial.svg)

The factorial operation is encountered in many areas of mathematics, notably in combinatorics, algebra, and mathematical analysis. Its most basic use counts the possible distinct sequences – the permutations – of **n** distinct objects: there are **n!**.

### Permutations (n)

### _**Pn** = **n!**_

Permutations differ from combinations, which are selections of some members of a set regardless of order. For example, written as tuples, there are six permutations of the set {1,2,3}, namely: (1,2,3), (1,3,2), (2,1,3), (2,3,1), (3,1,2), and (3,2,1). These are all the possible orderings of this three-element set. Anagrams of words whose letters are different are also permutations: the letters are already ordered in the original word, and the anagram is a reordering of the letters. The study of permutations of finite sets is an important topic in the fields of combinatorics and group theory.


### Placements (n)

![Card information](./pic/placements.jpg)

If you are already familiar with combinations, then you will easily notice that in order to find placements, you need to take all possible combinations, and then change the order in each one in all possible ways (that is, in fact, make more **permutations**). Therefore, the number of placements is also expressed in terms of the number of permutations and combinations


### Issues
[![GitHub issues](https://img.shields.io/github/issues/digital-technology-agency/math?color=0A0ECD)](https://github.com/digital-technology-agency/math/issues?q=is%3Aopen+is%3Aissue)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/digital-technology-agency/math?style=flat)](https://github.com/digital-technology-agency/math/issues?q=is%3Aissue+is%3Aclosed)


## Usage
### [golang/cmd/go](https://golang.org/cmd/go/)

```bash
go get github.com/digital-technology-agency/math/pkg/combinatorics
```

## Test

```bash
go test -run ''
```


## Quickstart

**C from n to k**

```go

import "github.com/digital-technology-agency/math/pkg/combinatorics"

func Example() {
    n := 100
    k := 3
    value := combinatorics.CnkUint(n, k)
	fmt.Printf("Result: %d", value)
    /*
        Result: 161700    
    */ 
}
```

**Factorial (n!)**

```go

import "github.com/digital-technology-agency/math/pkg/combinatorics"

func Example() {
    n := 4
    value := combinatorics.FactorialInt(n)
	fmt.Printf("Result: %d", value)
    /*
        Result: 24    
    */ 
}
```

**Permutations (n)**

```go

import "github.com/digital-technology-agency/math/pkg/combinatorics"

func Example() {
    n := 4
    value := combinatorics.PermutationsInt(n)
	fmt.Printf("Result: %d", value)
    /*
        Result: 24    
    */ 
}
```

**Placements (A from n to k)**

```go

import "github.com/digital-technology-agency/math/pkg/combinatorics"

func Example() {
    n := 10
    k := 10
    value := combinatorics.PlacementInt(n,k)
	fmt.Printf("Result: %d", value)
    /*
        Result: 3628800    
    */ 
}
```

## Run in Terminal

* Go to main page and click **Releases**

![Card information](./pic/cli-1.png)

* Choose a version for your operating system and click the link

![Card information](./pic/cli-2.png)

* Download and unzip the **math** file

![Card information](./pic/cli-3.png)
![Card information](./pic/cli-4.png)

* Open terminal in unzip folder
 
![Card information](./pic/cli-5.png)

* Run unzip binary file  **math** and show list **types**

```bash
 ./math
```

![Card information](./pic/cli-6.png)

* Comand argument list 

```bash
 ./math -h
```

![Card information](./pic/cli-7.png)

* Set arguments
```bash
 ./math -n=10000 -k=10
```

## Contributing
Pull requests and Github issues are welcome.  Please read our [contributing guide](./CONTRIBUTING.md) for more information.

[Агентство цифровых технологий](https://dta.agency/)
