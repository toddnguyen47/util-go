package convertslicestomap

func Convert(input []string) map[string]string {
	resultsMap := make(map[string]string)
	for _, val := range input {
		resultsMap[val] = val
	}
	return resultsMap
}
