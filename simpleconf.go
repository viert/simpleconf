package simpleconf

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strings"
	"strconv"
)

type ConfigParser struct {
	options			map[string]string
}

type KeyError struct {
	Key 	string
}

func (ke *KeyError) Error() string {
	return fmt.Sprintf("Option '%s' not found in config file", ke.Key)
}

func ParseConfig(filename string) (*ConfigParser, error) {

	parser := &ConfigParser{ options : make(map[string]string) }
	optionRE := regexp.MustCompile(`^([\w\-\.]+)\s*=\s*(\S+)$`)

	file, err := os.Open(filename)
	if (err != nil) { return nil, err }
	defer file.Close()
	
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if (err != nil && line == "") { break }

		// removing comments
		found := strings.Index(line, "#")
		if (found != -1) { line = line[0:found] }

		// chomping line
		line = strings.Trim(line, " \n")

		match := optionRE.FindStringSubmatch(line)
		if (len(match) == 3) {
			name, value := match[1], match[2]
			parser.options[name] = value
		}
	}
	return parser, nil
}

func (p *ConfigParser) GetString(name string) (string, error) {
	v, ok := p.options[name]
	if (ok) { return v, nil }
	return "", &KeyError{ Key: name }
}

func (p *ConfigParser) GetInt(name string) (int, error) {
	v, ok := p.options[name]
	if (!ok) { 
		return 0, &KeyError{ Key: name } 
	}
	i, err := strconv.ParseInt(v, 0, 64)
	if (err != nil) { return 0, err }
	return int(i), nil
}

func (p *ConfigParser) Keys() []string {
	result := make([]string, 0)
	for k, _ := range p.options {
		result = append(result, k)
	}
	return result
}
