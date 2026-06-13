# Templates_GoLang

<p align="center">
  <a href="https://pkg.go.dev/github.com/roonil03/templates_golang"><img src="https://pkg.go.dev/badge/github.com/roonil03/templates_golang.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/roonil03/templates_golang"><img src="https://goreportcard.com/badge/github.com/roonil03/templates_golang" alt="Go Report Card"></a>
  <a href="https://github.com/Roonil03/Templates_GoLang/blob/main/LICENSE"><img src="https://img.shields.io/github/license/Roonil03/Templates_GoLang" alt="License"></a>
  <a href="https://github.com/Roonil03/Templates_GoLang/releases"><img src="https://img.shields.io/github/v/release/Roonil03/Templates_GoLang?include_prereleases&style=flat-square" alt="Latest Release"></a>
</p>

To navigate or import the specific components of this repository directly, utilize the primary reference index below. Each structure links directly to its source implementation.

### Repository Navigation Index
- **Core Collections & Utilities**
  - [Deque Implementation](./deque/deque.go)
  - [LRU Cache Implementation](./lru_cache/lru_cache.go)
  - [Priority Queue / Heap Implementation](./heap/heap.go)
- **Advanced Tree Structures**
  - [Segment Tree Implementation](./segment_tree/segment_tree.go)
  - [Fenwick Tree (1D BIT) Implementation](./fenwick/fenwick.go)
  - [2D Fenwick Tree Implementation](./fenwick2d/fenwick2d.go)
  - [Treap Implementation](./treap/treap.go)
- **Graph & Advanced Relations**
  - [Disjoint Set Union (DSU) Implementation](./union_find/union_find.go)
  - [Rope Data Structure](./rope/rope.go)
  - [Euler Tour Implementation](./euler_tour/euler_tour.go)
- **Backward Compatibility**
  - [Legacy Modules (Go <= 1.20)](./legacy/)

---

## Repository Metadata & Verification

| Property | Value / Status | Description |
| :--- | :--- | :--- |
| **Module Path** | `github.com/Roonil03/Templates_GoLang` | Canonical Go module import string |
| **Go Compliance** | `>= go 1.21` | Supports generic type parameters and `cmp.Ordered` |
| **Legacy Support** | `go 1.20` and below | Contained inside isolated local subdirectory |
| **Ecosystem Status** | Stable / Production-Ready | Verified algorithms optimized for compilation footprint |
| **Stability Level** | v1.x.x Series | Strict SemVer adherence across structural interfaces |

A formal, high-performance library of advanced data structures optimized for competitive programming and complex algorithmic tasks in Go. Every root structure is fully type-safe, leveraging Go Generics to allow flexible map-like and collection-level templating.

## Algorithmic Complexity Reference

| Structure | Time Complexity | Space Complexity |
| :--- | :--- | :--- |
| **Deque** | $O(1)$ push/pop | $O(n)$ |
| **Segment Tree** | $O(\log n)$ query/update | $O(n)$ |
| **Fenwick Tree (BIT)** | $O(\log n)$ query/update | $O(n)$ |
| **2D Fenwick Tree** | $O(\log^2 n)$ query/update | $O(n^2)$ |
| **Treap** | $O(\log n)$ expected | $O(n)$ |
| **Disjoint Set Union (DSU)** | $O(\alpha(n))$ amortized | $O(n)$ |
| **Rope** | $O(\log n)$ split/concat | $O(n)$ |
| **Euler Tour** | $O(1)$ ancestor check | $O(n)$ |
| **LRU Cache** | $O(1)$ get/put | $O(\text{capacity})$ |
| **Priority Queue / Heap** | $O(\log n)$ push/pop | $O(n)$ |

## Competitive Programming Integration

### Standard Import
For local modular development or automated build test runners:
```bash
go get github.com/Roonil03/Templates_GoLang/segment_tree
```

### Direct Submission Strategy (Online Judges)

For platforms requiring a unified source submission file (e.g., Codeforces, LeetCode, CodeChef), copy the targeted structural definition blocks and associated pointer receiver methods directly into your main compilation file. The removal of unnecessary boilerplate ensures compatibility with code limits.

```go
package main

import (
	"cmp"
	"fmt"
)

// Example copy-paste deployment for an Online Judge
type PriorityQueue[T cmp.Ordered] struct {
	data []T
}

func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{data: make([]T, 0)}
}

func main() {
	pq := NewPriorityQueue[int]()
	// Competitive execution loop goes here
	fmt.Println(pq)
}
```
