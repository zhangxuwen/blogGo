package controller

import (
	"blogWeb/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
	--- 社区模块 ---
*/

// CommunityHandler 获取community社区列表
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

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {

	// 获取社区id
	idStr := c.Param("id")                     // 获取 URL 参数
	id, err := strconv.ParseInt(idStr, 10, 64) // 转换后的进制类型为10， 转换后的整数位数为int64
	if err != nil {
		ResponseError(c, CodeIvalidParam)
		return
	}

	// 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
