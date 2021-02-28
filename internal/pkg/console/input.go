package console

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Console ...
type Console struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

// New returns a new Console
func New(reader *bufio.Reader, writer *bufio.Writer) *Console {
	c := &Console{Reader: reader, Writer: writer}
	return c
}

// ReadString ...
func (c *Console) ReadString(prompt string) string {
	c.Writer.WriteString(fmt.Sprintf("%s >> ", prompt))
	c.Writer.Flush()
	text, err := c.Reader.ReadString('\n')
	if err != nil {
		c.Writer.WriteString("could not read the input\n")
		c.Writer.Flush()
		return ""
	}
	return strings.TrimSpace(text)
}

// ReadInteger ...
func (c *Console) ReadInteger(prompt string, limits ...int) int {
	// TODO: check length of limits and ensure is 0 or 2
	for {
		input := c.ReadString(prompt)
		result, err := strconv.Atoi(input)
		if err != nil {
			c.Writer.WriteString("input must be a whole number\n")
			c.Writer.Flush()
			continue
		}
		if len(limits) == 2 {
			min := limits[0]
			max := limits[1]
			if result < min || result > max {
				c.Writer.WriteString(
					fmt.Sprintf("input must be in the range: %d to %d\n", min, max))
				c.Writer.Flush()
				continue
			}
		}
		return result
	}
}

// ReadFloat64 ...
func (c *Console) ReadFloat64(prompt string, limits ...float64) float64 {
	// TODO: check length of limits and ensure is 0 or 2
	for {
		input := c.ReadString(prompt)
		result, err := strconv.ParseFloat(input, 64)
		if err != nil {
			c.Writer.WriteString("input must be a number\n")
			c.Writer.Flush()
			continue
		}
		if len(limits) == 2 {
			min := limits[0]
			max := limits[1]
			if result < min || result > max {
				c.Writer.WriteString(
					fmt.Sprintf("input must be in the range %.1f to %.1f\n", min, max))
				c.Writer.Flush()
				continue
			}
		}
		return result
	}
}
