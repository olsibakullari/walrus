package entities

import "errors"

var (
	ErrTopicNotFound      = errors.New("there isn't any subscriber listening to this topic")
	ErrTopicAlreadyExists = errors.New("this topic already exists")
)
