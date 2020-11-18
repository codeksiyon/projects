import requests
import json

twitter_post_url  = "https://twitter.com/i/api/1.1/statuses/update.json"

myTwitterHeaders = {
    "User-Agent":"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0",
    "Content-Type": "application/x-www-form-urlencoded",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "tr",
    "x-twitter-active-user": "yes",
    "x-csrf-token": "03d77a8cee17195593014ce",
    "Origin": "https://twitter.com",
    "DNT": "1",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAA8xnZz4puTs%3D1fk8LF81IUq16cHjhL33AGWWjCpTnA",
}

myTwitterCookies={
    "personalization_id":"\"v1_bRX14Td2KJkC/NA==\"",
    "guest_id":"v1%3A160182411",
    "external_referer":"pad7zjgzgv1mFWxJ12Owit7owX|0|8d8A2w%3D",
    "ct0":"03d77a8c49f4e593014c9a2e",
    "gt":"1322591481382",
    "_twitter_sess":"BAh7CiIKZmxhQzonQWN0NjoJdXNlcmwrCQCAV3ZkxOkO--a7dd38527303e6e692d4eef5258cd",
    "dnt":"1",
    "ads_prefs":"HBISAAA=",
    "kdt":"vF2DFRjScEjLre6Nbtm44Armv2Qk9",
    "remember_checked_on":"1",
    "twid":"u%3D107460593536",
    "auth_token":"dbd5ae9d9aa4668ae6182bc0a6",
}

myTwitterData={
    "include_profile_interstitial_type":"1",
    "include_blocking":"1",
    "include_blocked_by":"1",
    "include_followed_by":"1",
    "include_want_retweets":"1",
    "include_mute_edge":"1",
    "include_can_dm":"1",
    "include_can_media_tag":"1",
    "skip_status":"1",
    "cards_platform":"Web-12",
    "include_cards":"1",
    "include_ext_alt_text":"true",
    "include_quote_count":"true",
    "include_reply_count":"1",
    "tweet_mode":"extended",
    "simple_quoted_tweet":"true",
    "trim_user":"false",
    "include_ext_media_color":"true",
    "include_ext_media_availability":"true",
    "auto_populate_reply_metadata":"false",
    "batch_mode":"off",
    "status":"t.me/raifBlog" # @raifpy

}

def sendPost(text):
    
    myTwitterData["status"] = text
    response = requests.post(url=twitter_post_url,headers=myTwitterHeaders,data=myTwitterData,cookies=myTwitterCookies)
    if response.status_code != 200:
        exit(response.status_code)
    
    jsonVeri = json.loads(response.text)
    for key in jsonVeri.keys():
        print(key," : ",jsonVeri[key])


if __name__ == "__main__":
    text = input("Text : ")
    sendPost(text)