
// Usage:
// for i := 0; i < len(b); {
//     x, i = nextInt(b, i)
//     fmt.Println(x)
// }

func nextInt(b []byte, i int) (val, ni int) {
	for ; i < len(b) && !('0' <= b[i] && b[i] <= '9'); i++ {
	}
	x := 0
	for ; i < len(b) && '0' <= b[i] && b[i] <= '9'; i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}
