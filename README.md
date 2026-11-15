# Go Competitive Programming Templates

<p align="center">
  <a href="https://pkg.go.dev/github.com/Roonil03/Templates_GoLang"><img src="https://pkg.go.dev/badge/github.com/Roonil03/Templates_GoLang.svg" alt="Go Reference"></a>
  <a href="https://github.com/Roonil03/Templates_GoLang/blob/main/go.mod"><img src="https://img.shields.io/badge/Go-1.23-00ADD8?logo=go" alt="Go 1.23"></a>
  <a href="https://github.com/Roonil03/Templates_GoLang/blob/main/LICENSE"><img src="https://img.shields.io/github/license/Roonil03/Templates_GoLang" alt="License"></a>
  <a href="https://pkg.go.dev/github.com/Roonil03/Templates_GoLang?tab=versions"><img src="https://img.shields.io/github/v/tag/Roonil03/Templates_GoLang?filter=v%2A&amp;label=module" alt="Tagged module version"></a>
</p>

This module provides competitive programming templates for Go 1.23 and later. Each link below opens the source file. For Go 1.20, use the [legacy catalog](./legacy/README.md).

## Repository Navigation Index

- **Core Collections & Utilities**
  - [Fast I/O submission template](./io/template_io.go)
  - [Deque Implementation](./deque/Deque.go)
  - [LRU Cache Implementation](./lru_cache/lru_cache.go)
- **Advanced Tree Structures**
  - [Segment Tree Implementation](./segment_tree/segment_tree.go)
  - [Fenwick Tree (1D BIT) Implementation](./fenwick/fenwick.go)
  - [2D Fenwick Tree Implementation](./fenwick2d/fenwick2d.go)
  - [Treap Implementation](./treap/Treap.go)
- **Graph & Advanced Relations**
  - [Disjoint Set Union (DSU) Implementation](./union_find/union_find.go)
  - [Rope Data Structure](./rope/Rope.go)
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
- **Tree queries**
  - [Binary-lifting LCA](./lca/lca.go) — Find the lowest common ancestor or the k-th ancestor in a rooted forest.
  - [Centroid decomposition](./centroid_decomposition/centroid_decomposition.go) — Build the centroid-parent array of an undirected forest.
- **Graphs**
  - [Hopcroft–Karp matching](./hopcroft_karp/hopcroft_karp.go) — Return maximum matching partners on both sides of a bipartite graph.
  - [Euler path endpoints](./euler_path/euler_path.go) — Check connectivity and degrees in a directed or undirected multigraph.
  - [Hierholzer’s algorithm](./hierholzer/hierholzer.go) — Return a trail that uses every edge exactly once.
  - [2-SAT](./two_sat/two_sat.go) — Add two-literal clauses and recover a satisfying assignment.
  - [Maximum clique with coloring bounds](./maximum_clique/maximum_clique.go) — Find an exact maximum clique with a greedy-coloring upper bound.
  - [Dominator tree](./dominator_tree/dominator_tree.go) — Compute immediate dominators from a source in a directed graph.
- **Number theory**
  - [Sieve of Eratosthenes](./sieve/sieve.go) — Build a primality table or list all primes through an inclusive bound.
  - [Smallest prime factor](./smallest_prime_factor/smallest_prime_factor.go) — Build a linear sieve and factor integers using its table.
  - [Chinese remainder theorem](./chinese_remainder/chinese_remainder.go) — Combine congruences, including compatible non-coprime moduli.
  - [XOR linear basis](./xor_basis/xor_basis.go) — Insert unsigned vectors, test representability, and maximize a subset XOR.
  - [Number theoretic transform](./ntt/ntt.go) — Transform coefficients and convolve polynomials modulo 998244353.
- **Range queries and collections**
  - [Wavelet tree](./wavelet_tree/wavelet_tree.go) — Find order statistics and count values in a half-open range.
  - [Wavelet matrix](./wavelet_matrix/wavelet_matrix.go) — Answer the same static range queries using level-by-level partitions.
  - [Monotonic stack](./monotonic_stack/monotonic_stack.go) — Maintain ordered candidates and find next greater or smaller elements.
  - [Monotonic queue](./monotonic_queue/monotonic_queue.go) — Maintain a FIFO minimum or compute sliding-window minima.
  - [Min heap](./heap/heap.go) — Push and pop the smallest value using the binary heap.

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
| **Trie** | $O(L)$ insertion / search | $O(L \cdot \Sigma)$ |
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
| **KMP Matcher** | $O(n + m)$ | $O(m)$ |
| **Modular Power** | $O(\log \text{power})$ | $O(1)$ |
| **Rabin-Karp Hashing** | $O(n + m)$ average | $O(1)$ |
| **Suffix Automaton** | $O(n \cdot \Sigma)$ | $O(n \cdot \Sigma)$ |
| **Mo's Algorithm** | $O((n + q) \sqrt{n})$ | $O(q)$ |
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
| **Min heap** | O(log n) push or pop | O(n) |

Here, σ counts distinct values, w is the XOR word width, C counts clauses, and M is the running CRT period. Wavelet ranks and ranges are zero-based.

## Competitive Programming Integration

Install the package you need.

```sh
go get github.com/Roonil03/Templates_GoLang/heap
```

```go
package main

import (
    "fmt"
    "github.com/Roonil03/Templates_GoLang/heap"
)

func main() {
    h := heap.NewMinHeap(7, 2, 5)
    fmt.Println(h.Pop()) // 2
}
```

For a single-file judge, copy the implementation and its imports into your submission under `package main`. For packages that import another template, include its implementation and adjust the qualified calls as well.
