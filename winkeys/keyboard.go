package winkeys

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

const (
	KeyDown = 0
	KeyUp   = 0x0002
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
			DwFlags:     KeyDown,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, KEYBD_INPUT{
		Type: 1,
		Ki: KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     KeyUp,
			Time:        0,
			DwExtraInfo: 0,
		},
	})
	cbSize := int32(unsafe.Sizeof(KEYBD_INPUT{}))

	for _, inp := range inputs {
		send(1, unsafe.Pointer(&inp), cbSize)
	}
}
