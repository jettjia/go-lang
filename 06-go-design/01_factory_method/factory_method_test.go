package factory_method

import "testing"

func Test_ReadBook(t *testing.T) {
	NewBookReader(IPAD).ReadBook()
	NewBookReader(COMPUTER).ReadBook()
}
