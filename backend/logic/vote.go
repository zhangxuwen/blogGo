package logic

import (
	"blogWeb/dao/redis"
	"blogWeb/models"
	"strconv"

	"go.uber.org/zap"
)

// 投票功能

// VoteForPost 为帖子投票的汉书
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", int8(p.Direction)))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
