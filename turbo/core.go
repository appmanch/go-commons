package turbo

import (
	"log"
	"strings"
)

type TurboRouteMethod struct {
	registeredMethod string
}

type TurboRouteNode struct {
	key     string
	data    *TurboRouteNode
	methods []*TurboRouteMethod
}

type TurboRouteMap struct {
	root *TurboRouteNode
}

func (turboRouteMap *TurboRouteMap) insertNode(val string) *TurboRouteMap {
	if turboRouteMap.root == nil {
		log.Println("first insert")
		turboRouteMap.root = &TurboRouteNode{key: "/", data: nil, methods: nil}
	} else {
		log.Println("consecutive insert")
		turboRouteMap.root.insert(val)
	}
	return turboRouteMap
}

func (turboRouteNode *TurboRouteNode) insert(val string) {
	log.Printf("key : %s\n", turboRouteNode.key)
	log.Printf("data : %v", turboRouteNode.data)
	if turboRouteNode == nil {
		return
	} else if turboRouteNode.data == nil {
		log.Println("if key is nil")
		turboRouteNode.data = &TurboRouteNode{key: val, data: nil, methods: nil}
	} else if turboRouteNode.key == val {
		log.Println("key matches	")
		// TODO how to proceed when a route key matches to the already existing node key
		// keep calling insert if there is a match found
		turboRouteNode.insert(val)
	}
}

func printHierarchy(node *TurboRouteNode) {
	if node == nil {
		return
	}
	log.Print(node.key)
	printHierarchy(node.data)
}

func createMapping(path string) {
	// root to be kep "/", and rest of the path blocks to be added on top of that
	// for each route getting registered, it goes through this step and updates the nodes
	// the methods becomes the last part
	hierarchy := &TurboRouteMap{}

	log.Printf("CreateMapping: %s\n", path)
	pathArr := strings.Split(path, "/")
	log.Printf("CreateMapping: %s\n", pathArr)
	for _, val := range pathArr {

		log.Printf("CreateMapping: %s\n", val)
		hierarchy = hierarchy.insertNode(val)

		//insertNode(val)
	}
	printHierarchy(hierarchy.root)
}
