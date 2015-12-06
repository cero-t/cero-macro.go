package vjoy

import (
	"errors"
	"time"
)

func Acuire(vjoyId uint) error {
	loadDll()

	status := getVJDStatus(vjoyId)
	if status != 1 {
		return errors.New("Vjoy [" + string(vjoyId) + "] is not free")
	}

	acquired := acquireVJD(vjoyId)
	if !acquired {
		return errors.New("Vjoy [" + string(vjoyId) + "] cannot be acquired")
	}

	time.Sleep(50 * time.Millisecond)
	resetVJD(vjoyId)

	return nil
}

func axis(vjoyId uint, x uint, y uint) bool {
	return setAxis(x, vjoyId, HID_USAGE_X) && setAxis(y, vjoyId, HID_USAGE_Y)
}

func push(vjoyId uint, button uint) bool {
	return setBtn(1, vjoyId, button)
}

func release(vjoyId uint, button uint) bool {
	return setBtn(0, vjoyId, button)
}

func Close(vjoyId uint) {
	resetVJD(vjoyId)
	relinquishVJD(vjoyId)
}
