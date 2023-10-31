package simple_factory

type MatchResult struct {
	leftPlayers  []Player
	rightPlayers []Player

	matchCode int32
}

type IMatchHandler interface {
	doMatch(t int32) MatchResult
}

type TeamMatchHandler struct {
	IMatchHandler
}

type PVPMatchHandler struct {
	IMatchHandler
}

func (this TeamMatchHandler) doMatch(t int32) MatchResult {

	return MatchResult{
		matchCode: 1,
	}
}
func (this PVPMatchHandler) doMatch(t int32) MatchResult {
	return MatchResult{
		matchCode: 2,
	}
}

func getMatchHandler(t int32) IMatchHandler {

	switch t {
	case 1:
		return TeamMatchHandler{}
	case 2:
		return PVPMatchHandler{}
	default:
		return nil
	}

}
