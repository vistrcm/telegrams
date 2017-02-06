# telegrams
Simple library to work with [Telegram Bot API](https://core.telegram.org/bots).

Initialize bot:
```go
bot := telegrams.NewTelegramBot("bot_token")
```

GetMe. A simple method for testing your bot's auth token.
```go
var me telegrams.User = bot.GetMe()
```
