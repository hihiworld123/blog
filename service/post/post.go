package postservice

import (
	"blog/common"
	"blog/service/domain"
	"time"
)

func Add(post domain.PostRequest) error {

	post.Post.CreatedAt = time.Now()
	post.Post.UpdatedAt = time.Now()
	tx := common.Db.Create(&post.Post)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
