package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", UI)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/subtract", Subtract)
	http.HandleFunc("/multiply", Multiply)
	http.HandleFunc("/divide", Divide)
	fmt.Println("listening...")
	err := http.ListenAndServe(":8024", nil)
	if err != nil {
		panic(err)
	}
}

func Add(res http.ResponseWriter, req *http.Request) {
	a, b, err := ParseArgs(req)
	if err != nil {
		RespondError(res, err.Error())
		return
	}
	result := a + b
	RespondResult(res, result)
}

func Subtract(res http.ResponseWriter, req *http.Request) {
	a, b, err := ParseArgs(req)
	if err != nil {
		RespondError(res, err.Error())
		return
	}
	result := a - b
	RespondResult(res, result)
}

func Multiply(res http.ResponseWriter, req *http.Request) {
	a, b, err := ParseArgs(req)
	if err != nil {
		RespondError(res, err.Error())
		return
	}
	result := a * b
	RespondResult(res, result)
}

func Divide(res http.ResponseWriter, req *http.Request) {
	a, b, err := ParseArgs(req)
	if err != nil {
		RespondError(res, err.Error())
		return
	}
	result := a / b
	RespondResult(res, result)
}

func UI(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		RespondError(res, "no operation selected")
		return
	}
	RespondHTML(res, "index.html")
}

func ParseArgs(req *http.Request) (float64, float64, error) {
	a, err := strconv.ParseFloat(req.FormValue("a"), 64)
	if err != nil {
		return 0, 0, errors.New("bad first argument: " + req.FormValue("a"))
	}
	b, err := strconv.ParseFloat(req.FormValue("b"), 64)
	if err != nil {
		return 0, 0, errors.New("bad second argument: " + req.FormValue("b"))
	}
	return a, b, nil
}

func RespondError(res http.ResponseWriter, message string) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(res, message)
}

func RespondResult(res http.ResponseWriter, result float64) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, result)
}

func RespondHTML(res http.ResponseWriter, file string) {
	f, err := os.Open(file)
	if err != nil {
		RespondError(res, err.Error())
		return
	}
	defer f.Close()
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	io.Copy(res, f)
}
