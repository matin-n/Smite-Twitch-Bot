# Smite-Twitch-Bot
Used to lookup ranked player statistics by connecting to HiRez API

## Configuration
1. Replace `devId` and `authKey`; given from HiRez 
2. Replace `channelName` which is the channel in which you want the bot to connect
3. Replace `BotUsername` with the twitch bot username & `oauth:CODE` with the twitch bot oauth code; this can be located inside of  `twitch.NewClient("BotUsername", "oauth:CODE")`
4. Run the bot

## Usage
* `!ranked playername` to lookup ranked statistics

----
## Libraries Used
* [go-twitch-irc](https://github.com/gempir/go-twitch-irc)
* [GJSON](https://github.com/tidwall/gjson)
