
// Usage:
// for i := 0; i < len(b); {
//     x, i = nextInt(b, i)
//     fmt.Println(x)
// }

func nextInt(b []byte, i int) (val, ni int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}

func isDigit(c byte) bool {
    switch {
    case '0' <= c && c <= '9':
		return true
    }
    return false
}
