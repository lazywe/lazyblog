package models

type Recomment struct {
	Id    int
	Title string
}

//
// 获取comment
// 
//
func (this *Recomment) getComment() []*Recomment {
	var recomment []*Recomment
	recomment = append(recomment, &Recomment{Id: 1, Title: "精品推荐"})
	return recomment
}
