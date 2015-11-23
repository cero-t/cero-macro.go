package vjoy

import (
	"log"
	"time"
)

func Acuire(vjoyId uint) {
	loadDll()

	status := getVJDStatus(vjoyId)
	if status != 1 {
		log.Println("Vjoy is not free")
		log.Fatal(status)
	}

	acquired := acquireVJD(vjoyId)
	if !acquired {
		log.Fatal("Vjoy cannot be acquired")
	}

	time.Sleep(50 * time.Millisecond)
	resetVJD(vjoyId)
}

func axis(vjoyId uint, x uint, y uint) bool {
	return setAxis(x, vjoyId, HID_USAGE_X) && setAxis(y, vjoyId, HID_USAGE_Y);
}

func push(vjoyId uint, button uint) bool {
	return setBtn(1, vjoyId, button);
}

func release(vjoyId uint, button uint) bool {
	return setBtn(0, vjoyId, button);
}

func Close(vjoyId uint) {
	resetVJD(vjoyId)
	relinquishVJD(vjoyId)
}
