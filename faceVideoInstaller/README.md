# Facebook Video Downloader With Cookies

How:

    Replace www.facebook.com to mbasic.facebook.com
    Find Video redirect url and get redirect url
    Download video ( Android: Download/H:M:S.mp4  Desktop:PWD/H:M:S.mp4 )
    

<h2>Android</h2>
Compile:
 
    fyne package -os android/arm -appID "org.codeksiyon.faceinstaller" .
    adb install ./*.apk

<img src="https://github.com/codeksiyon/projects/blob/master/faceVideoInstaller/img/face-android2.jpg" height=400>

<h2>Desktop</h2>
Compile:

    go build .
    
<img src="https://github.com/codeksiyon/projects/blob/master/faceVideoInstaller/img/face-desktop2.png" height=400>
