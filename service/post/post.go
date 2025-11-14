package postservice

import (
	"blog/common"
	"blog/entity"
	"blog/service/domain"
	"time"

	log "github.com/sirupsen/logrus"
)

func Add(post domain.PostRequest) error {

	post.Post.CreatedAt = time.Now()
	post.Post.UpdatedAt = time.Now()
	tx := common.Db.Create(&post.Post)
	if tx.Error != nil {
		log.Error("postservice Add err: ", tx.Error)
		return tx.Error
	}

	return nil
}

func View(post domain.PostRequest) (entity.Post, error) {

	post1 := entity.Post{}
	tx := common.Db.Where("id = ?", post.Id).First(&post1)
	if tx.Error != nil {
		log.Error("postservice View err: ", tx.Error)
		return post1, tx.Error
	}

	return post1, nil
}

func Query(post domain.PostRequest) ([]entity.Post, error) {

	posts := []entity.Post{}
	tx := common.Db.Where("1=1")
	if post.Title != "" {
		tx.Where("title LIKE ?", "%"+post.Title+"%")
	}
	if post.UserId != 0 {
		tx.Where("user_id = ?", post.UserId)
	}

	find := tx.Find(&posts)
	if find.Error != nil {
		log.Error("postservice Query err: ", tx.Error)
		return posts, find.Error
	}

	return posts, nil
}

func Update(post domain.PostRequest) error {

	post.UpdatedAt = time.Now()
	post1 := entity.Post{
		Title:     post.Title,
		Content:   post.Content,
		UpdatedAt: post.UpdatedAt,
	}
	tx := common.Db.Model(&post.Post).Updates(&post1)
	if tx.Error != nil {
		log.Error("postservice Update err: ", tx.Error)
		return tx.Error
	}

	return nil
}

func Delete(post domain.PostRequest) error {

	tx := common.Db.Delete(&post.Post)
	if tx.Error != nil {
		log.Error("postservice Delete err: ", tx.Error)
		return tx.Error
	}

	return nil
}
