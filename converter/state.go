package converter

type State struct {
	direction uint8
	lp        bool
	mp        bool
	hp        bool
	lk        bool
	mk        bool
	hk        bool
	pause     bool
	save      bool
	reload    bool
	frames    uint16
}
