package invoicer

func CreateInvoices(invoices *[]Invoice) {
	for i, _ := range *invoices {
		parseLineBreaks(&(*invoices)[i].FromInfo)
		parseLineBreaks(&(*invoices)[i].ToInfo)
		parseLineBreaks(&(*invoices)[i].Notes)
		parseLineBreaks(&(*invoices)[i].Terms)
		parseLineBreaks(&(*invoices)[i].EmailInfos.Message)
		generateAutoCalculations(&(*invoices)[i])
	}
}

func SendInvoice(invoice *Invoice) {
	parsedHtmlInvoice, err := parseTemplateToString("$INVOICER_TEMPLATES_PATH/template_invoice_01.html", &invoice)
	if (err == nil) {
		pdfFileContent := getPDFContent(&parsedHtmlInvoice)
		attachmentName := "invoice.pdf"
		if (invoice.EmailInfos.AttachmentName != "") {
			attachmentName = invoice.EmailInfos.AttachmentName + ".pdf"
		}
		attachment := CreateAttachment(&pdfFileContent, &attachmentName, "application/pdf")
		_ = SendEmail(&invoice.EmailInfos, &*attachment)
	}
}

func generateAutoCalculations(invoice *Invoice) {
	for _, item := range invoice.Items {
		invoice.Calculations.Subtotal += item.Quantity * item.UnitCost
	}
	invoice.Calculations.Total = invoice.Calculations.Subtotal
	calculateSubTotal(&*invoice, invoice.Tax, invoice.Fields.Tax, &invoice.Calculations.Tax, false)
	calculateSubTotal(&*invoice, invoice.Discount, invoice.Fields.Discount, &invoice.Calculations.Discount, true)
	calculateSubTotal(&*invoice, invoice.Shipping, invoice.Fields.Shipping, &invoice.Calculations.Shipping, false)
	invoice.Calculations.Balance = invoice.Calculations.Total - invoice.AmountPaid
}

func calculateSubTotal(invoice *Invoice, value float64, valueType string, calculation *float64, isDiscount bool) {
	if (value > 0) {
		calc := &invoice.Calculations
		*calculation = value
		if (valueType == "%") {
			*calculation = calc.Total * (value / 100)
		}
		if (isDiscount) {
			*calculation = *calculation * -1
		}
		calc.Total = calc.Total + *calculation
	}
}
