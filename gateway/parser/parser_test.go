package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"code.sajari.com/docconv"
)

func TestParseDoc(t *testing.T) {
	f, err := os.Open("/Users/houmingyu/Documents/web5/filesearch/testfile/test.docx")
	if err != nil {
		t.Fatal(err)
	}
	data, _ := ioutil.ReadAll(f)
	f.Close()
	r := bytes.NewReader(data)
	s, err := ParseDoc(r, "test.docx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
}

func TestParsePDF(t *testing.T) {
	f, err := os.Open("/Users/houmingyu/Documents/web5/filesearch/testfile/test.pdf")
	if err != nil {
		t.Fatal(err)
	}
	data, _ := ioutil.ReadAll(f)
	f.Close()
	r := bytes.NewReader(data)
	s, err := ParseDoc(r, "test.pdf")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
}

func TestDocx(t *testing.T) {
	res, err := docconv.ConvertPath("../testfile/test.docx")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Body)
}

func TestPdf(t *testing.T) {
	res, err := docconv.ConvertPath("../testfile/test.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Body)
}
