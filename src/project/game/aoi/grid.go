package aoi

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	playerIDs map[int]bool
	pIDLock   sync.RWMutex
}

func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {

	return &Grid{
		GID:       gID,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: map[int]bool{},
	}
}

func (g *Grid) Add(playerId int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerId] = true
}

func (g *Grid) Remove(playerId int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerId)
}

func (g *Grid) getPlayerIDs() (playerIds []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	for k, _ := range g.playerIDs {
		playerIds = append(playerIds, k)
	}
	return
}

func (this *Grid) String() string {
	return fmt.Sprintf("Grid id :%d,minX:%d,maxX :%d ,minY :%d ,maxY:%d ,playerIds:%v",
		this.GID, this.MinX, this.MaxX, this.MinY, this.MaxY, this.playerIDs)
}
