package processor

import (
	"../vjoy"
	"time"
)

const frame int64 = 1000000000 / 60
var stopped bool

func ProcessPair(vjoyId1 uint, operations1 *[]Operation, vjoyId2 uint, operations2 *[]Operation) {
	stopped = false
	start := time.Now().UnixNano() + 1000000000
	go doProcess(vjoyId1, operations1, start)
	go doProcess(vjoyId2, operations2, start)
}

func ProcessSingle(vjoyId uint, operations *[]Operation) {
	stopped = false
	go doProcess(vjoyId, operations, 0)
}

func Stop()  {
	stopped = true
}


func doProcess(vjoyId uint, operations *[]Operation, start int64) {
	preSleep := start - time.Now().UnixNano()
	if preSleep > 0 {
		time.Sleep(time.Duration(preSleep))
	}

	var delta int64 = 0
	afterLastSleep := time.Now().UnixNano()

	for _, ope := range *operations {
		if stopped {
			break
		}

		if &ope == nil {
			continue
		}

		processOperation(vjoyId, ope);

		// adjust 60fps
		beforeSleep := time.Now().UnixNano()
		var sleepTime int64 = int64(ope.Frames) * frame - (beforeSleep - afterLastSleep) - delta;

		time.Sleep(time.Duration(sleepTime))
		afterLastSleep = time.Now().UnixNano()
		delta = afterLastSleep - beforeSleep - sleepTime;
	}

	vjoy.ReleaseLP(vjoyId)
	vjoy.ReleaseMP(vjoyId)
	vjoy.ReleaseHP(vjoyId)
	vjoy.ReleaseLK(vjoyId)
	vjoy.ReleaseMK(vjoyId)
	vjoy.ReleaseHK(vjoyId)
	vjoy.ReleasePause(vjoyId)
	vjoy.ReleaseSave(vjoyId)
	vjoy.ReleaseReload(vjoyId)
	vjoy.Direction5(vjoyId)
}

func processOperation(vjoyId uint, operation Operation) {
	if (operation.Direction != 0) {
		if operation.Direction == 1 {
			vjoy.Direction1(vjoyId)
		} else if operation.Direction == 2 {
			vjoy.Direction2(vjoyId)
		} else if operation.Direction == 3 {
			vjoy.Direction3(vjoyId)
		} else if operation.Direction == 4 {
			vjoy.Direction4(vjoyId)
		} else if operation.Direction == 5 {
			vjoy.Direction5(vjoyId)
		} else if operation.Direction == 6 {
			vjoy.Direction6(vjoyId)
		} else if operation.Direction == 7 {
			vjoy.Direction7(vjoyId)
		} else if operation.Direction == 8 {
			vjoy.Direction8(vjoyId)
		} else if operation.Direction == 9 {
			vjoy.Direction9(vjoyId)
		}
	}

	if (operation.Lp != 0) {
		if (operation.Lp == 1) {
			vjoy.PushLP(vjoyId)
		} else {
			vjoy.ReleaseLP(vjoyId)
		}
	}

	if (operation.Mp != 0) {
		if (operation.Mp == 1) {
			vjoy.PushMP(vjoyId)
		} else {
			vjoy.ReleaseMP(vjoyId)
		}
	}

	if (operation.Hp != 0) {
		if (operation.Hp == 1) {
			vjoy.PushHP(vjoyId)
		} else {
			vjoy.ReleaseHP(vjoyId)
		}
	}

	if (operation.Lk != 0) {
		if (operation.Lk == 1) {
			vjoy.PushLK(vjoyId)
		} else {
			vjoy.ReleaseLK(vjoyId)
		}
	}

	if (operation.Mk != 0) {
		if (operation.Mk == 1) {
			vjoy.PushMK(vjoyId)
		} else {
			vjoy.ReleaseMK(vjoyId)
		}
	}

	if (operation.Hk != 0) {
		if (operation.Hk == 1) {
			vjoy.PushHK(vjoyId)
		} else {
			vjoy.ReleaseHK(vjoyId)
		}
	}

	if (operation.Pause != 0) {
		if (operation.Pause == 1) {
			vjoy.PushPause(vjoyId)
		} else {
			vjoy.ReleasePause(vjoyId)
		}
	}

	if (operation.Save != 0) {
		if (operation.Save == 1) {
			vjoy.PushSave(vjoyId)
		} else {
			vjoy.ReleaseSave(vjoyId)
		}
	}

	if (operation.Reload != 0) {
		if (operation.Reload == 1) {
			vjoy.PushReload(vjoyId)
		} else {
			vjoy.ReleaseReload(vjoyId)
		}
	}
}
