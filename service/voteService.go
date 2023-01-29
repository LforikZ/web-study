// @Author Zihao_Li 2023/1/29 16:32:00
package service

import (
	"strconv"
	"web-study/dao/redis"
	"web-study/entity"
)

func PostVote(userID int64, p *entity.ParamVoteData) error {
	//zap.L().Debug("service.PostVote(userID, p) failed",
	//	zap.Int64("userID", userID),
	//	zap.String("postID", p.PostID),
	//	zap.Int("direction", p.Direction),
	//)
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
