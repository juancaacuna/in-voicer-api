package invoicer

import (
	"github.com/stretchr/testify/assert"
	"html/template"
  "testing"
)

func TestParseTemplateToString(t *testing.T) {
	data := struct{
    FromName	string
		Message		template.HTML
	}{
		"Juan Carlos Acuña",
		template.HTML("Hello in-voicer!"),
	}
	result, _ := parseTemplateToString("/templates/template_email.html", data)
	expected := 2000
	assert.True(t, len(result) > expected, "Should parse a template file, and return a string with the parsed content")

	result, _ = parseTemplateToString("/templates/template_email.html", nil)
	assert.True(t, len(result) >= expected, "Should not failed with empty data.")

	result, _ = parseTemplateToString("", data)
	expected = 0
	assert.True(t, len(result) == expected, "Should not failed with empty or invalid template.")
}

func TestParseLineBreaks(t *testing.T) {
	expected := "Hello<br>in-voicer!<br>Have a nice day"
	result := "Hello\nin-voicer!\nHave a nice day"
	parseLineBreaks(&result)
	assert.True(t, result == expected, "Should parse line breaks and replace them with <br>.")

	expected = "<br><br>"
	result = "\n\n"
	parseLineBreaks(&result)
	assert.True(t, result == expected, "Should parse line breaks and replace them with <br>.")

	expected = "Hello in-voicer"
	result = "Hello in-voicer"
	parseLineBreaks(&result)
	assert.True(t, result == expected, "Should parse string with no line breaks and return same string.")

	expected = ""
	result = ""
	parseLineBreaks(&result)
	assert.True(t, result == expected, "Should parse an empty string and return same empty string.")
}

func TestParseHtmlContent(t *testing.T) {
	expected := "<div style=\"font: 'Open Sans'\">Hello in-voicer!</div>"
	result := "<div style=\"font: &#34;Open Sans&#34;\">Hello in-voicer!</div>"
	parseHtmlContent(&result)
	assert.True(t, result == expected, "Should parse string and replace &#34; with 'single quotes'.")
	
	expected = ""
	result = ""
	parseHtmlContent(&result)
	assert.True(t, result == expected, "Should parse an empty string and return same empty string.")
}