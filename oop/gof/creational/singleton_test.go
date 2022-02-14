package creational

import "testing"

// Client code
func TestSingleton(t *testing.T) {
	var singleton1 Singleton = SingletonInstance()
	singleton1.SingletonOperation()
	var singleton2 Singleton = SingletonInstance()
	singleton2.SingletonOperation()

	if singleton1 != singleton2 { // test that both pointers point to the same instance
		t.Error("both singleton variables must be pointers that points to the same instance")
	}
	if singleton1.GetSingletonData() != singleton2.GetSingletonData() { // test that both pointers point to the same instance
		t.Error("both singleton variables must have same data")
	}
}
