package mysql

import (
	"gorm.io/gorm"
	"web-study/entity"
)

type Community struct {
	gorm.Model
	CommunityID   int    `json:"communityID" gorm:"communityId,uniqueIndex"`
	CommunityName string `json:"communityName" gorm:"communityName,uniqueIndex"`
	Introduction  string `json:"introduction" gorm:"introduction"`
}

func InsertData(result *entity.ParamInsertCommunity) (err error) {
	community := Community{
		CommunityID:   result.CommunityID,
		CommunityName: result.CommunityName,
		Introduction:  result.Introduction,
	}
	if result := db.Create(&community); result.Error != nil {
		return result.Error
	}
	return
}
