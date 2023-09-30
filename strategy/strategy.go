package strategy

type Strategy interface {
	Route(startPoint int, endPoint int)
}
