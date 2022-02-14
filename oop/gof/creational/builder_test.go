package creational

import (
	"testing"
)

// Client code
func TestBuilder(t *testing.T) {
	var concreteBuilder ConcreteBuilder = NewConcreteBuilder()
	// Director can accept different types of concrete builders
	var director Director = NewDirector(concreteBuilder)

	director.Construct()
	var product Product = concreteBuilder.GetResult()

	if !(product.partAIsBuilt && product.partBIsBuilt && product.partCIsBuilt) {
		t.Error("All product's parts should be built")
	}
}
