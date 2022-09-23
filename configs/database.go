package configs

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	logger "github.com/rs/zerolog/log"

	"github.com/louistwiice/go/fripclose/ent"
)

func NewDBConnection() *ent.Client {
	conf := LoadConfigEnv() //Load .env settings

	// Start by connecting to database
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbName)
	db, err := open(dbSource)
	if err!= nil {
		log.Panic("Database error: ... ", err.Error())
	}
	logger.Info().Msg("Database creation ...")
	return db
}

func NewRedisConnection() *redis.Client {
	conf := LoadConfigEnv() //Load .env settings

	opt, err := redis.ParseURL(conf.RedisBaseUrl)
	if err!= nil {
		log.Panic("Redis error: ... ", err.Error())
	}
	rdb := redis.NewClient(opt)
	logger.Info().Msg("Redis connection successful ...")
	
	return rdb
}

// To create new database connection
func open(source string) (*ent.Client, error) {
    db, err := sql.Open("mysql", source)
    if err != nil {
        return nil, err
    }
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(time.Hour)
    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB("mysql", db)
    return ent.NewClient(ent.Driver(drv)), nil
}