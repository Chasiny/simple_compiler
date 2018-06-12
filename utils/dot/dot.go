package dot

import (
	"fmt"
	"sync"
)

type node struct {
	id    int
	label string
	color string
}
type line struct {
	src  string
	dest string
}

type Dot struct {
	nodes []node
	lines []line
	id    int
	sync.Mutex
}

func (d *Dot) CreateNode(id int, label string) string {
	d.Lock()
	defer d.Unlock()

	d.nodes = append(d.nodes, node{
		label: label,
	})
	return fmt.Sprintf(`%d[style="filled",label="%s", fillcolor="grey"];`, id, label)
}

func (d *Dot) CreateLine(src, dest int) string {
	return fmt.Sprintf("%d->%d;", src, dest)
}
