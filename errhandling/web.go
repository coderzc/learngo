package main

import (
	"bufio"
	"errors"
	"github.com/gpmgo/gopm/log"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error
type httpRespFunc func(writer http.ResponseWriter, request *http.Request)

func handlerFileList(writer http.ResponseWriter, request *http.Request) error {
	if !strings.HasPrefix(request.URL.Path, "/list/") {
		return errors.New("path must be start with /list/")
	}
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	bw := bufio.NewWriter(writer)
	defer bw.Flush()
	bw.Write(all)
	return nil
}

func errWrapper(handler appHandler) httpRespFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Warn("painc:%v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}
		}()

		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			log.Warn("Err http request:%s", http.StatusText(code))
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(handlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
