package vjoy
import (
	"log"
	"strconv"
	"time"
)

const (
	HID_USAGE_X uint = 0x30
	HID_USAGE_Y uint = 0x31
)

var (
	timer = time.Now().UnixNano()
)

func loadDll() {
	log.Println("loadDll")
}

func getVJDStatus(vjoyId uint) uintptr {
	log.Println("getVJDStatus : " + strconv.Itoa(int(vjoyId)))
	return 1;
}

func acquireVJD(vjoyId uint) bool {
	log.Println("acquireVJD : " + strconv.Itoa(int(vjoyId)))
	return true
}

func setAxis(value uint, vjoyId uint, usage uint) bool {
	log.Println("setAxis : " + strconv.Itoa(int(value)) + "," + strconv.Itoa(int(vjoyId)) + "," + strconv.Itoa(int(usage)))
	return true
}

func setBtn(value uint, vjoyId uint, button uint) bool {
	log.Println("setBtn : " + strconv.Itoa(int(value)) + "," + strconv.Itoa(int(vjoyId)) + "," + strconv.Itoa(int(button)))
	return true
}

func resetVJD(vjoyId uint) bool {
	log.Println("resetVJD : " + strconv.Itoa(int(vjoyId)))
	return true
}

func resetAll() bool {
	log.Println("resetAll")
	return true
}

func relinquishVJD(vjoyId uint) {
	log.Println("relinquishVJD : " + strconv.Itoa(int(vjoyId)))
}
