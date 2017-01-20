package main

import (
	"database/sql"
    //_ "github.com/ziutek/mymysql/godrv"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    //"time"
)

//-------- Status enum ----------
type Status int

const (
	Active Status = 1 + iota
	Blocked
)

var statusStr = [...]string {
	"Active",
	"Blocked",
}

func (status Status) String() string {
	return statusStr[status - 1]
}
//------------------------------------


func readDB(info *Info) error {
	//listDrivers := sql.Drivers()

	db, err := sql.Open("mysql", "adserver:adserver@tcp(adserver-db-mysql-dev:3306)/adserver4")
	defer db.Close()
	if err != nil {
		fmt.Println("Cannot connect to db: ", err.Error())
	}

	rows, err := db.Query("Select id, name, ts_begin, ts_end, status+0 AS status from AdCampaign")
	//rows, err := db.Query("Select id, name, status from AdCampaign")
	defer rows.Close()
	if err != nil {
		fmt.Println("Cannot execute query: ", err.Error())
	}

	for rows.Next() {
		var id int
		var name sql.NullString
		var timeBegin NullTime
		var timeEnd NullTime
		var status Status

		rows.Scan(&id, &name, &timeBegin, &timeEnd, &status)
		fmt.Println(": ", id, name, status)
		if timeBegin.Valid {
			fmt.Println("begin ", timeBegin)
		}

		if timeEnd.Valid {
			fmt.Println("end ", timeEnd)
		}
	}

	//fmt.Println("db: ", listDrivers)

	//var adblockType AdblockType 
	//adblockType = 1
	//fmt.Println("type ", adblockType)

	//var blockList AdBlockList
	//blockList = make(AdBlockList)
	//readAdBlocks(db, blockList)
	readAdBlocks(db, info.adBlockList)
	readBanners(db, info.adBlockList)
	return nil
}
	
func readAdBlocks(db *sql.DB, adBlockList AdBlockList) error {
	rows, err := db.Query("SELECT id, type+0 AS type, place, size FROM AdBlock")
	defer rows.Close()
	if err != nil {
		fmt.Println("Error occured while select from AdBlock: ", err.Error())
		return err
	}

	// Заполняем список рекламных блоков
	var block AdBlock
	for rows.Next() {
		err = rows.Scan(&block.id, &block.adblockType, &block.place, &block.size)
		if err != nil {
			fmt.Println("Error occured while scan AdBlock: ", err.Error())
			return err
		}

		block.bannerList = make(BannerList)
		adBlockList[block.id] = block //search _, ok := m["route"]
	}

	//for _, value := range adBlockList {
	//	fmt.Println("id = ", value.id, " ", value.adblockType, " ", value.place)
	//}

	return nil
}

func readBanners(db *sql.DB, adBlockList AdBlockList) error {
	rows, err := db.Query("SELECT id, adcampaign_id, adblock_id, type, labels, content, url, status+0 AS status FROM Banner")
	defer rows.Close()

	if err != nil {
		fmt.Println("Error occured while select from Banner: ", err.Error())
		return err
	}

	var banner Banner
	var status Status

	var labels sql.NullString
	var content sql.NullString
	var url sql.NullString

	for rows.Next() {
		err = rows.Scan(&banner.id, &banner.adcampaign_id, &banner.adblock_id, &banner.bannerType, &labels, 
						&content, &url, &status)

		if err != nil {
			fmt.Println("Error occured while scan Banner: ", err.Error())
			return err
		}

		if status != Active {
			fmt.Println("Banner(", banner.id, ") status is not Active ")
			continue
		}

		adBlock, ok := adBlockList[banner.adblock_id]
		if !ok {
			fmt.Println("There is no adBlock with id: ", banner.adblock_id)
			continue
		}

		if labels.Valid {
			banner.labels = labels.String
		}

		if content.Valid {
			banner.content = content.String
		}

		if url.Valid {
			banner.url = url.String
		}

		adBlock.bannerList[banner.id] = banner
	}

	//-->
	for k, v := range adBlockList {
		fmt.Println("block id: ", k)

		for kBanner, vBanner := range v.bannerList {
			fmt.Println("banner id: ", kBanner, " cont=", vBanner.content)
		}
	}
	//<--

	return nil
}