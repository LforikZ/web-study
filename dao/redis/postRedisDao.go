// @Author Zihao_Li 2023/1/30 14:01:00
package redis

import "web-study/entity"

func GetPostIDsInOrder(p *entity.ParamPostDataPlus) ([]string, error) {
	// 从redis中获取 ids
	key := GetRedisKey(KeyPostTimeZset)
	if p.Order == entity.OrderScore {
		key = GetRedisKey(KeyPostScoreZset)
	}
	start := (p.Page * p.Size) - p.Size
	end := start + p.Size - 1
	//根据key值从大到小查询数量
	return client.ZRevRange(key, start, end).Result()
}
