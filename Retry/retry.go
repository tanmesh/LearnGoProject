package Retry

import (
	"fmt"
	"gopkg.in/retry.v1"
	"time"
)

func Retry() {
	strategy := retry.Exponential{
		Initial:  100 * time.Millisecond,
		Factor:   2,
		MaxDelay: 1 * time.Second,
		Jitter:   true,
	}

	for i, attempt := 0, retry.Start(strategy, nil); attempt.Next(); i++ {
		if i == 5 {
			break
		}
		fmt.Println(attempt.Count())
	}
}
