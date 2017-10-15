package invoicer

import (
	"bytes"
	"fmt"
	"github.com/aymerick/douceur/inliner"
	"github.com/joiggama/money"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func parseTemplateToString(templateName string, data interface{}) (string, error) {
	funcMap := template.FuncMap{
		"FormatNumber": func(value float64) string {
			return fmt.Sprintf("%.2f", value)
		},
		"FormatPercentage": func(valueType string, value string) string {
			if (valueType == "%") {
				return value + "%"
			}
			return valueType + value
		},
		"FormatMoney": func(currency string, value float64) string {
			return string(money.Format(value, money.Options{"currency": strings.ToUpper(currency)}))
		},
		"GetItemIndex": func(index int) int {
			return index + 1
		},
		"GetItemTotal": func(quantity float64, unitCost float64) float64 {
			return quantity * unitCost
		},
		"GetHTML": func(value string) template.HTML {
			return template.HTML(value)
		},
	}
	root, _ := os.Getwd()
	templatePaths := strings.Split(templateName, "/")
	t, err := template.New(templatePaths[len(templatePaths)-1]).Funcs(funcMap).ParseFiles(filepath.Join( root, templateName ))
	if (err != nil) {
		fmt.Println("ParseFiles", err)
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		fmt.Println("Execute", err)
		return "", err
	}
	parsedHtml := tpl.String()
	inlinedHtml, err := inliner.Inline(parsedHtml)
	if err != nil {
		fmt.Println("Inline", err)
		return "", err
	}
	return inlinedHtml, nil
}

func parseLineBreaks(content *string) {
	*content = strings.Replace(*content, "\n", "<br>", -1)
}

func parseHtmlContent(htmlContent *string) {
	*htmlContent = strings.Replace(*htmlContent, "&#34;", "'", -1)
}
