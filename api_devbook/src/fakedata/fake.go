package fakedata

import (
	models_post "api/src/models/post"
	models_user "api/src/models/user"
	post_repository "api/src/repositories/post"
	user_repository "api/src/repositories/user"

	"github.com/brianvoe/gofakeit/v6"
)

const RANGE_FAKE_DATA = 50

func FakeData() {
	for i := 0; i < RANGE_FAKE_DATA; i++ {

		userID := fakeUser()

		fakeFollow(userID)

		fakePost(userID)
	}
}

func fakeUser() uint {
	user := models_user.User{Name: gofakeit.Name(), Nick: gofakeit.Username(), Email: gofakeit.Email(), Pswrd: "123456"}
	repository, err := user_repository.NewRepository()
	if err != nil {
		panic(err)
	}
	if errs := user.Prepare(models_user.Signup); errs != nil {
		panic(errs)
	}

	userID, err := repository.Insert(user)
	if err != nil {
		panic(err)
	}

	return userID
}

func fakeFollow(userID uint) {
	repository, err := user_repository.NewRepository()
	if err != nil {
		panic(err)
	}
	userToFollow := gofakeit.Number(1, RANGE_FAKE_DATA)
	if err = repository.FollowUser(uint(userID), uint(userToFollow)); err != nil {
		panic(err)
	}
}

func fakePost(userID uint) {
	repository, err := post_repository.NewRepository()
	if err != nil {
		panic(err)
	}
	post := models_post.Post{Title: gofakeit.Quote(), Content: gofakeit.Phrase(), AuthorID: userID}
	if errs := post.Prepare(); errs != nil {
		panic(errs)
	}
	if _, err := repository.Insert(post); err != nil {
		panic(err)
	}
}
