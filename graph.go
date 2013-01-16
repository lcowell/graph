package graph

type edge struct {
	vertexes [2]*Vertex
	value    int
}

// A vertex
type Vertex struct {
	edges map[*Vertex]*edge
	value interface{}
}

// Returns all adjacent vertexes as a slice, which may be empty.
func (v *Vertex) GetNeighbors() []*Vertex {
	neighbors := []*Vertex{}

	for k, _ := range v.edges {
		neighbors = append(neighbors, k)
	}

	return neighbors
}

// The graph
type Graph struct {
	Vertexes map[string]*Vertex // A map of all the vertexes in this graph, indexed by their key.
}

// Initializes a new graph.
func New() *Graph {
	return &Graph{make(map[string]*Vertex)}
}

// Sets the value of the vertex with the specified key.
func (g *Graph) Set(key string, value interface{}) {
	v := &Vertex{make(map[*Vertex]*edge), value}
	g.Vertexes[key] = v
}

// Deletes the vertex with the specified key.
func (g *Graph) Delete(key string) bool {
	// get vertex in question
	v := g.Get(key)
	if v == nil {
		return false
	}

	// iterate over edges, remove edges from neighboring vertexes
	for _, e := range v.edges {
		ends := e.vertexes

		// choose other node, not v
		otherV := ends[0]
		if v == ends[0] {
			otherV = ends[1]
		}

		// delete edge to the to-be-deleted vertex
		delete(otherV.edges, v)
	}

	// delete vertex
	delete(g.Vertexes, key)

	return true
}

// Returns the vertex with this key, or nil if there is no vertex with this key.
func (g *Graph) Get(key string) *Vertex {
	return g.Vertexes[key]
}

// Creates an edge between the vertexes specified by the keys. Returns false if one or both of the keys are invalid or if they are the same.
func (g *Graph) Connect(key string, otherKey string, value int) bool {
	// recursive edges are forbidden
	if key == otherKey {
		return false
	}

	// get vertexes and check for validity of keys
	v := g.Get(key)
	if v == nil {
		return false
	}

	otherV := g.Get(otherKey)
	if otherV == nil {
		return false
	}

	// make a new edge
	e := &edge{[2]*Vertex{v, otherV}, value}

	// add it to both vertexes
	v.edges[otherV] = e
	otherV.edges[v] = e

	// success
	return true
}

// Removes an edge connecting the two vertexes. Returns false if one or both of the keys are invalid or if they are the same.
func (g *Graph) Disconnect(key string, otherKey string) bool {
	// recursive edges are forbidden
	if key == otherKey {
		return false
	}

	// get vertexes and check for validity of keys
	v := g.Get(key)
	if v == nil {
		return false
	}

	otherV := g.Get(otherKey)
	if otherV == nil {
		return false
	}

	delete(v.edges, otherV)
	delete(otherV.edges, v)

	return true
}