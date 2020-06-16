package winkeys

import (
	"testing"
	"time"
)

func TestSendkeyVk(t *testing.T) {
	time.Sleep(time.Second * 2)

	SendkeyVk(VkReturn)
	SendkeyVk(VkReturn)
}

func TestSendkeys(t *testing.T) {
	time.Sleep(time.Second * 2)

	Sendkeys("abcdefg")
	SendShortCutkeyVk(VkLControl, VkA) // Select All
	SendShortCutkeyVk(VkLControl, VkX) // Cut
	SendShortCutkeyVk(VkLControl, VkV) // Paste
	SendShortCutkeyVk(VkLControl, VkV) // Paste
}
