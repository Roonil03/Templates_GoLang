package suffix_automaton

type Node struct {
	Len  int
	Link int
	Next [26]int
}

type SuffixAutomaton struct {
	Nodes []Node
	Last  int
}

func NewSuffixAutomaton(capacity int) *SuffixAutomaton {
	sa := &SuffixAutomaton{
		Nodes: make([]Node, 1, capacity*2),
		Last:  0,
	}
	sa.Nodes[0].Link = -1
	for i := range sa.Nodes[0].Next {
		sa.Nodes[0].Next[i] = -1
	}
	return sa
}

func (sa *SuffixAutomaton) Extend(c byte) {
	idx := int(c - 'a')
	cur := len(sa.Nodes)
	newNode := Node{
		Len:  sa.Nodes[sa.Last].Len + 1,
		Link: -1,
	}
	for i := range newNode.Next {
		newNode.Next[i] = -1
	}
	sa.Nodes = append(sa.Nodes, newNode)

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
			cloneNode := Node{
				Len:  sa.Nodes[p].Len + 1,
				Link: sa.Nodes[q].Link,
				Next: sa.Nodes[q].Next,
			}
			sa.Nodes = append(sa.Nodes, cloneNode)

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
