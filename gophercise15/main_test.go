package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMain(t *testing.T){
	defer func(){
		time.Sleep(1 * time.Second)	
	}()
	go main()
		
}


func TestHandleDebug(t *testing.T) {
	urlquery :="/debug/?path=/home/tanvik/gocode/src/gophercises/gophercise15/main.go"
	req, _ := http.NewRequest("GET",urlquery,nil)
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(DebugCode)
	handler.ServeHTTP(res,req)
	expected := 200
	status := res.Code
	if status != expected{
		t.Errorf("handler returned unexpected status: got %v want %v",status, expected)
	}else{
		fmt.Println("Successfully passed test... ")

	}	
}

func TestFilePath(t *testing.T){
	wrongPath := "/debug/?path=demo"
	_,err := os.Open(wrongPath)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print("\n")
	req,_ := http.NewRequest("GET","",nil)
	res := httptest.NewRecorder()
	handler := (http.HandlerFunc(DebugCode))
	handler.ServeHTTP(res,req)
	
}


func TestHandleExc(t *testing.T){
	req, err := http.NewRequest("GET","/panic",nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := RecoverFunc(http.HandlerFunc(HandleException))
	handler.ServeHTTP(res,req)
}


