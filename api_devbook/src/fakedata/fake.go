package fakedata

import (
	models_post "api/src/models/post"
	models_user "api/src/models/user"
	post_repository "api/src/repositories/post"
	user_repository "api/src/repositories/user"
	aws_devbook_s3 "api/src/services/aws"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"strings"

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
	user := models_user.User{Name: gofakeit.Name(), Nick: gofakeit.Username(), Email: gofakeit.Email(), Pswrd: "123456", ImageBase64: genRandomImage(64, 64)}
	repository, err := user_repository.NewRepository()
	if err != nil {
		panic(err)
	}
	if errs := user.Prepare(models_user.Signup); errs != nil {
		panic(errs)
	}

	result, err := repository.Insert(user)
	if err != nil {
		panic(err)
	}

	if user.ImageBase64 != "" {

		base64Data := user.ImageBase64
		if strings.Contains(base64Data, ",") {
			base64Data = strings.Split(base64Data, ",")[1]
		}
		decodeBase64Img, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			panic(err)
		}
		s3Service, err := aws_devbook_s3.NewS3Service()
		if err != nil {
			panic(err)
		}

		upload := aws_devbook_s3.UploadInput{
			File_name: fmt.Sprintf("user/profile/imgs/user_%d.png", result),
			File_type: "image/png",
			File_body: decodeBase64Img,
		}
		if err = s3Service.Upload(upload); err != nil {
			panic(err)
		}
	}

	return result
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

	post := models_post.Post{Title: gofakeit.Quote(), Content: gofakeit.Phrase(), AuthorID: userID, ImageBase64: genRandomImage(80, 80)}
	if errs := post.Prepare(); errs != nil {
		panic(errs)
	}
	result, err := repository.Insert(post)
	if err != nil {
		panic(err)
	}

	if post.ImageBase64 != "" {
		base64Data := post.ImageBase64
		if strings.Contains(base64Data, ",") {
			base64Data = strings.Split(base64Data, ",")[1]
		}
		decodeBase64Img, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			panic(err)
		}
		s3Service, err := aws_devbook_s3.NewS3Service()
		if err != nil {
			panic(err)
		}
		upload := aws_devbook_s3.UploadInput{
			File_name: fmt.Sprintf("user/post/imgs/post_%d.png", result),
			File_type: "image/png",
			File_body: decodeBase64Img,
		}
		if err = s3Service.Upload(upload); err != nil {
			panic(err)
		}

	}

}

func genRandomImage(width, height int) string {
	img := gofakeit.Image(width, height)
	var buffer bytes.Buffer
	if err := png.Encode(&buffer, img); err != nil {
		panic(err)
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return imgBase64Str
}
