package main

import (
	//"database/sql"
    _ "github.com/go-sql-driver/mysql"
    //"fmt"
)
	
//------------- AdblockType enum ----------
type AdblockType int

const (
	TextAdBlock AdblockType = 1 + iota
	ImageAdBlock
	TextImageAdBlock
)

var adblockTypeStr = [...] string {
	"text",
	"image",
	"text+image",
}

func (adblockType AdblockType) String() string{
	return adblockTypeStr[adblockType - 1]
}
//----------------------------------------

type AdBlock struct {
	id int					//TODO int64 ?
	adblockType AdblockType
	place string
	size string

	//Cписок баннеров данного блока
	bannerList BannerList
}

type AdBlockList map[int]AdBlock

