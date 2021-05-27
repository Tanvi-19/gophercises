package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

 
func main(){
	rout := http.NewServeMux()
	rout.HandleFunc("/panic/",HandleException)
	rout.HandleFunc("/debug/",DebugCode)
	
	fmt.Println("Server is running on 8000 port....")
	log.Fatal(http.ListenAndServe(":8000",RecoverFunc(rout)))
}


func DebugCode(res http.ResponseWriter, req *http.Request){
	Fpath := req.FormValue("path")
	lineStr := req.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	fmt.Println("Visited line -",lineStr)
	if err != nil {
		line = -1
	}	
	
	file, err := os.Open(Fpath)
	if err != nil {
		http.Error(res,err.Error(),http.StatusInternalServerError)
		return
	}
	byteData := bytes.NewBuffer(nil)
	io.Copy(byteData,file)
	
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}

	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, byteData.String())
	
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(true), html.LineNumbersInTable(true), html.HighlightLines(lines))

	res.Header().Set("Content-Type", "text/html")
	formatter.Format(res, styles.GitHub, iterator)
	if err != nil{
		http.Error(res,err.Error(),http.StatusInternalServerError)
	}
	
} 


func HandleException(res http.ResponseWriter, req *http.Request){
	panicMsg()

	
//	_ = quick.Highlight(res,byteData.String(),"go","html","monokai")
}

func RecoverFunc(app http.Handler) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request){
		defer func(){
			if err := recover(); err != nil{
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				fmt.Fprintf(res,"<h1>panic : %v</h1><pre>%s</pre><pre>%d</pre>",err,links(string(stack)),http.StatusInternalServerError)
			}
		}()
		
		app.ServeHTTP(res,req)
	}
}

func links(stack string) string{
	lines := strings.Split(stack,"\n")
	for ind, line := range lines{
		if len(line) == 0 || line[0] !='\t'{
			continue
		}
		fileL :=""
		for j , ch :=range line{
			if ch == ':'{
				fileL = line[1:j]
				break
			}
		}
		var lineNum strings.Builder
		for i := len(fileL) + 2; i < len(line); i++{
			if line[i] < '0' || line[i] > '9'{
				break
			}
			lineNum.WriteByte(line[i]) 
		}
		V := url.Values{}
		V.Set("path",fileL)
		V.Set("line",lineNum.String())
		lines[ind] = "\t<a href=\"/debug?" + V.Encode() + "\">" + fileL + ":" + lineNum.String() + "</a>" + line[len(fileL)+2+len(lineNum.String()):]
	}
	return strings.Join(lines,"\n")
}
func panicMsg(){
	panic("Oh no!! Some panic error occurred")
}	