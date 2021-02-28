package main

import (
	"os"

	"github.com/pscompsci/pmtools/internal/pkg/console"
	"github.com/pscompsci/pmtools/internal/pmtools"
)

const (
	new      = iota
	edit     = iota
	sow      = iota
	pcalc    = iota
	generate = iota
	exit     = iota
)

func run(c *console.Console, p *pmtools.Project) {
	for {
		displayMenu(c)
		choice := c.ReadInteger("Choice", 1, 6)
		action(c, choice-1, p)
	}
}

func action(c *console.Console, choice int, p *pmtools.Project) {
	switch choice {
	case new:
		newProjectAction(c, p)
	case edit:
		c.Writer.WriteString("Edit project chosen\n")
		c.Writer.Flush()
	case sow:
		c.Writer.WriteString("Add SOW chosen\n")
		c.Writer.Flush()
	case pcalc:
		c.Writer.WriteString("Add pCalc chosen\n")
		c.Writer.Flush()
	case generate:
		c.Writer.WriteString("Generate chosen\n")
		c.Writer.Flush()
	case exit:
		c.Writer.WriteString("Goodbye\n")
		c.Writer.Flush()
		os.Exit(0)

	}
}

func newProjectAction(c *console.Console, p *pmtools.Project) {
	displayHeader(c, "Create New Project")
	p.ProjectCode = c.ReadString("Project Code")
	p.ProjectName = c.ReadString("Project Name")
	p.ClientName = c.ReadString("Client Name")
	p.ClientAbbreviation = c.ReadString("Client Abbr")
	p.ProjectManager = c.ReadString("Project Manager")
	p.BaseFolder = c.ReadString("Base Folder")
	p.SOW.Filepath = c.ReadString("SOW Location")
	p.PCalc.Filepath = c.ReadString("pCalc Location")
	displayFooter(c)
}
