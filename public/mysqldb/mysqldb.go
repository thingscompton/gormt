package mysqldb

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/thingscompton/gormt/public/dev"
	myerrors "github.com/thingscompton/gormt/public/errors"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/thingscompton/gormt/public/mylog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySqlDB ...
type MySqlDB struct {
	*gorm.DB
}

// OnInitDBOrm init MySqlDB
func OnInitDBOrm(dataSourceName string, maxIdleConns, maxOpenConns int, ignoreRecordNotFoundError bool) (orm *MySqlDB) {
	orm = new(MySqlDB)
	orm.OnGetDBOrm(dataSourceName, maxIdleConns, maxOpenConns, ignoreRecordNotFoundError)
	return
}

// OnGetDBOrm get gorm.db
func (i *MySqlDB) OnGetDBOrm(dataSourceName string, maxIdleConns, maxOpenConns int, ignoreRecordNotFoundError bool) *gorm.DB {
	if i.DB == nil {
		Default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError,
			Colorful:                  true,
		}) //
		dialector := mysql.Open(dataSourceName)
		var err error
		i.DB, err = gorm.Open(dialector, &gorm.Config{PrepareStmt: false,
			NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 全局禁用表名复数
			Logger:         Default})                                   // logger.Default
		if err != nil {
			mylog.Error(myerrors.Wrap(err, "Got error when connect database:"+dataSourceName))
			i.DB = nil
			return nil
		}

		sqlDB, err := i.DB.DB()
		if err != nil || sqlDB == nil {
			mylog.Error(myerrors.Wrap(err, "Got error when get db pool"))
			i.DB = nil
			return nil
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(maxIdleConns) // 空闲连接池中最大连接数

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(maxOpenConns) // 设置数据库的最大打开连接数。
	}

	// i.DB.SingularTable(true) //全局禁用表名复数
	if dev.IsDev() {
		i.DB = i.DB.Debug()
	} else {
		i.DB.Logger = GetDBlog(ignoreRecordNotFoundError)
	}

	return i.DB
}

// OnDestoryDB destorydb
func (i *MySqlDB) OnDestoryDB() {
	if i.DB != nil {
		if sqldb, err := i.DB.DB(); err == nil && sqldb != nil {
			sqldb.Close()
		}
		i.DB = nil
	}
}

// IsNotFound 判断错误是否未找到
func (i *MySqlDB) IsNotFound(errs ...error) bool {
	if len(errs) > 0 {
		for _, err := range errs {
			if err == gorm.ErrRecordNotFound {
				return true
			}
		}
	}
	return i.RecordNotFound()
}

// RecordNotFound check if returning ErrRecordNotFound error
func (i *MySqlDB) RecordNotFound() bool {
	return !errors.Is(i.Error, gorm.ErrRecordNotFound)
}

// Commit 自动提交(如果有错，Rollback)
func (i *MySqlDB) Commit(db *gorm.DB) {
	if db.Error != nil {
		db.Rollback() // 回滚
	} else {
		db.Commit()
	}
}
