package databases

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/lzyyauto/auto4play/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// InitDB create database instance and connect to database
func init() {
	var err error
	cfg := config.MysqlCfg
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err = sqlx.Connect("mysql", uri)
	if err != nil {
		log.Fatalf("Connect database failed: %s\n", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot connect mysql, error: %s", err)
	}
	log.Printf("[INFO] Connected to database.")
}

func preArgs(fields interface{}) string {
	fieldsVal := reflect.ValueOf(fields)

	var qmSize int

	switch fieldsVal.Kind() {
	case reflect.String:
		s := fields.(string)
		qmSize = len(strings.Split(s, ","))
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		qmSize = fieldsVal.Len()
	default:
		panic(fmt.Sprintf("Unsupport data type: %T", fields))
	}

	p := []string{}
	for i := 0; i < qmSize; i++ {
		p = append(p, "?")
	}
	return strings.Join(p, ",")
}

func preData(fields string, data interface{}) []interface{} {
	fieldSlice := strings.Split(fields, ",")
	dataSlice := make([]interface{}, 0)

	types := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	if types.Kind() == reflect.Ptr {
		types = types.Elem()
		values = values.Elem()
	}

	if values.Kind() != reflect.Struct {
		return dataSlice
	}

	for i := 0; i < len(fieldSlice); i++ {
		expectField := strings.TrimSpace(fieldSlice[i])
		for j := 0; j < types.NumField(); j++ {
			field := types.Field(j)
			dbtags := strings.Split(field.Tag.Get("gpo"), ",")
			fieldTag := dbtags[0]
			if fieldTag == expectField {
				fieldData := values.Field(j).Interface()
				var isNilStruct bool
				if values.Field(j).Kind() == reflect.Ptr {
					isNilStruct = values.Field(j).IsNil()
				}

				if len(dbtags) > 1 && dbtags[1] == "json" && !isNilStruct {
					fieldData, _ = json.Marshal(fieldData)
				}
				dataSlice = append(dataSlice, fieldData)
				break
			}
		}
	}

	return dataSlice
}
