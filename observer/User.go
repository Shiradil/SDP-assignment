package observer

type User interface {
	update(nsName string)
	getEmail() string
}
