package ktrans

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Options struct {
	Main        string   // Haupt-Karel-Datei
	Output      string   // Ausgabedatei (z.B. .pc)
	Prefix      string   // Installationspfad (optional)
	IncludeDirs []string // Include-Verzeichnisse
	RobotConfig string   // .cfg Datei, falls vorhanden

	Flags []string // zusätzliche Flags (/r, /d, …)
}

// Run startet ktrans mit den gegebenen Optionen
func Run(opts Options) error {
	ktransPath := FindKtransPath()
	if ktransPath == "" {
		return fmt.Errorf("ktrans binary not found in PATH")
	}

	args := []string{}

	// Flags
	args = append(args, opts.Flags...)

	// Output-Datei
	if opts.Output != "" {
		args = append(args, "/o", opts.Output)
	}

	// Robot Config
	if opts.RobotConfig != "" {
		args = append(args, "/config", opts.RobotConfig)
	}

	// Includes
	for _, inc := range opts.IncludeDirs {
		args = append(args, "/I", inc)
	}

	// Main-Datei
	args = append(args, opts.Main)

	// Exec
	cmd := exec.Command(ktransPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("▶ Running: %s %v\n", ktransPath, args)
	return cmd.Run()
}

// FindKtransPath sucht ktrans.exe im PATH
func FindKtransPath() string {
	path, err := exec.LookPath("ktrans.exe")
	if err != nil {
		return ""
	}
	abs, _ := filepath.Abs(path)
	return abs
}
