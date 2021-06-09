package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Mode int

const (
	ComboMode Mode = iota
	TriangleMode
	RectangleMode
	EllipseMode
	CircleMode
	RotatedrectMode
	BeziersMode
	RotatedellipseMode
	PolygonMode
)


func ImgMode(mode Mode) func() []string {

	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

func tempfile(prefix, suffix string) (*os.File, error) {
	infile, _ := ioutil.TempFile("", prefix)
	defer os.Remove(infile.Name())
	fileName := fmt.Sprintf("%s.%s", infile.Name(), suffix)
	return os.Create(fileName)
}


func TransformImage(imgreadr io.Reader, ext string, numOfShapes int, options ...func() []string) (io.Reader, error) {
	var args []string
	for _, option := range options {
		args = append(args, option()...)
	}
	transformreader := bytes.NewBuffer(nil)
	infile, _ := tempfile("inp_", ext)
	
	defer os.Remove(infile.Name())

	opfile, _ := tempfile("op_", ext)
	

	defer os.Remove(opfile.Name())

	_, err := io.Copy(infile, imgreadr)
	if err != nil {
		return nil, err 
	}
	op, err := primitive(infile.Name(), opfile.Name(), numOfShapes, args...)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(op))
	io.Copy(opfile, transformreader)
	
	return opfile, nil
}

func primitive(input, output string, numShapes int, args ...string) ([]byte, error) {

	CmdString := fmt.Sprintf("-i %s -o %s -n %d ", input, output, numShapes)
	args = append(strings.Fields(CmdString), args...)
	cmd := exec.Command("primitive", args...)
	return cmd.CombinedOutput()
}
