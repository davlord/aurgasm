package conf

import (
	"bufio"
	"os"
	"strings"
)

type ParsedConf struct {
	sections map[string]section
}

func (parsedConf ParsedConf) HasOption(section string, option string) (ok bool) {
	_, ok = parsedConf.sections[section].options[option]
	return
}

func (parsedConf ParsedConf) Option(section string, option string) (value string) {
	value, _ = parsedConf.sections[section].options[option]
	return
}

type section struct {
	options map[string]string
}

type parsingContext struct {
	parsedConf     ParsedConf
	currentSection section
}

func ParseConfFile(path string) (*ParsedConf, error) {
	var parsingContext *parsingContext = &parsingContext{
		parsedConf: ParsedConf{
			sections: make(map[string]section),
		},
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parseLine(parsingContext, scanner.Text())
	}

	return &parsingContext.parsedConf, nil
}

func parseLine(parsingContext *parsingContext, originalLine string) {

	// trim spaces at start and end of line
	line := strings.Trim(originalLine, " ")

	// skip empty lines
	if len(line) == 0 {
		return
	}

	// skip comments
	if strings.HasPrefix(line, "#") {
		return
	}

	// parse
	fields := strings.Fields(line)
	switch len(fields) {
	case 1:
		if strings.HasPrefix(line, "[") {
			// process section
			sectionName := line[1 : len(line)-1]
			section := section{
				options: make(map[string]string),
			}
			parsingContext.parsedConf.sections[sectionName] = section
			parsingContext.currentSection = section
		} else {
			// process single option
			parsingContext.currentSection.options[line] = ""
		}
	case 3:
		if fields[1] == "=" {
			// process key/value option
			parsingContext.currentSection.options[fields[0]] = fields[2]
		}
	}
}
