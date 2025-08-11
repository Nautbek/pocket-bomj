package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"pocket-bomj/src/bomj"
)

var (
	bot     *tgbotapi.BotAPI
	storage *bomj.Storage
	logger  *logrus.Logger
)

func main() {
	// Инициализируем логгер
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		logger.Warn("Файл .env не найден, используем системные переменные")
	}

	// Получаем токен бота
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		logger.Fatal("TELEGRAM_BOT_TOKEN не установлен")
	}

	// Инициализируем бота
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		logger.Fatal(err)
	}

	bot.Debug = false
	logger.Infof("Бот %s запущен", bot.Self.UserName)

	// Инициализируем хранилище
	storage = bomj.NewStorage()

	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Настраиваем обновления
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	// Обрабатываем сообщения
	for update := range updates {
		if update.Message == nil {
			continue
		}

		go handleMessage(update.Message)
	}
}

func handleMessage(message *tgbotapi.Message) {
	userID := message.From.ID
	text := message.Text

	logger.Infof("Сообщение от %d: %s", userID, text)

	// Проверяем, есть ли у пользователя бомж
	if !storage.HasBomj(userID) {
		// Первое обращение - предлагаем выбрать класс
		if text == "/start" || text == "start" {
			sendClassSelection(userID)
		} else if isClassSelection(text) {
			handleClassSelection(userID, text)
		} else {
			sendWelcomeMessage(userID)
		}
		return
	}

	// Обрабатываем команды для существующего бомжа
	switch strings.ToLower(text) {
	case "/start", "start":
		sendMainMenu(userID)
	case "/stats", "статистика", "стат":
		sendStats(userID)
	case "/work", "работа", "работать":
		handleWork(userID)
	case "/feed", "кормить", "еда":
		sendFoodMenu(userID)
	case "/event", "событие", "приключение":
		handleRandomEvent(userID)
	case "/help", "помощь", "help":
		sendHelp(userID)
	default:
		// Проверяем, не является ли это выбором еды
		if strings.HasPrefix(text, "еда_") {
			handleFoodSelection(userID, text)
		} else {
			sendMainMenu(userID)
		}
	}
}

func isClassSelection(text string) bool {
	text = strings.ToLower(text)
	return text == "тарелочник" || text == "тарелка" || text == "tarelka" ||
		text == "ветеран" || text == "veteran" ||
		text == "гик" || text == "geek" ||
		text == "доставщик" || text == "delivery"
}

func sendWelcomeMessage(userID int64) {
	text := `🎭 Добро пожаловать в игру "Карманный Бомж"!

Это социально острый проект с элементами черного юмора. 
В игре ты будешь заботиться о своём бомже, помогать ему в этом суровом мире и не давать ему умереть.

Напиши /start чтобы начать игру и выбрать класс своего бомжа! 🚀`

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func sendClassSelection(userID int64) {
	text := `🎯 Выбери класс своего бомжа:

🍽️ Тарелочник - мастер по сбору еды и тарелок
🎖️ Ветеран - уважаемый в бомжовом сообществе  
💻 Гик - знает все о технологиях
🚲 Доставщик - быстрый, но вечно голодный

Отправь название класса (например: "тарелочник")`

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func handleClassSelection(userID int64, className string) {
	var class bomj.BomjClass

	switch strings.ToLower(className) {
	case "тарелочник", "тарелка", "tarelka":
		class = bomj.ClassTarelka
	case "ветеран", "veteran":
		class = bomj.ClassVeteran
	case "гик", "geek":
		class = bomj.ClassGeek
	case "доставщик", "delivery":
		class = bomj.ClassDelivery
	default:
		sendClassSelection(userID)
		return
	}

	// Создаем нового бомжа
	newBomj := bomj.NewBomj(userID, class)
	storage.SaveBomj(newBomj)

	// Отправляем приветствие
	text := fmt.Sprintf(`🎉 Поздравляю! Ты создал бомжа класса "%s"!

%s

Теперь у тебя есть собственный бомж! Используй команды:
📊 /stats - посмотреть статистику
💼 /work - отправить на работу
🍽️ /feed - покормить
🎲 /event - случайное событие
❓ /help - помощь

Удачи в этом суровом мире! 🚀`,
		className, newBomj.GetClassDescription())

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func sendMainMenu(userID int64) {
	bomj := storage.GetBomj(userID)
	if bomj == nil {
		sendWelcomeMessage(userID)
		return
	}

	text := fmt.Sprintf(`🏠 Главное меню

%s

Что хочешь сделать?
💼 /work - отправить на работу
🍽️ /feed - покормить
🎲 /event - случайное событие
📊 /stats - посмотреть статистику
❓ /help - помощь`, bomj.GetStats())

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func sendStats(userID int64) {
	bomj := storage.GetBomj(userID)
	if bomj == nil {
		sendWelcomeMessage(userID)
		return
	}

	text := fmt.Sprintf(`📊 Статистика бомжа:

%s

🎒 Инвентарь:
`, bomj.GetStats())

	for _, item := range bomj.Inventory {
		text += fmt.Sprintf("• %s (x%d) - %d₽\n", item.Name, item.Quantity, item.Value)
	}

	if len(bomj.Inventory) == 0 {
		text += "Пусто 😢\n"
	}

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func handleWork(userID int64) {
	bomj := storage.GetBomj(userID)
	if bomj == nil {
		sendWelcomeMessage(userID)
		return
	}

	earnings, workType := bomj.Work()
	storage.SaveBomj(bomj)

	if earnings == 0 {
		text := fmt.Sprintf("❌ Работа не удалась: %s", workType)
		msg := tgbotapi.NewMessage(userID, text)
		bot.Send(msg)
		return
	}

	text := fmt.Sprintf(`💼 Работа завершена!

%s

💰 Заработано: %d₽
❤️ Здоровье: %d/%d
🍽️ Голод: %d/%d

Бомж устал, но доволен! 😊`,
		workType, earnings, bomj.Health, bomj.MaxHealth, bomj.Hunger, bomj.MaxHunger)

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func sendFoodMenu(userID int64) {
	text := `🍽️ Выбери еду для бомжа:

🍕 еда_пицца - Пицца (убирает голод на 30, +15 здоровья)
🍔 еда_бургер - Бургер (убирает голод на 25, +10 здоровья)
🥪 еда_бутерброд - Бутерброд (убирает голод на 20, +8 здоровья)
🍎 еда_яблоко - Яблоко (убирает голод на 15, +5 здоровья)
🥖 еда_хлеб - Хлеб (убирает голод на 10, +3 здоровья)
🍺 еда_пиво - Пиво (убирает голод на 5, -5 здоровья, но +10 настроения)`

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func handleFoodSelection(userID int64, foodChoice string) {
	bomjObj := storage.GetBomj(userID)
	if bomjObj == nil {
		sendWelcomeMessage(userID)
		return
	}

	var food bomj.Item

	switch foodChoice {
	case "еда_пицца":
		food = bomj.Item{ID: 101, Name: "Пицца", Type: "food", Value: 30, Quantity: 1}
	case "еда_бургер":
		food = bomj.Item{ID: 102, Name: "Бургер", Type: "food", Value: 25, Quantity: 1}
	case "еда_бутерброд":
		food = bomj.Item{ID: 103, Name: "Бутерброд", Type: "food", Value: 20, Quantity: 1}
	case "еда_яблоко":
		food = bomj.Item{ID: 104, Name: "Яблоко", Type: "food", Value: 15, Quantity: 1}
	case "еда_хлеб":
		food = bomj.Item{ID: 105, Name: "Хлеб", Type: "food", Value: 10, Quantity: 1}
	case "еда_пиво":
		food = bomj.Item{ID: 106, Name: "Пиво", Type: "food", Value: 5, Quantity: 1}
	default:
		sendFoodMenu(userID)
		return
	}

	// Кормим бомжа
	success := bomjObj.Feed(food)
	storage.SaveBomj(bomjObj)

	if !success {
		text := "😵 Бомж не может больше есть! Он уже сыт."
		msg := tgbotapi.NewMessage(userID, text)
		bot.Send(msg)
		return
	}

	// Специальные эффекты для пива
	extraMessage := ""
	if foodChoice == "еда_пиво" {
		bomjObj.Health = max(0, bomjObj.Health-5)
		extraMessage = "\n🍺 Бомж выпил пиво! Здоровье -5, но настроение +10! 🎉"
	}

	text := fmt.Sprintf(`🍽️ Бомж поел %s!

❤️ Здоровье: %d/%d
🍽️ Голод: %d/%d%s

Бомж доволен! 😋`,
		food.Name, bomjObj.Health, bomjObj.MaxHealth, bomjObj.Hunger, bomjObj.MaxHunger, extraMessage)

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func handleRandomEvent(userID int64) {
	bomj := storage.GetBomj(userID)
	if bomj == nil {
		sendWelcomeMessage(userID)
		return
	}

	description, moneyChange := bomj.GetRandomEvent()
	storage.SaveBomj(bomj)

	text := fmt.Sprintf(`🎲 Случайное событие!

%s

💰 Изменение денег: %+d₽
❤️ Здоровье: %d/%d
🍽️ Голод: %d/%d

Жизнь бомжа полна приключений! 🚀`,
		description, moneyChange, bomj.Health, bomj.MaxHealth, bomj.Hunger, bomj.MaxHunger)

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

func sendHelp(userID int64) {
	text := `❓ Помощь по игре "Карманный Бомж"

🎯 Основные команды:
/start - главное меню
/stats - статистика бомжа
/work - отправить на работу
/feed - покормить бомжа
/event - случайное событие
/help - эта справка

🍽️ Еда:
Отправь "еда_название" чтобы покормить бомжа
Доступные варианты: пицца, бургер, бутерброд, яблоко, хлеб, пиво

💡 Советы:
• Работать можно раз в час
• Кормить бомжа нужно регулярно
• Случайные события происходят каждый день
• Повышай уровень для улучшения характеристик

Удачи в заботе о своем бомже! 🚀`

	msg := tgbotapi.NewMessage(userID, text)
	bot.Send(msg)
}

// Helper function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
