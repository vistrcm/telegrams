# telegrams
Simple library to work with [Telegram Bot API](https://core.telegram.org/bots).

Initialize bot:
```go
bot := telegrams.NewTelegramBot("bot_token")
```

GetMe. A simple method for testing your bot's auth token.
```go
me, _ := bot.GetMe()
```


## Thanks
Thanks to [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) project for some ideas and code pieces.
