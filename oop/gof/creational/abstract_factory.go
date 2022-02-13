package creational

import "fmt"

// ----------------------------------------------------------------------------
// CREATIONAL PATTERN: ABSTRACT FACTORY
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Declares an interface for a type of product object
type AbstractProduct interface {
	String() string // Returns information about product's concrete type
}

// ----------------------------------------------------------------------------
// Defines a product object to be created by the ConcreteFactory1 type
// Implements the AbstractProduct interface
type ConcreteProduct1 struct{}

func (concreteProduct ConcreteProduct1) String() string {
	return "Hi, I am a concrete product of type 1"
}

// ----------------------------------------------------------------------------
// Defines a product object to be created by the ConcreteFactory2 type
// Implements the AbstractProduct interface
type ConcreteProduct2 struct{}

func (concreteProduct ConcreteProduct2) String() string {
	return "Hi, I am a concrete product of type 2"
}

// ----------------------------------------------------------------------------
// Declares an interface for operations that create AbstractProduct objects
type AbstractFactory interface {
	CreateProduct() AbstractProduct
	// Other optional methods to produce other unrelated types of abstract products:
	/* CreateProductB() AbstractProductB */
	/* CreateProductC() AbstractProductC */
	/* ... */
}

// ----------------------------------------------------------------------------
// Implements the operation to create ConcreteProduct1 objects
type ConcreteFactory1 struct{}

func (concreteFactory ConcreteFactory1) CreateProduct() AbstractProduct {
	return &ConcreteProduct1{}
}

// ----------------------------------------------------------------------------
// Implements the operation to create ConcreteProduct2 objects
type ConcreteFactory2 struct{}

func (concreteFactory ConcreteFactory2) CreateProduct() AbstractProduct {
	return &ConcreteProduct2{}
}

// ----------------------------------------------------------------------------
// Uses only AbstractFactory and AbstractProduct interfaces
type Client struct {
	factory AbstractFactory
}

func NewClient(factory AbstractFactory) *Client {
	// Dependency Injection
	return &Client{factory}
}

// Instanciate a product and prints it
func (client Client) ShowAProduct() {
	fmt.Println(client.factory.CreateProduct())
}

// ----------------------------------------------------------------------------
// Clients in action
func main() {
	var client1 *Client = NewClient(&ConcreteFactory1{})
	client1.ShowAProduct()

	var client2 *Client = NewClient(&ConcreteFactory2{})
	client2.ShowAProduct()
}
