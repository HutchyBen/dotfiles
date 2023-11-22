package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

type Block struct {
	FullText            string `json:"full_text,omitempty"`
	ShortText           string `json:"short_text,omitempty"`
	Color               string `json:"color,omitempty"`
	Background          string `json:"background,omitempty"`
	Border              string `json:"border,omitempty"`
	BorderTop           int    `json:"border_top,omitempty"`
	BorderRight         int    `json:"border_right,omitempty"`
	BorderBottom        int    `json:"border_bottom,omitempty"`
	BorderLeft          int    `json:"border_left,omitempty"`
	MinWidth            int    `json:"min_width,omitempty"`
	Align               string `json:"align,omitempty"`
	Urgent              bool   `json:"urgent,omitempty"`
	Name                string `json:"name,omitempty"`
	Instance            string `json:"instance,omitempty"`
	Separator           bool   `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	Markup              string `json:"markup,omitempty"`
}

type Input struct {
	Name      string   `json:"name"`
	Instance  string   `json:"instance"`
	Button    int      `json:"button"`
	Modifiers []string `json:"modifiers"`
	X         int      `json:"x"`
	Y         int      `json:"y"`
	RelativeX int      `json:"relative_x"`
	RelativeY int      `json:"relative_y"`
	OutputX   int      `json:"output_x"`
	OutputY   int      `json:"output_y"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
}

type Config struct {
	MonzoPlaygroundKey string `json:"monzo_playground_key"`
}

var config Config

func getMonzo(b *Block) {
	var accounts struct {
		Accounts []struct {
			ID string `json:"id"`
		} `json:"accounts"`
	}
	for len(accounts.Accounts) == 0 {
		req, _ := http.NewRequest("GET", "https://api.monzo.com/accounts", nil)
		req.Header.Add("Authorization", "Bearer "+config.MonzoPlaygroundKey)
	
		res, _ := http.DefaultClient.Do(req)
		json.NewDecoder(res.Body).Decode(&accounts)
		time.Sleep(time.Minute)
	}
	for {
		var Balance struct {
			TotalBalance int `json:"total_balance"`
		}

		req, _ := http.NewRequest("GET", "https://api.monzo.com/balance?account_id="+accounts.Accounts[0].ID, nil)
		req.Header.Add("Authorization", "Bearer "+config.MonzoPlaygroundKey)
		res, _ := http.DefaultClient.Do(req)
		json.NewDecoder(res.Body).Decode(&Balance)

		b.Name = "money"
		b.Align = "center"
		b.SeparatorBlockWidth = 21
		b.FullText = fmt.Sprintf("Money: £%.2f", float32(Balance.TotalBalance)/100)
		time.Sleep(time.Minute * 5)
	}
}

func getBattery(b *Block) {
	if b == nil {
		b = new(Block)
	}
	for {
		capa, err := os.ReadFile("/sys/class/power_supply/BAT1/capacity")
		if err != nil {
			return
		}

		status, err := os.ReadFile("/sys/class/power_supply/BAT1/status")
		if err != nil {
			return
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
		blend := c1.BlendHsv(c2, float64(cInt)/100)

		b.Color = blend.Hex()
		b.MinWidth = len("- Battery 100%") // this will surely get optimized
		b.Name = "battery"
		b.Align = "center"
		b.SeparatorBlockWidth = 21
		time.Sleep(time.Second * 5)
	}
}
func getVolume(b *Block) {
	if b == nil {
		b = new(Block)
	}
	for {

		var volume string

		cmd, err := exec.Command("pactl", "get-sink-volume", "@DEFAULT_SINK@").Output()
		if err != nil {
			return
		}

		volume = strings.Split(string(cmd), "/")[1]

		volume = strings.TrimSpace(volume)
		b.FullText = fmt.Sprintf("Volume %s", volume)
		b.Name = "volume"
		b.MinWidth = len("Volume 100%")
		b.Align = "center"
		b.SeparatorBlockWidth = 21
		time.Sleep(250 * time.Millisecond)
	}
}

func inputVolume(input Input) {
	if input.Button == 4 {
		err := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "+5%").Run()
		if err != nil {
			log.Println(err)
		}
	} else if input.Button == 5 {
		err := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "-5%").Run()
		if err != nil {
			log.Println(err)
		}
	}
}

func getBrightness(b *Block) {
	if b == nil {
		b = new(Block)
	}
	for {
		brightness, err := os.ReadFile("/sys/class/backlight/intel_backlight/brightness")
		maxBrightness, err := os.ReadFile("/sys/class/backlight/intel_backlight/max_brightness")
		if err != nil {
			return
		}

		brightnessInt, _ := strconv.Atoi(strings.TrimSpace(string(brightness)))
		maxBrightnessInt, _ := strconv.Atoi(strings.TrimSpace(string(maxBrightness)))

		b.FullText = fmt.Sprintf("Brightness %d%%", int((float32(brightnessInt)/float32(maxBrightnessInt))*100))
		b.Name = "brightness"
		b.MinWidth = len("Brightness 100%")
		b.Align = "center"
		b.SeparatorBlockWidth = 21
		time.Sleep(500 * time.Millisecond)
	}
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

func getTime(b *Block) {
	if b == nil {
		b = new(Block)
	}
	for {
		now := time.Now()

		b.FullText = now.Format("15:04:05")
		b.Name = "time"
		b.Align = "center"
		b.SeparatorBlockWidth = 21
		time.Sleep(time.Second)
	}
}

func ProcessInput(inputStr string) {
	var input Input
	log.Println(inputStr)
	err := json.Unmarshal([]byte(inputStr), &input)
	if err != nil {
		log.Println(err)
	}
	switch input.Name {
	case "brightness":
		inputBrightness(input)
		break
	case "volume":
		inputVolume(input)
		break
	}
}

func main() {
	logFile, _ := os.OpenFile("/tmp/status.log", os.O_RDWR|os.O_CREATE, 0666)
	log.SetOutput(logFile)
	configFile, _ := os.Open("/home/ben/.config/sway/status/config.json")
	json.NewDecoder(configFile).Decode(&config)

	fmt.Println("{\"version\":1, \"click_events\": true}")
	fmt.Print("[")

	go func() {
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

	// Async Shitstains
	var Monzo Block
	var Volume Block
	var Brightness Block
	var Battery Block
	var Time Block

	go getMonzo(&Monzo)
	go getVolume(&Volume)
	go getBrightness(&Brightness)
	go getBattery(&Battery)
	go getTime(&Time)
	for {
		blocks := []Block{Volume, Brightness, Battery, Time}
		if Monzo.FullText != "" {
			blocks = append([]Block{Monzo}, blocks...)
		}

		data, _ := json.MarshalIndent(blocks, "", "    ")
		fmt.Println(string(data))

		fmt.Print(",")
		time.Sleep(50 * time.Millisecond)
	}
}
