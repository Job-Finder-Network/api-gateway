package entity

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var DB *sql.DB

type DBconf struct {
	Name     string `yaml:"name"`
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func InitDb() {
	conf := ParseDbYaml()
	DB, _ = sql.Open("mysql", conf.Username+":"+conf.Password+"@tcp("+conf.Ip+":"+conf.Port+")/"+conf.Name)
}
func ParseDbYaml() DBconf {
	var c DBconf
	yamlFile, err := ioutil.ReadFile("db.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
