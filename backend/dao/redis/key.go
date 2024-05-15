package redis

// redis key

// redis key 注意使用命名空间的方式，方便查询和拆分

const (
	Prefix                = "bulubell:"
	KeyPostTimeZSet       = "post:time"   // zest;帖子以及发帖时间
	KeyPostScoreZSet      = "post:score"  // zset;帖子以及投票分数
	KeyPostVotedZetPrefix = "post:voted:" // zest;记录用户以及投票类型;参数是post id
	KeyCommunitySetPrefix = "community:"  // set;保存每个分区下帖子的id
)

// getRedisKey
func getRedisKey(key string) string {
	return Prefix + key
}
