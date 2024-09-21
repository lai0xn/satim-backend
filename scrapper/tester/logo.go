package tester

import (
	"errors"
	"regexp"
)

var logo_image = "satim.cdn.dz/satim_logo.svg"

func TestLogo(htmlContent string) error {
	imgTagPattern := `<img[^>]*\ssrc=["']?(` + regexp.QuoteMeta(logo_image) + `)["']?[^>]*>`

	re := regexp.MustCompile(imgTagPattern)
	if !re.MatchString(htmlContent) {
		return errors.New("Didn't find the green number image in the ")
	}
	return nil
}
