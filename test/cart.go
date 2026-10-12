package cart

type Cart struct {
	item     map[string]*LineItem
	promo    map[string]*CashPromotion
	discount float64
	price    float64
}

type LineItem struct {
	Name      string
	UnitPrice float64
	Quantity  float64
}

type CashPromotion struct {
	Name  string
	Price float64
}

type SpecialPromotion struct {
	Discount float64
}

func (c *Cart) TotalPrice() float64 {
	totalprice := 0.0
	for _, v := range c.item {
		discount := 0.0
		quantity := 1.0
		if v.Quantity > 1 {
			quantity = v.Quantity
		}
		dc, ok := c.promo[v.Name]
		if ok {
			discount = dc.Price
		}
		totalprice += quantity * (v.UnitPrice - discount)
	}
	totalprice -= c.discount
	return totalprice
}

func (c *Cart) AddItems(item LineItem) {
	if c.item == nil {
		c.item = make(map[string]*LineItem)
	}
	_, ok := c.item[item.Name]
	if !ok {
		c.item[item.Name] = &item
	}
}

func (c *Cart) AddPromotion(v interface{}) {
	switch val := v.(type) {
	case CashPromotion:
		if c.promo == nil {
			c.promo = make(map[string]*CashPromotion)
		}
		_, ok := c.promo[val.Name]
		if !ok {
			c.promo[val.Name] = &val
		}
	case SpecialPromotion:
		c.discount = val.Discount
	default:

	}

}
