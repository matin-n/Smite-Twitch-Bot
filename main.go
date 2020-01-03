package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gempir/go-twitch-irc"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var sessionId = "" // leave empty
var timestamp = "" // leave empty

var devId = "" // replace devId
var authKey = "" // replace authkey



var channelName = "" // replace with channel username in which you want the bot to connect

func main() {

	go createSession()

	fmt.Println("Started...")


	client := twitch.NewClient("", "") // replace with twitch bot username and twitch bot oauth code

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if strings.Contains(message.Message, "!ranked ") {
			parts := strings.Split(message.Message, "!ranked ")
			rankedStats := gjson.Get(getPlayer(parts[1]), "0").String()
			// account level / hours played
			// win / loss
			// ranked / mmr

			if rankedStats != "" {

				accountLvl := gjson.Get(getPlayer(parts[1]), "0.Level").String()
				hoursPlayed := gjson.Get(getPlayer(parts[1]), "0.HoursPlayed").String()
				wins := gjson.Get(getPlayer(parts[1]), "0.RankedConquest.Wins").String()
				losses := gjson.Get(getPlayer(parts[1]), "0.RankedConquest.Losses").String()
				mmr := gjson.Get(getPlayer(parts[1]), "0.RankedConquest.Rank_Stat").String()
				rank := createRank(gjson.Get(getPlayer(parts[1]), "0.RankedConquest.Tier").String())

				client.Say(channelName, parts[1]+": Account Level "+accountLvl+" / "+"Hours Played "+hoursPlayed)
				time.Sleep(1 * time.Second)
				client.Say(channelName, parts[1]+": Wins "+wins+" / "+"Losses "+losses)
				time.Sleep(1 * time.Second)
				client.Say(channelName, parts[1]+": Rank "+rank+" / "+"MMR "+mmr)
			}else{
				client.Say(channelName, parts[1]+": Error on lookup / hidden profile")
			}
		}

	})

	client.Join(channelName)

	err := client.Connect()
	if err != nil {
		panic(err)
	}

}

//func createSession(devId, authKey, timestamp string
func createSession() {
	for {
		timestamp = GetCurrentTime()
		signature := createSignature(devId, "createsession", authKey, GetCurrentTime())

		resp, err := http.Get("http://api.smitegame.com/smiteapi.svc/createsessionJson/" + devId + "/" + signature + "/" + GetCurrentTime())
		if err != nil {
			// handle error
		}

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		sessionId = gjson.Get(bodyString, "session_id").String()
		time.Sleep(15 * time.Minute)
	}
}

func getPlayer(username string) string {
	signature := createSignature(devId, "getplayer", authKey, GetCurrentTime())

	resp, err := http.Get("http://api.smitegame.com/smiteapi.svc/getplayerjson/" + devId + "/" + signature + "/" + sessionId + "/" + GetCurrentTime() + "/" + username)

	if err != nil {
		// handle error
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return bodyString
}

func createSignature(devId, functionName, authKey, timestamp string) string {
	return GetMD5Hash(devId + functionName + authKey + timestamp)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetCurrentTime() string {
	// string timestamp = DateTime.UtcNow.ToString("yyyyMMddHHmmss");
	t := time.Now().UTC()
	return t.Format("20060102150405") // https://stackoverflow.com/a/20234207
}

func createRank(rank string) string {

	switch {
	case rank == "1":
		return "Bronze V"
	case rank == "2":
		return "Bronze IV"
	case rank == "3":
		return "Bronze III"
	case rank == "4":
		return "Bronze II"
	case rank == "5":
		return "Bronze I"
	case rank == "6":
		return "Silver V"
	case rank == "7":
		return "Silver IV"
	case rank == "8":
		return "Silver III"
	case rank == "9":
		return "Silver II"
	case rank == "10":
		return "Silver I"
	case rank == "11":
		return "Gold V"
	case rank == "12":
		return "Gold IV"
	case rank == "13":
		return "Gold III"
	case rank == "14":
		return "Gold II"
	case rank == "15":
		return "Gold I"
	case rank == "16":
		return "Platinum V"
	case rank == "17":
		return "Platinum IV"
	case rank == "18":
		return "Platinum III"
	case rank == "19":
		return "Platinum II"
	case rank == "20":
		return "Platinum I"
	case rank == "21":
		return "Diamond V"
	case rank == "22":
		return "Diamond IV"
	case rank == "23":
		return "Diamond III"
	case rank == "24":
		return "Diamond II"
	case rank == "25":
		return "Diamond I"
	case rank == "26":
		return "Masters"
	case rank == "27":
		return "Grandmasters"
	}

	return "0"
}
