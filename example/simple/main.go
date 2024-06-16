package simple

import "github.com/humangrass/sup"

// nolint:unused // example
func main() {
	sup.SetProjectName("Yo!")
	sup.SetLogLevel(sup.InfoLevel) // by default

	sup.Info("Wazzaaaaa!!!") // 2006-01-31 22:33:44 [Yo!] Wazzaaaaa!!!
}
