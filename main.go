package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var file string
	if len(os.Args) > 1 {
		f, _ := os.Open("options.txt")
		raw, _ := io.ReadAll(f)
		f.Close()
		file = string(raw)
	} else if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		raw, _ := io.ReadAll(os.Stdin)
		file = string(raw)
	} else {
		os.Exit(1)
	}

	var fq_list [][]string

	for fq := range strings.SplitSeq(file, "\n") {
		fq_list = append(fq_list, strings.Split(fq, "."))
	}

	tree := node{
		label: ".",
	}

	for _, sep_fq := range fq_list {
		current_node := &tree
		for _, name := range sep_fq {
			if new_node, err := current_node.find_child(name); err == nil {
				current_node = new_node
			} else {
				new_node := &node{ label: name }
				current_node.children = append(current_node.children, new_node)
				current_node = new_node
			}
		}
	}

	tree.print_tree(0)
}

type node struct {
	label string
	children []*node
}

func (n node) find_child(name string) (*node, error) {
	for _, i := range n.children {
		if i.label == name {
			return i, nil
		}
	}
	return &node{}, fmt.Errorf("doesn't exist")
}

func (n node) print_tree(indent int) {
	fmt.Println(strings.Repeat(" ", indent*2) + n.label)
	for _, child := range n.children {
		child.print_tree(indent+1)
	}
}

