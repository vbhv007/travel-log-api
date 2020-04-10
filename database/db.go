package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type logEntity struct {
	gorm.Model
	Title  string
	Description string
	Rating int
	ImageUrl string
	Latitude float64
	Longitude float64
}

var LogEntityDao LogEntityDaoT

type LogEntityDaoT interface {
	Save(newData *logEntity) error
	Migrate() error
	Find(data *logEntity) (*logEntity, error)
	Update(oldData *logEntity, newData *logEntity) error
}

func init() {
	dao := &LogEntityDaoImpl{}
	dao.init()
	LogEntityDao = dao
}

type LogEntityDaoImpl struct {
	log *logEntity
	db  *gorm.DB
}

func (dao *LogEntityDaoImpl) init() {
	dao.log = &logEntity{}
	db, err := gorm.Open("sqlite3", "travelDB.db")
	if err != nil {
		panic("failed to connect database")
	}
	dao.db = db
}

func (dao *LogEntityDaoImpl) Migrate() error {
	dao.db.AutoMigrate(&logEntity{})
	return nil
}

func (dao *LogEntityDaoImpl) Save(newData *logEntity) error {
	if dao.db.NewRecord(newData) {
		dao.db.Create(&newData)
		return nil
	}
	return fmt.Errorf("failed to insert data. Already exist")
}

func (dao *LogEntityDaoImpl) Find(data *logEntity) (*logEntity, error){
	var log *logEntity
	dao.db.First(&log, data.ID)
	return log, nil
	//dao.db.First(&log, "code = ?", "L1212") // find product with code l1212
}

func (dao *LogEntityDaoImpl) Update(oldData *logEntity, newData *logEntity) error {
	dao.db.Model(&oldData).Update(newData)
	return nil
}
