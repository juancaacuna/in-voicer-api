package invoicer

type Invoice struct {
	Header             string `json:"header,omitempty"`
	ToTitle            string `json:"to_title,omitempty"`
	InvoiceNumberTitle string `json:"invoice_number_title,omitempty"`
	DateTitle          string `json:"date_title,omitempty"`
	PaymentTermsTitle  string `json:"payment_terms_title,omitempty"`
	CurrencyTitle      string `json:"currency_title,omitempty"`
	DueDateTitle       string `json:"due_date_title,omitempty"`
	PurchaseOrderTitle string `json:"purchase_order_title,omitempty"`
	// Items
	QuantityHeader     string `json:"quantity_header,omitempty"`
	ItemHeader         string `json:"item_header,omitempty"`
	UnitCostHeader     string `json:"unit_cost_header,omitempty"`
	AmountHeader       string `json:"amount_header,omitempty"`
	
	SubtotalTitle      string `json:"subtotal_title,omitempty"`
	DiscountsTitle     string `json:"discounts_title,omitempty"`
	TaxTitle           string `json:"tax_title,omitempty"`
	ShippingTitle      string `json:"shipping_title,omitempty"`
	TotalTitle         string `json:"total_title,omitempty"`
	AmountPaidTitle    string `json:"amount_paid_title,omitempty"`
	BalanceTitle       string `json:"balance_title,omitempty"`
	TermsTitle         string `json:"terms_title,omitempty"`
	NotesTitle         string `json:"notes_title,omitempty"`

	Currency 			string 		`json:"currency,omitempty"`

	From          string  	`json:"from"`
	FromInfo      string  	`json:"from_info"`
	To            string  	`json:"to"`
	ToInfo      	string  	`json:"to_info"`
	Logo          string  	`json:"logo,omitempty"`
	Number        string  	`json:"number,omitempty"`
	PurchaseOrder string  	`json:"purchase_order,omitempty"`
	Date          string  	`json:"date"`
	DueDate       string  	`json:"due_date,omitempty"`
	PaymentTerms  string  	`json:"payment_terms,omitempty"`
	Items         []Item  	`json:"items"`
	Fields        Field   	`json:"fields,omitempty"`
	Discount      float64 	`json:"discount,omitempty"`
	Tax           float64 	`json:"tax,omitempty"`
	Shipping      float64 	`json:"shipping,omitempty"`
	AmountPaid    float64 	`json:"amount_paid,omitempty"`
	Notes         string  	`json:"notes,omitempty"`
	Terms         string  	`json:"terms,omitempty"`
	EmailInfos   	EmailInfo `json:"email_info"`
	Calculations	Calculation
}

type Item struct {
	Name        string  `json:"name"`
	Quantity    float64 `json:"quantity"`
	UnitCost   float64  `json:"unit_cost"`
}

type Field struct {
	Tax       string  `json:"tax"`
	Discount  string	`json:"discount"`
	Shipping  string 	`json:"shipping"`
}

type EmailInfo struct {
  FromName        string    `json:"from_name"`
	FromEmail				string    `json:"email_from"`
  ToEmails				[]string  `json:"emails_to"`
  Subject 				string    `json:"subject"`
	Message 				string    `json:"message,omitempty"`
	AttachmentName 	string    `json:"attachment_name,omitempty"`
}

type Calculation struct {
	Subtotal	float64
	Discount	float64
	Tax				float64
	Shipping 	float64
	Total			float64
	Balance		float64
}

/*
[{
   "header":"FACTURA",
   "to_title":"Facturar a:",
   "date_title":"Fecha:",
   "payment_terms_title":"Términos de pago:",
   "due_date_title":"Fecha de expiración:",
   "currency_title":"Moneda:",
   "purchase_order_title":"Purchase Order",
   "item_header":"Descripción",
   "quantity_header":"Cantidad",
   "unit_cost_header":"Costo Unitario",
   "amount_header":"Subtotal",
   "subtotal_title":"Subtotal:",
   "discounts_title":"Descuento:",
   "tax_title":"Impuesto:",
   "shipping_title":"Envío:",
   "total_title":"Total:",
   "amount_paid_title":"Monto pagado:",
   "balance_title":"Balance:",
   "terms_title":"Términos",
   "notes_title":"Notas",
   "logo":null,
	 "from":"Importadora TresVeinte S.A.",
	 "from_info":"Ced. 1-3455-6432\nTel: 2344-5621",
	 "to":"Almacén Don Julio S.A.",
	 "to_info":"Ced. 1-4355-5437\nTel: 4567-3221",
   "number":"242351",
   "date":"Oct 11, 2017",
   "payment_terms":"A 3 meses",
   "due_date":"Oct 25, 2017",
   "items":[
      {
         "quantity":10,
         "name":"Cocina eléctrica",
         "unit_cost":299.75
      },
      {
         "quantity":40,
         "name":"Plancha",
         "unit_cost":30.50
      },
      {
         "quantity":8,
         "name":"Coffee Maker",
         "unit_cost":33.99
      },
      {
         "quantity":10,
         "name":"Tostadora",
         "unit_cost":45.25
      },
      {
         "quantity":35,
         "name":"TV 40'",
         "unit_cost":400.65
      },
      {
         "quantity":14,
         "name":"Platilla eléctrica",
         "unit_cost":400.99
      },
      {
         "quantity":3,
         "name":"Extractor",
         "unit_cost":200.25
      },
      {
         "quantity":11,
         "name":"TV 65'",
         "unit_cost":700.99
      },
      {
         "quantity":15,
         "name":"Secadora platos",
         "unit_cost":300.25
      }
   ],
   "currency":"USD",
   "fields":{
      "discount":"%",
      "tax":"%",
      "shipping":"$"
   },
   "discount":10,
   "tax":13.3,
   "shipping":500,
   "amount_paid":5000,
   "notes":"Favor indicarnos cada cuánto hará una orden de compra. Gracias.",
   "terms":"Factura sujeta a legislación costarricense.",
   "email_info":{
      "from_name":"Importadora 3-20",
   		"emails_to":["jcarlos323@hotmail.com"],
  	 	"email_from":"info@importadora320.com",
   		"message":"Buenos días,\n\nAdjunto podrá encontrar la facturación mensual. Favor realizar pago antes del 15 del mes.\n\nGracias,\nImportadora 3-20",
   		"attachment_name":"factura"
   }
}]

*/