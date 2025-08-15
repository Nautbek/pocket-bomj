package src

type Bomj struct {
	id     int64
	health uint8   //percent 0-100
	money  float32 //Rubles
}

type BomjInterface interface {
	GetId() int64
	SetHealth(uint8)
	GetHealth() uint8
	SetId(int64)
	PlusMoney(float32)
	MinusMoney(float32)
	GetMoney() float32
	SetMoney(float32)
}

func (b *Bomj) GetId() int64 {
	return b.id
}

func (b *Bomj) SetHealth(health uint8) {
	b.health = health
}

func (b *Bomj) GetHealth() uint8 {
	return b.health
}

func (b *Bomj) SetId(id int64) {
	b.id = id
}

func (b *Bomj) PlusMoney(money float32) {
	b.money += money
}

func (b *Bomj) MinusMoney(money float32) {
	b.money -= money
}

func (b *Bomj) GetMoney() float32 {
	return b.money
}

func (b *Bomj) SetMoney(money float32) {
	b.money = money
}
