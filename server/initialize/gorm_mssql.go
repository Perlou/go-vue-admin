package initialize

import (
	"github.com/perlou/go-vue-admin/server/config"
	"github.com/perlou/go-vue-admin/server/global"
	"github.com/perlou/go-vue-admin/server/initialize/internal"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GormMssql 初始化Mssql数据库
func GormMssql() *gorm.DB {
	m := global.GVA_CONFIG.Mssql
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMssqlByConfig 初始化Mysql数据库用过传入配置
func GormMssqlByConfig(m config.Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
