package observer

type Subject interface {
	Subscribe(User)
	CancelSubscribe(User)
	Notify()
}
