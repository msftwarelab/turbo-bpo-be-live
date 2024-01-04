package pdf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/template"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	str "github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func Create(input *models.Iform) (*string, error) {

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	htmlStr, err := parseHtml(input)
	if err != nil {
		fmt.Println("error %v", err)
		return nil, err
	}
	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(*htmlStr)))
	//pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))
	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	//Your Pdf Name
	cdt := fmt.Sprintf("%s", utils.Int64ToStr(millis.NowInMillis()))
	addressOnFilename := ""
	if input.TxtSubjectAddress != nil {
		addressOnFilename = *input.TxtSubjectAddress
	}
	filePath := fmt.Sprintf("/tmp/turboTmpData/%s%s.pdf", cdt, addressOnFilename)
	err = pdfg.WriteFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Done")
	return str.ToObject(filePath), nil

}

func parseHtml(input *models.Iform) (*string, error) {

	//TODO, refactor to use dynamic file location
	tpl, err := template.New("page1.html").ParseFiles(config.AppConfig.GetString("pdfTemplateDirPage1"))
	var strData bytes.Buffer

	err = tpl.Execute(&strData, *input)
	if err != nil {
		return nil, err
	}
	return ToObject(strData.String()), nil
}

func ToObject(i string) *string {
	return &i
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
