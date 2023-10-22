package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	// 言語ファイルを読み込む
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("locales/ja.json")
	bundle.MustLoadMessageFile("locales/en.json")

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		// クライアントの言語を検出
		acceptLang := r.Header.Get("Accept-Language")
		langMatcher := language.NewMatcher(bundle.LanguageTags())
		tag, _, _ := langMatcher.Match(language.Make(acceptLang))

		// ローカライズされたメッセージを取得
		localizer := i18n.NewLocalizer(bundle, tag.String())
		greeting := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "greeting",
		})

		// ローカライズされたメッセージをクライアントに送信
		fmt.Fprintf(w, "%v\n", greeting)
	})

	http.ListenAndServe(":8080", nil)
}
