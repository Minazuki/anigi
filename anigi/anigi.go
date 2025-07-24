package anigi

import (
	"anigi/config"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Minazuki/systray"
)

type iconFile struct {
	FileName string `json:"fileName"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Delay    int    `json:"delay"`
}

type iconSummary struct {
	Name  string     `json:"name"`
	Files []iconFile `json:"files"`
}

type Anigi struct {
	icons []string
	delay []int
}

func NewAnigi(cfg config.AnigiCfg) (*Anigi, error) {

	fmt.Printf("initializing anigi with icon: %s and source provider: %s \n", cfg.Icon, cfg.SourceProvider)
	workDir, _ := os.Getwd()
	jsonFile := filepath.Join(workDir, cfg.Icon, cfg.Icon+"_ico.json")

	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("icon file not found: %s", jsonFile)
	}

	file, err := os.Open(jsonFile)
    if err != nil {
        return nil, fmt.Errorf("error opening icon file: %s", jsonFile)
    }
    defer file.Close()

	byteValue, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("error reading icon file: %s", jsonFile)
    }

	iconSummary := iconSummary{}
	err = json.Unmarshal(byteValue, &iconSummary)
    if err != nil {
        return nil, fmt.Errorf("error unmarshaling icon file: %s", jsonFile)
    }

	if len(iconSummary.Files) == 0 {
		return nil, fmt.Errorf("no icons found in file: %s", jsonFile)
	}

	a := &Anigi{
		icons: make([]string, 0, len(iconSummary.Files)),
		delay: make([]int, 0, len(iconSummary.Files)),
	}
	
	for _, icon := range iconSummary.Files {
		iconPath := filepath.Join(workDir, cfg.Icon, icon.FileName)
		if _, err := os.Stat(iconPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("icon file not found: %s", iconPath)
		}
		a.icons = append(a.icons, iconPath)
		a.delay = append(a.delay, icon.Delay)
	}

	return a, nil
}

func (a *Anigi) Run() {
	systray.Run(a.onReady, a.onExit)
}

func (a *Anigi) onReady() {
	if len(a.icons) == 0 {
		panic("No icons available to display.")
	}
	for _, iconPath := range a.icons {
		systray.CacheIcon(iconPath)
	}
	systray.SetIconFromCache(a.icons[0])
}

func (a *Anigi) onExit() {
	// Cleanup if necessary
	fmt.Println("Exiting anigi...")
}