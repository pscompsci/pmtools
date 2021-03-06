package main

import (
	"fmt"

	"github.com/pscompsci/pmtools/internal/pkg/console"
)

func displayMenu(c *console.Console) {
	c.Writer.WriteString("\n====================")
	c.Writer.WriteString("\n|       Menu       |")
	c.Writer.WriteString("\n====================")
	c.Writer.WriteString("\n| 1. New Project   |")
	c.Writer.WriteString("\n| 2. Edit Project  |")
	c.Writer.WriteString("\n| 3. Add SOW       |")
	c.Writer.WriteString("\n| 4. Add pCalc     |")
	c.Writer.WriteString("\n| 5. Generate      |")
	c.Writer.WriteString("\n| 6. Exit          |")
	c.Writer.WriteString("\n====================\n")
	c.Writer.Flush()
}

func displayHeader(c *console.Console, title string) {
	c.Writer.WriteString("\n========================================")
	c.Writer.WriteString(fmt.Sprintf("\n%s", title))
	c.Writer.WriteString("\n========================================\n")
	c.Writer.Flush()
}

func displayFooter(c *console.Console) {
	c.Writer.WriteString("========================================\n")
	c.Writer.Flush()
}
