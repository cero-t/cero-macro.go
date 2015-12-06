package vjoy

func Direction1(vjoyId uint) bool {
	return axis(vjoyId, 0, 0x7FFF)
}

func Direction2(vjoyId uint) bool {
	return axis(vjoyId, 0x4000, 0x7FFF)
}

func Direction3(vjoyId uint) bool {
	return axis(vjoyId, 0x7FFF, 0x7FFF)
}

func Direction4(vjoyId uint) bool {
	return axis(vjoyId, 0, 0x4000)
}

func Direction5(vjoyId uint) bool {
	return axis(vjoyId, 0x4000, 0x4000)
}

func Direction6(vjoyId uint) bool {
	return axis(vjoyId, 0x7FFF, 0x4000)
}

func Direction7(vjoyId uint) bool {
	return axis(vjoyId, 0, 0)
}

func Direction8(vjoyId uint) bool {
	return axis(vjoyId, 0x4000, 0)
}

func Direction9(vjoyId uint) bool {
	return axis(vjoyId, 0x7FFF, 0)
}

func PushLP(vjoyId uint) bool {
	return push(vjoyId, 4)
}

func ReleaseLP(vjoyId uint) bool {
	return release(vjoyId, 4)
}

func PushMP(vjoyId uint) bool {
	return push(vjoyId, 1)
}

func ReleaseMP(vjoyId uint) bool {
	return release(vjoyId, 1)
}

func PushHP(vjoyId uint) bool {
	return push(vjoyId, 6)
}

func ReleaseHP(vjoyId uint) bool {
	return release(vjoyId, 6)
}

func PushLK(vjoyId uint) bool {
	return push(vjoyId, 3)
}

func ReleaseLK(vjoyId uint) bool {
	return release(vjoyId, 3)
}

func PushMK(vjoyId uint) bool {
	return push(vjoyId, 2)
}

func ReleaseMK(vjoyId uint) bool {
	return release(vjoyId, 2)
}

func PushHK(vjoyId uint) bool {
	return push(vjoyId, 8)
}

func ReleaseHK(vjoyId uint) bool {
	return release(vjoyId, 8)
}

func PushPause(vjoyId uint) bool {
	return push(vjoyId, 10)
}

func ReleasePause(vjoyId uint) bool {
	return release(vjoyId, 10)
}

func PushSave(vjoyId uint) bool {
	return push(vjoyId, 12)
}

func ReleaseSave(vjoyId uint) bool {
	return release(vjoyId, 12)
}

func PushReload(vjoyId uint) bool {
	return push(vjoyId, 11)
}

func ReleaseReload(vjoyId uint) bool {
	return release(vjoyId, 11)
}
