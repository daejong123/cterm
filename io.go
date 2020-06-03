package cterm

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
)

func getUserHomeDir() string {
	var home string
	switch runtime.GOOS {
	case "windows":
		home, _ = os.LookupEnv("LOCALAPPDATA")
	case "linux":
		home, _ = os.LookupEnv("HOME")
	case "darwin":
		home, _ = os.LookupEnv("HOME")
	}
	return home
}

func getCtermDir() string {
	ctermDir := getUserHomeDir() + "/.cterm/"
	if _, err := os.Stat(ctermDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(ctermDir, 0644)
			if err != nil {
				panic(err)
			}
		}
	}
	return ctermDir
}

func dataFilePath() string {
	ctermDir := getCtermDir()
	ctermDataFilePath := ctermDir + "data.json"
	return ctermDataFilePath
}

func writeToDatafile(data []IndexListSourceDataType) {
	if v, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		ioutil.WriteFile(dataFilePath(), v, 0644)
	}
}

func readFromDataFile() []IndexListSourceDataType {
	data := []IndexListSourceDataType{}
	if v, err := ioutil.ReadFile(dataFilePath()); err != nil {
		return data
	} else {
		if e := json.Unmarshal(v, &data); e != nil {
			return data
		}
	}
	return data
}
