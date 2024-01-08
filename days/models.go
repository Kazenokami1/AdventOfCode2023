package days

type pipe struct {
	shape     string
	position  [2]int
	neighbors []*pipe
}

func (p *pipe) addNeighbor(pipe *pipe) {
	p.neighbors = append(p.neighbors, pipe)
}

type box struct {
	slots []string
}

type grid struct {
	heatLossFromStart int
	heatValue         int
	neighbors         []*grid
}

func (g *grid) addNeighbor(grid *grid) {
	g.neighbors = append(g.neighbors, grid)
}

type part struct {
	xVal      int
	mVal      int
	aVal      int
	sVal      int
	partValue int
}

func (p *part) calcPartValue() {
	p.partValue = p.xVal + p.mVal + p.aVal + p.sVal
}
