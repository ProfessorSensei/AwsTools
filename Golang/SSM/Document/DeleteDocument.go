package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var DocName, DocReg string

func main() {
	// user input Document name
	DocNreader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter Document name to be deleted:")
	DocName, _ = DocNreader.ReadString('\n')
	// remove new line
	DocName = strings.Replace(DocName, "\n", "", -1)

	// user input document region
	DocRegreader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter region the document is located in:")
	DocReg, _ = DocRegreader.ReadString('\n')
	// remove new line
	DocReg = strings.Replace(DocReg, "\n", "", -1)
	// ssm session
	// session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(DocReg)},
	)
	svc := ssm.New(sess)
	// use for document content
	doc := &ssm.DeleteDocumentInput{
		Name: &DocName,
	}
	// Create Document
	DocResult, err := svc.DeleteDocument(doc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Document '%v' deleted", DocName)
	fmt.Println(DocResult)

}
