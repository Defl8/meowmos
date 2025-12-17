package utils

func MenuWrap(min, max int) int {
	return (min%max + max) % max
}
