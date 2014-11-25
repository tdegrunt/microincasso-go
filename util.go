package microincasso

func ConcatenateBytes(s1, s2 [][]byte) [][]byte {

	union := s1
	for _, v := range s2 {
		union = append(union, v)
	}

	return union

}
