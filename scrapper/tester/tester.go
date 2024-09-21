package tester

import (
	"errors"
	"regexp"
)

type Config struct {
	GreenIconLink     string
	TermsCheckboxName string
	PaymentMethod     string
}

type Tester struct {
	Content string
	Config  Config
}

func (t *Tester) TestNumber() error {
	imgTagPattern := `<img[^>]*\ssrc=["']?(` + regexp.QuoteMeta(t.Config.GreenIconLink) + `)["']?[^>]*>`

	re := regexp.MustCompile(imgTagPattern)
	if !re.MatchString(t.Content) {
		return errors.New("Image tag is not present in the HTML content.")
	}
	return nil
}

func (t *Tester) TestPaymentMethod() error {
	paymentMethod := regexp.QuoteMeta(t.Config.PaymentMethod)
	paymentMethodPattern := `(?i)<input[^>]*(type=["']?(checkbox|radio)["']?)[^>]*>\s*<label[^>]*>(.*?)` + paymentMethod + `(.*?)<\/label>`

	re := regexp.MustCompile(paymentMethodPattern)
	if !re.MatchString(t.Content) {
		return errors.New("Checkbox or radio button with specified payment method is not present.")
	}
	return nil
}

func (t *Tester) TestTermsAndConditions() error {
	termsPattern := `(?i)<label[^>]*>(I accept.*?|I agree.*?|J'accepte.*?)<\/label>`

	re := regexp.MustCompile(termsPattern)
	if !re.MatchString(t.Content) {
		return errors.New("Label for 'I accept', 'I agree', or 'J'accepte' is not present.")
	}
	return nil
}
