package model
type Section struct {
	List []Good//商品列表
	Rank []Good//商品排行
}
type SectionMerchant struct {
	List []Merchant
	Rank []Merchant
}