package logic

import (
	"blogWeb/dao/mysql"
	"blogWeb/models"
)

// GetcommunityList 获取社区列表的逻辑
func GetcommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的 community 并返回
	return mysql.GetCommunityList()

}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
