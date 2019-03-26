package conf

const (
	defaultConfFile = "/etc/pacman.conf"
)

type Config struct {
	colors Colors
}

func (config Config) TerminalColors() Colors {
	return config.colors
}

func LoadConfig() (*Config, error) {

	// parse config file
	parsedConf, err := ParseConfFile(defaultConfFile)
	if err != nil {
		return nil, err
	}
	colorsEnabled := parsedConf.HasOption("options", "Color")

	// build config
	var config *Config = new(Config)
	config.colors = buildColors(colorsEnabled)

	return config, nil
}

func buildColors(colorsEnabled bool) Colors {
	if colorsEnabled {
		return Colors{
			Title:   bold,
			Repo:    boldMagenta,
			Version: boldGreen,
			Meta:    boldCyan,
			NoColor: noColor,
		}
	} else {
		return Colors{
			Title:   "",
			Repo:    "",
			Version: "",
			Meta:    "",
			NoColor: "",
		}

	}
}
