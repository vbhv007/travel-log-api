package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vbhv007/travel-log-api/dto"
)

const (
	DBName = "travelDB.db"
	DBType = "sqlite3"
)

var LogEntityDao LogEntityDaoT

type LogEntityDaoT interface {
	Save(newData *dto.LogEntity) error
	Migrate() error
	Find(condition interface{}) ([]*dto.LogEntity, error)
	Update(oldData *dto.LogEntity, newData *dto.LogEntity) error
}

func init() {
	dao := &LogEntityDaoImpl{}
	dao.init()
	LogEntityDao = dao
}

type LogEntityDaoImpl struct {
	log *dto.LogEntity
	db  *gorm.DB
}

func (dao *LogEntityDaoImpl) init() {
	dao.log = &dto.LogEntity{}
	db, err := gorm.Open(DBType, DBName)
	if err != nil {
		panic("failed to connect database")
	}
	dao.db = db
	dbErr := dao.Migrate()
	if dbErr != nil {
		panic("Unable to migrate DB")
	}
}

func (dao *LogEntityDaoImpl) Migrate() error {
	dao.db.Debug().AutoMigrate(&dto.LogEntity{})
	return nil
}

func (dao *LogEntityDaoImpl) Save(newData *dto.LogEntity) error {
	if dao.db.NewRecord(newData) {
		dao.db.Create(&newData)
		return nil
	}
	return fmt.Errorf("failed to insert data. Already exist")
}

func (dao *LogEntityDaoImpl) Find(condition interface{}) ([]*dto.LogEntity, error) {
	var logs []*dto.LogEntity
	dao.db.Where(condition).Find(&logs)
	return logs, nil
}

func (dao *LogEntityDaoImpl) Update(oldData *dto.LogEntity, newData *dto.LogEntity) error {
	dao.db.Model(&oldData).Update(newData)
	return nil
}
