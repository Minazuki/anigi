package anigi

import (
	"anigi/config"
	"anigi/sourceProvider"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

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
	name string
	icons []string
	delay []int
	provider sourceProvider.SourceProvider
	ctx  context.Context
}

var (
	ErrJsonFileNotExist         = errors.New("error jsonFile file not Exist")
	ErrOpeningJsonFile          = errors.New("error opening jsonFile file")
	ErrReadingJsonFile          = errors.New("error reading jsonFile file")
	ErrUnmarshalingJsonFile     = errors.New("error unmarshaling jsonFile file")
	ErrNoIconsFoundInJsonFile   = errors.New("error no icons found in jsonFile file")
	ErrCantCreateSourceProvider = errors.New("error cant create source provider")
	ErrIconFileNotFound         = errors.New("error ico file not found")
)

func NewAnigi(ctx context.Context, cfg config.AnigiCfg) (*Anigi, error) {

	fmt.Printf("initializing anigi with icon: %s and source provider: %s \n", cfg.Icon, cfg.SourceProvider)
	workDir, _ := os.Getwd()
	jsonFile := filepath.Join(workDir, cfg.Icon, cfg.Icon+"_ico.json")

	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		return nil, ErrJsonFileNotExist
	}

	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, ErrOpeningJsonFile
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, ErrReadingJsonFile
	}

	iconSummary := iconSummary{}
	err = json.Unmarshal(byteValue, &iconSummary)
	if err != nil {
		return nil, ErrUnmarshalingJsonFile
	}

	if len(iconSummary.Files) == 0 {
		return nil, ErrNoIconsFoundInJsonFile
	}
	p, err := sourceProvider.NewSourceProvider(cfg.SourceProvider, ctx, 1*time.Second)
	if err != nil {
		return nil, ErrCantCreateSourceProvider
	}

	a := &Anigi{
		name: cfg.Tittle,
		icons: make([]string, 0, len(iconSummary.Files)),
		delay: make([]int, 0, len(iconSummary.Files)),
		provider: p,
		ctx:  ctx,
	}
	
	for _, icon := range iconSummary.Files {
		iconPath := filepath.Join(workDir, cfg.Icon, icon.FileName)
		if _, err := os.Stat(iconPath); os.IsNotExist(err) {
			return nil, ErrIconFileNotFound
		}
		a.icons = append(a.icons, iconPath)
		a.delay = append(a.delay, icon.Delay)        
	}

	return a, nil
}

func (a *Anigi) Run() {
	systray.Run(a.onReady, a.onExit)
}

func (a *Anigi) Stop() {
	systray.Quit()
}

func (a *Anigi) onReady() {
	if len(a.icons) == 0 {
		panic("No icons available to display.")
	}
	for _, iconPath := range a.icons {
		systray.CacheIcon(iconPath)
	}
	systray.SetIconFromCache(a.icons[0])
	systray.SetTitle(a.name)
	systray.SetTooltip("0%")

	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	pace := 1.0
	ticker := time.NewTicker(gifDelayToDuration(a.delay[0], pace))
	defer ticker.Stop()

	idx := 0
	maxIdx := len(a.icons) - 1
	for {
		select {
		case <-ticker.C:
			idx++
			if idx > maxIdx {
				idx = 0
			}
			systray.SetIconFromCache(a.icons[idx])
			t := gifDelayToDuration(a.delay[idx], pace)
			ticker.Reset(t)
			//fmt.Printf("Icon changed to: %s, next update in %v\n", a.icons[idx], t)
		case value := <-a.provider.ValueChan():
			valueStr := fmt.Sprintf("%s: %.2f%%", a.provider.Name(), value)
			systray.SetTooltip(valueStr)
			neoPace := percentToPace(value)
			if neoPace != pace {
				pace = neoPace
				//fmt.Printf("%s Pace changed to: %.2f\n", valueStr, pace)
			}
		case <-a.ctx.Done():
			systray.Quit()
		case <-quit.ClickedCh:
			systray.Quit()
			return
		}
	}
}

func (a *Anigi) onExit() {
	// Cleanup if necessary
	fmt.Println("Exiting anigi...")
}


func gifDelayToDuration(delay int, pace float64) time.Duration {
	// https://www.w3.org/Graphics/GIF/spec-gif89a.txt
	// vii) Delay Time - If not 0, this field specifies the number of
	// hundredths (1/100) of a second to wait before continuing with the
	// processing of the Data Stream. The clock starts ticking immediately
	// after the graphic is rendered. This field may be used in
	// conjunction with the User Input Flag field.

	// Convert GIF delay from hundredths of a second to time.Duration
	t := time.Duration(delay) * 10 * time.Millisecond
	return time.Duration(float64(t) * pace)
}

func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func percentToPace(percent float64) float64 {
	// 50  1.0
	// 0~50  2.0~1.0
	// 50~100 1.0~0
	if percent < 0 {
		percent = 0
	}
	if percent >= 100 {
		percent = 100
	}
	outPace := 1.0
	if percent == 50 {
		outPace = 1.0
	} else if percent < 50 {
		outPace = 2.0 - (percent/50.0)
	} else {		
		outPace = 1.0 - ((percent-50)/50.0)
	}
	// guard 0.01 to 2.0
	return clamp(outPace, 0.02, 2.0)
}
