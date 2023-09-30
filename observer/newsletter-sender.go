package observer

import "fmt"

type NewsletterSender struct {
	Name       string
	UsersEmail map[string]User
}

func (ns *NewsletterSender) Subscribe(user User) {
	ns.UsersEmail[user.getEmail()] = user
}

func (ns *NewsletterSender) UnSubscribe(user User) {
	delete(ns.UsersEmail, user.getEmail())
	fmt.Println("Current number of subscribers: ", len(ns.UsersEmail))
}

func (ns *NewsletterSender) Notify() {
	for _, user := range ns.UsersEmail {
		user.update(ns.Name)
	}
}
