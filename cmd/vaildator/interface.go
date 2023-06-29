package vaildator

type Vaildator interface {
	// Vaildate action is runned by calling VailateString function
	VaildateString(input string) error
}
