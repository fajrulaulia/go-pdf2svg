package svgpdf

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

type App struct{}

func Exporter(source string, newFile string) error {
	var c *App
	if !c.isValidFile(source) {
		return errors.New("Invalid Source File")
	}
	if len(newFile) < 1 {
		return errors.New("Invalid New File")
	}
	c.command(`inkscape --file="` + source + `" --export-plain-svg="` + newFile + `.svg" --export-margin=100 --export-margin=200 --export-margin=300`)
	return nil
}

func (c *App) command(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(stdout), nil
}

func (c *App) isValidFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
