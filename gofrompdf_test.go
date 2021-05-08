package svgpdf

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	var app App
	str, err := app.command("apt | grep apt  | head -n 1 -c 3")
	assert.Equal(t, err, nil, "Expected not error")
	assert.Equal(t, str, "apt")
}
func TestIsValidFile(t *testing.T) {
	var app App
	err := ioutil.WriteFile("kucing.pdf", []byte("Hello"), 0755)
	assert.Equal(t, err, nil, "Expected not error")
	result := app.isValidFile("kucing.pdf")
	assert.Equal(t, result, true, "expect true")
	err = os.Remove("kucing.pdf")
	assert.Equal(t, err, nil, "expect true")
	result = app.isValidFile("kucing.pdf")
	assert.Equal(t, result, false, "expect false")
}

func TestCreateFile(t *testing.T) {
	err := Exporter("contoh.pdf", "newSvgLetter")
	assert.Equal(t, err, nil, "Expected not error")
	var app App
	result := app.isValidFile("newSvgLetter.svg")
	assert.Equal(t, result, true, "expect false")
	err = os.Remove("newSvgLetter.svg")
	assert.Equal(t, err, nil, "expect true")
}
