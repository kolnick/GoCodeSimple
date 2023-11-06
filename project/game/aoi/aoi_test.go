package aoi

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	manager := NewAOIManager(100, 300, 4, 200, 450, 5)
	fmt.Println(manager)
}

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	manager := NewAOIManager(0, 250, 5, 0, 250, 5)

	for k, _ := range manager.grids {
		// 获取当前格子周边的九宫格
		grids := manager.GetSurroundGridsByGid(k)
		// 得到九宫格所有的ids
		fmt.Println("gid :,k ,grids len=", len(grids))
		gids := make([]int, 0, len(grids))
		for _, grid := range gids {
			gids = append(gids, grid)
		}
		fmt.Printf(" grid ID:%d ,surrounding grid IDs are %v\n", k, gids)
	}

}
