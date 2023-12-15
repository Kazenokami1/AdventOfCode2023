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
