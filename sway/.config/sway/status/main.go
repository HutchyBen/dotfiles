package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"github.com/lucasb-eyer/go-colorful"
)

type Block struct {
	FullText           string `json:"full_text,omitempty"`
	ShortText          string `json:"short_text,omitempty"`
	Color              string `json:"color,omitempty"`
	Background         string `json:"background,omitempty"`
	Border             string `json:"border,omitempty"`
	BorderTop          int    `json:"border_top,omitempty"`
	BorderRight        int    `json:"border_right,omitempty"`
	BorderBottom       int    `json:"border_bottom,omitempty"`
	BorderLeft         int    `json:"border_left,omitempty"`
	MinWidth           int    `json:"min_width,omitempty"`
	Align              string `json:"align,omitempty"`
	Urgent             bool   `json:"urgent,omitempty"`
	Name               string `json:"name,omitempty"`
	Instance           string `json:"instance,omitempty"`
	Separator          bool   `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	Markup             string `json:"markup,omitempty"`
}

type Input struct {
	Name       string   `json:"name"`
	Instance   string   `json:"instance"`
	Button     int      `json:"button"`
	Modifiers  []string `json:"modifiers"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
	RelativeX  int      `json:"relative_x"`
	RelativeY  int      `json:"relative_y"`
	OutputX    int      `json:"output_x"`
	OutputY    int      `json:"output_y"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
}

func getBattery() Block {
	b := Block{}
	capa, err := os.ReadFile("/sys/class/power_supply/BAT1/capacity")
	if err != nil {
		return Block{}
	}

	status, err := os.ReadFile("/sys/class/power_supply/BAT1/status")
	if err != nil {
		return Block{}
	}

	statusStr := strings.TrimSpace(string(status))
	capacityStr := strings.TrimSpace(string(capa))
	
	if capacityStr == "100" {
		b.FullText = "Battery Full"
	} else {
		switch statusStr {
		case "Charging":
			b.FullText = fmt.Sprintf("+ Battery %s%%", capacityStr)
			break
		case "Discharging":
			b.FullText = fmt.Sprintf("- Battery %s%%", capacityStr)
			break
		default:
			b.FullText = fmt.Sprintf("%s %s%%", statusStr, capacityStr)
			break
		}
	}

	// calculate color
	cInt, _ := strconv.Atoi(capacityStr)
    	c1, _ := colorful.Hex("#f38ba8")
    	c2, _ := colorful.Hex("#a6e3a1")
	blend := c1.BlendHsv(c2, float64(cInt) / 100)

	b.Color = blend.Hex()
	b.MinWidth = len("- Battery 100%") // this will surely get optimized
	b.Name = "battery"
	b.Align = "center"
	b.SeparatorBlockWidth = 21
	return b
}

func getBrightness() Block {
	b := Block{}
	brightness, err := os.ReadFile("/sys/class/backlight/intel_backlight/brightness")
	maxBrightness, err := os.ReadFile("/sys/class/backlight/intel_backlight/max_brightness")
	if err != nil {
		return Block{}
	}

	brightnessInt, _ := strconv.Atoi(strings.TrimSpace(string(brightness)))
	maxBrightnessInt, _ := strconv.Atoi(strings.TrimSpace(string(maxBrightness)))

	b.FullText =  fmt.Sprintf("Brightness %d%%", int((float32(brightnessInt) / float32(maxBrightnessInt)) * 100))
	b.Name = "brightness"
	b.MinWidth = len("Brightness 100%")
	b.Align = "center"
	b.SeparatorBlockWidth = 21
	return b
}

func inputBrightness(input Input) {
	if input.Button == 4 {
		err := exec.Command("/home/ben/.config/sway/systool/systool", "brightness", "+1000").Run()
		if err != nil {
			log.Println(err)
		}
	} else if input.Button == 5 {
		err := exec.Command("/home/ben/.config/sway/systool/systool", "brightness", "-1000").Run()
		if err != nil {
			log.Println(err)
		}
	}
}

func getTime() Block {
	b := Block{}
	now := time.Now()
	
	b.FullText = now.Format("15:04:05")
	b.Name = "time"
	b.Align = "center"
	b.SeparatorBlockWidth = 21
	return b
}

func ProcessInput(inputStr string) {
	var input Input
	log.Println(inputStr)
	err := json.Unmarshal([]byte(inputStr), &input)
	if err != nil {
		log.Println(err)
	}
	switch (input.Name) {
	case "brightness":
		inputBrightness(input)
		break
	}
}

func main() {
	logFile, _ := os.OpenFile("/tmp/status.log", os.O_RDWR | os.O_CREATE, 0666)
	log.SetOutput(logFile)


	fmt.Println("{\"version\":1, \"click_events\": true}")
	fmt.Print("[")
	
	go func () {
		brackets := 0
		json := ""
		for {
			var str string
			fmt.Scan(&str)
			str = strings.TrimPrefix(str, ",") // new inputs start with a comma which ignoring
			str = strings.TrimPrefix(str, "[") // ignore start of the list
			json += str

			if str == "{" {
				brackets++
			} else if str == "}" {
				brackets--
				if brackets == 0 {
					ProcessInput(json)
					json = ""
				}
			}
		}
	}()
	for {	
		blocks := []Block{getBrightness(), getBattery(), getTime()}
		data, _ := json.MarshalIndent(blocks, "", "    ")
		fmt.Println(string(data))

		fmt.Print(",")
		time.Sleep(500 * time.Millisecond)
	}
}
