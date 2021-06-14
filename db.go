package hcpairing

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	setupTags()
	Start() error
	AppendRecord(state string, tags []string)
	GetRecordsByZipcode(zipcode string) []APIRecord
}

type database struct {
	connection *gorm.DB
	dsn        string
}

type Record struct {
	gorm.Model
	Zipcode string `sql:"DEFAULT:''" json:"zipcode"`
	Tags    []Tag  `gorm:"many2many:record_tags;" json:"tags"`
}

type APIRecord struct {
	Zipcode string   `json:"zipcode"`
	Tags    []string `json:"tags"`
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

func recordToAPISpec(records []Record) []APIRecord {

	apiRecords := []APIRecord{}
	for _, item := range records {
		tags := []string{}
		for _, tag := range item.Tags {
			tags = append(tags, tag.Name)
		}
		apiRecords = append(
			apiRecords,
			APIRecord{
				Zipcode: item.Zipcode,
				Tags:    tags,
			},
		)
	}
	return apiRecords
}

func (d *database) Start() error {

	var err error
	d.connection, err = gorm.Open(postgres.Open(d.dsn), &gorm.Config{})
	d.connection.AutoMigrate(&Record{}, &Tag{})
	d.setupTags()
	return err
}

func (d *database) setupTags() {

	for _, tag := range allTags {
		result := d.connection.Where("name = ?", tag).First(&Tag{})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			d.connection.Create(&Tag{Name: tag})
		} else {
			Logger.Debug("tag " + tag + " already exists")
		}
	}
}

func (d *database) AppendRecord(zipcode string, tags []string) {

	recordTags := []Tag{}
	for _, tag := range tags {
		result := Tag{}
		d.connection.Where("name = ?", tag).First(&result)
		recordTags = append(recordTags, result)
	}
	d.connection.Create(
		&Record{
			Zipcode: zipcode,
			Tags:    recordTags,
		},
	)
}

func (d *database) GetRecordsByZipcode(zipcode string) []APIRecord {

	records := []Record{}
	if zipcode != "" {
		d.connection.Preload("Tags").Where("zipcode = ?", zipcode).Find(&records)
	} else {
		d.connection.Preload("Tags").Find(&records)
	}
	return recordToAPISpec(records)
}
