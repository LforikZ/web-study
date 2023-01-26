package service

import (
	"web-study/dao/mysql"
	"web-study/entity"
	"web-study/pkg/snowflake"
)

func CreatPost(data *entity.ParamPostData) (err error) {
	// 生成post id
	data.PostID = snowflake.GenID()
	// 保存数据库
	if err = mysql.InsertPostData(data); err != nil {
		return err
	}
	// 返回
	return nil
}
