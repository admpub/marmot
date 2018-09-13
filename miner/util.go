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
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/admpub/log"
)

// Wait some secord
func Wait(waittime int) {
	if waittime <= 0 {
		return
	}
	// debug
	log.Debugf("Wait %d Second.", waittime)
	time.Sleep(time.Duration(waittime))
}

// CopyM Header map[string][]string ,can use to copy a http header, so that they are not effect each other
func CopyM(h http.Header) http.Header {
	if h == nil || len(h) == 0 {
		return h
	}
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}

//TooSortSizes if a file size small than sizes(KB) ,it will be throw a error
func TooSortSizes(data []byte, sizes float64) error {
	if float64(len(data))/1000 < sizes {
		return fmt.Errorf("FileSize:%d bytes,%d kb < %f kb dead too sort", len(data), len(data)/1000, sizes)
	}
	return nil
}

// OutputMaps Just debug a map
func OutputMaps(info string, args map[string][]string) {
	s := "\n"
	for k, v := range args {
		s = s + fmt.Sprintf("%-25s| %-6s\n", k, strings.Join(v, "||"))
	}
	log.Debugf("[GoWorker] %s", s)
}
