package main

// t.me/Codeksiyon | raifpy
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne"

	"fyne.io/fyne/dialog"

	"fyne.io/fyne/theme"

	"fyne.io/fyne/widget"

	"fyne.io/fyne/app"

	"github.com/PuerkitoBio/goquery"

	"github.com/raifpy/Go/errHandler"
)

func gui() {
	uyg := app.New()
	uyg.Settings().SetTheme(theme.DarkTheme())
	win := uyg.NewWindow("Facebook Video Downloader ")

	if _, err := ioutil.ReadDir("/sdcard/Download"); err != nil {
		uyg.SendNotification(fyne.NewNotification("Access Denied !", "I Don't Have Write File Permission !"))
	} else {
		os.Chdir("/sdcard/Download")
	}

	entry := widget.NewEntry()
	entry.SetPlaceHolder("https://wwww.facebook.com/groups/codeksiyon/video...")

	buton := widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {
		if entry.Text == "" {
			dialog.ShowInformation("Something went wrong", "URL Entry can't be empty  :))", win)
			return
		}
		url := urlformat(entry.Text)
		resp, err := request(url)
		if err != nil {
			dialog.ShowInformation("Something went wrong", err.Error(), win)
			return
		}
		url = getVideo(resp.Body)
		if url == "" {
			dialog.ShowInformation("Something went wrong", "I can't find video.\nMaybe your cookies are wrong.\n"+resp.Request.URL.String(), win)
			return
		}
		dosya := time.Now().Format("15:04:05") + ".mp4"
		err = indirVideo(url, dosya)
		if err != nil {
			dialog.ShowInformation("Someting went wrong", err.Error(), win)
			return
		}

		dialog.ShowInformation("Saved", fmt.Sprintf("Video saved \n%s\n location succesfully "+"Download"+"/"+dosya), win)

	})

	group := widget.NewGroup("@codeksiyon Facebook Video Downloader With Cookies", entry, buton)
	win.SetContent(group)
	win.ShowAndRun()
}

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			log.Println("PANIC : ", rec)
		}
	}()
	gui()
}
func indirVideo(url, isim string) error {
	url = "https://mbasic.facebook.com" + url

	videoResp, err := request(url)
	defer videoResp.Body.Close()
	if err != nil {
		return err
	}
	file, err := os.Create(isim)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(file, videoResp.Body)

	return err

}

func request(url string) (*http.Response, error) {
	//byteHeaders, err := json.Marshal(headers)
	//errHandler.HandlerExit(err)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))

	if err != nil {
		return &http.Response{}, err
	}
	errHandler.HandlerExit(err)
	req.AddCookie(&http.Cookie{Name: "sb", Value: "SIDOUFENWMRHLJKSDFUOSID"}) // Special
	req.AddCookie(&http.Cookie{Name: "datr", Value: "a7e_asduUIOPASD"})
	req.AddCookie(&http.Cookie{Name: "c_user", Value: "3284023984092834"})
	req.AddCookie(&http.Cookie{Name: "spin", Value: "r.1asdsad.trasdsa_t.1asdasdsa94sadsa_8_s.1_v.2_"})
	req.AddCookie(&http.Cookie{Name: "xs", Value: "50%3AXoI7VLYHanidTASUSDUIPAMNQWEUOIQW8MNScMh7v2nIHRavRq9bQUUuJSwpSSRH"})
	req.AddCookie(&http.Cookie{Name: "fr", Value: "1Wjpg6S7Xy.yetwqtnca80.BfXK4r.AWWnYaLK"})

	istek := http.Client{}
	response, err := istek.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return response, nil

}
func urlformat(url string) string {
	return strings.Replace(url, "www", "mbasic", -1)
}
func getVideo(responeBody io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(responeBody)
	//fmt.Println(doc.Html())
	errHandler.Handler(err)
	var alink string
	doc.Find("a").Each(func(_ int, eleman *goquery.Selection) {
		//fmt.Println(eleman)
		if aa, varmi := eleman.Attr("href"); varmi {
			//fmt.Println(aa)
			if strings.HasPrefix(aa, "/video") {
				alink = aa
			}
		}

	})

	return alink
}

// t.me/Codeksiyon | raifpy
