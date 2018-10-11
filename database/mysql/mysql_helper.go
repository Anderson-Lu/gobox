package mysql_helper

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
) 

func NewMysqlClient(host, dbname, user, pwd string) (gorm.DB*, error) {	
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbname))
	return &CRUD{
		client:      db,
		redisHelper: redisHelper,
	}, err
} 