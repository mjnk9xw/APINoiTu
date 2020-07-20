package models

import "sync"

type ModelsIf interface {
	CheckTu(preSeq, seq string) (bool, error)
}

var once sync.Once
var modelsIf ModelsIf

func GetModels() ModelsIf {
	once.Do(func() {
		modelsIf = newModels()
	})
	return modelsIf
}
