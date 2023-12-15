package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	cachelib "github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	redisLib "github.com/eko/gocache/store/redis/v4"
	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

const COUNTER_CACHE_KEY = "counter"

func main() {
	redisStore := redisLib.NewRedis(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}))
	cacheManager := cachelib.New[[]byte](redisStore)

	e := echo.New()
	counter := 1
	e.GET("/counts", func(c echo.Context) error {
		ctx := c.Request().Context()
		// try to get counter first
		cached, err := cacheManager.Get(ctx, COUNTER_CACHE_KEY)
		if err != nil && !errors.Is(err, store.NotFound{}) {
			return echo.ErrInternalServerError
		}
		if cached != nil {
			resp := fmt.Sprintf("counter: %s", cached)
			return c.String(http.StatusOK, resp)
		}
		cacheManager.Set(ctx, COUNTER_CACHE_KEY, []byte(strconv.Itoa(counter)))
		resp := fmt.Sprintf("counter: %d", counter)
		return c.String(http.StatusOK, resp)
	})

	e.POST("/counts", func(c echo.Context) error {
		ctx := c.Request().Context()
		cacheManager.Delete(ctx, COUNTER_CACHE_KEY)
		counter = counter + 1
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
