# Bloom Filters

## Introduction

In computing, a bloom filter is a space-efficient probabilistic data structure, which is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set". Elements can be added to the set, but not removed (though this can be addressed with the counting Bloom filter variant); the more items added, the larger the probability of false positives.

## Usage

When you would like to init a bloom filter, you need to know how many elements you have or will have, and what is the false positive rate you are willing to tolerate. Generally speaking, the more elements you are expected, the more memory is required. Similarly, the lower false positive you want, the more memory you will use.

Here is an example about creating a new bloom filter with a desired false positive rate of 1% and expected 1000 items

```go
bf := bloomfilters.New(0.01, 1000)
```

Then you can use `Add` function to add some elements to the Bloom Filter

```go
bf.Add("apple")
bf.Add("banana")
bf.Add("cherry")
```

`Contains` can check if an item might exist in the Bloom Filter.

```go
bf.Contains("apple")
```

`Size` returns the number of bits set to 1 in the bitmap.

```go
bf.Size()
```

`HashCount` HashCount returns the number of hash functions used by the Bloom Filter.

```go
bf.HashCount()
```

`Values` returns the indices of all bits set in the bitmap (for debugging or analysis purposes).

```go
bf.Values()
```

`Reset` can clear all bits in the Bloom Filter bitmap.

```go
bf.Reset()
```

## Reference

Wiki: https://en.wikipedia.org/wiki/Bloom_filter