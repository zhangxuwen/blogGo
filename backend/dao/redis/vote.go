package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

/*
投票的几种情况
derection=1时，有两种情况
	1. 之前没有投过票，现在投赞成票 --> 更新分数和投票记录 差值的绝对值：1 +432
	2. 之前投反对票，现在改投赞成票 --> 更新分数和投票记录 差值的绝对值：2 +432 * 2
derection=0时，有两种情况
	1. 之前投过赞成票，现在要取消投票 --> 更新分数和投票纪录 差值的绝对值：1 -432
	2. 之前投过反对票，现在要取消投票 --> 更新分数和投票记录 差值的绝对值：1 +432
derection=-1时，有两种情况
	1. 之前没有投过票，现在投反对票 --> 更新分数和投票记录 差值的绝对值：1 -432
	2. 之前投赞成票，现在改投反对票 --> 更新分数和投票记录 差值的绝对值：2 -432*2

投票的限制
每个帖子自发表之日起一个星期内允许用户投票，超过一个星期就不允许再投票了
	1. 到期之后将 redis 中保存的赞成票及反对票数存储到 mysql 表中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekInseconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepested   = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) error {

	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()

	// 更新：把帖子id加到社区的set
	ckey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(ckey, postID)

	_, err := pipeline.Exec()

	return err
}

func VoteForPost(userID, postID string, value float64) (err error) {
	// 判断投票限制
	// 去redis取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInseconds {
		return ErrVoteTimeExpire
	}

	// 更新帖子分数
	// 先查当前用户给当前帖子的投票纪录
	oValue := rdb.ZScore(getRedisKey(KeyPostVotedZetPrefix+postID), userID).Val()

	// 如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == oValue {
		return ErrVoteRepested
	}

	var dir float64
	if value > oValue {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(oValue - value) // 计算两次投票的差值

	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID).Result()
	// 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZetPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZetPrefix+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return
}
