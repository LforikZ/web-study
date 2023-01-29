// @Author Zihao_Li 2023/1/29 18:40:00
package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

// 基于用户投票的相关算法：http://ruanyifeng.com/blog/algorithm/
// 本项目使用简化版的投票分数
// 投一票就加432分  86400秒除以200 -> 200张赞成票可以给帖子续一天 -> 《redis实战》
/* 投票的几种情况
direction = 1
	1. 之前没投过票，现在投赞成 差值为 1
	2. 之前投反对票，现在投赞成	差值为 2
direction = 0
	1. 之前投赞成票，现在取消	差值为 -1
	2. 之前投反对票，现在取消	差值为 1
direction = -1
	1. 之前没投过票，现在投反对	差值为 -1
	2. 之前投赞成票，现在投反对	差值为 -2

投票的限制:
每个帖子自发表之日起一个星期内可以进行投票，之后不允许投票
并将截止票数存储到数据库中去
然后将 KeyPostVotedZsetPrefix 从redis中删除
*/
const (
	oneWeekSecond = 7 * 24 * 3600 // 一周的秒
	scoreVote     = 432           //每一票值多少分
)

var (
	ErrVoteTimePass = errors.New("投票时间已过")
)

// VoteForPost
// @Description  实现投票计数功能
// @Author Zihao_Li 2023-01-29 19:44:18
// @Param userID
// @Param postID
// @Param value
// @Return error
func VoteForPost(userID, postID string, value float64) error {
	// 1、判断投票限制
	// 去redis中取发布时间
	postTime := client.ZScore(GetRedisKey(KeyPostTimeZset), postID).Val()
	//time.Now().Unix() 时间戳
	if float64(time.Now().Unix())-postTime > oneWeekSecond {
		return ErrVoteTimePass
	}

	// 2、更新帖子分数
	//查询当前用户对于当前帖子的投票记录
	ov := client.ZScore(GetRedisKey(KeyPostVotedZsetPrefix+postID), userID).Val()
	diff := value - ov

	//加个事务，将2 和 3捆绑起来
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(GetRedisKey(KeyPostScoreZset), diff*scoreVote, postID)

	// 3、记录用户更新之后的分数
	if value == 0 {
		//如果是取消  直接移除数据
		pipeline.ZRem(GetRedisKey(KeyPostVotedZsetPrefix+postID), postID)

	} else {
		pipeline.ZAdd(GetRedisKey(KeyPostVotedZsetPrefix+postID), redis.Z{
			Score:  value,  //表示是赞成票还是反对票
			Member: userID, //表示那个用户
		})
	}
	_, err := pipeline.Exec()
	return err
}

// CreatPost
// @Description  将创建帖子的id 传到redis中
// @Author Zihao_Li 2023-01-29 19:44:47
// @Param postID
func CreatPost(postID int64) error {
	//加入事务操作
	pipeline := client.TxPipeline()
	//帖子时间
	pipeline.ZAdd(GetRedisKey(KeyPostTimeZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//帖子分数
	pipeline.ZAdd(GetRedisKey(KeyPostScoreZset), redis.Z{
		Score:  0,
		Member: postID,
	})

	_, err := pipeline.Exec()
	return err
}
