package util

const (
	NGN = "NGN"
	USD = "USD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case NGN, USD:
		return true
	}
	return false
}
