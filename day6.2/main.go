package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const santa = "SAN"
const you = "YOU"

type celestialBody struct {
	name          string
	bodiesInOrbit []*celestialBody
	inOrbitOf     []*celestialBody
}

func main() {
	bodiesLookup := make(map[string]*celestialBody)
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	inputArray := strings.Split(strings.Trim(inputString, "\n"), "\n")

	for _, celestialBodyConnection := range inputArray {
		celestialBodies := strings.Split(celestialBodyConnection, ")")
		bodyOrbited := createBodyIfNoneExists(celestialBodies[0], bodiesLookup)
		bodyOrbiting := createBodyIfNoneExists(celestialBodies[1], bodiesLookup)
		bodyOrbited.bodiesInOrbit = append(bodyOrbited.bodiesInOrbit, bodyOrbiting)
		bodyOrbiting.inOrbitOf = append(bodyOrbiting.inOrbitOf, bodyOrbited)
	}

	fmt.Println(findDistanceBetween(bodiesLookup[santa], bodiesLookup[you]))
}

func createBodyIfNoneExists(name string, bodiesLookup map[string]*celestialBody) *celestialBody {
	if bodiesLookup[name] == nil {
		newBody := celestialBody{
			name: name,
		}
		bodiesLookup[name] = &newBody
		return &newBody
	}
	return bodiesLookup[name]
}

func findDistanceBetween(bodyOne, bodyTwo *celestialBody) int {
	bodyOneConnectionsAndDistances := make(map[*celestialBody]int)
	bodyTwoConnectionsAndDistances := make(map[*celestialBody]int)
	bodyOneConnectionsAndDistances[bodyOne] = 0
	bodyTwoConnectionsAndDistances[bodyTwo] = 0
	bodyOneNext := []*celestialBody{bodyOne}
	bodyTwoNext := []*celestialBody{bodyTwo}

	for i := 0; i < len(bodyOneNext) && i < len(bodyTwoNext); i++ {
		currentBodyOneEntry := bodyOneNext[i]
		currentBodyTwoEntry := bodyTwoNext[i]

		bodyOneNodesToAdd, connectionAtOne := visit(currentBodyOneEntry, bodyOneConnectionsAndDistances, bodyOneConnectionsAndDistances[currentBodyOneEntry]+1, bodyTwoConnectionsAndDistances)
		if connectionAtOne != nil {
			return bodyOneConnectionsAndDistances[connectionAtOne] + bodyTwoConnectionsAndDistances[connectionAtOne] - 2
		}

		bodyTwoNodesToAdd, connectionAtTwo := visit(currentBodyTwoEntry, bodyTwoConnectionsAndDistances, bodyOneConnectionsAndDistances[currentBodyOneEntry]+1, bodyOneConnectionsAndDistances)
		if connectionAtTwo != nil {
			return bodyOneConnectionsAndDistances[connectionAtTwo] + bodyTwoConnectionsAndDistances[connectionAtTwo] - 2
		}

		bodyOneNext = append(
			bodyOneNext,
			bodyOneNodesToAdd...,
		)

		bodyTwoNext = append(
			bodyTwoNext,
			bodyTwoNodesToAdd...,
		)
	}

	return 0
}

func visit(body *celestialBody, bodiesVisited map[*celestialBody]int, distance int, bodiesToSearchIn map[*celestialBody]int) ([]*celestialBody, *celestialBody) {
	fmt.Println("Visited a node", body.name, distance)
	var adjacentBodies []*celestialBody
	var newBodiesToVisit []*celestialBody
	adjacentBodies = append(adjacentBodies, body.inOrbitOf...)
	adjacentBodies = append(adjacentBodies, body.bodiesInOrbit...)
	for _, adjacentBody := range adjacentBodies {
		if _, ok := bodiesVisited[adjacentBody]; !ok {
			bodiesVisited[adjacentBody] = distance
			newBodiesToVisit = append(newBodiesToVisit, adjacentBody)
		}
		if _, ok := bodiesToSearchIn[adjacentBody]; ok {
			fmt.Println("Found a path")
			return nil, adjacentBody
		}
	}
	return newBodiesToVisit, nil
}
