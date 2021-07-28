package wikisearch

type BotServise interface {
	Start() error
	Send(Message) error
	HandleMessage(func(Message))
}

type Message struct {
	User User
	Text string
}
