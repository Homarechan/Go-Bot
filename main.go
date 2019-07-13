package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"

	"./datas"
	"./prefcodes"
)

var helpMessage = `コマンド一覧
路線一覧:[都道府県]
駅一覧:[駅番号]
駅情報:[駅番号]
所属路線一覧:[駅番号]
隣接駅:[路線番号]

使い方
[コマンド]:[引数]
例:
路線一覧: 大阪府
駅情報:1130224`

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if event.ReplyToken == "00000000000000000000000000000000" {
						return
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(parse(message.Text))).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func parse(message string) string {
	if startsWith(message, "路線一覧:") {
		return fetchText(message[13:len(message)])
	} else if startsWith(message, "駅一覧:") {
		code, err := strconv.Atoi(message[10:len(message)])
		if err != nil {
			return "エラー:\n路線一覧で取得した路線番号を入力してください"
		}
		return getLineData(code)
	} else if startsWith(message, "駅情報:") {
		code, err := strconv.Atoi(message[10:len(message)])
		if err != nil {
			return "エラー:\n駅一覧で取得した駅番号を入力してください"
		}
		return getStationData(code)
	} else if startsWith(message, "所属路線一覧:") {
		code, err := strconv.Atoi(message[19:len(message)])
		if err != nil {
			return "エラー:\n駅一覧で取得した駅番号を入力してください"
		}
		return getGroupData(code)
	} else if startsWith(message, "隣接駅:") {
		code, err := strconv.Atoi(message[10:len(message)])
		if err != nil {
			return "エラー:\n駅一覧で取得した駅番号を入力してください"
		}
		return getJoinData(code)
	}
	return helpMessage
}

func startsWith(str string, text string) bool {
	return strings.HasPrefix(str, text)
}

func fetchText(message string) string {
	text := message
	if len(text)/3 == 2 {
		if text == "大阪" || text == "京都" {
			text += "府"
		} else if text == "東京" {
			text = "東京都"
		} else {
			text += "県"
		}
	} else if len(text)/3 == 3 {
		if text == "神奈川" || text == "和歌山" {
			text += "県"
		}
	}

	_, ok := prefcodes.NameToCode[text]
	if !(ok) {
		return "都道府県名が無効です"
	}
	return getPrefData(text)
}

func getPrefData(pref string) string {
	linesInterface, err := datas.GetPrefData(prefcodes.NameToCode[pref])
	if err != nil {
		return err.Error()
	}

	lines := linesInterface.Lines
	result := "[" + pref + "]\n"
	for _, line := range lines {
		result += "\n"
		result += strconv.Itoa(line.Code)
		result += ": "
		result += line.Name
	}
	return result
}

func getLineData(linecode int) string {
	stations, err := datas.GetLineData(linecode)
	if err != nil {
		return "エラー:\n路線一覧で取得した路線番号を入力してください"
	}

	result := "[" + stations.Line.Name + "]\n"
	for _, station := range stations.Stations {
		result += "\n"
		result += strconv.Itoa(station.Code)
		result += ": "
		result += station.Name
	}
	return result
}

func getStationData(stationcode int) string {
	station, err := datas.GetStationData(stationcode)
	if err != nil {
		return "エラー:\n駅一覧で取得した駅番号を入力してください"
	}

	result := "[" + station.Station.Name + "]\n\n"
	result += "駅コード: " + strconv.Itoa(station.Station.Code)
	result += "\n"
	result += "駅グループコード: " + strconv.Itoa(station.Station.GroupCode)
	result += "\n"
	result += "路線: " + station.Station.LineName + "(" + strconv.Itoa(station.Station.LineCode) + ")"
	result += "\n"
	result += "都道府県: " + prefcodes.CodeToName[station.Station.PrefCode]
	result += "\n"
	result += "緯度: " + fmt.Sprintf("%f", station.Station.Latitude)
	result += "\n"
	result += "経度: " + fmt.Sprintf("%f", station.Station.Longtitude)

	return result
}

func getGroupData(stationcode int) string {
	group, err := datas.GetGroupData(stationcode)
	if err != nil {
		return "エラー:\n駅一覧で取得した駅番号を入力してください"
	}

	result := "[" + group.Station.Name + "]\n"
	for _, line := range group.GroupStations {
		result += "\n"
		result += strconv.Itoa(line.LineCode)
		result += ": "
		result += line.LineName
		result += "\n名称: "
		result += line.Name
	}
	return result
}

func getJoinData(stationcode int) string {
	joins, err := datas.GetJoinData(stationcode)
	if err != nil {
		return "エラー:\n路線一覧で取得した路線番号を入力してください"
	}

	line, err := datas.GetLineData(stationcode)
	if err != nil {
		return err.Error()
	}

	result := "[" + line.Line.Name + "]\n\n"
	for _, join := range joins.StationJoins {
		result += "隣接駅1:"
		result += join.Name1
		result += "("
		result += strconv.Itoa(join.Code1)
		result += ")\n"
		result += "隣接駅2:"
		result += join.Name2
		result += "("
		result += strconv.Itoa(join.Code2)
		result += ")\n\n"
	}
	return result
}
