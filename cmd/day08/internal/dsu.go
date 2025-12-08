package internal

// DisjointSetUnion represents a disjoint set union data structure
// it is used to keep track of the connected components of a graph
// it is implemented as a forest of trees, where each tree represents a connected component
// the parent array is used to store the parent of each element
// the size array is used to store the size of each tree
// the sets variable is used to store the number of sets in the forest
type DisjointSetUnion struct {
	parent []int
	size   []int
	sets   int
}

// NewDisjointSetUnion creates a new disjoint set union data structure with n elements
func NewDisjointSetUnion(n int) *DisjointSetUnion {
	p := make([]int, n)
	s := make([]int, n)

	// make each element its own parent to start with
	for i := range n {
		p[i] = i
		s[i] = 1
	}

	return &DisjointSetUnion{parent: p, size: s, sets: n}
}

// Find finds the root of the tree that x is in
func (d *DisjointSetUnion) Find(x int) int {
	// if the parent of x is not x, then x is not the root of the tree
	// so we need to find the root of the tree
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

// Union unions the sets that a and b are in
// it returns true if a and b were in different sets, false otherwise
func (d *DisjointSetUnion) Union(a, b int) bool {
	rootA := d.Find(a)
	rootB := d.Find(b)

	// if the roots are the same, then a and b are already in the same set
	if rootA == rootB {
		return false
	}

	// make the smaller tree a child of the larger tree
	if d.size[rootA] < d.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	// set the parent of the smaller tree to the root of the larger tree
	d.parent[rootB] = rootA

	// increment the size of the larger tree by the size of the smaller tree
	d.size[rootA] += d.size[rootB]

	// decrement the number of sets
	d.sets--

	return true
}

// Size returns the size of the tree that x is in
func (d *DisjointSetUnion) Size(x int) int {
	// find the root of the tree and return the size of the tree
	return d.size[d.Find(x)]
}

// Sets returns the number of sets in the forest
func (d *DisjointSetUnion) Sets() int {
	return d.sets
}
