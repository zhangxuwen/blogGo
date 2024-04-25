package mysql

import (
	"blogWeb/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		// viper.GetString("mysql.user"),
		cfg.User,
		// viper.GetString("mysql.password"),
		cfg.Password,
		// viper.GetString("mysql.host"),
		cfg.Host,
		// viper.GetInt("mysql.port"),
		cfg.Port,
		// viper.GetString("mysql.dbname"),
		cfg.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB error", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.max_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	return
}

func Close() {
	_ = db.Close()
}
