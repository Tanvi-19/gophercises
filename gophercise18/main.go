package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/gophercises/gophercise18/primitive"
)

func main() {
	Rout := http.NewServeMux()

	Rout.HandleFunc("/", RoutHandler)

	Rout.HandleFunc("/upload", UploadHandler)

	Rout.HandleFunc("/modify/", ModifyHandler)

	fs := http.FileServer(http.Dir("./photos"))
	Rout.Handle("/photos/", http.StripPrefix("/photos", fs))
	fmt.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", Rout))

}


func RoutHandler(res http.ResponseWriter, req *http.Request) {
	html := `<html>
                         <h3>UPLOAD IMAGE</h3>
                         <form action="/upload" method="POST" enctype="multipart/form-data">
                         <input type="file" name ="img"/>
                         <button type="submit">UPLOAD</button>
                         </form>
                         </html>`

	fmt.Fprint(res, html)
}
func UploadHandler(res http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("img")
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()	
	ext := filepath.Ext(header.Filename)[1:]
	output, _ := tempFile("", ext)
	
	io.Copy(output, file)

	http.Redirect(res, req, "/modify/"+filepath.Base(output.Name()), http.StatusFound)
}
func ModifyHandler(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./photos/" + filepath.Base(req.URL.Path))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	rawmode := req.FormValue("mode")
	ext := filepath.Ext(file.Name())[1:]
	if rawmode == "" {
		
		ModeChoice(res, req, ext, file)
		return
	}
	mode, err := strconv.Atoi(rawmode)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	numStr := req.FormValue("n")
	if numStr == "" {
		NumChoice(res, req, ext, primitive.Mode(mode), file)
		return
	}
	strconv.Atoi(numStr) //temperorily
	
	res.Header().Set("Content-Type", "image/png")
	http.Redirect(res, req, "/photos/"+filepath.Base(file.Name()), http.StatusFound)

}


type OptStruct struct {
	num  int
	mode primitive.Mode
}

func generateImageList(ext string, fileSeeker io.ReadSeeker, opts ...OptStruct) ([]string, error) {
	opFileList := []string{}
	for _, value := range opts {
		fileSeeker.Seek(0, 0)
		opFileName, _ := generateImg(fileSeeker, ext, value.num, value.mode)
		opFileList = append(opFileList, opFileName)
	}
	return opFileList, nil
}

func generateImg(file io.Reader, ext string, numshapes int, mode primitive.Mode) (string, error) {
	output, err := primitive.TransformImage(file, ext, numshapes, primitive.ImgMode(mode))
	if err != nil {
		return "", err
	}

	saveout, _ := tempFile("", ext)

    io.Copy(saveout, output)       

	return saveout.Name(), err
}

func NumChoice(res http.ResponseWriter, req *http.Request, ext string, mode primitive.Mode, fileSeeker io.ReadSeeker) {

	op := []OptStruct{
		{20, mode},
		{30, mode},
		{40, mode},
		{50, mode},
	}
	opFileList, _ := generateImageList(ext, fileSeeker, op...)
	
	htmlist := `<html>
                    <body>
						{{range .}}
						<a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.Numshapes}}">
						<img style ="width 30%" src="/photos/{{.Name}}">
						{{end}}
                	</body>
                </html>
        `
	temp := template.Must(template.New("").Parse(htmlist))

	type Opts struct {
		Name      string
		Mode      primitive.Mode
		Numshapes int
	}
	var opts []Opts
	for index, val := range opFileList {
		opts = append(opts, Opts{Name: filepath.Base(val), Mode: op[index].mode, Numshapes: op[index].num})
	}

	temp.Execute(res, opts)
}

func ModeChoice(res http.ResponseWriter, req *http.Request, ext string, fileSeeker io.ReadSeeker) {
	op := []OptStruct{
		{22, primitive.TriangleMode},
		{22, primitive.CircleMode},
		{22, primitive.ComboMode},
		{22, primitive.PolygonMode},
	}
	opFileList, _ := generateImageList(ext, fileSeeker, op...)
	
	htmlist := `<html>
                    <body>
                		{{range .}}
                		<a href="/modify/{{.Name}}?mode={{.Mode}}">
                		<img style ="width 30%" src="/photos/{{.Name}}">
                		{{end}}
                	</body>
                </html>
        `
	temp := template.Must(template.New("").Parse(htmlist))
        
type Opts struct {
        Name string
        Mode primitive.Mode
}
	var opts []Opts
	for index, val := range opFileList {
		opts = append(opts, Opts{Name: filepath.Base(val), Mode: op[index].mode})
	}

	temp.Execute(res, opts)
        
}


func tempFile(prefix, suffix string) (*os.File, error) {
	infile, _ := ioutil.TempFile("./photos/", prefix)
	defer os.Remove(infile.Name())
	fileName := fmt.Sprintf("%s.%s", infile.Name(), suffix)
	fmt.Println(fileName)
	return os.Create(fileName)
}
