package main

import (
	"bufio"
	"os"

	"github.com/pscompsci/pmtools/internal/pkg/console"
	"github.com/pscompsci/pmtools/internal/pmtools"
)

func main() {
	r := *bufio.NewReader(os.Stdin)
	w := *bufio.NewWriter(os.Stdout)

	c := console.New(&r, &w)

	p := &pmtools.Project{SOW: &pmtools.SOW{}, PCalc: &pmtools.PCalc{}}

	// pn := c.ReadString("Project Name")
	// pc := c.ReadString("Project Code")
	// cn := c.ReadString("Client Name")
	// ca := c.ReadString("Client Abbreviation")
	// pm := c.ReadString("Project Manager")

	// fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n", pn, pc, cn, ca, pm)

	run(c, p)
}
