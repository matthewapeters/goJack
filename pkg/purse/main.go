package purse

import (
	"fmt"
)

type purse struct {
	Total int
	Bets  map[interface{}]int
}

type Purse *purse

func NewPurse(buyin int) Purse {
	return &purse{
		Total: buyin,
	}
}

func (p *purse) Pay(amount int) (out int, err error) {
	if amount >= p.Total {
		err = fmt.Errorf("purse is empty")
		out = p.Total
		p.Total = 0
		return
	}
	p.Total -= amount
	out = amount
	return
}

func (p *purse) Receive(amount int) {
	p.Total += amount
}

func (p *purse) Bet(h interface{}, amount int) (err error) {
	if amount > p.Total {
		err = fmt.Errorf("insufficient funds to bet %d", amount)
		return
	}
	p.Bets[h], err = p.Pay(amount)
	if err != nil {
		return
	}
	return
}

func (p *purse) Win(h interface{}) {
	p.Receive(p.Bets[h])
	delete(p.Bets, h)
}

func (p *purse) Lose(h interface{}) {
	delete(p.Bets, h)
}
