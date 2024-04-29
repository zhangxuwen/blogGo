package controller

import (
	"blogWeb/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
	--- 社区模块 ---
*/

func CommunityHandler(c *gin.Context) {
	// 查询到所有社区 (community_id, community_name) 以列表的形式返回
	data, err := logic.GetcommunityList()
	if err != nil {
		zap.L().Error("logic.GetcommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不把服务端报错暴露外面
		return
	}
	ResponseSuccess(c, data)
}
