package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/louistwiice/go/fripclose/ent"
)

type RouterBase struct {
	Database	*ent.Client
	Redis		*redis.Client
	OpenApp		*gin.RouterGroup
}

type Routers struct {
	RouterBase
	RestrictedApp 	*gin.RouterGroup
}