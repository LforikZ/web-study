package service

import (
	"web-study/dao/mysql"
	"web-study/entity"
)

func InsertComData(community *entity.ParamInsertCommunity) (err error) {
	if err = mysql.InsertData(community); err != nil {
		return err
	}
	return
}
