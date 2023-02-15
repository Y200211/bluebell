package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func GetIDsFromKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取ID
	// 根据用户请求中携带的 order参数确定要查询的 key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.Orderscore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.确定查询索引起始点
	return GetIDsFromKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedPF + id)
	//	v1 := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v1)
	//}
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIdsInOrder 按社区查找
func GetCommunityPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取ID
	// 根据用户请求中携带的 order参数确定要查询的 key
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.Orderscore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	if client.Exists(key).Val() < 1 {
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return GetIDsFromKey(key, p.Page, p.Size)
}
