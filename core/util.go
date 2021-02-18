package core

const hextable = "0123456789ABCDEF"

func EncodeToString(b []byte) string {
	l := len(b)
	h := make([]byte, l*2)
	for i, v := range b {
		h[i*2] = hextable[v>>4]
		h[i*2+1] = hextable[v&0x0f]
	}
	return string(h)
}

func Equals(a Address, b Address) bool {
	if a == nil || b == nil {
		return false
	}
	aa, err := a.GetAddress()
	if err != nil {
		return false
	}
	ba, err := a.GetAddress()
	if err != nil {
		return false
	}
	return aa == ba
}
