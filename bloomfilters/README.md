# Bloom Filters

## Introduction

In computing, a bloom filter is a space-efficient probabilistic data structure, which is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set". Elements can be added to the set, but not removed (though this can be addressed with the counting Bloom filter variant); the more items added, the larger the probability of false positives.

## Usage

When you would like to init a bloom filter, you need to know how many elements you have or will have, and what is the false positive rate you are willing to tolerate. Generally speaking, the more elements you are expected, the more memory is required. Similarly, the lower false positive you want, the more memory you will use.

Here is an example about creating a new bloom filter with a desired false positive rate of 1% and expected 1000 items

```go
bf := bloomfilters.New(0.01, 1000)
```

## Reference

Wiki: https://en.wikipedia.org/wiki/Bloom_filter