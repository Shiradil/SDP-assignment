package main

import (
	"awesomeProject/observer"
)

func main() {
	usr1 := &observer.Subscriber{Email: "1user1@gmail.com"}
	usr2 := &observer.Subscriber{Email: "2user2@gmail.com"}
	usr3 := &observer.Subscriber{Email: "3user3@gmail.com"}
	usr4 := &observer.Subscriber{Email: "4user4@gmail.com"}
	usr5 := &observer.Subscriber{Email: "5user5@gmail.com"}

	channel := observer.NewsletterSender{
		Name:       "channel",
		UsersEmail: map[string]observer.User{},
	}
	channel.Subscribe(usr1)
	channel.Subscribe(usr2)
	channel.Subscribe(usr3)
	channel.Subscribe(usr4)
	channel.Subscribe(usr5)
	channel.Notify()
	println("Unsubscribing usr3")
	channel.UnSubscribe(usr3)
	channel.Notify()

}
