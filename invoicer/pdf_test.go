package invoicer

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestPDF(t *testing.T) {
	expected := 5561
	template := "<html><body><div>Hello in-voicer!</div></body></html>"
	result := getPDFContent(&template)
	assert.True(t, len(result) == expected, "Should create a new PDF file, and return the bytes array.")
	
	expected = 1272
	template = ""
	result = getPDFContent(&template)
	assert.True(t, len(result) == expected, "Should work with empty template.")
	
	expected = 6433
	template = "This is not an HTML template"
	result = getPDFContent(&template)
  assert.True(t, len(result) == expected, "Should work with non-template.")
}