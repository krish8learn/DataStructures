package src

import "errors"

type Queue struct {
	Fqueue []string
}

//Enqueue
func (q *Queue) Enqueue(str string) {
	q.Fqueue = append(q.Fqueue, str)
}

//Dequeue
func (q *Queue) Dequeue() error {
	if len(q.Fqueue) > 0 {
		q.Fqueue = q.Fqueue[1:]
		return nil
	}
	return errors.New("Dequeue Error: Queue is empty")
}

//Front value of Q
func (q *Queue) Front() (string, error) {
	if len(q.Fqueue) > 0 {
		return q.Fqueue[0], nil
	}
	return "", errors.New("Queue is Empty")
}

//length of Queue
func (q Queue) Length() int {
	return len(q.Fqueue)
}

//Empty check
func (q *Queue) Empty() bool {
	return len(q.Fqueue) == 0
}
