package database

import (
	"bufio"
	"bytes"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

var tables = []string{
	"options",
	"polls",
	"user_polls",
	"users",
}

type Database struct {
	DB *gorm.DB
}

func InitDatabase() *Database {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_TEST_HOST"),
		os.Getenv("DB_TEST_USER"),
		os.Getenv("DB_TEST_PASS"),
		os.Getenv("DB_TEST_DBNAME"),
		os.Getenv("DB_TEST_PORT"),
	)
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: connectionString}), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Database{
		DB: db,
	}
}

func (d *Database) TruncateTables() {
	d.DB.Exec("SET FOREIGN_KEY_CHECKS=0")
	defer d.DB.Exec("SET FOREIGN_KEY_CHECKS=1")

	for i := range tables {
		err := d.DB.Table(tables[i]).Exec(fmt.Sprintf("TRUNCATE TABLE %s", tables[i])).Error
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (d *Database) ExecFixture(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	for scanner.Scan() {
		query := scanner.Text()
		if err = d.DB.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}
