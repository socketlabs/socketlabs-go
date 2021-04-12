package injectionapi

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const nano = 1000000000

const defaultNumberOfRetries = 0
const maximumAllowedNumberOfRetries = 5
const minimumRetryTime = time.Duration(1 * nano)
const maximumRetryTime = time.Duration(10 * nano)

type RetrySettings interface {
	GetMaximumNumberOfRetries() (int)
	GetNextWaitInterval(numberOfAttempts int) (time.Duration)
	GetRetryDelta(numberOfAttempts int) (int)
}

type retrySettings struct {
	MaximumNumberOfRetries int
}

func CreateRetrySettings(maximumNumberOfRetries int) RetrySettings{
	if maximumNumberOfRetries < 0 {
		panic("maximumNumberOfRetries must be greater than 0")
	}

	if maximumNumberOfRetries > maximumAllowedNumberOfRetries {
		panic(fmt.Sprint("The maximum number of allowed retries is ", maximumAllowedNumberOfRetries))
	}

	return &retrySettings{
		MaximumNumberOfRetries: maximumNumberOfRetries,
	}
}

func (retrySettings *retrySettings) GetMaximumNumberOfRetries() int {
	return retrySettings.MaximumNumberOfRetries
}

func (retrySettings *retrySettings) GetNextWaitInterval(numberOfAttempts int) time.Duration {
	var interval = int(math.Min(
								float64(minimumRetryTime.Milliseconds() + int64(retrySettings.GetRetryDelta(numberOfAttempts))),
								float64(maximumRetryTime.Milliseconds())))
	return time.Duration(interval)
}


func (retrySettings *retrySettings) GetRetryDelta(numberOfAttempts int) int {

	min := int(float64(time.Duration(1 * nano).Milliseconds()) * 0.8)
	max := int(float64(time.Duration(1 * nano).Milliseconds()) * 1.2)

	return int((math.Pow(2.0, float64(numberOfAttempts)) - 1.0) * float64(rand.Intn(max - min) + min))
}