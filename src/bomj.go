package src

type Bomj struct {
	Id     int64
	Health uint8   //percent 0-100
	Money  float32 //Rubles
}

type BomjInterface interface {
	GetId() int64
	SetHealth(uint8)
	GetHealth() uint8
	SetId(int64)
	PlusMoney(float32)
	MinusMoney(float32)
	GetMoney() float32
}

func (b *Bomj) GetId() int64 {
	return b.Id
}

func (b *Bomj) SetHealth(health uint8) {
	b.Health = health
}

func (b *Bomj) GetHealth() uint8 {
	return b.Health
}

func (b *Bomj) SetId(id int64) {
	b.Id = id
}

func (b *Bomj) PlusMoney(money float32) {
	b.Money += money
}

func (b *Bomj) MinusMoney(money float32) {
	b.Money -= money
}

func (b *Bomj) GetMoney() float32 {
	return b.Money
}
