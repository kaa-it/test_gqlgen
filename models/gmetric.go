package models

import "log"

type GMetric struct {
	ItemsIds []int
}

func (gm *GMetric) Count() int {
	log.Println("Count called")
	return len(gm.ItemsIds)
}
