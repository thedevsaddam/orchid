package config

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"os"
)

func InitializeSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionDriver := os.Getenv("SESSION_STORE_DRIVER")
		sessionName := os.Getenv("SESSION_STORE_NAME")
		appKey := os.Getenv("APP_KEY")
		switch sessionDriver {
		case "cookie":
			store := sessions.NewCookieStore([]byte(appKey))
			sessions.Sessions(sessionName, store)
		case "redis":
			store, _ := sessions.NewRedisStore(10, "tcp", os.Getenv("SESSION_HOST") + ":" + os.Getenv("SESSION_PORT"), "", []byte(appKey))
			sessions.Sessions(sessionName, store)
		case "memcache":
			store := sessions.NewMemcacheStore(memcache.New(os.Getenv("SESSION_HOST") + ":" + os.Getenv("SESSION_PORT")), "", []byte(appKey))
			sessions.Sessions(sessionName, store)
		default:
			store := sessions.NewCookieStore([]byte(appKey))
			sessions.Sessions(sessionName, store)
		}
		c.Next()
	}
}
