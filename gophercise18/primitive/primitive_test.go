package primitive

import (
	"errors"
	"io"
	"log"
	"os"
	"testing"
)

func TestMode(t *testing.T){
	mode := CircleMode
	_ = ImgMode(mode)
}

func TestTempFile(t *testing.T){
	tempfile("","")
}

func TestTransformImg(t *testing.T){
	img, err := os.Open("cartoon.png")
	if err != nil {
		log.Println(err)
		
	}
	output, err := TransformImage(img, "png", 4, ImgMode(BeziersMode))
	if err != nil {
		log.Println(err)
		
	}
	of, _ := os.OpenFile("pic.png", os.O_RDWR|os.O_CREATE, 0755)
	io.Copy(of, output)

	_, err = TransformImage(img, "txt", 4, ImgMode(EllipseMode))
	if err != nil {
		log.Println(err)
		
	}
}

func TestCopyFunc(t *testing.T){
	file, _ := os.Open("test.txt")
	TransformImage(file,"png",4,ImgMode(PolygonMode))
	CopyFunc := io.Copy
	newCopy := CopyFunc
	defer func(){
		CopyFunc = newCopy
	}()

	CopyFunc = func (w io.Writer,r io.Reader)(int64 , error)  {
		return 0,errors.New("Mocked Copy Function.")
	}
	file,_ = os.Open("cartoon.png")
	TransformImage(file,"png",2,ImgMode(TriangleMode))
	
}