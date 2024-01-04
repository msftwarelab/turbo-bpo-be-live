package excel

import (
	"context"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
)

func CreateBilling(ctx context.Context, billing *datastore.Billing) (string, error) {
	f, err := excelize.OpenFile(config.AppConfig.GetString("billingExcelTemplate"))
	if err != nil {
		println(err.Error())
		return "", err
	}
	// F10 - invoice number
	// F11 - date
	// B11 - Client full Name
	// B12 - Client Address
	// B13 - City / State / ZipCode

	EntryCount := len(billing.Entries)
	clientInfo := *billing.ClientInfo(ctx)
	clientAddress2 := fmt.Sprintf("%s, %s , %s", *clientInfo.City, *clientInfo.State, *clientInfo.Zipcode)

	f.SetCellValue("Invoice", "F10", *billing.InvoiceNumber)
	f.SetCellValue("Invoice", "F11", billing.FormatedDate())
	f.SetCellValue("Invoice", "B11", *billing.UserName)
	f.SetCellValue("Invoice", "B12", *clientInfo.Address)
	f.SetCellValue("Invoice", "b13", clientAddress2)
	//f.SetCellValue("Invoice", "C20", "Helslo world.")

	cellStyleB2O := f.GetCellStyle("Invoice", "B20")
	cellStyleG2O := f.GetCellStyle("Invoice", "G20")

	for i := 1; i < EntryCount; i++ {
		f.InsertRow("Invoice", 20)
		f.SetCellStyle("Invoice", "B21", "F21", cellStyleB2O)
		f.SetCellStyle("Invoice", "G21", "G21", cellStyleG2O)
		f.MergeCell("Invoice", "C21", "F21")
	}
	numberingStarterLocation := 20
	var totalAmount float64
	for i := 0; i < EntryCount; i++ {

		f.SetCellInt("Invoice", concatCel("B", numberingStarterLocation+i), i+1)
		//fill description
		description := *billing.Entries[i].Description
		f.SetCellValue("Invoice", concatCel("C", numberingStarterLocation+i), description)
		//fill amount
		amount := *billing.Entries[i].Amount
		totalAmount += amount
		f.SetCellValue("Invoice", concatCel("G", numberingStarterLocation+i), amount)
	}

	//render total amount
	f.SetCellValue("Invoice", concatCel("G", numberingStarterLocation+EntryCount), totalAmount)

	fileLocation := fmt.Sprintf("%s/%s%s", config.AppConfig.GetString("tmpDir"), *billing.InvoiceNumber, ".xlsx")
	if err := f.SaveAs(fileLocation); err != nil {
		return "", err
	}
	return fileLocation, nil

}

func concatCel(a string, b int) string {
	return fmt.Sprintf("%s%v", a, b)
}
