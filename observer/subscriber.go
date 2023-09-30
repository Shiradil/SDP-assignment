package observer

import "fmt"

type Subscriber struct {
	Email string
}

func (u *Subscriber) update(nsname string) {
	fmt.Printf("Sending to subscriber %s from newsletter sender %s\n", u.getEmail(), nsname)
}

func (u *Subscriber) getEmail() string {
	return u.Email
}
