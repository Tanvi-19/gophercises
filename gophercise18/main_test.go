package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gophercises/gophercise18/primitive"
)

type Function func(resp http.ResponseWriter, req *http.Request)
var File, _ = os.Open("./cartoon.png")


func TestMainFunc(test *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
}

func CheckFunc(hFunct Function, r *http.Request,imagbody bool, key string) string {
	var err error
	var method string
	var url string
	var query string 
	if imagbody {
		file := File
		body := &bytes.Buffer{}
		w := multipart.NewWriter(body)
		part, err := w.CreateFormFile(key, "test.png")
		if err != nil {
			log.Println(err)
		}
		io.Copy(part, file)
		w.Close()
		
		req, _ := http.NewRequest(method, url+query, body)
		req.Header.Set("Content-Type", w.FormDataContentType())
	} else {
		_, err = http.NewRequest(method, url+query, nil)
	}
	if err != nil {
		fmt.Println(err)
	}
	res := httptest.NewRecorder()
	
	return res.Body.String()
}	

func TestRoutHandler(test *testing.T) {
	req,_:= http.NewRequest("GET","/",nil)
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(RoutHandler)
	handler.ServeHTTP(res,req)
	outcome := 200
	if status := res.Code; status != outcome {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, outcome)
	}
	CheckFunc(RoutHandler,req,false,"")
}

func TestUpoadLink(test *testing.T) {
	req,_:= http.NewRequest("POST","/upload",nil)
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)
	handler.ServeHTTP(res,req)
	outcome := 200
	if status := res.Code; status != outcome {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, outcome)
	}
	CheckFunc(UploadHandler,req,true, "img")
	CheckFunc(UploadHandler, req, false, "")
}


func TestModifyHandler(test *testing.T) {
	req,_:= http.NewRequest("GET","/modify/GO.png",nil)
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ModifyHandler)
	handler.ServeHTTP(res,req)
	outcome := 400
	if status := res.Code; status != outcome {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, outcome)
	}
	CheckFunc(ModifyHandler,req,true, "img")
}


func TestModifyNegativeHandler(test *testing.T) {
	req,_:= http.NewRequest("GET","/modify/GO.png?mode=ux",nil)	
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ModifyHandler)
	handler.ServeHTTP(res,req)
	outcome := 400
	if status := res.Code; status != outcome {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, outcome)
	}
	CheckFunc(ModifyHandler,req,true, "img")
	GenImgFunc := generateImg
	oldGenImgFunc := GenImgFunc
	defer func() { GenImgFunc = oldGenImgFunc }()
	GenImgFunc = func(file io.Reader, ext string, numshapes int, mode primitive.Mode) (string, error) {
		return "", errors.New("error occurred")
	}
}

func TestImproperFile(test *testing.T) {
	req,_:= http.NewRequest("GET","/modify/emptydata.png?mode=ux",nil)	
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ModifyHandler)
	handler.ServeHTTP(res,req)
	outcome := 400
	if status := res.Code; status != outcome {
		fmt.Printf("handler returned wrong status code: got %v want %v\n",
			status, outcome)
	}
	CheckFunc(ModifyHandler,req,true, "img")
}


func TestGetExtfileErr(test *testing.T) {
	TempFileFunc := tempFile
	oldTempFileFunc := TempFileFunc

	defer func (){
		TempFileFunc = oldTempFileFunc
	}()
	TempFileFunc = func(a, b string) (*os.File, error) {
		return nil, errors.New("Mocked ExtFile")

	}
	tempFile("", "")
	f, _ := os.Open("z.png")
	_, err := generateImg(f, ".png", 5, primitive.TriangleMode)
	if err != nil {
		log.Println(err)
	}
	req,_:= http.NewRequest("POST","/upload",nil)	
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)
	handler.ServeHTTP(res,req)
	outcome := 400
	if status := res.Code; status != outcome {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, outcome)
	}
	CheckFunc(UploadHandler,req,true, "img")
	TempFileFunc = oldTempFileFunc
	outcome = 500
	if status := res.Code; status != outcome {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, outcome)
	}
	CheckFunc(UploadHandler,req,true, "img")

	_,err = generateImg(File,"txt",3,primitive.ComboMode)
	if err!=nil{
		log.Println(err)
	}


}


/*

func TestImproperFile(test *testing.T) {
	CheckLinks(ModifyHandler, "GET", "/modify/nodata.png?mode=2", "", 500, true, "img")
}
func TestGetExtfileErr(test *testing.T) {
	TempFileFunc := tempFile
	oldTempFileFunc := TempFileFunc

	TempFileFunc = func(a, b string) (*os.File, error) {
		return nil, errors.New("Mocked ExtFile")

	}
	tempFile("", "")
	f, _ := os.Open("samurai.png")
	_, err := generateImg(f, "png", 5, primitive.TriangleMode)
	if err != nil {
		log.Println(err)
	}
	CheckLinks(UploadHandler, "POST", "/upload", "", 200, true, "img")
	TempFileFunc = oldTempFileFunc
	
	CheckLinks(UploadHandler, "POST", "/upload", "", 500, true, "test")

	_,err = generateImg(File,"txt",3,primitive.ComboMode)
	if err!=nil{
		log.Println(err)
	}


}

func TestErrCopy(test *testing.T){
	CopyFunc := io.Copy
	oldCopyFunc := CopyFunc
	defer func(){
		CopyFunc = oldCopyFunc
	}()
	CopyFunc = func (w io.Writer,r io.Reader)(int64 , error){
		return 0,errors.New("Mocked Copy Function.")
	}
	f, _ := os.Open("samurai.png")
	_, err := generateImg(f, "png", 5, primitive.TriangleMode)
	if err != nil {
		log.Println(err)
	}

	CheckLinks(UploadHandler, "POST", "/upload", "", 200, true, "img")
	

}
*/