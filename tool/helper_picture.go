// Date: 18-2-9

package tool

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/admpub/marmot/expert"
	"github.com/admpub/marmot/miner"
)

// DownloadHTMLPictures Download one HTML page's all pictures
// @URL: http://image.baidu.com
// @SaveDir /home/images
// @ProxyAddress : "socks5://127.0.0.1:1080"
func DownloadHTMLPictures(URL string, SaveDir string, MinerNum int, ProxyAddress interface{}) error {

	// Check valid
	_, err := url.Parse(URL)
	if err != nil {
		return err
	}

	// New a worker to get url
	worker, _ := miner.New(ProxyAddress)

	result, err := worker.SetURL(URL).SetUserAgent(miner.RandomUserAgent()).Get()
	if err != nil {
		return err
	}

	// Find all picture
	pictures := expert.FindPicture(string(result))

	return DownloadURLPictures(pictures, SaveDir, MinerNum, ProxyAddress)
}

// DownloadURLPictures Download pictures faster!
func DownloadURLPictures(PictureUrls []string, SaveDir string, MinerNum int, ProxyAddress interface{}) error {
	// Empty, What a pity!
	if len(PictureUrls) == 0 {
		return errors.New("empty")
	}

	// Make dir!
	err := os.MkdirAll(SaveDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Divide pictures into several worker
	xxx, _ := DevideStringList(PictureUrls, MinerNum)

	// Chanel to info exchange
	chs := make(chan int, len(PictureUrls))

	// Go at the same time
	for num, pictureList := range xxx {

		// Get pool miner
		workerPicture, ok := miner.Pool.Get(strconv.Itoa(num))
		if !ok {
			// No? set one!
			workerTemp, _ := miner.New(ProxyAddress)
			workerPicture = workerTemp
			workerTemp.SetUserAgent(miner.RandomUserAgent())
			miner.Pool.Set(strconv.Itoa(num), workerTemp)
		}

		// Go save picture!
		go func(images []string, worker *miner.Worker, num int) {
			for _, img := range images {

				// Check, May be Pass
				_, err := url.Parse(img)
				if err != nil {
					continue
				}

				// Change Name of our picture
				filename := strings.Replace(invalidFileNameRelacer.Replace(img), "#", "_", -1)

				// Exist?
				if _, err := os.Stat(SaveDir + "/" + filename); os.IsExist(err) {
					fmt.Println("File Exist：" + SaveDir + "/" + filename)
					chs <- 0
				} else {

					// Not Exsit?
					imgsrc, e := worker.SetURL(img).Get()
					if e != nil {
						fmt.Println("Download " + img + " error:" + e.Error())
						chs <- 0
						return
					}

					// Save it!
					e = ioutil.WriteFile(SaveDir+"/"+filename, imgsrc, os.ModePerm)
					if e == nil {
						fmt.Printf("SP%d: Keep in %s/%s\n", num, SaveDir, filename)
					}
					chs <- 1
				}
			}
		}(pictureList, workerPicture, num)
	}

	// Every picture should return
	for i := 0; i < len(PictureUrls); i++ {
		<-chs
	}

	return nil
}
