package main
	
type Banner struct {
	id int
	adcampaign_id int
	adblock_id int
	bannerType string
	labels string
	url string
	content string
}

type BannerList map[int]Banner