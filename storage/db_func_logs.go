package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DBName = "travelDB.db"
	DBType = "sqlite3"
)

var LogEntityDao LogEntityDaoT

type LogEntityDaoT interface {
	Save(newData *LogEntity) error
	Migrate() error
	Find(condition interface{}) ([]*LogEntity, error)
	Update(oldData *LogEntity, newData *LogEntity) error
}

func init() {
	dao := &LogEntityDaoImpl{}
	dao.init()
	LogEntityDao = dao
}

type LogEntityDaoImpl struct {
	log *LogEntity
	db  *gorm.DB
}

func (dao *LogEntityDaoImpl) init() {
	dao.log = &LogEntity{}
	db, err := gorm.Open(DBType, DBName)
	if err != nil {
		panic("failed to connect storage")
	}
	dao.db = db
	dbErr := dao.Migrate()
	if dbErr != nil {
		panic("Unable to migrate DB")
	}
}

func (dao *LogEntityDaoImpl) Migrate() error {
	dao.db.Debug().AutoMigrate(&LogEntity{})
	return nil
}

func (dao *LogEntityDaoImpl) Save(newData *LogEntity) error {
	if dao.db.NewRecord(newData) {
		dao.db.Create(&newData)
		return nil
	}
	return fmt.Errorf("failed to insert data. Already exist")
}

func (dao *LogEntityDaoImpl) Find(condition interface{}) ([]*LogEntity, error) {
	var logs []*LogEntity
	dao.db.Where(condition).Find(&logs)
	return logs, nil
}

func (dao *LogEntityDaoImpl) Update(oldData *LogEntity, newData *LogEntity) error {
	dao.db.Model(&oldData).Update(newData)
	return nil
}
