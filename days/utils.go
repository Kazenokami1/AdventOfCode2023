package days

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type pipe struct {
	shape     string
	position  [2]int
	neighbors []*pipe
}

func (p *pipe) addNeighbor(pipe *pipe) {
	p.neighbors = append(p.neighbors, pipe)
}
