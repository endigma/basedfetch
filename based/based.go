package based

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/endigma/basedfetch/sys"
	"github.com/fatih/color"
)

var score int = 0

func Compute(s sys.System) {
	// spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	// spin.Start()                                                   // Start the spinner
	// time.Sleep(2 * time.Second)                                    // Run for some time to simulate work
	// spin.Stop()

	if strings.Contains(s.GPUModel, "GeForce") || strings.Contains(s.GPUModel, "Nvidia") {
		entry("eww nvidia gpu", -2)
	}

	if s.CPUVendor == "AuthenticAMD" {
		entry("amd cpu... based.", +5)
	}

	if strings.Contains(s.CPUModel, "Ryzen") {
		entry("ryzen based shintel btfoed", +5)
	}

	if s.Arch != "x86_64" {
		entry("non-x86 is based", +5)
	}

	switch {
	case strings.HasPrefix(s.Kernel, "5.") && s.OS == "linux":
		entry("latest major linux version. based.", +2)

	case strings.HasPrefix(s.Kernel, "4.") && s.OS == "linux":
		entry("linux 4 ??", -1)

	case strings.HasPrefix(s.Kernel, "3.") && s.OS == "linux":
		entry("WHY ARE YOU USING LINUX 3.X", -5)

	case strings.HasPrefix(s.Kernel, "2.") && s.OS == "linux":
		entry("WHAT THE SHITFUCK WENT WRONG FOR YOU TO BE RUNNING THIS ON LINUX 2.X??????", -10)

	case strings.HasPrefix(s.Kernel, "1.") && s.OS == "linux":
		entry("WHAT ARE YOU EVEN DOING? LINUX 1.X??? WHY?", -20)
	}

	switch s.OS {
	case "linux":
		entry("linux linux based based gnu stallman approves", +5)
	case "darwin":
		entry("tim apple approves of your pc", +2)
	case "windows":
		entry("EWWWWWWWWWWWWWWWWWWW WINDOWS :PUKE:", -15)
	default:
		entry("nonstandard OS. Based.", -15)
	}

	if s.MemoryTotal > 16000 {
		entry("16GB of ram is the bare minimum", +5)
	}

	if s.User == s.Hostname {
		entry("why is your username and hostname the same? what is wrong with you?", -5)
	}

	if s.User == "User" || s.User == "user" {
		entry("why is your username \"user\" ?? are you dumb?", -5)
	}

	if s.MemoryUsed < 3000 {
		entry("low memory usage, based.", +5)
	}

	if (s.MemoryTotal - s.MemoryUsed) < 1000 {
		entry("I BOUGHT MY RAM IM GONNA USE IT", +5)
	}

	fmt.Printf("\n%s: !!! %s !!!\n", color.New(color.FgHiBlack, color.Bold).Sprint("Based score"), color.New(func(points int) color.Attribute {
		if points >= 0 {
			return color.FgGreen
		}
		return color.FgRed
	}(score), color.Bold).Sprintf("%+d", score))

	switch {
	case score > 15:
		color.HiGreen("alright, you're pretty based.")
	case score > 10:
		color.Green("pretty alright score, but you could be more based.")
	case score > 5:
		color.HiYellow("scraping the barrel here buddy, you need to get these scores up.")
	case score > 0:
		color.Yellow("if you're any less based, the stallman might come get you...")
	default:
		color.Red("you should stop using computers")
	}

}

func entry(msg string, points int) {
	colors := []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
		color.FgWhite,
		color.FgHiBlack,
		color.FgHiRed,
		color.FgHiGreen,
		color.FgHiYellow,
		color.FgHiBlue,
		color.FgHiMagenta,
		color.FgHiCyan,
		color.FgHiWhite,
	}

	score += points

	rand.Seed(time.Now().UnixNano())

	c := colors[rand.Intn(len(colors))]

	fmt.Printf("%s (%s)\n",
		color.New(c).Sprint(msg),
		color.New(func(points int) color.Attribute {
			if points >= 0 {
				return color.FgGreen
			}
			return color.FgRed
		}(points)).Sprintf("%+d", points))
}
