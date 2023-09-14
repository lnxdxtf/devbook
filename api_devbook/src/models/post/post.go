package models_post

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint      `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint      `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint      `json:"likes,omitempty"`
	ImageBase64 string  `json:"image_base64,omitempty"` // Used only in create and update post images
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (post *Post) validate() []error {
	var errorsArr []error

	if post.Title == "" {
		errorsArr = append(errorsArr, errors.New("title is required"))
	}

	if post.Content == "" {
		errorsArr = append(errorsArr, errors.New("content is required"))
	}

	if len(errorsArr) == 0 {
		return nil
	}

	return errorsArr
}

func (post *Post) format() error {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	return nil
}

func (post *Post) Prepare() []error {
	errorsArr := []error{}
	if errs := post.validate(); errs != nil {
		errorsArr = errs
	}
	err := post.format()
	if err != nil {
		errorsArr = append(errorsArr, err)
	}

	if len(errorsArr) == 0 {
		return nil
	}
	return errorsArr
}
