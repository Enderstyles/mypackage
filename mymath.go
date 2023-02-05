package mypackage

func MaxFloat(a, b float64) float64 {
	if a > b {return a}
	return b
}
func MaxInt(a, b int) int{
	if a > b {return a}
	return b
}
func MinFloat(a, b float64) float64 {
	if a > b {return b}
	return a
}
func MinInt(a, b int) int{
	if a > b {return b}
	return a
}