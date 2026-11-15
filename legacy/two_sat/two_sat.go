// Package two_sat solves Boolean formulas in conjunctive normal form with two
// literals per clause, using implication graphs and Kosaraju's SCC algorithm.
package two_sat

// TwoSAT uses zero-based variable IDs. Solve takes O(V+C) time and space,
// where V is the number of variables and C is the number of clauses.
type TwoSAT struct{ graph, reverse [][]int }

func NewTwoSAT(n int) *TwoSAT { return &TwoSAT{make([][]int, 2*n), make([][]int, 2*n)} }
func literal(v int, value bool) int {
	if value {
		return 2*v + 1
	}
	return 2 * v
}
func (s *TwoSAT) imply(a, b int) {
	s.graph[a] = append(s.graph[a], b)
	s.reverse[b] = append(s.reverse[b], a)
}

// AddClause adds (variable a == aValue) OR (variable b == bValue).
// Invalid variable IDs panic. Clauses may be added between Solve calls.
func (s *TwoSAT) AddClause(a int, aValue bool, b int, bValue bool) {
	u, v := literal(a, aValue), literal(b, bValue)
	s.imply(u^1, v)
	s.imply(v^1, u)
}

// Force requires the chosen variable to equal value.
func (s *TwoSAT) Force(v int, value bool) { s.AddClause(v, value, v, value) }

// Solve returns a satisfying assignment, or (nil,false) for an unsatisfiable
// formula. Each call uses fresh traversal state and preserves the clauses.
func (s *TwoSAT) Solve() ([]bool, bool) {
	n := len(s.graph)
	seen := make([]bool, n)
	order := make([]int, 0, n)
	var visit func(int)
	visit = func(v int) {
		seen[v] = true
		for _, w := range s.graph[v] {
			if !seen[w] {
				visit(w)
			}
		}
		order = append(order, v)
	}
	for v := 0; v < n; v++ {
		if !seen[v] {
			visit(v)
		}
	}
	component := make([]int, n)
	for v := range component {
		component[v] = -1
	}
	var assign func(int, int)
	assign = func(v, id int) {
		component[v] = id
		for _, w := range s.reverse[v] {
			if component[w] == -1 {
				assign(w, id)
			}
		}
	}
	id := 0
	for i := len(order) - 1; i >= 0; i-- {
		if v := order[i]; component[v] == -1 {
			assign(v, id)
			id++
		}
	}
	values := make([]bool, n/2)
	for v := range values {
		if component[2*v] == component[2*v+1] {
			return nil, false
		}
		values[v] = component[2*v] < component[2*v+1]
	}
	return values, true
}
