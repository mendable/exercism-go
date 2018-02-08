package tree

import (
	"errors"
	"sort"
)

/*
FastBuild is my new improved version.

[Wed Feb 07][21:22:06] ~/exercism/go/tree-building$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkTwoTree-4              30    43125067 ns/op
BenchmarkTenTree-4             300     4741538 ns/op
BenchmarkShallowTree-4         300     4213447 ns/op
PASS
ok    _/Users/jason/exercism/go/tree-building 5.857s
*/
func FastBuild(records []Record) (root *Node, err error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	table := map[int]*Node{}
	for idx, record := range records {
		if table[record.ID] != nil {
			return nil, errors.New("duplicate node")
		}

		table[record.ID] = &Node{ID: record.ID}

		if idx == 0 {
			root = table[record.ID]

			if record.ID != record.Parent {
				return nil, errors.New("root node has parent")
			}
		}

		if record.ID > len(records)-1 {
			return nil, errors.New("non-continuous")
		}

		if idx != 0 && record.Parent == record.ID {
			return nil, errors.New("cycle directly")
		}

		if record.Parent > record.ID {
			return nil, errors.New("higher id parent of lower id")
		}

		if idx != 0 {
			table[record.Parent].Children = append(table[record.Parent].Children, table[record.ID])
		}
	}

	if root == nil {
		return nil, errors.New("no root node")
	}

	return root, nil
}
