package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"web-study/entity"
)

type Community struct {
	gorm.Model
	CommunityID   int    `json:"communityID" gorm:"community_id;index"`
	CommunityName string `json:"communityName" gorm:"community_name;index"`
	Introduction  string `json:"introduction" gorm:"introduction"`
}

func FindCommunityById(communityId int) (community Community, err error) {
	if result := db.Where("community_id=?", communityId).Find(&community); result.Error == sql.ErrNoRows {
		zap.L().Warn("this is no community in db")
		err = result.Error
	}
	return community, err
}

func FindCommunityByName(name string) Community {
	var community Community
	db.Where("community_name=?", name).Find(&community)
	return community
}

func FindCommunityList() (a []entity.ParamListCommunity, err error) {
	var communityList []Community
	if result := db.Select("community_id", "community_name").Find(&communityList); result.Error == sql.ErrNoRows {
		zap.L().Warn("this is no community in db")
		err = result.Error
	}
	for _, list := range communityList {
		var middleList entity.ParamListCommunity
		middleList.CommunityID = list.CommunityID
		middleList.CommunityName = list.CommunityName
		a = append(a, middleList)
	}
	return a, err
}

func InsertCommunityData(result *entity.ParamInsertCommunity) (err error) {
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
