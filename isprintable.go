package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] < 32 { // les charactères imprimables commencent au n°32 dans la table Ascii
			return false
		}
	}
	return true
}
