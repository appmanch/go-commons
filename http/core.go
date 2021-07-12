package http

import (
	"log"
	"strings"
)

type TurboRouteMethod struct {
	registeredMethod string
}

type TurboRouteNode struct {
	key string
	data *TurboRouteNode
	methods []*TurboRouteMethod
}

type TurboRouteMap struct {
	root *TurboRouteNode
}

func (turboRouteMap *TurboRouteMap) insertNode(val string) *TurboRouteMap {
	if turboRouteMap.root == nil {
		turboRouteMap.root = &TurboRouteNode{key: "/", data: nil, methods: nil}
	} else {
		turboRouteMap.root.insert(val)
	}
	return turboRouteMap
}

func (turboRouteNode *TurboRouteNode) insert(val string) {
	if turboRouteNode == nil {
		return
	} else if turboRouteNode.key == "" {
		turboRouteNode.data = &TurboRouteNode{key: val, data: nil, methods: nil}
	} else if turboRouteNode.key == val {
		// TODO how to proceed when a route key matches to the already existing node key
	}
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
		if val != "" {
			log.Printf("CreateMapping: %s\n", val)
			hierarchy.insertNode(val)
		}
		//insertNode(val)
	}

	log.Printf("%v", hierarchy)
}