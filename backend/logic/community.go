package logic

import (
	"blogWeb/dao/mysql"
	"blogWeb/models"
)

func GetcommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的 community 并返回
	return mysql.GetCommunityList()

}
