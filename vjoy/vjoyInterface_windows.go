package vjoy

import (
	"syscall"
	"log"
)

const (
	HID_USAGE_X uint = 0x30
	HID_USAGE_Y uint = 0x31
)

var (
	dll *syscall.DLL
	procGetVJDStatus *syscall.Proc
	procAcquireVJD *syscall.Proc
	procSetAxis *syscall.Proc
	procSetBtn *syscall.Proc
	procResetVJD *syscall.Proc
	procResetAll *syscall.Proc
	procRelinquishVJD *syscall.Proc
)

func loadDll() {
	if dll != nil {
		return
	}

	dll, _err := syscall.LoadDLL("vJoyInterface.dll");
	if _err != nil {
		log.Fatal("load dll vJoyInterface.dll", _err)
	}

	vJoyEnabled := initProc(dll, "vJoyEnabled")
	result, _, _err := vJoyEnabled.Call()
	if result == 0 {
		log.Fatal("vJoy is not enabed")
	}

	procGetVJDStatus = initProc(dll, "GetVJDStatus")
	procAcquireVJD = initProc(dll, "AcquireVJD")
	procSetAxis = initProc(dll, "SetAxis")
	procSetBtn = initProc(dll, "SetBtn")
	procResetVJD = initProc(dll, "ResetVJD")
	procResetAll = initProc(dll, "ResetAll")
	procRelinquishVJD = initProc(dll, "RelinquishVJD")
}

func initProc(dll *syscall.DLL, name string) *syscall.Proc {
	proc, _err := dll.FindProc(name)
	if _err != nil {
		log.Fatal("get proc vJoyEnabled", _err)
	}
	return proc
}

func getVJDStatus(vjoyId uint) uintptr {
	result, _code, _err := procGetVJDStatus.Call(uintptr(vjoyId))
	if _code != 0 {
		log.Fatal("error call proc GetVJDStatus", _err)
	}

	return result;
}

func acquireVJD(vjoyId uint) bool {
	result, _code, _err := procAcquireVJD.Call(uintptr(vjoyId))
	if _code != 0 {
		log.Fatal("error call proc AcquireVJD", _err)
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func setAxis(value uint, vjoyId uint, usage uint) bool {
	result, _code, _err := procSetAxis.Call(uintptr(value), uintptr(vjoyId), uintptr(usage))
	if _code != 0 {
		log.Fatal("error call proc SetAxis", _err)
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func setBtn(value uint, vjoyId uint, button uint) bool {
	result, _code, _err := procSetBtn.Call(uintptr(value), uintptr(vjoyId), uintptr(button))
	if _code != 0 {
		log.Fatal("error call proc SetBtn", _err)
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func resetVJD(vjoyId uint) bool {
	result, _code, _err := procResetVJD.Call(uintptr(vjoyId))
	if _code != 0 {
		log.Fatal("error call proc ResetVJD", _err)
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func resetAll() bool {
	result, _code, _err := procResetAll.Call()
	if _code != 0 {
		log.Fatal("error call proc ResetAll", _err)
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func relinquishVJD(vjoyId uint) {
	_, _code, _err := procRelinquishVJD.Call(uintptr(vjoyId))
	if _code != 0 {
		log.Fatal("error call proc RelinquishVJD", _err)
	}
}
