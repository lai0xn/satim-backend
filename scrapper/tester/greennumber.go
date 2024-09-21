package tester

import (
	"errors"
	"regexp"
)

var green_icon_link = "satim.cdn.dz/green_number.svg"

func TestNumber(htmlContent string) error {
	imgTagPattern := `<img[^>]*\ssrc=["']?(` + regexp.QuoteMeta(green_icon_link) + `)["']?[^>]*>`

	re := regexp.MustCompile(imgTagPattern)
	if !re.MatchString(htmlContent) {
		return errors.New("Didn't find the green number image in the ")
	}
	return nil
}
