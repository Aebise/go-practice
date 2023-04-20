package db

import (
	"fmt"
	"go-practice/models"
)

func GetPosts() (posts []models.Post, err error) {
	posts = []models.Post{
		{
			ID:          "post1",
			ImageURL:    "image1",
			VideoURL:    "this is video",
			Views:       50,
			Like:        43,
			PosterID:    "poster1",
			CategoryID:  "painting",
			Status:      "Approved",
			Description: "Look at my new drawing.",
		}, {
			ID:          "post2",
			ImageURL:    "image2",
			VideoURL:    "this is video2",
			Views:       50,
			Like:        43,
			PosterID:    "poster2",
			CategoryID:  "drawing",
			Status:      "Approved",
			Description: "Look at my new drawing.",
		},
	}

	return posts, nil

}

func AddPost(post models.Post) (models.Post, error) {
	// add to db
	// return the added post including the post ID.

	fmt.Println(post)
	post.ID = "post1"
	return post, nil
}

func GetPost(id string) (models.Post, error) {
	fmt.Println("id: ", id)
	post := models.Post{
		ID:          id,
		ImageURL:    "image2",
		VideoURL:    "this is video2",
		Views:       50,
		Like:        43,
		PosterID:    "poster2",
		CategoryID:  "drawing",
		Status:      "Approved",
		Description: "Look at my new drawing.",
	}

	return post, nil
}

func UpdatePost(post models.Post) (models.Post, error) {
	fmt.Println(post)
	// update the post on db
	// return updated data

	return post, nil
}
