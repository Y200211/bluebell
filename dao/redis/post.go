package redis

import (
	"bluebell/models"

	"github.com/go-redis/redis"
)

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取ID
	// 根据用户请求中携带的 order参数确定要查询的 key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.Orderscore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.确定查询索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return client.ZRevRange(key, start, end).Result()

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
