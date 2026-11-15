# Go Competitive Programming Templates

<p align="center">
  <a href="https://pkg.go.dev/github.com/Roonil03/Templates_GoLang"><img src="https://pkg.go.dev/badge/github.com/Roonil03/Templates_GoLang.svg" alt="Go Reference"></a>
  <a href="https://github.com/Roonil03/Templates_GoLang/blob/main/go.mod"><img src="https://img.shields.io/badge/Go-1.21-00ADD8?logo=go" alt="Go 1.21"></a>
  <a href="https://github.com/Roonil03/Templates_GoLang/blob/main/LICENSE"><img src="https://img.shields.io/github/license/Roonil03/Templates_GoLang" alt="License"></a>
  <a href="https://pkg.go.dev/github.com/Roonil03/Templates_GoLang?tab=versions"><img src="https://img.shields.io/github/v/tag/Roonil03/Templates_GoLang?filter=v%2A&amp;label=module" alt="Tagged module version"></a>
</p>

This module provides competitive programming templates for Go 1.21 and later. Each link below opens the source file. For Go 1.20, use the [legacy catalog](./legacy/README.md).

## Repository Navigation Index

### Collections and heaps

- [Deque Implementation](./deque/Deque.go)
- [LRU Cache Implementation](./lru_cache/lru_cache.go)
- [Monotonic stack](./monotonic_stack/monotonic_stack.go)
- [Monotonic queue](./monotonic_queue/monotonic_queue.go)
- [Min heap](./min_heap/min_heap.go)
- [Priority queue (max heap)](./priority_queue/priority_queue.go)
- [Fibonacci Heap](./fibonacci_heap/fib_heap.go)

### Ordered trees

- [Treap Implementation](./treap/Treap.go)
- [Red-Black Tree](./red_black_tree/rbt.go)
- [Policy-Based Data Structure (PBDS)](./pbds/pbds.go)

### Strings

- [Trie Implementation](./trie/trie.go)
- [KMP Matcher & Prefix Function](./kmp/kmp.go)
- [Rabin-Karp Rolling Hash](./rabin_karp/rabin_karp.go)
- [Suffix Array & LCP](./suffix_array/suffix_array.go)
- [Suffix Automaton](./suffix_automaton/suffix_automaton.go)
- [Rope Data Structure](./rope/Rope.go)

### Range queries

- [Segment Tree Implementation](./segment_tree/segment_tree.go)
- [Lazy Segment Tree](./lazy_segment_tree/lazy_st.go)
- [Persistent Segment Tree](./persistent_segment_tree/persistent_st.go)
- [Fenwick Tree (1D BIT) Implementation](./fenwick/fenwick.go)
- [2D Fenwick Tree Implementation](./fenwick2d/fenwick2d.go)
- [Sparse Table Implementation](./sparse_table/sparse_table.go)
- [Mo's Offline Block Decomposition](./mos_algorithm/mos.go)
- [Wavelet tree](./wavelet_tree/wavelet_tree.go)
- [Wavelet matrix](./wavelet_matrix/wavelet_matrix.go)

### Graph algorithms

#### Connectivity and ordering

- [Disjoint Set Union (DSU) Implementation](./union_find/union_find.go)
- [Tarjan's Strongly Connected Components](./tarjan_scc/tarjan.go)
- [Kahn's Topological Sort](./kahn/kahn.go)
- [Dominator tree](./dominator_tree/dominator_tree.go)

#### Shortest paths

- [Dijkstra's Shortest Path](./dijkstra/dijkstra.go)
- [Bellman-Ford Algorithm](./bellman_ford/bellman_ford.go)
- [Floyd-Warshall All-Pairs](./floyd_warshall/floyd_warshall.go)

#### Minimum spanning trees

- [Kruskal's Minimum Spanning Tree](./kruskal/kruskal.go)
- [Prim's Minimum Spanning Tree](./prim/prim.go)

#### Flow and matching

- [Dinic's Maxflow Algorithm](./maxflow/maxflow.go)
- [Hopcroft–Karp matching](./hopcroft_karp/hopcroft_karp.go)

#### Eulerian paths

- [Euler path endpoints](./euler_path/euler_path.go)
- [Hierholzer’s algorithm](./hierholzer/hierholzer.go)

#### Constraints and cliques

- [2-SAT](./two_sat/two_sat.go)
- [Maximum clique with coloring bounds](./maximum_clique/maximum_clique.go)

### Tree algorithms

- [Euler Tour Implementation](./euler_tour/euler_tour.go)
- [Binary-lifting LCA](./lca/lca.go)
- [Heavy-Light Decomposition (HLD)](./hld/hld.go)
- [Centroid decomposition](./centroid_decomposition/centroid_decomposition.go)

### Mathematics

- [Modular Arithmetic & Modular Power](./mod_arithmetic/mod_arith.go)
- [Sieve of Eratosthenes](./sieve/sieve.go)
- [Smallest prime factor](./smallest_prime_factor/smallest_prime_factor.go)
- [Chinese remainder theorem](./chinese_remainder/chinese_remainder.go)
- [XOR linear basis](./xor_basis/xor_basis.go)
- [Number theoretic transform](./ntt/ntt.go)
- [Matrix Multiplication (Divide & Conquer)](./matrix_multiply_dc/matrix_dc.go)
- [Gaussian Elimination](./gaussian_elimination/gaussian.go)
- [Convex Hull Trick](./convex_hull_trick/cht.go)

### Sequence algorithms

- [Kadane's Maximum Subarray Sum](./kadane/kadane.go)

### Utilities

- [Fast I/O submission template](./io/template_io.go)

## Algorithmic Complexity Reference

| Structure | Time Complexity | Space Complexity |
| :--- | :--- | :--- |
| **Deque** | $O(1)$ amortized push/pop | $O(n)$ |
| **Segment Tree** | $O(\log n)$ query/update | $O(n)$ |
| **Fenwick Tree (BIT)** | $O(\log n)$ query/update | $O(n)$ |
| **2D Fenwick Tree** | $O(\log^2 n)$ query/update | $O(n^2)$ |
| **Treap** | $O(\log n)$ expected | $O(n)$ |
| **Disjoint Set Union (DSU)** | $O(\alpha(n))$ amortized | $O(n)$ |
| **Rope** | $O(h^2)$ split, $O(h)$ concat (height $h$) | $O(n)$ |
| **Euler Tour** | $O(1)$ ancestor check | $O(n)$ |
| **LRU Cache** | $O(1)$ get/put | $O(\text{capacity})$ |
| **Sparse Table** | $O(n \log n)$ init / $O(1)$ query | $O(n \log n)$ |
| **Trie** | $O(L)$ insertion / search | $O(P \cdot \Sigma)$ total node storage |
| **PBDS (Order Stats Tree)** | $O(\log n)$ expected insert / query | $O(n)$ |
| **Heavy-Light Decomposition** | $O(n)$ build / $O(\log^2 n)$ query | $O(n)$ |
| **Suffix Array (with LCP)** | $O(n \log^2 n)$ build | $O(n)$ |
| **Dijkstra's Algorithm** | $O((V + E) \log V)$ | $O(V + E)$ |
| **Bellman-Ford Algorithm** | $O(V \cdot E)$ | $O(V + E)$ |
| **Floyd-Warshall Algorithm** | $O(V^3)$ | $O(V^2)$ |
| **Kruskal's MST** | $O(E \log E)$ | $O(V + E)$ |
| **Prim's MST** | $O(E \log V)$ | $O(V + E)$ |
| **Tarjan's SCC** | $O(V + E)$ | $O(V)$ |
| **Dinic's Maxflow** | $O(V^2 E)$ | $O(V + E)$ |
| **Kahn's Algorithm** | $O(V + E)$ | $O(V + E)$ |
| **Kadane's Algorithm** | $O(n)$ | $O(1)$ |
| **KMP Matcher** | $O(n + m)$ | $O(m)$ auxiliary + $O(z)$ returned matches |
| **Modular Power** | $O(\log \text{power})$ | $O(1)$ |
| **Rabin-Karp Hashing** | $O(n + m + Hm)$; $O(nm)$ worst case | $O(1)$ auxiliary + $O(z)$ returned matches |
| **Suffix Automaton** | $O(n \cdot \Sigma)$ | $O(n \cdot \Sigma)$ |
| **Mo's Algorithm** | $O(q \log q + (qB + n\lceil n/B \rceil)F + qA)$ | $O(q)$ answers, excluding callback state |
| **Convex Hull Trick** | $O(\log n)$ query, $O(1)$ amortized insert | $O(n)$ |
| **Persistent Segment Tree** | $O(\log n)$ query / update | $O(n + q \log n)$ |
| **Lazy Segment Tree** | $O(\log n)$ range query / update | $O(n)$ |
| **Matrix Multiplication (D&C)** | $O(rkc)$ for $r × k$ and $k × c$ inputs | $O(rc)$ output plus logarithmic recursion |
| **Gaussian Elimination** | $O(n^3)$ | $O(n^2)$ |
| **Fibonacci Heap** | $O(1)$ insert / minimum lookup | $O(n)$ |
| **Red-Black Tree** | $O(\log n)$ search / insert | $O(n)$ |
| **Binary-lifting LCA** | O(V log V) build, O(log V) query | O(V log V) |
| **Centroid decomposition** | O(V log V) | O(V) |
| **Hopcroft–Karp matching** | O((V + E) √V) | O(V + E) |
| **Euler path endpoints** | O(V + E) | O(V + E) |
| **Hierholzer’s algorithm** | O(V + E) | O(V + E) |
| **2-SAT** | O(V + C) | O(V + C) |
| **Maximum clique with coloring bounds** | O(V² 2^V) worst case | O(V²) |
| **Dominator tree** | O((V + E) log V) | O(V + E) |
| **Sieve of Eratosthenes** | O(n log log n) | O(n) |
| **Smallest prime factor** | O(n) build, O(log x) factor | O(n) |
| **Chinese remainder theorem** | O(k log M) Euclidean steps | O(1) for int64 inputs |
| **XOR linear basis** | O(w) per insert or query | O(w) |
| **Number theoretic transform** | O(n log n) | O(n) convolution |
| **Wavelet tree** | O(n log n) build, O(log(σ + 1)) query | O(n log(σ + 1)) |
| **Wavelet matrix** | O(n log n) build, O(log(σ + 1)) query | O(n log(σ + 1)) |
| **Monotonic stack** | O(1) amortized push, O(n) all neighbors | O(n) |
| **Monotonic queue** | O(1) amortized per operation | O(n) |
| **Min heap** | O(n) build, O(log n) push or pop, O(1) peek | O(n) |
| **Priority queue (max heap)** | O(n) build, O(log n) push or pop, O(1) peek | O(n) |

Here, σ counts distinct values, w is the XOR word width, C counts clauses, and M is the running CRT period. Wavelet ranks and ranges are zero-based.

For the trie, L is one word's length, P counts all allocated nodes (the root plus distinct prefixes across all inserted words), and Σ = 26. For KMP and Rabin–Karp, n and m are the text and pattern lengths, z is the number of returned matches, and H counts Rabin–Karp hash hits that require verification. These bounds assume constant-time element comparisons and value mapping; repetitive inputs can require $O(nm)$ verification even without hash collisions.

For Mo's algorithm, n is the array length, q is the query count, B > 0 is the caller's block size, F bounds the time per add/remove callback, and A bounds the time per answer callback. With B ≈ √n and constant-time callbacks, the total is $O(q \log q + (n + q)\sqrt{n})$, including query sorting. Space assumes fixed-size answers and excludes storage owned by the callbacks.

## Competitive Programming Integration

Install the package you need.

```sh
go get github.com/Roonil03/Templates_GoLang/min_heap
```

```go
package main

import (
    "fmt"
    "github.com/Roonil03/Templates_GoLang/min_heap"
)

func main() {
    h := min_heap.NewMinHeap(7, 2, 5)
    fmt.Println(h.Pop())
}
```

For a single-file judge, copy the implementation and its imports into your submission under `package main`. For packages that import another template, include its implementation and adjust the qualified calls as well.
