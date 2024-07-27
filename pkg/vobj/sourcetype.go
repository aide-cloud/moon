package vobj

// Status 数据状态
//
//go:generate go run ../../cmd/server/stringer/cmd.go -type=SourceType -linecomment
type SourceType int

const (
	// SourceTypeUnknown 未知
	SourceTypeUnknown SourceType = iota // 未知

	SourceSystem // 系统来源

	SourceTeam // 团队来源
)

func GetSourceType(sourceType string) SourceType {
	switch sourceType {
	case "System":
		return SourceSystem
	case "Team":
		return SourceTeam
	default:
		return SourceTypeUnknown
	}
}
