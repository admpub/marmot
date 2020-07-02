//
// 	Copyright 2017 by marmot author: gdccmcm14@live.com.
// 	Licensed under the Apache License, Version 2.0 (the "License");
// 	you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
// 		http://www.apache.org/licenses/LICENSE-2.0
// 	Unless required by applicable law or agreed to in writing, software
// 	distributed under the License is distributed on an "AS IS" BASIS,
// 	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// 	See the License for the specific language governing permissions and
// 	limitations under the License
//

package miner

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/admpub/gohttp"
	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	// GLOBAL TIMEOUT
	SetGlobalTimeout(3)

	// a new spider without proxy
	// NewWorker(nil)
	worker, err := NewWorker(nil)

	//proxy := "http://smart:smart2016@104.128.121.46:808"
	//worker, err := NewWorker(proxy)

	if err != nil {
		panic(err)
	}
	// method can be get and post
	worker.SetMethod(GET)
	// wait times,can zero
	worker.SetWaitTime(1)
	// which url fetch
	worker.SetURL("http://www.cjhug.me")

	//worker.SetUserAgent(spider.RandomUserAgent())

	// go!fetch url --||
	body, err := worker.Go()
	if err != nil {
		log.Println(err.Error())
	} else {
		// bytes get!
		// fmt.Printf("%s", string(body))
	}

	// if filesize small than 500KB
	err = TooShortSizes(body, 500)
	log.Println(err.Error())
}

//var testURL = `http://127.0.0.1/ping.php`
var testURL = `http://127.0.0.1:9998/`
var testData = []byte(`[1,2,3,4,5]`)

func postByWorker(t *testing.T) {
	worker := NewAPI()
	worker.SetURL(testURL)
	body, err := worker.SetBinary(testData).PostJSON()
	//body, err := worker.SetForm(url.Values{`data`: []string{string(testData)}}).Post()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, string(testData), string(body))
	assert.Equal(t, http.StatusOK, worker.StatusCode)

	worker.SetURL(testURL + `?post=1`)
	body, err = worker.SetForm(url.Values{`data`: []string{string(testData)}}).Post()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, string(testData), string(body))
	assert.Equal(t, http.StatusOK, worker.StatusCode)
}

func postByHTTP(t *testing.T) {
	client := gohttp.New()
	resp, errs := client.Post(testURL).Send(string(testData)).End()
	if len(errs) > 0 {
		var errStr string
		for i, e := range errs {
			if i > 0 {
				errStr += ";\n"
			}
			errStr += e.Error()
		}
		panic(errStr)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(testData), string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func server() {
	http.HandleFunc(`/`, func(resp http.ResponseWriter, req *http.Request) {
		var b []byte
		req.ParseForm()
		if req.URL.Query().Get(`post`) == `1` {
			b = []byte(req.PostForm.Get(`data`))
		} else {
			b, _ = ioutil.ReadAll(req.Body)
			req.Body.Close()
		}
		resp.WriteHeader(http.StatusOK)
		resp.Write(b)
	})
	go http.ListenAndServe(`:9998`, nil)
}

func TestPostJSON(t *testing.T) {
	server()
	postByWorker(t)
	postByHTTP(t)
}
