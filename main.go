package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// @TODO: Check the color output of vscode and terminal. Color space conversion needed?

const defaultColor = "#FF0000"

type VSCodeTheme struct {
	Name   string `json:"name"`
	Colors struct {
		EditorForeground                           string `json:"editor.foreground"`                           // Editor background color.
		EditorBackground                           string `json:"editor.background"`                           // Editor default foreground color.
		TerminalBackground                         string `json:"terminal.background"`                         // The background of the Integrated Terminal's viewport.
		TerminalBorder                             string `json:"terminal.border"`                             // The color of the border that separates split panes within the terminal. This defaults to panel.border.
		TerminalForeground                         string `json:"terminal.foreground"`                         // The default foreground color of the Integrated Terminal.
		TerminalAnsiBlack                          string `json:"terminal.ansiBlack"`                          // 'Black' ANSI color in the terminal.
		TerminalAnsiBlue                           string `json:"terminal.ansiBlue"`                           // 'Blue' ANSI color in the terminal.
		TerminalAnsiBrightBlack                    string `json:"terminal.ansiBrightBlack"`                    // 'BrightBlack' ANSI color in the terminal.
		TerminalAnsiBrightBlue                     string `json:"terminal.ansiBrightBlue"`                     // 'BrightBlue' ANSI color in the terminal.
		TerminalAnsiBrightCyan                     string `json:"terminal.ansiBrightCyan"`                     // 'BrightCyan' ANSI color in the terminal.
		TerminalAnsiBrightGreen                    string `json:"terminal.ansiBrightGreen"`                    // 'BrightGreen' ANSI color in the terminal.
		TerminalAnsiBrightMagenta                  string `json:"terminal.ansiBrightMagenta"`                  // 'BrightMagenta' ANSI color in the terminal.
		TerminalAnsiBrightRed                      string `json:"terminal.ansiBrightRed"`                      // 'BrightRed' ANSI color in the terminal.
		TerminalAnsiBrightWhite                    string `json:"terminal.ansiBrightWhite"`                    // 'BrightWhite' ANSI color in the terminal.
		TerminalAnsiBrightYellow                   string `json:"terminal.ansiBrightYellow"`                   // 'BrightYellow' ANSI color in the terminal.
		TerminalAnsiCyan                           string `json:"terminal.ansiCyan"`                           // 'Cyan' ANSI color in the terminal.
		TerminalAnsiGreen                          string `json:"terminal.ansiGreen"`                          // 'Green' ANSI color in the terminal.
		TerminalAnsiMagenta                        string `json:"terminal.ansiMagenta"`                        // 'Magenta' ANSI color in the terminal.
		TerminalAnsiRed                            string `json:"terminal.ansiRed"`                            // 'Red' ANSI color in the terminal.
		TerminalAnsiWhite                          string `json:"terminal.ansiWhite"`                          // 'White' ANSI color in the terminal.
		TerminalAnsiYellow                         string `json:"terminal.ansiYellow"`                         // 'Yellow' ANSI color in the terminal.
		TerminalSelectionBackground                string `json:"terminal.selectionBackground"`                // The selection background color of the terminal.
		TerminalCursorBackground                   string `json:"terminalCursor.background"`                   // The background color of the terminal cursor. Allows customizing the color of a character overlapped by a block cursor.
		TerminalCursorForeground                   string `json:"terminalCursor.foreground"`                   // The foreground color of the terminal cursor.
		TerminalDropBackground                     string `json:"terminal.dropBackground"`                     // The background color when dragging on top of terminals. The color should have transparency so that the terminal contents can still shine through.
		TerminalTabActiveBorder                    string `json:"terminal.tab.activeBorder"`                   // Border on the side of the terminal tab in the panel. This defaults to tab.activeBorder.
		TerminalCommandDecorationDefaultBackground string `json:"terminalCommandDecoration.defaultBackground"` // The default terminal command decoration background color.
		TerminalCommandDecorationSuccessBackground string `json:"terminalCommandDecoration.successBackground"` // The terminal command decoration background color for successful commands.
		TerminalCommandDecorationErrorBackground   string `json:"terminalCommandDecoration.errorBackground"`   // The terminal command decoration background color for error commands
	} `json:"colors"`
}

type TerminalTheme struct {
	Name                string `json:"name"`
	Background          string `json:"background"`
	Black               string `json:"black"`
	Blue                string `json:"blue"`
	BrightBlack         string `json:"brightBlack"`
	BrightCyan          string `json:"brightCyan"`
	BrightGreen         string `json:"brightGreen"`
	BrightPurple        string `json:"brightPurple"`
	BrightRed           string `json:"brightRed"`
	BrightWhite         string `json:"brightWhite"`
	BrightYellow        string `json:"brightYellow"`
	CursorColor         string `json:"cursorColor"`
	Cyan                string `json:"cyan"`
	Foreground          string `json:"foreground"`
	Green               string `json:"green"`
	Purple              string `json:"purple"`
	Red                 string `json:"red"`
	SelectionBackground string `json:"selectionBackground"`
	White               string `json:"white"`
	Yellow              string `json:"yellow"`
}

func toTerminalTheme(vs *VSCodeTheme) *TerminalTheme {
	result := new(TerminalTheme)

	choose := func(colors ...string) string {
		for _, c := range colors {
			if c != "" {
				return c
			}
		}
		fmt.Println("no color found, using default color")
		return defaultColor
	}

	result.Name = vs.Name
	result.Background = choose(vs.Colors.TerminalBackground)
	result.Black = choose(vs.Colors.TerminalAnsiBlack)
	result.Blue = choose(vs.Colors.TerminalAnsiBlue)
	result.BrightBlack = choose(vs.Colors.TerminalAnsiBrightBlack)
	result.BrightCyan = choose(vs.Colors.TerminalAnsiBrightCyan)
	result.BrightGreen = choose(vs.Colors.TerminalAnsiBrightGreen)
	result.BrightPurple = choose(vs.Colors.TerminalAnsiBrightMagenta)
	result.BrightRed = choose(vs.Colors.TerminalAnsiBrightRed)
	result.BrightWhite = choose(vs.Colors.TerminalAnsiBrightWhite)
	result.BrightYellow = choose(vs.Colors.TerminalAnsiBrightYellow)
	result.CursorColor = choose(vs.Colors.TerminalCursorForeground)
	result.Cyan = choose(vs.Colors.TerminalAnsiCyan)
	result.Foreground = choose(vs.Colors.TerminalForeground)
	result.Green = choose(vs.Colors.TerminalAnsiGreen)
	result.Purple = choose(vs.Colors.TerminalAnsiMagenta)
	result.Red = choose(vs.Colors.TerminalAnsiRed)
	result.SelectionBackground = choose(vs.Colors.TerminalSelectionBackground)
	result.White = choose(vs.Colors.TerminalAnsiWhite)
	result.Yellow = choose(vs.Colors.TerminalAnsiYellow)

	return result
}

func VSCodeToTerminal(vscode []byte) (*TerminalTheme, error) {
	vs := new(VSCodeTheme)
	err := json.Unmarshal(vscode, vs)
	if err != nil {
		return nil, err
	}

	return toTerminalTheme(vs), nil
}

func printUsageAndExit() {
	w := os.Stdout
	fmt.Fprintf(w, "terco converts Visual Studio Code color theme to Windows Terminal color theme\n\n")
	fmt.Fprintf(w, "USAGE:\n")
	fmt.Fprintf(w, "\tterco <vscode-theme-file>\n\n")
	os.Exit(1)
}

func main() {
	argc := len(os.Args)
	if argc != 2 {
		printUsageAndExit()
	}

	inBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	theme, err := VSCodeToTerminal(inBytes)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	outBytes, err := json.MarshalIndent(theme, "", "  ")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Fprintln(os.Stdout, string(outBytes))
}
