package creational

import "testing"

// Clients in action
func TestAbstractFactory(t *testing.T) {
	var client1 *Client = NewClient(&ConcreteFactory1{})
	client1.ShowAProduct()

	var client2 *Client = NewClient(&ConcreteFactory2{})
	client2.ShowAProduct()
}
