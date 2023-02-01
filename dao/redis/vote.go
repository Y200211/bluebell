package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票的限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2. 更新帖子分数
	//
	ov := client.ZScore(getRedisKey(KeyPostVotedPF+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	_, err := client.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID).Result()
	if ErrVoteTimeExpire != nil {
		return err
	}
	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		_, err = client.ZRem(getRedisKey(KeyPostVotedPF+postID), userID).Result()
	} else {
		_, err = client.ZAdd(getRedisKey(KeyPostVotedPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		}).Result()
	}
	return err

}
