package main

// t.me/Codeksiyon | raifpy
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/headzoo/surf"

	"fyne.io/fyne"

	"fyne.io/fyne/dialog"

	"fyne.io/fyne/theme"

	"fyne.io/fyne/widget"

	"fyne.io/fyne/app"

	"github.com/PuerkitoBio/goquery"

	"github.com/raifpy/Go/errHandler"
)

func gui() {

	errHandler.HandlerExit(os.Chdir("/data/data/org.codeksiyon.faceinstall"))

	uyg := app.New()
	uyg.Settings().SetTheme(theme.DarkTheme())
	win := uyg.NewWindow("Facebook Video Downloader ")

	if _, err := ioutil.ReadDir("/sdcard/Download"); err != nil {
		uyg.SendNotification(fyne.NewNotification("Access Denied !", "I Don't Have Read-Write File Permission !"))

	}

	entry := widget.NewEntry()
	entry.SetPlaceHolder("https://wwww.facebook.com/groups/codeksiyon/video...")

	buton := widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {
		cookies := getCookies(win)
		if cookies == nil {
			return
		}
		if entry.Text == "" {
			dialog.ShowInformation("Something went wrong", "URL Entry can't be empty  :))", win)
			return
		}
		url := urlformat(entry.Text)
		resp, err := request(url, cookies)
		if err != nil {
			dialog.ShowInformation("Something went wrong", err.Error(), win)
			return
		}
		url = getVideo(resp.Body)
		if url == "" {
			dialog.ShowInformation("Something went wrong", "I can't find video.\nMaybe your cookies are wrong.\n"+resp.Request.URL.String(), win)
			return
		}
		dosya := "/sdcard/Download/" + time.Now().Format("15:04:05") + ".mp4"

		err = indirVideo(url, dosya, cookies)
		if err != nil {
			uyg.SendNotification(fyne.NewNotification("Error to save video!", err.Error()))
			dialog.ShowInformation("Someting went wrong", err.Error(), win)
			return
		}

		dialog.ShowInformation("Saved", fmt.Sprintf("%s\n location succesfully ", dosya), win)

	})

	clearCookies := widget.NewButtonWithIcon("Delete Cookies", theme.DeleteIcon(), func() {
		err := os.Remove("./cookies")
		if errHandler.HandlerBool(err) {
			dialog.ShowError(err, win)
			return
		}
		dialog.ShowInformation("Removed", "Your cookies was romoved", win)
	})
	about := widget.NewButtonWithIcon("About", theme.InfoIcon(), func() {
		author := widget.NewLabel("Codeksiyon - Facebook Installer")
		howWork := widget.NewTextGridFromString("login facebook & get cookies\nreplace www.face to mbasic.face\nrequests with cookies\nfind a-href startswith '/video'\ndownload video")
		url, _ := url.Parse("https://t.me/Codeksiyon")
		hyper := widget.NewHyperlink("t.me/Codeksiyon", url)

		dialog.ShowCustom("About", "  Back  ", widget.NewVBox(author, howWork, hyper), win)
	})

	group := widget.NewGroup("@codeksiyon Facebook Video", entry, buton, clearCookies, about)

	win.SetContent(group)
	win.ShowAndRun()
}

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			log.Println("PANIC : ", rec)
			f, _ := os.Create("/sdcard/Download/codeksiyon_panic.txt")
			f.WriteString(fmt.Sprintln(rec))
			f.Close()
		}
	}()
	gui()
}

func showError(win fyne.Window, err error) {
	raifpy, _ := url.Parse("https://t.me/raifpy")
	link := widget.NewHyperlink("Click to pm me", raifpy)
	shit := widget.NewTextGridFromString("UnExcepted Error Catched !\n" + err.Error() + "\n")
	shitLayout := widget.NewVBox(shit, link)

	dialog.ShowCustom("Error", "Ok", shitLayout, win)
}

func takeCookies(win fyne.Window) {
	dialog.ShowConfirm("Login", "I need your cookies\nfor download facebook videos.\nThis operation is for one use.\nDo you want to continue?", func(ok bool) {
		if ok {
			faceUsername := widget.NewEntry()
			facePassword := widget.NewPasswordEntry()
			labelText := widget.NewTextGrid()
			faceLoginbtn := widget.NewButton("Login", func() {

				if strings.Trim(faceUsername.Text, " ") == "" || strings.Trim(facePassword.Text, " ") == "" {
					dialog.ShowError(fmt.Errorf("Entrys must be full"), win)
					//dialogKanal.Hide()
					return
				}
				browser := surf.NewBrowser()
				//browser.SetUserAgent("Mozilla/5.0 (Linux; Android 9; SM-J730F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36")
				if err := browser.Open("https://mbasic.facebook.com"); errHandler.HandlerBool(err) {
					showError(win, err)
				}
				query, err := goquery.NewDocumentFromReader(bytes.NewBufferString(browser.Body()))
				if errHandler.HandlerBool(err) {
					showError(win, err)
					return
				}
				class, exists := query.Find("form").Attr("class")
				if !exists {
					showError(win, fmt.Errorf("cannot find <form class='x'>"))
					return
				}
				class = strings.Replace(class, " ", ".", -1)
				sub, err := browser.Form("form." + class)
				if errHandler.HandlerBool(err) {
					showError(win, err)
					return
				}
				sub.Input("email", faceUsername.Text)
				sub.Input("pass", facePassword.Text)
				labelText.SetText("Requesting ..")
				if err = sub.Submit(); errHandler.HandlerBool(err) {
					showError(win, err)
					return
				}
				labelText.SetText("Checking ..")
				browser.Open("https://mbasic.facebook.com/me")
				if urlBrowse := browser.Url().String(); urlBrowse == "https://mbasic.facebook.com/?refsrc=https%3A%2F%2Fmbasic.facebook.com%2Fme&_rdr" || urlBrowse == "https://mbasic.facebook.com/me" {
					showError(win, fmt.Errorf("I Think Your Password Wrong :D\nor someting wrong !\n\nCheck\n/Download/codeksiyon_login_log.txt"))
					labelText.SetText("Your Password Wrong!")
					return
				}
				fmt.Println(browser.Url())
				fmt.Println(browser.Title())

				labelText.SetText("Success !")
				var cokstr string
				for _, islem := range browser.SiteCookies() {
					cokstr += islem.Name + " " + islem.Value + "\n"
				}
				file, err := os.Create("cookies")
				if errHandler.HandlerBool(err) {
					showError(win, err)
					return
				}
				_, err = file.WriteString(cokstr)
				if errHandler.HandlerBool(err) {
					showError(win, err)
					return
				}
				query, _ = goquery.NewDocumentFromReader(bytes.NewBufferString(browser.Body()))
				dialog.ShowInformation("Success!", "Your cookies were saved\nWelcome "+query.Find("title").Text(), win)

				return

			})

			faceUsername.SetPlaceHolder("number or email")
			facePassword.SetPlaceHolder("password")
			group := widget.NewVBox(labelText, faceUsername, facePassword, faceLoginbtn)
			dialog.ShowCustom("Login Facebook", "            Back            ", group, win)
		}
	}, win)
}
func getCookies(win fyne.Window) map[string]string {
	/*var cookies = map[string]string{
		"sb":     "",
		"datr":   "",
		"c_user": "",
		"spin":   "",
		"xs":     "",
		"fr":     "",
	}*/
	file, err := ioutil.ReadFile("./cookies")
	if errHandler.HandlerBool(err) {
		if os.IsNotExist(err) {
			takeCookies(win)
			return nil

		} else {
			dialog.ShowError(err, win)
			return nil
		}
	} else {
		kuk := map[string]string{}
		fileStr := string(file)
		for _, value := range strings.Split(fileStr, "\n") {
			if value != "" {
				key := strings.Split(value, " ")
				kuk[key[0]] = key[1]
			}
		}

		return kuk

	}
}

func indirVideo(url, isim string, cookies map[string]string) error {

	url = "https://mbasic.facebook.com" + url

	videoResp, err := request(url, cookies)
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

func request(url string, cookies map[string]string) (*http.Response, error) {
	//byteHeaders, err := json.Marshal(headers)
	//errHandler.HandlerExit(err)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))

	if err != nil {
		return &http.Response{}, err
	}
	errHandler.HandlerExit(err)

	for key, value := range cookies {
		req.AddCookie(&http.Cookie{Name: key, Value: value})
	}
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 9; SM-J730F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36")
	/*
		req.AddCookie(&http.Cookie{Name: "sb", Value: "SIDOUFENWMRHLJKSDFUOSID"}) // Special
		req.AddCookie(&http.Cookie{Name: "datr", Value: "a7e_asduUIOPASD"})
		req.AddCookie(&http.Cookie{Name: "c_user", Value: "3284023984092834"})
		req.AddCookie(&http.Cookie{Name: "spin", Value: "r.1asdsad.trasdsa_t.1asdasdsa94sadsa_8_s.1_v.2_"})
		req.AddCookie(&http.Cookie{Name: "xs", Value: "50%3AXoI7VLYHanidTASUSDUIPAMNQWEUOIQW8MNScMh7v2nIHRavRq9bQUUuJSwpSSRH"})
		req.AddCookie(&http.Cookie{Name: "fr", Value: "1Wjpg6S7Xy.yetwqtnca80.BfXK4r.AWWnYaLK"})
	*/

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
	if alink == "" {
		file, _ := os.Create("/sdcard/Download/codeksiyon_login_log.txt")
		//file, _ := os.Create("log_codeksiyon.txt")
		file.WriteString(doc.Text())
		file.Close()
	}
	return alink
}

// t.me/Codeksiyon | raifpy
