package commands

type VhsCommand string

const (
	Undefined   VhsCommand = ""
	CreateSite  VhsCommand = "create"
	EnableSite  VhsCommand = "enable"
	DisableSite VhsCommand = "disable"
	DeleteSite  VhsCommand = "delete"
)

func ToVhsCommand(s string) VhsCommand {
	switch s {
	case "create":
		return CreateSite
	case "enable":
		return EnableSite
	case "disable":
		return DisableSite
	case "delete":
		return DeleteSite
	default:
		return Undefined
	}
}
