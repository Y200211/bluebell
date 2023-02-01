package redis

const (
	KeyPrefix        = "bluebell:"
	KeyPostTimeZSet  = "post:time"
	KeyPostScoreZSet = "post:score"
	KeyPostVotedPF   = "post:voted:"
)

// 给 redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
