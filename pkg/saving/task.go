package saving

import (
	"context"
	"time"
)

func StartTask(ss Service) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	doEvery(ctx, 20*time.Second, SaveData, ss)
}

func doEvery(ctx context.Context, d time.Duration, f func(time.Time, Service), ss Service) error {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case x := <-ticker.C:
			go f(x, ss)
		}
	}
}

func SaveData(t time.Time, ss Service) {
	ss.Save()
}
