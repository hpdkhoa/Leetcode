package _11_69_Sqrt_x_

func mySqrt(x int) int {
	g := x
	for g*g > x {
		g = (g+x/g)/2
	}
	return g
}