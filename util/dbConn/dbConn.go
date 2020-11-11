package dbConn

/**
数据库连接封装
*/
import (
	"database/sql"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	//User "gitlab.ziguangcn.com/sellerhome/advertising/controller/user"
)

/**
casbin pgsql 权限管理连接数据库
*/
var CasbinEnforcer *casbin.Enforcer

/**
pgsql gorm 连接数据库
*/
var PgGormDb *gorm.DB

/**
pgsql 连接数据库
*/
var PgsqlConn *sql.DB

/**
mysql sqlx 连接数据库
*/
var MysqlConn *gorm.DB

/**
redis 连接
*/
var RedisConn redis.Conn

func init() {
	initViper()

	//initPostgreSQL(
	//	viper.GetString("postgreSQL.address"),
	//	viper.GetInt("postgreSQL.port"),
	//	viper.GetString("postgreSQL.dbname"),
	//	viper.GetString("postgreSQL.account"),
	//	viper.GetString("postgreSQL.password"))
	//initCasbinPGSQL(
	//	viper.GetString("postgreSQL.address"),
	//	viper.GetInt("postgreSQL.port"),
	//	viper.GetString("postgreSQL.dbname"),
	//	viper.GetString("postgreSQL.account"),
	//	viper.GetString("postgreSQL.password"))

	//initSQLite3()
	//initCasbinSQLite()

	//mysql
	inintMysql(viper.GetString("mysql.address"),
				viper.GetString("mysql.dbname"),
				viper.GetString("mysql.account"),
				viper.GetString("mysql.password"))

	//redis
	//inintRedis(viper.GetString("redis.address"))
}

//casbin sqlite 连接数据库
func initCasbinSQLite() {
	a, err := gormadapter.NewAdapter("sqlite3", "gorm.db", true)
	e, err := casbin.NewEnforcer("conf/auth_model.conf", a)
	if err != nil {
		panic(err)
	}
	CasbinEnforcer = e
}

//sqlite连接
func initSQLite3() {
	//gorm
	gdb, gerr := gorm.Open("sqlite3", "gorm.db")
	if gerr != nil {
		panic(gerr)
	}
	//全局禁用表名复数形式
	gdb.SingularTable(true)
	//全局启用gorm sql debug模式
	gdb.LogMode(viper.GetBool("gorm.debug"))
	PgGormDb = gdb
	PgGormDb.DB().SetMaxIdleConns(3)
	PgGormDb.DB().SetMaxOpenConns(30)
	fmt.Println(">>>>>>>>>>>SQLite init success")
}

//casbin pgsql 连接数据库
func initCasbinPGSQL(address string, port int, dbname string, user string, password string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", address, port, user, password, dbname)
	a, err := gormadapter.NewAdapter("postgres", psqlInfo, true)
	e, err := casbin.NewEnforcer("conf/auth_model.conf", a)
	if err != nil {
		panic(err)
	}
	CasbinEnforcer = e
}

//PostgreSQL连接
func initPostgreSQL(address string, port int, dbname string, user string, password string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", address, port, user, password, dbname)

	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//PgsqlConn = db
	//PgsqlConn.SetMaxIdleConns(3)
	//PgsqlConn.SetMaxOpenConns(30)

	//gorm
	gdb, gerr := gorm.Open("postgres", psqlInfo)
	if gerr != nil {
		panic(gerr)
	}
	//全局禁用表名复数形式
	gdb.SingularTable(true)
	//全局启用gorm sql debug模式
	gdb.LogMode(viper.GetBool("gorm.debug"))
	PgGormDb = gdb
	PgGormDb.DB().SetMaxIdleConns(3)
	PgGormDb.DB().SetMaxOpenConns(30)
	fmt.Println(">>>>>>>>>>>PostgreSQL init success")
}

//mysql数据库连接
func inintMysql(address string, dbname string, account string, passwort string) {
	d, err := gorm.Open("mysql", ""+account+":"+passwort+"@tcp("+address+")/"+dbname+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	MysqlConn = d
	//全局启用gorm sql debug模式
	MysqlConn.LogMode(viper.GetBool("gorm.debug"))
	//连接池设置
	MysqlConn.DB().SetMaxOpenConns(30) //设置最大打开的连接数
	MysqlConn.DB().SetMaxIdleConns(3)  //设置最大空闲连接数
	fmt.Println(">>>>>>>>>>>mysql init success")
}

//redis连接设置
func inintRedis(address string) {
	Pool := redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	RedisConn = Pool.Get()
	fmt.Println(">>>>>>>>>>>redis init success")
}

//初始化配置文件读取
func initViper() {
	viper.SetConfigName("config")             // name of config file (without extension)
	viper.SetConfigType("yaml")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/123/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.123") // call multiple times to add many search paths
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
