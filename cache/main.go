package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	cachelib "github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/eko/gocache/store/go_cache/v4"
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
)

const COUNTER_CACHE_KEY = "counter"

func main() {
	// expire in 1 minute, purge interval 2 minutes
	client := cache.New(1*time.Minute, 2*time.Minute)
	cacheStore := go_cache.NewGoCache(client)
	cacheManager := cachelib.New[[]byte](cacheStore)

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
