package hcpairing

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Start() error
	SetupTags()
	AppendRecord(state string, tags []string)
	GetRecords() []Record
}

type database struct {
	connection *gorm.DB
	dsn        string
}

type Record struct {
	gorm.Model
	State string `json:"state"`
	Tags  []Tag  `gorm:"many2many:record_tags;" json:"tags"`
}

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

var (
	DBConn Database = NewDatabase()
)

func NewDatabase() Database {

	instance := database{
		connection: nil,
		dsn: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			Config.GetPostgresHost(),
			Config.GetPostgresUser(),
			Config.GetPostgresPassword(),
			Config.GetPostgresDBName(),
			Config.GetPostgresPort(),
		),
	}
	return &instance
}

func (d *database) Start() error {

	var err error
	d.connection, err = gorm.Open(postgres.Open(d.dsn), &gorm.Config{})
	d.connection.AutoMigrate(&Record{}, &Tag{})
	return err
}

func (d *database) SetupTags() {

	for _, tag := range allTags {
		result := d.connection.Where("name = ?", tag).First(&Tag{})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			d.connection.Create(&Tag{Name: tag})
		} else {
			Logger.Debug("tag " + tag + " already exists")
		}
	}
}

func (d *database) AppendRecord(state string, tags []string) {

	recordTags := []Tag{}
	for _, tag := range tags {
		result := Tag{}
		d.connection.Where("name = ?", tag).First(&result)
		recordTags = append(recordTags, result)
	}
	d.connection.Create(
		&Record{
			State: state,
			Tags:  recordTags,
		},
	)
}

func (d *database) GetRecords() []Record {

	records := []Record{}
	d.connection.Preload("Tags").Find(&records)
	return records
}
