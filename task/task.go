package task

import (
	"context"
	"fmt"
	"time"
)

func StartTask() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	doEvery(ctx, 2*time.Second, SaveData)
}

func doEvery(ctx context.Context, d time.Duration, f func(time.Time)) error {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case x := <-ticker.C:
			go f(x)
		}
	}
}

func SaveData(t time.Time) {
	fmt.Println("tick")
}
