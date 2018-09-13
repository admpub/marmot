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
	"log"
	"testing"
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
	err = TooSortSizes(body, 500)
	log.Println(err.Error())
}
