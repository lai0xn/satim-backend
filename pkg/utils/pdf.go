package utils

import (
	"log"
	"strings"
	"yugioh-decks/pkg/types"

	"github.com/signintech/gopdf"
)

func wrapText(pdf *gopdf.GoPdf, text string, maxWidth float64) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		testLine := currentLine + word + " "
		testWidth, err := pdf.MeasureTextWidth(testLine)
		if err != nil {
			log.Fatalf("Error measuring text width: %s", err)
		}
		if testWidth <= maxWidth {
			currentLine = testLine
		} else {
			lines = append(lines, currentLine)
			currentLine = word + " "
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

func Genratepdf(checks types.CheckList) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595, H: 842}})
	pdf.AddPage()

	err := pdf.AddTTFFont("Arial", "./Arial.ttf")
	if err != nil {
		log.Fatalf("Error adding font: %s", err)
	}

	pdf.SetFont("Arial", "", 24)
	title := "Sample PDF Title"
	titleWidth, err := pdf.MeasureTextWidth(title)
	if err != nil {
		log.Fatalf("Error measuring text width: %s", err)
	}
	pdf.SetX((595 - titleWidth) / 2)
	pdf.SetY(40)
	pdf.Cell(nil, title)

	pdf.Br(200)

	pdf.SetFont("Arial", "", 14)
	description := "This is a sample description for the PDF. It provides context and details about the content. The description may be long and need to be wrapped accordingly to fit within ."
	pdf.SetX(20)
	pdf.SetY(100)

	lines := wrapText(&pdf, description, 575)
	for _, line := range lines {
		pdf.SetX(20)
		pdf.Cell(nil, line)
		pdf.Br(20)
	}

	pdf.Br(40)

	pdf.SetFont("Arial", "", 14)
	tasks := []string{
		"1. Check SSL certificate & expiration date. []",
		"2. check Logos of Banks in website And green number []",
		"3. Check Captcha in website []",
		"4. Check the request from the website to the server using proxy []",
	}

	pdf.SetX(40)
	pdf.Cell(nil, "Tasks:")
	pdf.Br(30)

	for _, task := range tasks {
		lines := wrapText(&pdf, task, 575)
		for _, line := range lines {
			pdf.SetX(60)
			pdf.Cell(nil, line)
			pdf.Br(20)
		}
	}

	pdf.SetX(40)

	pdf.Cell(nil, "Dev Badge:")
	pdf.Br(30)

	for _, task := range tasks {
		lines := wrapText(&pdf, task, 575)
		for _, line := range lines {
			pdf.SetX(60)
			pdf.Cell(nil, line)
			pdf.Br(20)
		}
	}

	pdf.SetStrokeColor(232, 48, 56)
	pdf.SetLineWidth(2)
	pdf.RectFromUpperLeft(10, 10, 575, 822)

	err = pdf.WritePdf("output.pdf")
	if err != nil {
		log.Fatalf("Error writing PDF: %s", err)
	}

	log.Println("PDF generated successfully!")
}
