# Go Competitive Programming Templates

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
- **Advanced Query & Sequence Processing**
  - [Sparse Table Implementation](./sparse_table/sparse_table.go)
  - [Trie Implementation](./trie/trie.go)
  - [Policy-Based Data Structure (PBDS)](./pbds/pbds.go)
  - [Heavy-Light Decomposition (HLD)](./hld/hld.go)
  - [Suffix Array & LCP](./suffix_array/suffix_array.go)
- **Graph & Network Optimization**
  - [Dijkstra's Shortest Path](./dijkstra/dijkstra.go)
  - [Bellman-Ford Algorithm](./bellman_ford/bellman_ford.go)
  - [Floyd-Warshall All-Pairs](./floyd_warshall/floyd_warshall.go)
  - [Kruskal's Minimum Spanning Tree](./kruskal/kruskal.go)
  - [Prim's Minimum Spanning Tree](./prim/prim.go)
- **String Processing & Automata**
  - [KMP Matcher & Prefix Function](./kmp/kmp.go)
  - [Rabin-Karp Rolling Hash](./rabin_karp/rabin_karp.go)
  - [Suffix Automaton](./suffix_automaton/suffix_automaton.go)
- **Graph Flow & Combinatorial Topology**
  - [Tarjan's Strongly Connected Components](./tarjan_scc/tarjan.go)
  - [Dinic's Maxflow Algorithm](./maxflow/maxflow.go)
  - [Kahn's Topological Sort](./kahn/kahn.go)
- **Mathematical & Specialized Query Optimizations**
  - [Kadane's Maximum Subarray Sum](./kadane/kadane.go)
  - [Modular Arithmetic & Modular Power](./mod_arithmetic/mod_arith.go)
  - [Mo's Offline Block Decomposition](./mos_algorithm/mos.go)
- **Geometric Optimizations & Persistent Queries**
  - [Convex Hull Trick](./convex_hull_trick/cht.go)
  - [Persistent Segment Tree](./persistent_segment_tree/persistent_st.go)
  - [Lazy Segment Tree](./lazy_segment_tree/lazy_st.go)
- **Advanced Algebraic & Matrix Computations**
  - [Matrix Multiplication (Divide & Conquer)](./matrix_multiply_dc/matrix_dc.go)
  - [Gaussian Elimination](./gaussian_elimination/gaussian.go)
- **Advanced Self-Balancing Systems**
  - [Fibonacci Heap](./fibonacci_heap/fib_heap.go)
  - [Red-Black Tree](./red_black_tree/rbt.go)
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
| **Sparse Table** | $O(n \log n)$ init / $O(1)$ query | $O(n \log n)$ |
| **Trie** | $O(L)$ insertion / search | $O(L \cdot \Sigma)$ |
| **PBDS (Order Stats Tree)** | $O(\log n)$ insert / query | $O(n)$ |
| **Heavy-Light Decomposition** | $O(n)$ build / $O(\log^2 n)$ query | $O(n)$ |
| **Suffix Array (with LCP)** | $O(n \log n)$ or $O(n \log^2 n)$ build | $O(n)$ |
| **Dijkstra's Algorithm** | $O((V + E) \log V)$ | $O(V + E)$ |
| **Bellman-Ford Algorithm** | $O(V \cdot E)$ | $O(V + E)$ |
| **Floyd-Warshall Algorithm** | $O(V^3)$ | $O(V^2)$ |
| **Kruskal's MST** | $O(E \log E)$ | $O(V + E)$ |
| **Prim's MST** | $O(E \log V)$ | $O(V + E)$ |
| **Tarjan's SCC** | $O(V + E)$ | $O(V)$ |
| **Dinic's Maxflow** | $O(V^2 E)$ | $O(V + E)$ |
| **Kahn's Algorithm** | $O(V + E)$ | $O(V + E)$ |
| **Kadane's Algorithm** | $O(n)$ | $O(1)$ |
| **KMP Matcher** | $O(n + m)$ | $O(m)$ |
| **Modular Power** | $O(\log \text{power})$ | $O(1)$ |
| **Rabin-Karp Hashing** | $O(n + m)$ average | $O(1)$ |
| **Suffix Automaton** | $O(n \cdot \Sigma)$ | $O(n \cdot \Sigma)$ |
| **Mo's Algorithm** | $O((n + q) \sqrt{n})$ | $O(q)$ |
| **Convex Hull Trick** | $O(\log n)$ query / insert | $O(n)$ |
| **Persistent Segment Tree** | $O(\log n)$ query / update | $O(n + q \log n)$ |
| **Lazy Segment Tree** | $O(\log n)$ range query / update | $O(n)$ |
| **Matrix Multiplication (D&C)** | $O(n^3)$ or $O(n^{\log_2 7})$ block scaling | $O(n^2)$ |
| **Gaussian Elimination** | $O(n^3)$ | $O(n^2)$ |
| **Fibonacci Heap** | $O(1)$ insert/decrease-key / $O(\log n)$ delete | $O(n)$ |
| **Red-Black Tree** | $O(\log n)$ search / insert / delete | $O(n)$ |

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
