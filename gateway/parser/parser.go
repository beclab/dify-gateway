package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"code.sajari.com/docconv"
)

var ParseAble = map[string]bool{
	".doc":      true,
	".docx":     true,
	".pdf":      true,
	".txt":      true,
	".md":       true,
	".markdown": true,
}

//func IsParseAble(filename string) bool {
//	fileType := GetTypeFromName(filename)
//	_, ok := ParseAble[fileType]
//	return ok
//}

func GetTypeFromName(filename string) string {
	return strings.ToLower(path.Ext(filename))
}

//func outputPdfText(inputPath string) (data string, err error) {
//	f, err := os.Open(inputPath)
//	if err != nil {
//		return "", err
//	}
//
//	defer f.Close()
//
//	pdfReader, err := model.NewPdfReader(f)
//	if err != nil {
//		return "", err
//	}
//
//	numPages, err := pdfReader.GetNumPages()
//	if err != nil {
//		return "", err
//	}
//
//	fmt.Printf("--------------------\n")
//	fmt.Printf("PDF to text extraction:\n")
//	fmt.Printf("--------------------\n")
//	data = ""
//	for i := 0; i < numPages; i++ {
//		pageNum := i + 1
//
//		page, err := pdfReader.GetPage(pageNum)
//		if err != nil {
//			return "", err
//		}
//
//		ex, err := extractor.New(page)
//		if err != nil {
//			return "", err
//		}
//
//		text, err := ex.ExtractText()
//		if err != nil {
//			return "", err
//		}
//
//		data += text
//		//fmt.Println("------------------------------")
//		//fmt.Printf("Page %d:\n", pageNum)
//		//fmt.Printf("\"%s\"\n", text)
//		//fmt.Println("------------------------------")
//	}
//
//	return data, nil
//}
//
//func outputDocText(filename string) (data string, err error) {
//	//docFile, err := os.Open(filename)
//	//if err != nil {
//	//	return "", err
//	//}
//	//defer docFile.Close()
//
//	docData, err := document.Open(filename)
//	if err != nil {
//		return "", err
//	}
//
//	data = ""
//	for _, para := range docData.Paragraphs() {
//		for _, run := range para.Runs() {
//			fmt.Print(run.Text())
//			data += run.Text()
//		}
//		data += "\n"
//		fmt.Println()
//	}
//	return data, nil
//}

func ParseDoc(f io.Reader, filename string) (string, error) {
	fileType := GetTypeFromName(filename)
	if _, ok := ParseAble[fileType]; !ok {
		return "", nil
	}
	if fileType == ".txt" || fileType == ".md" || fileType == ".markdown" {
		data, err := ioutil.ReadAll(f)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	//if fileType == ".pdf" {
	//	data, err := outputPdfText(filename)
	//	if err != nil {
	//		return "", err
	//	}
	//	return data, nil
	//}
	//if fileType == ".doc" {
	//	data, err := outputDocText(filename)
	//	if err != nil {
	//		return "", err
	//	}
	//	return data, nil
	//}
	// parsing pdf seems to always fail below...
	mimeType := docconv.MimeTypeByExtension(filename)
	fmt.Println(mimeType)
	fmt.Println(filename)
	res, err := docconv.Convert(f, mimeType, true)
	if err != nil {
		return "", err
	}
	return res.Body, nil
}
