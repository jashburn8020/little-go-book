package c04codeorginterfaces

import (
	// Go does not allow cyclical imports
	"github.com/jashburn8020/little-go-book/c04codeorginterfaces/db"
	"testing"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}

	return item.Price, true
}

func TestPriceCheckFound(t *testing.T) {
	price, ok := PriceCheck(1)

	if !(ok && price == 9.001) {
		t.Error()
	}
}

func TestPriceCheckNotFound(t *testing.T) {
	price, ok := PriceCheck(2)

	if ok || price != 0 {
		t.Error()
	}
}
