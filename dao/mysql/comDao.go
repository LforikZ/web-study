package mysql

import (
	"gorm.io/gorm"
	"web-study/entity"
)

type Community struct {
	gorm.Model
	CommunityID   int    `json:"communityID" gorm:"communityId;index"`
	CommunityName string `json:"communityName" gorm:"communityName;index"`
	Introduction  string `json:"introduction" gorm:"introduction"`
}

func FindById(communityId int) Community {
	var community Community
	db.Where("community_id=?", communityId).Find(&community)
	return community
}

func FindByName(name string) Community {
	var community Community
	db.Where("community_name=?", name).Find(&community)
	return community
}

func FindList() []Community {
	var communityList []Community
	db.Table("communities").Select("community_id", "community_name").Scan(&communityList)
	return communityList
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
