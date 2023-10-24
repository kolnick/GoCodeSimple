package aoi

import (
	"fmt"
)

type AOIManager struct {
	MinX   int           // 区域左边界坐标
	MaxX   int           // 区域右边界坐标
	MinY   int           // 区域上边界坐标
	MaxY   int           // 区域下边界坐标
	CountX int           // x方向格子数量
	CountY int           // y方向的格子数量
	grids  map[int]*Grid // 当前区域中有那些格子，key=鸽子 ID，value =格子对象
}

// 初始化一个AOI区域
func NewAOIManager(minX, maxX, minY, maxY, countX, countY int) *AOIManager {
	manager := &AOIManager{
		MinX:   minX,
		MaxY:   maxY,
		MaxX:   maxX,
		MinY:   minY,
		CountX: countX,
		CountY: countY,
		grids:  make(map[int]*Grid),
	}

	// 初始化 AOI区域中所有的格子
	for y := 0; y < countY; y++ {
		for x := 0; x < countX; x++ {
			// 格子编号  = idy*nx+idx 利用格子坐标得到格子编号
			gid := y*countX + x
			manager.grids[gid] = NewGrid(gid,
				manager.MinX+x*manager.gridWidth(),
				manager.MinX+(x+1)*manager.gridWidth(),
				manager.MinY*manager.gridLength(),
				manager.MinY+(y+1)*manager.gridLength(),
			)
		}
	}
	return manager
}

// 得到每个格子在x轴方向的宽度

func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CountX
}

// 得到每个格子在x轴方向的长度
func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CountY
}

func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOiManager :\n minX: %d ,maxX %d, countX:%d,  minY:%d, maxY:%d ,countY:%d \n Grids in AOI Manager:\n",
		m.MinX, m.MaxX, m.CountX, m.MinY, m.MaxY, m.CountY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

// 根据gid获取周围的格子
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	// 判断gid是否存在
	if _, ok := m.grids[gID]; !ok {
		return
	}
	// 将当前gid添加到九宫格中
	grids = append(grids, m.grids[gID])
	// 根据gid得到当前格子所在的x轴编号
	idx := gID % m.CountX
	// 判断当前idx的左边是否还有格子

	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}
	if idx < m.CountX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	// 将x轴当前的格子都取出，进行遍历，再分别得到每个格子的上下是否有格子
	// 得到当前x轴的格子id集合
	gidXs := make([]int, 0, len(grids))
	for _, v := range grids {
		gidXs = append(gidXs, v.GID)
	}
	// 遍历x轴格子
	for _, v := range gidXs {
		// 计算格子处于第几列
		idy := v / m.CountX

		// 判断当前idy的上边是否还有格子
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CountX])
		}
		// 判断当前idy的下边是否还有格子
		if idy < m.CountY-1 {
			grids = append(grids, m.grids[v+m.CountX])
		}
	}
	return
}

// 通过坐标获取对应的格子ID
func (m *AOIManager) GetGIDByPos(x, y float32) int {
	gx := (int(x) - m.MinX) / m.gridWidth()
	gy := (int(x) - m.MinY) / m.gridLength()
	return gy*m.CountX + gx
}

// 通过坐标得到周边九宫格内的全部PlayerIds
func (m *AOIManager) GetPIDsByPos(x, y float32) (playerIDs []int) {
	// 根据当前坐标得到是哪个格子
	gid := m.GetGIDByPos(x, y)
	// 根据格子ID的达到周边九宫格信息
	grids := m.GetSurroundGridsByGid(gid)
	for _, v := range grids {
		playerIDs = append(playerIDs, v.getPlayerIDs()...)
		fmt.Printf("===>grid ID: %d ,pids: %v===", v.GID, v.getPlayerIDs())
	}
	return
}

//通过gid获取当前格子的全部PlayerId
func (m *AOIManager) GetPLayerIdsByGid(gid int) (playerIds []int) {
	return m.grids[gid].getPlayerIDs()
}

// 移除一个格子中的PlayerID
func (m *AOIManager) RemovePidFromGrid(pId, gId int) {
	m.grids[gId].Remove(pId)
}

// 通过坐标将一个player添加到一个格子中
func (m *AOIManager) AddToGridByPos(pId int, x, y float32) {
	gid := m.GetGIDByPos(x, y)
	grid := m.grids[gid]
	grid.Add(pId)
}

// 通过坐标把一个player从对应的格子中删除
func (m *AOIManager) RemoveFromGridByPos(pid int, x, y float32) {
	gid := m.GetGIDByPos(x, y)
	grid := m.grids[gid]
	grid.Remove(pid)
}
