// Package redis @Author Zihao_Li 2023/1/29 16:29:00
package redis

// redis ket注意使用命名空间的方式，方便业务查询和拆分
const (
	KeyPrefix              = "web-study:"
	KeyPostTimeZset        = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZset       = "post:score"  // zset; 帖子及投票的分数
	KeyPostVotedZsetPrefix = "post:voted:" // zset: 记录用户及投票类型:(参数是 post id)
)

func GetRedisKey(key string) string {
	return KeyPrefix + key
}
