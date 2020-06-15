package winkeys

import (
	"testing"
	"time"
)

func TestSendkeyVk(t *testing.T) {
	time.Sleep(time.Second * 5)

	SendkeyVk(13)
	SendkeyVk(13)
}
