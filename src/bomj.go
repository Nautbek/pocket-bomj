package src

type Bomj struct {
	Id     int64
	Health uint8 `json:"health"` //percent 0-100
}

type BomjInterface interface {
	SetHealth(uint8)
	GetHealth() uint8
	SetId(int64)
}

func (br *Bomj) SetHealth(health uint8) {
	br.Health = health
}

func (br *Bomj) GetHealth() uint8 {
	return br.Health
}

func (br *Bomj) SetId(id int64) {
	br.Id = id
}
