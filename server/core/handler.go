package core

type Comment struct {
	Id       string `json:"_id"`
	TopicId  string `json:"topic_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Censored bool   `json:"censored"`
}

func CreateComment(comment *Comment) {
	InsertComment(comment)
}
