package commentservice

import (
	"blog/common"
	"blog/entity"
	"blog/service/domain"
	"time"

	log "github.com/sirupsen/logrus"
)

func PostComment(request domain.CommentRequest) error {

	request.Comment.CreatedAt = time.Now()
	tx := common.Db.Create(&request.Comment)
	if tx.Error != nil {
		log.Error("commentservice PostComment err: ", tx.Error)
		return tx.Error
	}

	return nil
}

func AllComment(request domain.CommentRequest) ([]entity.Comment, error) {

	comments := []entity.Comment{}
	tx := common.Db.Where("1=1")
	if request.Content != "" {
		tx.Where("content LIKE ?", "%"+request.Content+"%")
	}
	if request.PostId != 0 {
		tx.Where("post_id = ?", request.PostId)
	}

	find := tx.Find(&comments)
	if find.Error != nil {
		log.Error("commentservice AllComment err: ", tx.Error)
		return comments, find.Error
	}

	return comments, nil
}
