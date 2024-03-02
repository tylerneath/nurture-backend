package common

func Xor(a, b bool) bool {
	return (a || b) && !(a && b)
}
