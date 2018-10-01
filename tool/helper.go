package tool

import (
	"errors"
	"fmt"
	"strings"
)

func Input(say, defaults string) string {
	fmt.Println(say)
	var str string
	fmt.Scanln(&str)
	if strings.TrimSpace(str) == "" {
		if strings.TrimSpace(defaults) != "" {
			return defaults
		}
		fmt.Println("can not empty")
		return Input(say, defaults)
	}
	//fmt.Println("--" + str + "--")
	return str
}

var (
	invalidFileNameRelacer = strings.NewReplacer([]string{
		" ", "#01#",
		"\\", "#02#",
		"/", "#03#",
		":", "#04#",
		"\"", "#05#",
		"?", "#06#",
		"<", "#07#",
		">", "#08#",
		"|", "#09#",
	}...)
)

// DevideStringList change by python
func DevideStringList(files []string, num int) (map[int][]string, error) {
	length := len(files)
	split := map[int][]string{}
	if num <= 0 {
		return split, errors.New("num must not negtive")
	}
	if num > length {
		num = length
	}
	process := length / num
	for i := 0; i < num; i++ {
		// slice inside has a refer, so must do this append
		//split[i]=files[i*process : (i+1)*process] wrong!
		split[i] = append(split[i], files[i*process:(i+1)*process]...)
	}
	remain := files[num*process:]
	for i := 0; i < len(remain); i++ {
		split[i%num] = append(split[i%num], remain[i])
	}
	return split, nil
}
