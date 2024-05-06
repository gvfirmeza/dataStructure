package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

type Queue struct {
    front *Node
    rear  *Node
}

func (q *Queue) enqueue(value int) {
    newNode := &Node{value, nil}

    if q.rear == nil {
        q.front = newNode
        q.rear = newNode
    } else {
        q.rear.next = newNode
        q.rear = newNode
    }

    fmt.Println("Elemento inserido:", value)
}

func (q *Queue) dequeue() int {
    if q.front == nil {
        fmt.Println("A fila está vazia")
        return -1
    }

    value := q.front.value
    q.front = q.front.next

    if q.front == nil {
        q.rear = nil
    }

    return value
}

func (q *Queue) traverse() {
    if q.front == nil {
        fmt.Println("A fila está vazia")
        return
    }

    current := q.front
    for current != nil {
        fmt.Println(current.value)
        current = current.next
    }
}

func main() {
    queue := Queue{}

    queue.enqueue(1)
    queue.enqueue(2)
    queue.enqueue(3)

	fmt.Println("Elementos na fila:")
    queue.traverse()

    fmt.Println("Elemento removido:", queue.dequeue())

    fmt.Println("Elementos na fila:")
    queue.traverse()

	queue.enqueue(4)

	fmt.Println("Elementos na fila:")
    queue.traverse()

    fmt.Println("Elemento removido:", queue.dequeue())

    fmt.Println("Elementos na fila:")
    queue.traverse()
}