package main

import (
	global "gonovel/internal"
	"gonovel/pkg/node"
)

func main() {
	node := node.CreateNode(global.Volume)
	println(node.Brief())
}
