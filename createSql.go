package main

import (
	"fmt"
	"net/url"
)

type formData struct {
	nation, area, ideology, team, sea, nf string
}

func (fd *formData) setFromData(data url.Values) {
	fd.nation = data.Get("nation_name")
	fd.area = data.Get("area")
	fd.ideology = data.Get("ideology")
	fd.team = data.Get("team")
	fd.sea = data.Get("sea")
	fd.nf = data.Get("nf")
}

func createSql(fd *formData) {
	sqlStr := "insert into nations(nation_name,area_id,ideology_id,team_id,sea,nf) values(\"" + fd.nation + "\"," + fd.area + "," + fd.ideology + "," + fd.team + "," + fd.sea + "," + fd.nf + ");"
	fmt.Println(sqlStr)
}
