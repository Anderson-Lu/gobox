package retry_helper

import "fmt"

func Retry(retryTimes int, errHandler func(error), job func() error) error {
	for i := 0; i < retryTimes; i++ {
		if e := job(); e != nil {
			errHandler(e)
			continue
		}
		return nil
	}
	return fmt.Errorf("retry job failed with %d times", retryTimes)
}
