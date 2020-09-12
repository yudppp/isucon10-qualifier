package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	// iconv "github.com/djimenez/iconv-go"
)

var (
	chairFeatures = []string{
		"ヘッドレスト付き",
		"肘掛け付き",
		"キャスター付き",
		"アーム高さ調節可能",
		"リクライニング可能",
		"高さ調節可能",
		"通気性抜群",
		"メタルフレーム",
		"低反発",
		"木製",
		"背もたれつき",
		"回転可能",
		"レザー製",
		"昇降式",
		"デザイナーズ",
		"金属製",
		"プラスチック製",
		"法事用",
		"和風",
		"中華風",
		"西洋風",
		"イタリア製",
		"国産",
		"背もたれなし",
		"ラテン風",
		"布貼地",
		"スチール製",
		"メッシュ貼地",
		"オフィス用",
		"料理店用",
		"自宅用",
		"キャンプ用",
		"クッション性抜群",
		"モーター付き",
		"ベッド一体型",
		"ディスプレイ配置可能",
		"ミニ机付き",
		"スピーカー付属",
		"中国製",
		"アンティーク",
		"折りたたみ可能",
		"重さ500g以内",
		"24回払い無金利",
		"現代的デザイン",
		"近代的なデザイン",
		"ルネサンス的なデザイン",
		"アームなし",
		"オーダーメイド可能",
		"ポリカーボネート製",
		"フットレスト付き",
	}

	estateFeatures = []string{
		"最上階",
		"防犯カメラ",
		"ウォークインクローゼット",
		"ワンルーム",
		"ルーフバルコニー付",
		"エアコン付き",
		"駐輪場あり",
		"プロパンガス",
		"駐車場あり",
		"防音室",
		"追い焚き風呂",
		"オートロック",
		"即入居可",
		"IHコンロ",
		"敷地内駐車場",
		"トランクルーム",
		"角部屋",
		"カスタマイズ可",
		"DIY可",
		"ロフト",
		"シューズボックス",
		"インターネット無料",
		"地下室",
		"敷地内ゴミ置場",
		"管理人有り",
		"宅配ボックス",
		"ルームシェア可",
		"セキュリティ会社加入済",
		"メゾネット",
		"女性限定",
		"バイク置場あり",
		"エレベーター",
		"ペット相談可",
		"洗面所独立",
		"都市ガス",
		"浴室乾燥機",
		"インターネット接続可",
		"テレビ・通信",
		"専用庭",
		"システムキッチン",
		"高齢者歓迎",
		"ケーブルテレビ",
		"床下収納",
		"バス・トイレ別",
		"駐車場2台以上",
		"楽器相談可",
		"フローリング",
		"オール電化",
		"TVモニタ付きインタホン",
		"デザイナーズ物件",
	}
	chairFeaturesMap  map[string]int
	estateFeaturesMap map[string]int
)

func init() {
	chairFeaturesMap = make(map[string]int, len(chairFeatures))
	for i, cf := range chairFeatures {
		chairFeaturesMap[cf] = i + 1
	}

	estateFeaturesMap = make(map[string]int, len(estateFeatures))
	for i, ef := range estateFeatures {
		estateFeaturesMap[ef] = i + 1
	}
}

func main() {
	// chair
	{

		file, err := os.Open("chair_features.tsv")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader := csv.NewReader(file)
		reader.Comma = '\t'
		data := ""
		// remove header
		_, err = reader.Read() // 1行読み出し
		if err != nil {
			panic(err)
		}
		for {
			line, err := reader.Read() // 1行読み出し
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			if line[1] == "" {
				continue
			}
			for _, feature := range strings.Split(line[1], ",") {
				featureID, ok := chairFeaturesMap[feature]
				if !ok {
					panic(feature)
				}
				data += fmt.Sprintf("(%v, %v),", line[0], featureID)
			}
		}
		fmt.Printf(
			"INSERT INTO isuumo.chair_features (chair_id, feature_id) VALUES %s;\n",
			strings.TrimSuffix(data, ","),
		)
	}

	{
		file, err := os.Open("estate_features.tsv")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader := csv.NewReader(file)
		reader.Comma = '\t'
		data := ""
		// remove header
		_, err = reader.Read() // 1行読み出し
		if err != nil {
			panic(err)
		}
		for {
			line, err := reader.Read() // 1行読み出し
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			if line[1] == "" {
				continue
			}
			for _, feature := range strings.Split(line[1], ",") {
				featureID, ok := estateFeaturesMap[feature]
				if !ok {
					panic(feature)
				}
				data += fmt.Sprintf("(%v, %v),", line[0], featureID)
			}
		}
		fmt.Printf(
			"INSERT INTO isuumo.estate_features (estate_id, feature_id) VALUES %s;",
			strings.TrimSuffix(data, ","),
		)
	}
}
