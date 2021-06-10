# Mathematical experiments

# Intro
[![Go](https://github.com/digital-technology-agency/math/actions/workflows/go.yml/badge.svg)](https://github.com/digital-technology-agency/math/actions/workflows/go.yml)
[![License](http://img.shields.io/badge/Licence-MIT-brightgreen.svg)](LICENSE)
[![Website dta.agency](https://img.shields.io/website-up-down-green-red/http/shields.io.svg)](https://dta.agency)
[![GitHub release](https://img.shields.io/github/v/release/digital-technology-agency/math)](https://github.com/digital-technology-agency/math/releases/latest)

---
![Card information](./pic/c-from-n-to-k.png)

Let there be **n** different objects. To find the number of combinations of **n** objects by **k**, we will choose combinations of **m** objects in all possible ways, while paying attention to the different composition of the combinations, but not the order (it is not important here, unlike the placements).


### Issues
[![GitHub issues](https://img.shields.io/github/issues/digital-technology-agency/math?color=0A0ECD)](https://github.com/digital-technology-agency/math/issues?q=is%3Aopen+is%3Aissue)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/digital-technology-agency/math?style=flat)](https://github.com/digital-technology-agency/math/issues?q=is%3Aissue+is%3Aclosed)


# Usage
### [golang/cmd/go](https://golang.org/cmd/go/)

```bash
go get github.com/digital-technology-agency/math
```

## Quickstart

```go

import "github.com/digital-technology-agency/math"

func Example() {
    n := 100
    k := 3
    value := math.CnkUint(n, k)
	fmt.Printf("Result: %d\n", value)
    /*
        Result: 161700    
    */ 
}
```


## Contributing
Pull requests and Github issues are welcome.  Please read our [contributing guide](./CONTRIBUTING.md) for more information.
