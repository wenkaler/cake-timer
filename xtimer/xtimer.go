package xtimer

import (
	"fmt"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wenkaler/kit/log"
	"github.com/wenkaler/kit/log/level"
)

const (
	developer = 123
	admin     = iota + 1
	user
)

type XTimer struct {
	UPD     tgbotapi.UpdatesChannel
	BOT     *tgbotapi.BotAPI
	Storage Storage
}

type Storage interface {
	GetUsers()
	GetUser(id int) (User, error)
}

type User struct {
	ID         int64
	Name       string
	Permission int
	Date       time.Time
}

func New(token string, t int, logger log.Logger) (*XTimer, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		level.Error(logger).Log("msg", "failed create bot", "err", err)
		os.Exit(1)
	}
	level.Info(logger).Log("msg", "Authorized on account", "bot-name", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = t
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		level.Error(logger).Log("msg", "failed ged update bot chanel", "err", err)
		os.Exit(1)
	}
	xt := &XTimer{
		UPD: updates,
		BOT: bot,
	}
	return xt, nil
}

func (xt *XTimer) Run() {
	for u := range xt.UPD {
		if u.Message == nil {
			continue
		}
		if u.Message.IsCommand() {

		}
	}
}

func (xt *XTimer) command(u tgbotapi.Update) {
	switch u.Message.Command() {
	case "start":
	case "help":
	case "add":
	case "del":
	}
}

func (xt *XTimer) permision(user tgbotapi.User) bool {
	u, err := xt.Storage.GetUser(user.ID)
	if err != nil {
		xt.alert(fmt.Errorf("failed get user by id: %v", err))
	}
	if u.Permision == admin {
		return true
	}
	return false
}

func (xt *XTimer) alert(err error) {
	duty, err := xt.Storage.GetDuty()
	if err != nil {
		xt.BOT.Send()
	}
	msg := tgbotapi.NewMessage()
	xt.BOT.Send()
}
