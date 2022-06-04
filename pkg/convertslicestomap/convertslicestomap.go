package convertslicestomap

func Convert(input []string) map[string]string {
	resultsMap := make(map[string]string)
	for i := 0; i < len(input); i++ {
		val := input[i]
		resultsMap[val] = val
	}
	return resultsMap
}
