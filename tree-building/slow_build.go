package tree

import (
	"errors"
	"fmt"
)

/*
SlowBuild was original version that needed to be refactored.

Requires commenting "duplicate node" failure test before old/bad/slow/original code can be benchmarked.

[Wed Feb 07][21:07:36] ~/exercism/go/tree-building$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkTwoTree-4               1  5865085880 ns/op
BenchmarkTenTree-4              10   136178492 ns/op
BenchmarkShallowTree-4           1  1102719305 ns/op
PASS
ok    _/Users/jason/exercism/go/tree-building 8.484s
*/
type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func SlowBuild(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	root := &Node{}
	todo := []*Node{root}
	n := 1
	for {
		if len(todo) == 0 {
			break
		}
		newTodo := []*Node(nil)
		for _, c := range todo {
			for _, r := range records {
				if r.Parent == c.ID {
					if r.ID < c.ID {
						return nil, errors.New("a")
					} else if r.ID == c.ID {
						if r.ID != 0 {
							return nil, fmt.Errorf("b")
						}
					} else {
						n++
						switch len(c.Children) {
						case 0:
							nn := &Node{ID: r.ID}
							c.Children = []*Node{nn}
							newTodo = append(newTodo, nn)
						case 1:
							nn := &Node{ID: r.ID}
							if c.Children[0].ID < r.ID {
								c.Children = []*Node{c.Children[0], nn}
								newTodo = append(newTodo, nn)
							} else {
								c.Children = []*Node{nn, c.Children[0]}
								newTodo = append(newTodo, nn)
							}
						default:
							nn := &Node{ID: r.ID}
							newTodo = append(newTodo, nn)
						breakpoint:
							for range []bool{false} {
								for i, cc := range c.Children {
									if cc.ID > r.ID {
										a := make([]*Node, len(c.Children)+1)
										copy(a, c.Children[:i])
										copy(a[i+1:], c.Children[i:])
										copy(a[i:i+1], []*Node{nn})
										c.Children = a
										break breakpoint
									}
								}
								c.Children = append(c.Children, nn)
							}
						}
					}
				}
			}
		}
		todo = newTodo
	}
	if n != len(records) {
		return nil, Mismatch{}
	}
	if err := chk(root, len(records)); err != nil {
		return nil, err
	}
	return root, nil
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
