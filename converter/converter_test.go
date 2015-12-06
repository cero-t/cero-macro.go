package converter
import (
	"testing"
	"../converter"
)

func TestSimple(t *testing.T) {
	list := []string{"2 2", "3 2", "6,lp 2", "5 2"}
	operaions := converter.LinesToOperations(&list)
	size := len(*operaions)

	if size != 4 {
		t.Error(size, 4)
	}
}

func TestLoop(t *testing.T) {
	loop := []string{"<loop 2>", "lp 2", "mp 2", "</loop>"}
	direct := []string{"lp 2", "mp 2", "lp 2", "mp 2"}

	actual := *converter.LinesToOperations(&loop)
	expect := *converter.LinesToOperations(&direct)

	// assert
	if len(actual) != len(expect) {
		t.Error(len(actual), len(expect))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expect[i] {
			t.Error(i, actual[i], expect[i])
		}
	}
}

func TestLoop2(t *testing.T) {
	loop := []string{"2,lk 2", "<loop 2>", "4,lp 2", "mp 2", "6,hp 2", "</loop>", "2,lp 3", "<loop 2>", "2,mp 2", "8,mk 2", "</loop>"}
	direct := []string{"2,lk 2", "4,lp 2", "mp 2", "6,hp 2", "4,lp 2", "mp 2", "6,hp 2", "2,lp 3", "2,mp 2", "8,mk 2", "2,mp 2", "8,mk 2"}

	actual := *converter.LinesToOperations(&loop)
	expect := *converter.LinesToOperations(&direct)

	// assert
	if len(actual) != len(expect) {
		t.Error(len(actual), len(expect))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expect[i] {
			t.Error(i, actual[i], expect[i])
		}
	}
}
