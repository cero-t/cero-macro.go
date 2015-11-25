package processor

type Operation struct {
	Direction uint8
	Lp        int8
	Mp        int8
	Hp        int8
	Lk        int8
	Mk        int8
	Hk        int8
	Pause     int8
	Save      int8
	Reload    int8
	Frames    uint16
}
