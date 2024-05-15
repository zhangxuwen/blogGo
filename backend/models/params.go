package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// 定义请求的参数结构体

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamtypeVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前用户
	PostID    string `json:"post_id,string" binding:"required"`       // 帖子id
	Direction int    `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票（1）还是反对票（-1）
}

// ParmaPostList 获取帖子列表query参数
type ParmaPostList struct {
	CommunityID int64  `json:"Community_id" from:"community_id"`
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}

// ParamCommunityPostList 按社区获取帖子列表 query string 参数
// type ParamCommunityPostList struct {
//	CommunityID int64 `json:"Community_id" from:"community_id"`
//	*ParmaPostList
// }
