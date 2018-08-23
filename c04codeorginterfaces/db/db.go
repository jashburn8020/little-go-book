package db

type Item struct {
	Price float64
}

func LoadItem(id int) *Item {
	if id == 1 {
		return &Item{Price: 9.001}
	}

	return nil
}
