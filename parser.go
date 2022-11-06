package gonf

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func fileToTree(filepath string) (*node, error) {
	scanner, err := fileAsScanner(filepath)
	if err != nil {
		return nil, err
	}
	root_node := newRootNode()
	current_node := root_node
	for scanner.Scan() {
		l := scanner.Text()
		if !isLineValid(l) {
			continue
		}
		splitted := strings.Split(l, ":")
		var n *node
		switch getLineType(splitted) {
		case block:
			n = newBlockNode(splitted[0])
		case value:
			n = newValueNode(splitted[0], splitted[1])
		}
		if n == nil {
			continue
		}
		new_node := newNodeFromLine(l)
		current_node.children = append(current_node.children, &new_node)
	}
	return root_node, nil
}

func getLineType(s_line []string) int8 {
	if len(s_line) == 1 {
		return block
	}
	return value
}

func newNodeFromLine(l string) node {
	splitted := strings.Split(l, ":")
	new_node := node{
		key:      splitted[0],
		value:    "",
		children: make([]*node, 0),
	}
	if len(splitted) > 1 {
		new_node.value = splitted[1]
	}
	return new_node
}

func fileAsScanner(filepath string) (*bufio.Scanner, error) {
	block, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(block)
	return bufio.NewScanner(reader), nil
}

func printTree(n *node) {
	fmt.Printf("[%s]:{%s}\n", n.key, n.value)
	for _, n := range n.children {
		printTree(n)
	}
}

func isLineValid(l string) bool {
	return l != "" && l != "\n"
}
