package collection

type Collection[T any] struct {
	items []T
}

func New[T any]() Collection[T] {
	collection := Collection[T]{}
	return collection
}

func Collect[T any](items []T) Collection[T] {
	c := Collection[T]{}
	c.items = items
	return c
}

func (c *Collection[T]) Add(item T) {
	c.items = append(c.items, item)
}

func (c *Collection[T]) Items() []T {
	return c.items
}

func (c *Collection[T]) Delete(item T) {
}
