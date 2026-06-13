package suffix_automaton

type Node struct {
	Len  int
	Link int
	Next []int
}

type SuffixAutomaton[T comparable] struct {
	Nodes []Node
	Last  int
	sigma int
	val   func(T) int
}

func NewSuffixAutomaton[T comparable](capacity, sigma int, val func(T) int) *SuffixAutomaton[T] {
	sa := &SuffixAutomaton[T]{
		Nodes: make([]Node, 1, capacity*2),
		Last:  0,
		sigma: sigma,
		val:   val,
	}
	sa.Nodes[0].Link = -1
	sa.Nodes[0].Next = make([]int, sigma)
	for i := range sa.Nodes[0].Next {
		sa.Nodes[0].Next[i] = -1
	}
	return sa
}

func (sa *SuffixAutomaton[T]) Extend(c T) {
	idx := sa.val(c)
	cur := len(sa.Nodes)
	nextArr := make([]int, sa.sigma)
	for i := range nextArr {
		nextArr[i] = -1
	}
	sa.Nodes = append(sa.Nodes, Node{
		Len:  sa.Nodes[sa.Last].Len + 1,
		Link: -1,
		Next: nextArr,
	})

	p := sa.Last
	for p != -1 && sa.Nodes[p].Next[idx] == -1 {
		sa.Nodes[p].Next[idx] = cur
		p = sa.Nodes[p].Link
	}

	if p == -1 {
		sa.Nodes[cur].Link = 0
	} else {
		q := sa.Nodes[p].Next[idx]
		if sa.Nodes[p].Len+1 == sa.Nodes[q].Len {
			sa.Nodes[cur].Link = q
		} else {
			clone := len(sa.Nodes)
			cloneNext := make([]int, sa.sigma)
			copy(cloneNext, sa.Nodes[q].Next)
			sa.Nodes = append(sa.Nodes, Node{
				Len:  sa.Nodes[p].Len + 1,
				Link: sa.Nodes[q].Link,
				Next: cloneNext,
			})

			for p != -1 && sa.Nodes[p].Next[idx] == q {
				sa.Nodes[p].Next[idx] = clone
				p = sa.Nodes[p].Link
			}
			sa.Nodes[q].Link = clone
			sa.Nodes[cur].Link = clone
		}
	}
	sa.Last = cur
}
