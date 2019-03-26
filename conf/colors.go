package conf

const (
	noColor     = "\033[0m"
	bold        = "\033[0;1m"
	black       = "\033[0;30m"
	red         = "\033[0;31m"
	green       = "\033[0;32m"
	yellow      = "\033[0;33m"
	blue        = "\033[0;34m"
	magenta     = "\033[0;35m"
	cyan        = "\033[0;36m"
	white       = "\033[0;37m"
	boldBlack   = "\033[1;30m"
	boldRed     = "\033[1;31m"
	boldGreen   = "\033[1;32m"
	boldYellow  = "\033[1;33m"
	boldBlue    = "\033[1;34m"
	boldMagenta = "\033[1;35m"
	boldCyan    = "\033[1;36m"
	boldWhite   = "\033[1;37m"
)

type Colors struct {
	Title   string
	Repo    string
	Version string
	NoColor string
	Meta    string
}
