package tester

import (
	"fmt"
	"regexp"
)

func TestPaymentMethod(htmlContent string) {
	paymentMethodPattern := `<input[^>]*(type=["']?(checkbox|radio)["']?)[^>]*(value=["']?(CIB|EDAHABIA)["']?)[^>]*>`

	re := regexp.MustCompile(paymentMethodPattern)
	if re.MatchString(htmlContent) {
		fmt.Println("Checkbox or radio button with CIB/EDAHABIA as a payment method is present.")
	} else {
		fmt.Println("Checkbox or radio button with CIB/EDAHABIA as a payment method is not present.")
	}
}
