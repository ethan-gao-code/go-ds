# Bloom Filters

## Introduction

In computing, a bloom filter is a space-efficient probabilistic data structure, which is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set". Elements can be added to the set, but not removed (though this can be addressed with the counting Bloom filter variant); the more items added, the larger the probability of false positives.

## Reference

Wiki: https://en.wikipedia.org/wiki/Bloom_filter