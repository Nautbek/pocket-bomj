package bomj

import (
	"fmt"
	"math/rand"
	"time"
)

// BomjClass представляет класс бомжа
type BomjClass string

const (
	ClassTarelka  BomjClass = "tarelka"  // Тарелочник
	ClassVeteran  BomjClass = "veteran"  // Ветеран
	ClassGeek     BomjClass = "geek"     // Гик
	ClassDelivery BomjClass = "delivery" // Доставщик
)

// Bomj представляет бомжа
type Bomj struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Class       BomjClass `json:"class"`
	Name        string    `json:"name"`
	Level       int       `json:"level"`
	Experience  int       `json:"experience"`
	Health      int       `json:"health"`
	MaxHealth   int       `json:"max_health"`
	Hunger      int       `json:"hunger"`
	MaxHunger   int       `json:"max_hunger"`
	Money       int       `json:"money"`
	Reputation  int       `json:"reputation"`
	Inventory   []Item    `json:"inventory"`
	LastFed     time.Time `json:"last_fed"`
	LastWorked  time.Time `json:"last_worked"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
}

// Item представляет предмет в инвентаре
type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    int    `json:"value"`
	Quantity int    `json:"quantity"`
}

// NewBomj создает нового бомжа указанного класса
func NewBomj(userID int64, class BomjClass) *Bomj {
	now := time.Now()

	// Генерируем случайное имя
	names := []string{
		"Васян", "Петька", "Колян", "Серый", "Бородач", "Косой", "Хромой", "Кривой",
		"Длинный", "Короткий", "Толстый", "Худой", "Быстрый", "Медленный", "Умный", "Тупой",
	}

	bomj := &Bomj{
		UserID:      userID,
		Class:       class,
		Name:        names[rand.Intn(len(names))],
		Level:       1,
		Experience:  0,
		Health:      100,
		MaxHealth:   100,
		Hunger:      0,
		MaxHunger:   100,
		Money:       0,
		Reputation:  0,
		Inventory:   []Item{},
		LastFed:     now,
		LastWorked:  now,
		CreatedAt:   now,
		LastUpdated: now,
	}

	// Устанавливаем характеристики в зависимости от класса
	bomj.setClassStats()

	return bomj
}

// setClassStats устанавливает характеристики в зависимости от класса
func (b *Bomj) setClassStats() {
	switch b.Class {
	case ClassTarelka:
		b.MaxHealth = 120
		b.Health = 120
		b.MaxHunger = 80
		b.Reputation = 10
		b.Inventory = append(b.Inventory, Item{ID: 1, Name: "Тарелка", Type: "weapon", Value: 5, Quantity: 1})

	case ClassVeteran:
		b.MaxHealth = 150
		b.Health = 150
		b.MaxHunger = 120
		b.Reputation = 20
		b.Inventory = append(b.Inventory, Item{ID: 2, Name: "Орден", Type: "accessory", Value: 10, Quantity: 1})

	case ClassGeek:
		b.MaxHealth = 80
		b.Health = 80
		b.MaxHunger = 60
		b.Reputation = 5
		b.Inventory = append(b.Inventory, Item{ID: 3, Name: "Смартфон", Type: "device", Value: 15, Quantity: 1})

	case ClassDelivery:
		b.MaxHealth = 100
		b.Health = 100
		b.MaxHunger = 150
		b.Reputation = 0
		b.Inventory = append(b.Inventory, Item{ID: 4, Name: "Велосипед", Type: "vehicle", Value: 20, Quantity: 1})
	}
}

// GetClassDescription возвращает описание класса
func (b *Bomj) GetClassDescription() string {
	switch b.Class {
	case ClassTarelka:
		return "🍽️ Бомж-тарелочник - мастер по сбору еды и тарелок. Может найти пропитание в самых неожиданных местах!"
	case ClassVeteran:
		return "🎖️ Бомж-ветеран - уважаемый в бомжовом сообществе. Имеет связи и авторитет среди коллег."
	case ClassGeek:
		return "💻 Бомж-гик - знает все о технологиях, но не может их себе позволить. Мастер по поиску зарядок."
	case ClassDelivery:
		return "🚲 Бомж-доставщик - быстрый, но вечно голодный. Может доставить что угодно куда угодно!"
	default:
		return "🤔 Неизвестный класс бомжа"
	}
}

// GetStatus возвращает текущий статус бомжа
func (b *Bomj) GetStatus() string {
	healthStatus := "❤️"
	if b.Health < 30 {
		healthStatus = "💔"
	} else if b.Health < 70 {
		healthStatus = "🖤"
	}

	hungerStatus := "😋"
	if b.Hunger > 80 {
		hungerStatus = "😵"
	} else if b.Hunger > 50 {
		hungerStatus = "😫"
	}

	return healthStatus + " " + hungerStatus
}

// Feed кормит бомжа
func (b *Bomj) Feed(food Item) bool {
	if b.Hunger >= b.MaxHunger {
		return false
	}

	hungerReduction := food.Value
	if food.Type == "food" {
		hungerReduction *= 2
	}

	b.Hunger = max(0, b.Hunger-hungerReduction)
	b.Health = min(b.MaxHealth, b.Health+food.Value/2)
	b.LastFed = time.Now()
	b.LastUpdated = time.Now()

	return true
}

// Work заставляет бомжа работать
func (b *Bomj) Work() (int, string) {
	if time.Since(b.LastWorked) < time.Hour {
		return 0, "Бомж еще не отдохнул от предыдущей работы! 😴"
	}

	if b.Hunger > 80 {
		return 0, "Бомж слишком голоден для работы! 😫"
	}

	if b.Health < 30 {
		return 0, "Бомж слишком болен для работы! 🤒"
	}

	// Разные виды работы для разных классов
	var earnings int
	var workType string

	switch b.Class {
	case ClassTarelka:
		workType = "собирал тарелки и еду"
		earnings = rand.Intn(15) + 5
		b.Hunger = min(b.MaxHunger, b.Hunger+10)

	case ClassVeteran:
		workType = "давал советы молодым бомжам"
		earnings = rand.Intn(20) + 10
		b.Reputation += 2

	case ClassGeek:
		workType = "помогал с техникой"
		earnings = rand.Intn(25) + 15
		b.Experience += 5

	case ClassDelivery:
		workType = "доставлял посылки"
		earnings = rand.Intn(30) + 20
		b.Hunger = min(b.MaxHunger, b.Hunger+20)
	}

	b.Money += earnings
	b.Experience += earnings / 5
	b.Health = max(0, b.Health-5)
	b.LastWorked = time.Now()
	b.LastUpdated = time.Now()

	// Проверяем повышение уровня
	b.checkLevelUp()

	return earnings, workType
}

// checkLevelUp проверяет повышение уровня
func (b *Bomj) checkLevelUp() {
	requiredExp := b.Level * 100

	if b.Experience >= requiredExp {
		b.Level++
		b.Experience -= requiredExp
		b.MaxHealth += 10
		b.Health = b.MaxHealth
		b.MaxHunger += 5
		b.Reputation += 5
	}
}

// GetRandomEvent возвращает случайное событие для бомжа
func (b *Bomj) GetRandomEvent() (string, int) {
	events := []struct {
		description  string
		moneyChange  int
		healthChange int
		hungerChange int
	}{
		{"Нашел брошенную еду! 🍕", 0, 5, -20},
		{"Помог старушке донести сумки! 👵", 5, 0, 10},
		{"Попал под дождь! ☔", 0, -10, 0},
		{"Нашел монетки на улице! 🪙", 3, 0, 0},
		{"Поделился едой с другом! 🤝", -2, 5, 15},
		{"Попал в драку! 👊", 0, -15, 0},
		{"Выиграл в карты! 🃏", 8, 0, 0},
		{"Проиграл в карты! 😭", -5, 0, 0},
		{"Нашел зарядку для телефона! 🔌", 0, 0, 0},
		{"Помылся в фонтане! 🚿", 0, 10, 0},
	}

	event := events[rand.Intn(len(events))]

	b.Money += event.moneyChange
	b.Health = max(0, min(b.MaxHealth, b.Health+event.healthChange))
	b.Hunger = max(0, min(b.MaxHunger, b.Hunger+event.hungerChange))

	return event.description, event.moneyChange
}

// GetStats возвращает статистику бомжа
func (b *Bomj) GetStats() string {
	return fmt.Sprintf(`
🏆 %s (Уровень %d)
💰 Деньги: %d₽
❤️ Здоровье: %d/%d
🍽️ Голод: %d/%d
⭐ Репутация: %d
📈 Опыт: %d/%d
%s
`, b.Name, b.Level, b.Money, b.Health, b.MaxHealth, b.Hunger, b.MaxHunger, b.Reputation, b.Experience, b.Level*100, b.GetStatus())
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
