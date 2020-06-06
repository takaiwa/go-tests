package initialisation

/*
　構造体や配列の初期化に関するメモ
*/
type Struct1 struct {
	x string
	y []string
}

func GetStruct1() *Struct1 {
	s1 := new(Struct1)

	s1.x = "A"
	s1.y = append(s1.y, "B")
	s1.y = append(s1.y, "C")

	return s1
}
