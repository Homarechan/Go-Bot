package prefcodes

const (
	_ = iota
	Hokkaido
	Aomori
	Iwate
	Miyagi
	Akita
	Yamagata
	Fukushima
	Ibaraki
	Tochigi
	Gunma
	Saitama
	Chiba
	Tokyo
	Kanagawa
	Nigata
	Toyama
	Ishikawa
	Fukui
	Yamanashi
	Nagano
	Gifu
	Shizuoka
	Aichi
	Mie
	Shiga
	Kyoto
	Osaka
	Hyogo
	Nara
	Wakayama
	Tottori
	Shimane
	Okayama
	Hiroshima
	Yamaguchi
	Tokushima
	Kagawa
	Ehime
	Kochi
	Fukuoka
	Saga
	Nagasaki
	Kumamoto
	Oita
	Miyazaki
	Kagoshima
	Okinawa
)

var NameToCode = map[string]int{
	"北海道":  Hokkaido,
	"青森県":  Aomori,
	"岩手県":  Iwate,
	"宮城県":  Miyagi,
	"秋田県":  Akita,
	"山形県":  Yamagata,
	"福島県":  Fukushima,
	"茨城県":  Ibaraki,
	"栃木県":  Tochigi,
	"群馬県":  Gunma,
	"埼玉県":  Saitama,
	"千葉県":  Chiba,
	"東京都":  Tokyo,
	"神奈川県": Kanagawa,
	"新潟県":  Nigata,
	"富山県":  Toyama,
	"石川県":  Ishikawa,
	"福井県":  Fukui,
	"山梨県":  Yamanashi,
	"長野県":  Nagano,
	"岐阜県":  Gifu,
	"静岡県":  Shizuoka,
	"愛知県":  Aichi,
	"三重県":  Mie,
	"滋賀県":  Shiga,
	"京都府":  Kyoto,
	"大阪府":  Osaka,
	"兵庫県":  Hyogo,
	"奈良県":  Nara,
	"和歌山県": Wakayama,
	"鳥取県":  Tottori,
	"島根県":  Shimane,
	"岡山県":  Okayama,
	"広島県":  Hiroshima,
	"山口県":  Yamaguchi,
	"徳島県":  Tokushima,
	"香川県":  Kagawa,
	"愛媛県":  Ehime,
	"高知県":  Kochi,
	"福岡県":  Fukuoka,
	"佐賀県":  Saga,
	"長崎県":  Nagasaki,
	"熊本県":  Kumamoto,
	"大分県":  Oita,
	"宮崎県":  Miyazaki,
	"鹿児島県": Kagoshima,
	"沖縄県":  Okinawa,
}

var CodeToName = map[int]string{
	Hokkaido:  "北海道",
	Aomori:    "青森県",
	Iwate:     "岩手県",
	Miyagi:    "宮城県",
	Akita:     "秋田県",
	Yamagata:  "山形県",
	Fukushima: "福島県",
	Ibaraki:   "茨城県",
	Tochigi:   "栃木県",
	Gunma:     "群馬県",
	Saitama:   "埼玉県",
	Chiba:     "千葉県",
	Tokyo:     "東京都",
	Kanagawa:  "神奈川県",
	Nigata:    "新潟県",
	Toyama:    "富山県",
	Ishikawa:  "石川県",
	Fukui:     "福井県",
	Yamanashi: "山梨県",
	Nagano:    "長野県",
	Gifu:      "岐阜県",
	Shizuoka:  "静岡県",
	Aichi:     "愛知県",
	Mie:       "三重県",
	Shiga:     "滋賀県",
	Kyoto:     "京都府",
	Osaka:     "大阪府",
	Hyogo:     "兵庫県",
	Nara:      "奈良県",
	Wakayama:  "和歌山県",
	Tottori:   "鳥取県",
	Shimane:   "島根県",
	Okayama:   "岡山県",
	Hiroshima: "広島県",
	Yamaguchi: "山口県",
	Tokushima: "徳島県",
	Kagawa:    "香川県",
	Ehime:     "愛媛県",
	Kochi:     "高知県",
	Fukuoka:   "福岡県",
	Saga:      "佐賀県",
	Nagasaki:  "長崎県",
	Kumamoto:  "熊本県",
	Oita:      "大分県",
	Miyazaki:  "宮崎県",
	Kagoshima: "鹿児島県",
	Okinawa:   "沖縄県",
}
