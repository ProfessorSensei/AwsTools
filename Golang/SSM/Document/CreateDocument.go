package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var tpl *template.Template
var DocName, DocReg, DocMesg string

func init() {
	tpl = template.Must(template.ParseFiles("./htmlfolder/SSMDocument.gohtml"))
}

func main() {
	// user input Document name
	DocNreader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter Document name:")
	DocName, _ = DocNreader.ReadString('\n')
	// remove new line
	DocName = strings.Replace(DocName, "\n", "", -1)

	// user input document region
	DocRegreader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter region for document to be created in:")
	DocReg, _ = DocRegreader.ReadString('\n')
	// remove new line
	DocReg = strings.Replace(DocReg, "\n", "", -1)

	// user input message to go on instances
	DocMsgreader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter message to go on instances:")
	DocMesg, _ = DocMsgreader.ReadString('\n')
	// remove new line
	DocMesg = strings.Replace(DocMesg, "\n", "", -1)

	var test bytes.Buffer
	err := tpl.ExecuteTemplate(&test, "SSMDocument.gohtml", DocMesg)
	if err != nil {
		log.Fatalln(err)
	}
	// read stdout output as string
	result := test.String()
	fmt.Println(result)
	fmt.Printf("%T\n", result)
	DocFormat := "YAML"
	DocType := "Command"
	TargType := "/"
	// ssm session
	// session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(DocReg)},
	)
	svc := ssm.New(sess)
	// use for document content
	doc := &ssm.CreateDocumentInput{
		Content:        &result,
		DocumentFormat: &DocFormat,
		DocumentType:   &DocType,
		Name:           &DocName,
		TargetType:     &TargType,
	}
	// Create Document
	DocResult, err := svc.CreateDocument(doc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(DocResult)
}
