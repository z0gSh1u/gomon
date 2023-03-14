package core

import (
	"errors"
)

type Comment struct {
	Id       string `json:"_id"`
	TopicId  string `json:"topic_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Censored bool   `json:"censored"`
}

func CreateComment(comment *Comment) error {
	// check the comment is legal
	if len(comment.TopicId) == 0 {
		return errors.New("invalid TopicId")
	}
	if len(comment.Name) == 0 {
		return errors.New("empty Name")
	}
	if len(comment.Content) == 0 {
		return errors.New("empty Content")
	}
	comment.Censored = true

	daoResponse := DAOInsertComment(comment)
	if daoResponse.Status == DAOFailed {
		return errors.New(daoResponse.Msg)
	}

	return nil
}

func SelectComments(topicId string) ([]Comment, error) {
	if len(topicId) == 0 {
		return nil, errors.New("invalid TopicId")
	}
	daoResponse := DAOSelectCommentsByTopicId(topicId)

	if daoResponse.Status == DAOFailed {
		return nil, errors.New(daoResponse.Msg)
	}

	return daoResponse.Payload.([]Comment), nil
}
