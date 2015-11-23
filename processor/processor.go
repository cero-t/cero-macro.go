package processor

import (
	"../vjoy"
	"time"
	"log"
)

const frame int64 = 1000000000 / 60
var vjoyId uint = 1

func Process(operations *[]Operation) {
	log.Println("START")
	var delta int64 = 0
	var afterLastSleep int64 = time.Now().UnixNano()

	for _, ope := range *operations {
		if &ope == nil {
			continue
		}

		processOperation(ope);

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
	vjoy.Direction5(vjoyId)
	log.Println("END")
}

func processOperation(operation Operation) {
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
}
