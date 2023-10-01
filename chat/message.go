package chat

type Message struct {
	Text string `db:"message"`
	Id   int    `db:"id"`
}
