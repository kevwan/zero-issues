package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func Conn(dns string) *gorm.DB {
	//dsn1 := "zhangzhuang:Zhang123456@tcp(rm-uf63x44czp1i998i41o.mysql.rds.aliyuncs.com:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		DisableAutomaticPing:                     true, //在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: true, //在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
	})

	dbs, _ := db.DB()
	dbs.SetMaxIdleConns(10)           //连接池中最大的空闲连接数
	dbs.SetMaxOpenConns(100)          //连接池中最大可以容纳多少链接数
	dbs.SetConnMaxLifetime(time.Hour) //连接池中链接最大可可复用时间

	return db
}
