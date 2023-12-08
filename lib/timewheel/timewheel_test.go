package timewheel

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeWheel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	tw := New(time.Second, 60)

	for i := 0; i < 10; i++ {
		n := i + 1
		now := time.Now()

		tw.AddTask(context.Background(), "task#"+strconv.Itoa(n), func(ctx context.Context, taskID string) error {
			ch <- fmt.Sprintf("%s - %ds", taskID, int64(time.Since(now).Seconds()))

			return nil
		}, WithDefer(func(attempts uint16) time.Duration {
			return time.Second * time.Duration(n+i)
		}))
	}

	tw.Run()

	ret := make([]string, 0, 10)

	for v := range ch {
		ret = append(ret, v)
		if len(ret) == 10 {
			break
		}
	}

	assert.Equal(t, []string{
		"task#1 - 1s",
		"task#2 - 3s",
		"task#3 - 5s",
		"task#4 - 7s",
		"task#5 - 9s",
		"task#6 - 11s",
		"task#7 - 13s",
		"task#8 - 15s",
		"task#9 - 17s",
		"task#10 - 19s",
	}, ret)
}
