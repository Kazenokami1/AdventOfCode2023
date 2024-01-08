package days

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func isGreaterThan(first int, second int) bool {
	return first > second
}
