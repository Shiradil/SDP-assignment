package main

import "awesomeProject/strategy"

var (
	start      = 10
	end        = 150
	strategies = []strategy.Strategy{
		&strategy.PublicTransportStrategy{},
		&strategy.RoadStrategy{},
		&strategy.WalkStrategy{},
	}
)

func main() {
	nav := strategy.Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
