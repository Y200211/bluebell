package redis

import "bluebell/models"

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
