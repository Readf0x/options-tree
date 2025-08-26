package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var html = false

func main() {
	var file string
	offset := 1
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		html = true
		if len(os.Args) > 2 {
			offset = 2
		}
	}
	if len(os.Args) > 1 && os.Args[1] != "-h" {
		f, _ := os.Open(os.Args[offset])
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
		for i, name := range sep_fq {
			if new_node, err := current_node.find_child(name); err == nil {
				current_node = new_node
			} else {
				new_node := &node{ label: name, fq: "#opt-" + strings.Join(sep_fq[:i+1], ".") }
				current_node.children = append(current_node.children, new_node)
				current_node = new_node
			}
		}
	}

	fmt.Println("<ul>")
	for _, i := range tree.children {
		i.print_tree(1)
	}
	fmt.Println("</ul>")
}

type node struct {
	label string
	fq string
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
	if n.label == "" {
		return
	}
	i := strings.Repeat(" ", indent*2)
	if html && len(n.children) == 0 {
		fmt.Printf("%s<li><a href=\"options.html%s\">%s</a></li>\n", i, n.fq, n.label)
	} else if html && len(n.children) > 0 {
		fmt.Printf("%s<details><summary><a href=\"options.html%s\">%s</a></summary>\n", i, n.fq, n.label)
	} else {
		fmt.Println(strings.Repeat(" ", indent*2) + n.label)
	}
	if html && len(n.children) > 0 {
		fmt.Printf("%s<ul>\n", i)
	}
	for _, child := range n.children {
		child.print_tree(indent+1)
	}
	if html && len(n.children) > 0 {
		fmt.Printf("%s</ul></details>\n", i)
	}
}

