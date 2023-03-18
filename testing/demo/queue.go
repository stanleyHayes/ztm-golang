package queue

type Queue struct {
	Items    []int
	Capacity int
}

func (q *Queue) New(capacity int) {
	q.Items = make([]int, 0, capacity)
	q.Capacity = capacity
}

func (q *Queue) Append(item int) bool {
	if len(q.Items) == q.Capacity {
		return false
	}
	q.Items = append(q.Items, item)
	return true
}

func (q *Queue) Next() (int, bool) {
	if len(q.Items) > 0 {
		next := q.Items[0]
		q.Items = q.Items[1:]
		return next, true
	}
	return 0, false
}

func main() {

}
