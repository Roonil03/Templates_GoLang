package maximum_clique

type MaximumClique struct{ adjacent [][]bool }

func NewMaximumClique(n int) *MaximumClique {
	g := &MaximumClique{make([][]bool, n)}
	for v := range g.adjacent {
		g.adjacent[v] = make([]bool, n)
	}
	return g
}

func (g *MaximumClique) AddEdge(u, v int) {
	if u != v {
		g.adjacent[u][v] = true
		g.adjacent[v][u] = true
	}
}

func (g *MaximumClique) Solve() []int {
	best := []int{}
	current := []int{}
	colorSort := func(candidates []int) (order, colors []int) {
		remaining := append([]int(nil), candidates...)
		for color := 1; len(remaining) > 0; color++ {
			class, next := []int{}, []int{}
			for _, v := range remaining {
				fits := true
				for _, w := range class {
					if g.adjacent[v][w] {
						fits = false
						break
					}
				}
				if fits {
					class = append(class, v)
					order = append(order, v)
					colors = append(colors, color)
				} else {
					next = append(next, v)
				}
			}
			remaining = next
		}
		return
	}
	var expand func([]int)
	expand = func(candidates []int) {
		order, colors := colorSort(candidates)
		for i := len(order) - 1; i >= 0; i-- {
			if len(current)+colors[i] <= len(best) {
				return
			}
			v := order[i]
			current = append(current, v)
			next := []int{}
			for _, w := range order[:i] {
				if g.adjacent[v][w] {
					next = append(next, w)
				}
			}
			if len(next) == 0 {
				if len(current) > len(best) {
					best = append([]int(nil), current...)
				}
			} else {
				expand(next)
			}
			current = current[:len(current)-1]
		}
	}
	candidates := make([]int, len(g.adjacent))
	for v := range candidates {
		candidates[v] = v
	}
	expand(candidates)
	return best
}
