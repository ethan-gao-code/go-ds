[![Go Report Card](https://goreportcard.com/badge/github.com/ethan-gao-code/go-ds)](https://goreportcard.com/report/github.com/ethan-gao-code/go-ds)
[![build-and-test](https://github.com/ethan-gao-code/go-ds/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/ethan-gao-code/go-ds/actions/workflows/build-and-test.yml)
[![codecov](https://codecov.io/gh/ethan-gao-code/go-ds/graph/badge.svg?token=RR7ZSMPTR1)](https://codecov.io/gh/ethan-gao-code/go-ds)
[![release](https://img.shields.io/github/v/release/ethan-gao-code/go-ds)](https://github.com/ethan-gao-code/go-ds/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

# Go-DS (Go Data Structures)

Implementation of various data structures in Go.

`go-ds` is an open-source Go library that provides implementations of various data structures. This library aims to simplify the usage and understanding of common data structures in Go for developers and is designed to be extensible and efficient.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation
To install the library, you can use `go get`:
```shell
go get github.com/ethan-gao-code/go-ds
```

## Usage
Below are examples of how to use the various data structures in this library. You can find more detailed examples in the [examples folder](https://github.com/ethan-gao-code/go-ds/tree/main/examples).

### Bloom Filters
A Bloom filter is a space-efficient probabilistic data structure, that is used to test whether an element is a member of a set.

For more details and usages about Bloom Filters, you can read [Bloom Filter README](https://github.com/ethan-gao-code/go-ds/bloomfilters/README.md)
```go
package main

import (
	"fmt"

	"github.com/ethan-gao-code/go-ds/bloomfilters"
)

func main() {
	bf := bloomfilters.New(0.01, 1000) // Create a new BloomFilter with a desired false positive rate of 0.01 and expected 1000 items
	
	bf.Add("apple") // Use Add to add element into the Bloom Filter
	
	bf.Contains("apple") // Contains can check if some elements are in the Bloom Filter. It returns true or false
	
	bf.Size() // Size return the size of the Bloom Filter (number of bits set to 1)

	bf.HashCount() // HashCount gets the number of hash functions used by the Bloom Filter

	bf.Reset() // Reset the Bloom Filter
}
```

### Doubly Linked List (WIP)

### Skip List (WIP)

### Queue (WIP)

### Sets (WIP)

### Stack (WIP)

### AVL Tree (WIP)

## Contributing
We encourage you to use `go-ds` in your projects and share your feedback with us. Your input will help guide future improvements and feature additions to the library.

We also welcome contributions! If you'd like to contribute to the development of `go-ds`, please fork the repository and submit a pull request. Be sure to follow the coding standards and include tests where applicable.

## License
`go-ds` is open source and available under the MIT License. See the [LICENSE](https://github.com/ethan-gao-code/go-ds/blob/main/LICENSE) file for more information.
