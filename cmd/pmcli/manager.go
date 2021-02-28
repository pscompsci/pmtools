package main

import (
	"os"

	"github.com/pscompsci/pmtools/internal/pkg/console"
)

const (
	new      = iota
	edit     = iota
	sow      = iota
	pcalc    = iota
	generate = iota
	exit     = iota
)

func run(c *console.Console) {
	for {
		displayMenu(c)
		choice := c.ReadInteger("Choice", 1, 6)
		action(c, choice-1)
	}
}

func action(c *console.Console, choice int) {
	switch choice {
	case new:
		c.Writer.WriteString("New project chosen\n")
		c.Writer.Flush()
	case edit:
		c.Writer.WriteString("Edit project chosen\n")
		c.Writer.Flush()
	case sow:
		c.Writer.WriteString("Edit project chosen\n")
		c.Writer.Flush()
	case pcalc:
		c.Writer.WriteString("Add SOW chosen\n")
		c.Writer.Flush()
	case generate:
		c.Writer.WriteString("Add pCalc chosen\n")
		c.Writer.Flush()
	case exit:
		c.Writer.WriteString("Goodbye\n")
		c.Writer.Flush()
		os.Exit(0)

	}
}
