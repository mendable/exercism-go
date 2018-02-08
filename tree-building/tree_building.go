package tree

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

const useNewCode = true

func Build(records []Record) (*Node, error) {
	if useNewCode {
		return FastBuild(records)
	} else {
		return SlowBuild(records)
	}
}
