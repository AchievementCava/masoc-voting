package guildScraper

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/CSSUoB/society-voting/internal/config"
	"github.com/carlmjohnson/requests"
)

var (
	activeSessionToken string
	tokenMutex         sync.RWMutex
	tokenInitOnce      sync.Once
)

func getActiveToken() string {
	tokenInitOnce.Do(func() {
		activeSessionToken = config.Get().Guild.SessionToken
	})

	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	return activeSessionToken
}

func updateActiveToken(newToken string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	activeSessionToken = newToken
}

func AutoRefreshAdminToken(ctx context.Context) {
	timer := time.NewTicker(time.Minute * 10)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := refreshAdminToken(); err != nil {
				slog.Error("refresh admin token", "error", err)
			}
		}
	}
}


func refreshCookie(ctx context.Context) error {
	conf := config.Get().Guild

	return requests.URL("https://www.guildofstudents.com/profile").
		Cookie(".AspNet.SharedCookie", conf.AdminToken).
		AddValidator(func(res *http.Response) error {
			for _, cookie := range res.Cookies() {
				if cookie.Name == ".AspNet.SharedCookie" && cookie.Value != conf.AdminToken {
					conf.AdminToken = cookie.Value
				}
			}
			return nil
		}).
		Fetch(ctx)
}

