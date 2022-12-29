package main

import (
	"fmt"
)

// Queue  - Array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// New method initialises with Order with priority, quantity, product, customerName
func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

// Add method adds the order to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

// main method
func main() {
	var queue Queue
	queue = make(Queue, 0)

	var order1 *Order
	order1 = &Order{}
	var priority1 int
	priority1 = 2
	var quantity1 int
	quantity1 = 20
	var product1 string
	product1 = "Computer"
	var customerName1 string
	customerName1 = "Greg White"

	order1.New(priority1, quantity1, product1, customerName1)

	var order2 *Order
	order2 = &Order{}
	var priority2 int
	priority2 = 1
	var quantity2 int
	quantity2 = 10
	var product2 string
	product2 = "Monitor"
	var customerName2 string
	customerName2 = "John Smith"

	order2.New(priority2, quantity2, product2, customerName2)

	queue.Add(order1)
	queue.Add(order2)

	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
