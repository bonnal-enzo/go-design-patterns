package creational

// ----------------------------------------------------------------------------
// CREATIONAL PATTERN: Builder
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Represents the complex object under construction
type Product struct {
	partAIsBuilt bool
	partBIsBuilt bool
	partCIsBuilt bool
}

// ----------------------------------------------------------------------------
// Specify an abstract interface for creating parts of a Product object.
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
}

// ----------------------------------------------------------------------------
// Conbstructs and assembles parts of the product by implementing the Builder interface
type ConcreteBuilder struct {
	// Defines and keeps track of the representation it creates
	product *Product // pointer to product that will be constructed inplace
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{&Product{}}
}

func (builder ConcreteBuilder) BuildPartA() {
	builder.product.partAIsBuilt = true
}
func (builder ConcreteBuilder) BuildPartB() {
	builder.product.partBIsBuilt = true
}
func (builder ConcreteBuilder) BuildPartC() {
	builder.product.partCIsBuilt = true
}

// Provides an interface for retrieving the product
func (builder ConcreteBuilder) GetResult() Product {
	return *(builder.product)
}

// ----------------------------------------------------------------------------
// Construct an object using the Builder interface
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) Director {
	return Director{builder}
}
func (director Director) Construct() {
	director.builder.BuildPartA()
	director.builder.BuildPartB()
	director.builder.BuildPartC()
}
