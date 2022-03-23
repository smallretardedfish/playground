package downcase

func ToLower(str string) string {
	var res []byte = []byte(str)

	for i := range res {
		if res[i] <= 90 && res[i] >= 65 {
			res[i] += 32
		}
	}
	return string(res)
}
