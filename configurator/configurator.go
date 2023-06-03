package configurator

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type configurator struct {
	dirPath string
}

type Configurator interface {
	ApplyTheme(string) error
	ListThemes() ([]string, error)
	SelectTheme() (string, error)
}

func (c *configurator) getConfigDirPath() error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	c.dirPath = userHome + "/.config/alacritty"
	return nil
}

func (c *configurator) readYML(filePath string) (map[string]interface{}, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *configurator) readConfiguration() (map[string]interface{}, error) {
	filePath := fmt.Sprintf("%s/alacritty.yml", c.dirPath)
	return c.readYML(filePath)
}

func (c *configurator) readTheme(themeName string) (map[string]interface{}, error) {
	filePath := fmt.Sprintf("%s/themes/%s.yml", c.dirPath, themeName)
	return c.readYML(filePath)
}

func (c *configurator) mergeThemeInConf(conf map[string]interface{}, theme map[string]interface{}) error {
	conf["colors"] = theme["colors"]

	file, err := os.Create(fmt.Sprintf("%s/alacritty.yml", c.dirPath))
	if err != nil {
		return err
	}

	defer file.Close()

	content, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	file.Write(content)

	return nil
}

func (c *configurator) ApplyTheme(themeName string) error {
	if themeName == "" {
		return errors.New("empty theme name")
	}

	cleanRgx := regexp.MustCompile(`[\n\r]`)
	themeName = cleanRgx.ReplaceAllString(themeName, "")

	conf, err := c.readConfiguration()
	if err != nil {
		return err
	}

	theme, err := c.readTheme(themeName)
	if err != nil {
		return err
	}

	err = c.mergeThemeInConf(conf, theme)
	if err != nil {
		return err
	}

	return nil
}

func (c *configurator) SelectTheme() (string, error) {
	availableThemes, err := c.ListThemes()
	if err != nil {
		return "", err
	}

	cmdOutput := &bytes.Buffer{}
	command := fmt.Sprintf("echo -e '%s' | fzf", strings.Join(availableThemes, "\n"))
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 130" {
		return "", err
	}

	return cmdOutput.String(), nil
}

func (c *configurator) ListThemes() ([]string, error) {
	files, err := os.ReadDir(fmt.Sprintf("%s/themes", c.dirPath))
	if err != nil {
		return nil, err
	}

	nameRgx := regexp.MustCompile(".yml")
	filesName := []string{}
	for _, file := range files {
		name := nameRgx.ReplaceAllString(file.Name(), "")
		filesName = append(filesName, name)
	}

	return filesName, nil
}

func NewConfigurator() (Configurator, error) {
	instance := configurator{}
	err := instance.getConfigDirPath()
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
