package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const twitterPostURL = "https://twitter.com/i/api/1.1/statuses/update.json"

var myTwitterHeaders = map[string]string{
	"User-Agent":                "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0",
	"Content-Type":              "application/x-www-form-urlencoded",
	"x-twitter-auth-type":       "OAuth2Session",
	"x-twitter-client-language": "tr",
	"x-twitter-active-user":     "yes",
	"x-csrf-token":              "03d77a8c144e17593014e",
	"Origin":                    "https://twitter.com",
	"DNT":                       "1",
	"authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAnNwIzUejRC5E6I8xnpuTs%3D1Zv7k8LF81I6cHjhLTv4FA33AGWWjCpTnA",
}

var myTwitterCookies = map[string]string{
	"personalization_id":  "\"v1_bRd2KJkC/NA==\"",
	"guest_id":            "v1%3A1601811420",
	"external_referer":    "padhuUpFWxJ12Ozwit7owX|0|8d8A2w%3D",
	"ct0":                 "03d77a8ce51495593014c9a2e",
	"gt":                  "132259148213592",
	"_twitter_sess":       "BAh7CiSGFzaHsABjQHVzZWR7DoPY3JlYXRlZF9hdGwrCMjMtH91AToMY3NyZl9p%250AZCIlMWM4OWU1ZGQ3MmNkMjM4OTI4YjVhZjA5NjFlYjQxYmQ6B2lkI---id",
	"dnt":                 "1",
	"ads_prefs":           "HBISAAA=",
	"kdt":                 "vF2DFRiLLre6Nbtm57a3h44Armv2kQk9",
	"remember_checked_on": "1",
	"twid":                "u%3D1005921865793536",
	"auth_token":          "dbd5ae9d9aa4182bc0a6",
}

var myTwitterData = url.Values{
	"include_profile_interstitial_type": {"1"},
	"include_blocking":                  {"1"},
	"include_blocked_by":                {"1"},
	"include_followed_by":               {"1"},
	"include_want_retweets":             {"1"},
	"include_mute_edge":                 {"1"},
	"include_can_dm":                    {"1"},
	"include_can_media_tag":             {"1"},
	"skip_status":                       {"1"},
	"cards_platform":                    {"Web-12"},
	"include_cards":                     {"1"},
	"include_ext_alt_text":              {"true"},
	"include_quote_count":               {"true"},
	"include_reply_count":               {"1"},
	"tweet_mode":                        {"extended"},
	"simple_quoted_tweet":               {"true"},
	"trim_user":                         {"false"},
	"include_ext_media_color":           {"true"},
	"include_ext_media_availability":    {"true"},
	"auto_populate_reply_metadata":      {"false"},
	"batch_mode":                        {"off"},
	"status":                            {"t.me/raifBlog"}, // @raifpy
}

func sendPost(text string) {
	myTwitterData["status"][0] = text
	request, err := http.NewRequest("POST", twitterPostURL, strings.NewReader(myTwitterData.Encode()))
	if err != nil {
		panic(err)
	}
	for key, value := range myTwitterHeaders {
		request.Header.Set(key, value)
	}
	for key, value := range myTwitterCookies {
		request.AddCookie(&http.Cookie{Name: key, Value: value})
	}

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		panic(response.StatusCode)
	}

	byteResponse, _ := ioutil.ReadAll(response.Body)
	//stringResponse := string(byteResponse)

	fmt.Println(string(byteResponse))

}

func main() {
	var text string

	fmt.Print("Text : ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()

	}
	sendPost(text)

}
