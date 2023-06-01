# Linear vs Binary Search Benchmark

Linear search(O(N)) vs Binary search O(log2N) on small size array less than 128.

## Benchmark Results

```
goos: linux
goarch: amd64
pkg: linear-vs-binary-search
cpu: AMD Ryzen 9 7950X 16-Core Processor
BenchmarkLinearSearch-32        49509034                24.23 ns/op
BenchmarkBinarySearch-32        34832584                34.37 ns/op
PASS
ok      linear-vs-binary-search 2.463s

```

What interesting here is linear search performs better because of  branch prediction and a predictable memory-access pattern.


https://en.wikipedia.org/wiki/Branch_predictor

https://en.wikipedia.org/wiki/Memory_access_pattern

https://par.nsf.gov/servlets/purl/10187649
