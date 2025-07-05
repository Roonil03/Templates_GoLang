# Go Programming Data Structures

A comprehensive collection of advanced data structures for competitive programming in Go, including efficient implementations of structures not available in Go's standard library.

## 📋 Table of Contents

- [Quick Start](#-quick-start)
- [Method 1: Direct GitHub Import](#method-1-direct-github-import)
- [Method 2: Copy-Paste Technique](#method-2-copy-paste-technique)
- [Available Data Structures](#-available-data-structures)
- [Usage Examples](#-usage-examples)
- [Performance Characteristics](#-performance-characteristics)
- [Best Practices](#%EF%B8%8F-best-practices)

## 🚀 Quick Start

This repository provides two main approaches for using these data structures in competitive programming:

1. **Direct GitHub Import** - Import specific modules directly from GitHub
2. **Copy-Paste Technique** - Copy individual implementations into your solution file

## Method 1: Direct GitHub Import

### Prerequisites
- Go 1.16 or later
- Internet connection during compilation

### Step 1: Initialize Your Project
```bash
mkdir my-competitive-programming
cd my-competitive-programming
go mod init competitive-programming
```

### Step 2: Create Repository Structure
```
your-repo/
├── go.mod
├── go.sum
├── ds/
│   ├── deque.go
│   ├── segment_tree.go
│   ├── fenwick.go
│   ├── treap.go
│   ├── union_find.go
│   ├── rope.go
│   ├── euler_tour.go
│   ├── lru_cache.go
│   ├── heap.go
│   ├── fenwick2d.go
│   └── utils.go
└── main.go
```

### Step 3: Import in Your Code
```go
package main

import (
    "fmt"
    "github.com/roonil03/Templates_GoLang/ds"
    "github.com/roonil03/Templates_GoLang/SegmentTree
)

func main() {
    // Using Segment Tree
    arr := []int{1, 3, 5, 7, 9, 11}
    st := ds.NewSegmentTree(arr)
    
    // Using Fenwick Tree
    fenwick := SegmentTree.NewFenwick(10)
    fenwick.Update(1, 5)
    
    // Using DSU
    dsu := ds.NewDSU(10)
    dsu.Union(1, 2)
}
```

### Step 4: Add Dependencies
```bash
go get github.com/roonil03/Templates_GoLang
```

### Alternative: Use `replace` for Local Development
Add to your `go.mod`:
```
replace github.com/roonil03/Templates_GoLang => ./structures
```

## Method 2: Copy-Paste Technique

For online judges and contests where external imports might not work, use the copy-paste approach:

### Step 1: Template Structure
Create a main template file:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// ==================== DATA STRUCTURES ====================

// [PASTE DATA STRUCTURE IMPLEMENTATIONS HERE]

// ==================== UTILITY FUNCTIONS ====================

// [PASTE UTILITY FUNCTIONS HERE]

// ==================== MAIN SOLUTION ====================

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    // Your solution here
}

// ==================== HELPER FUNCTIONS ====================

func readInts(scanner *bufio.Scanner) []int {
    scanner.Scan()
    strs := strings.Fields(scanner.Text())
    ints := make([]int, len(strs))
    for i, s := range strs {
        ints[i], _ = strconv.Atoi(s)
    }
    return ints
}

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    return n
}
```

### Step 2: Copy Individual Structures
The `ds` folder contains the following:
- [2DFenwick](https://github.com/Roonil03/Templates_GoLang/tree/main/fenwick2D)
- [BIT](https://github.com/Roonil03/Templates_GoLang/tree/main/bit)
- [Deque](https://github.com/Roonil03/Templates_GoLang/tree/main/deque)
- [EulerTour](https://github.com/Roonil03/Templates_GoLang/tree/main/eulertour)
- [LRUCache](https://github.com/Roonil03/Templates_GoLang/tree/main/lrucache)
- [PriorityQueueMax](https://github.com/Roonil03/Templates_GoLang/tree/main/priorityqueue)
- [PriorityQueueMin](https://github.com/Roonil03/Templates_GoLang/tree/main/priorityqueue)
- [Rope](https://github.com/Roonil03/Templates_GoLang/tree/main/rope)
- [SegmentTree](https://github.com/Roonil03/Templates_GoLang/tree/main/segmenttree)
- [Treap](https://github.com/Roonil03/Templates_GoLang/tree/main/treap)
- [UnionFind](https://github.com/Roonil03/Templates_GoLang/tree/main/unionfind)

*[Utilities](https://github.com/Roonil03/Templates_GoLang/tree/main/u) is also present in this folder, which contains `min`, `max` and `abs` functions*
## 📚 Available Data Structures

| Structure | Time Complexity | Space Complexity | Use Cases |
|-----------|----------------|------------------|-----------|
| **Deque** | O(1) push/pop | O(n) | Sliding window, BFS |
| **Segment Tree** | O(log n) query/update | O(n) | Range queries, RMQ |
| **Fenwick Tree** | O(log n) query/update | O(n) | Prefix sums, frequency |
| **Treap** | O(log n) expected | O(n) | Dynamic sets, implicit keys |
| **Union Find** | O(α(n)) amortized | O(n) | Connected components |
| **Rope** | O(log n) split/concat | O(n) | String manipulation |
| **Euler Tour** | O(1) ancestor check | O(n) | Tree queries, LCA |
| **LRU Cache** | O(1) get/put | O(capacity) | Caching, recent access |
| **2D Fenwick** | O(log²n) query/update | O(n²) | 2D prefix sums |
| **Priority Queue** | O(log n) push/pop | O(n) | Dijkstra, scheduling |

## 💡 Usage Examples

### Example 1: Range Sum Queries
```go
func solveProblem() {
    // Input: array and queries
    arr := []int{1, 3, 5, 7, 9, 11}
    st := NewSegmentTree(arr)
    
    // Query sum in range [1, 3]
    result := st.Query(1, 3) // Returns 15 (3+5+7)
    
    // Update position 2 to value 10
    st.Update(2, 10)
    
    // Query again
    result = st.Query(1, 3) // Returns 20 (3+10+7)
}
```

### Example 2: Union Find for Connected Components
```go
func countComponents(n int, edges [][]int) int {
    dsu := NewDSU(n)
    
    for _, edge := range edges {
        dsu.Union(edge[0], edge[1])
    }
    
    components := 0
    for i := 0; i < n; i++ {
        if dsu.Find(i) == i {
            components++
        }
    }
    return components
}
```

### Example 3: Sliding Window Maximum with Deque
```go
func slidingWindowMaximum(nums []int, k int) []int {
    dq := NewDeque(len(nums))
    result := make([]int, 0, len(nums)-k+1)
    
    for i, num := range nums {
        // Remove elements outside window
        for dq.size > 0 && i-dq.data[dq.head] >= k {
            dq.PopFront()
        }
        
        // Remove smaller elements
        for dq.size > 0 && nums[dq.data[dq.tail-1]] <= num {
            dq.PopBack()
        }
        
        dq.PushBack(i)
        
        if i >= k-1 {
            result = append(result, nums[dq.data[dq.head]])
        }
    }
    
    return result
}
```

## ⚡ Performance Characteristics

### Time Complexities Summary
- **Fenwick Tree**: Prefix sum in O(log n), update in O(log n)
- **Segment Tree**: Range query in O(log n), point update in O(log n)
- **Treap**: All operations in O(log n) expected
- **DSU**: Union/Find in O(α(n)) amortized
- **Deque**: Push/Pop at both ends in O(1) amortized

### Space Complexities
- Most structures use **O(n)** space
- **2D Fenwick**: O(n × m) space
- **LRU Cache**: O(capacity) space

## 🛠️ Best Practices

### For Online Judges
1. **Copy-paste approach** is more reliable
2. Keep implementations **compact** and **comment-free**
3. Test locally before submitting
4. Use **consistent naming** across problems

### For Local Development
1. Use **GitHub imports** for better code organization
2. Write **unit tests** for each structure
3. Profile performance on large inputs
4. Keep a **template file** ready
