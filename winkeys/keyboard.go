package winkeys

import (
	"golang.org/x/sys/windows"
	"syscall"
	"time"
	"unsafe"
)

const (
	keyDown    = 0
	KeyExtend  = 0x0001
	keyUp      = 0x0002
	keyUnicode = 0x0004
)

/*
参考：
https://docs.microsoft.com/ja-jp/windows/win32/inputdev/virtual-key-codes?redirectedfrom=MSDN
*/
const (
	VkReturn   = 13
	VkA        = 0x41
	VkX        = 0x58
	VkV        = 0x56
	VkLShift   = 0xA0
	VkRShift   = 0xA1
	VkLControl = 0xA2
	VkRControl = 0xA3
	VkLMenu    = 0xA4
	VkRMenu    = 0xA5
	VkF1       = 0x70
	VkF2       = 0x71
	VkF3       = 0x72
	VkF4       = 0x73
	VkF5       = 0x74
	VkF6       = 0x75
	VkF7       = 0x76
	VkF8       = 0x77
	VkF9       = 0x78
	VkF10      = 0x79
	VkF11      = 0x7A
	VkF12      = 0x7B
)

type KEYBD_INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
	Unused      [8]byte
}

var (
	libuser32 *windows.LazyDLL
	sendInput *windows.LazyProc
)

func init() {
	// Library
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	sendInput = libuser32.NewProc("SendInput")
}

func send(nInputs uint32, pInputs unsafe.Pointer, cbSize int32) uint32 {
	ret, _, _ := syscall.Syscall(sendInput.Addr(), 3,
		uintptr(nInputs),
		uintptr(pInputs),
		uintptr(cbSize))

	return uint32(ret)
}

func SendkeyVk(vk uint16) {
	var inputs []KEYBD_INPUT
	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     keyDown,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     keyUp,
			Time:        0,
			DwExtraInfo: 0,
		},
	})
	cbSize := int32(unsafe.Sizeof(KEYBD_INPUT{}))

	for _, inp := range inputs {
		send(1, unsafe.Pointer(&inp), cbSize)
	}
}

func SendShortCutkeyVk(vk1 uint16, vk2 uint16) {

	var inputs []KEYBD_INPUT
	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         vk1,
			WScan:       0,
			DwFlags:     keyDown,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         vk2,
			WScan:       0,
			DwFlags:     keyDown,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         vk2,
			WScan:       0,
			DwFlags:     keyUp,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         vk1,
			WScan:       0,
			DwFlags:     keyUp,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	cbSize := int32(unsafe.Sizeof(KEYBD_INPUT{}))

	for _, inp := range inputs {
		send(1, unsafe.Pointer(&inp), cbSize)
		time.Sleep(time.Millisecond * 100)
	}
}

func Sendkey(c uint16) {
	var inputs []KEYBD_INPUT
	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         0,
			WScan:       c,
			DwFlags:     keyDown | keyUnicode,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         0,
			WScan:       c,
			DwFlags:     keyUp | keyUnicode,
			Time:        0,
			DwExtraInfo: 0,
		},
	})
	cbSize := int32(unsafe.Sizeof(KEYBD_INPUT{}))

	for _, inp := range inputs {
		send(1, unsafe.Pointer(&inp), cbSize)
	}
}

func Sendkeys(str string) {
	for _, c := range str {
		Sendkey(uint16(c))
		time.Sleep(time.Millisecond * 300)
	}
}
