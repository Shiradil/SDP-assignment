package strategy

import "fmt"

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("Public Transport A:[%d] to B:[%d] Avg speed:[%d] Total:[%d] Total time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
