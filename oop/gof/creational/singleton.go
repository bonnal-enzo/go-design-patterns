package creational

import (
	"fmt"
	"math/rand"
)

// ----------------------------------------------------------------------------
// CREATIONAL PATTERN: SINGLETON
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Public interface
type Singleton interface {
	SingletonOperation()
	GetSingletonData() float64
}

// Private concrete type
type concreteSingleton struct {
	singletonData float64
}

func (singleton concreteSingleton) SingletonOperation() {
	fmt.Printf("Performs singleton operation on '%f' data\n", singleton.singletonData)
}
func (singleton concreteSingleton) GetSingletonData() float64 {
	return singleton.singletonData
}

var singletonUniqueInstance *concreteSingleton

// Public function that lets clients access a unique singleton instance
func SingletonInstance() Singleton {
	if singletonUniqueInstance == nil {
		// in a concurrent context we should use a mutex lock here
		singletonUniqueInstance = &concreteSingleton{rand.Float64()}
	}
	return singletonUniqueInstance
}
