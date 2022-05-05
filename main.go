package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

// @TODO: Check the color output of vscode and terminal. Color space conversion needed?

type VSCodeTheme struct {
	Name   string `json:"name"`
	Colors struct {
		EditorForeground                           string `json:"editor.foreground"`                           // Editor background color.
		EditorBackground                           string `json:"editor.background"`                           // Editor default foreground color.
		EditorCursorForeground                     string `json:"editorCursor.foreground"`                     // Color of the editor cursor.
		EditorSelectionBackground                  string `json:"editor.selectionBackground"`                  // Color of the editor selection.
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
	BrightBlue          string `json:"brightBlue"`
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

func getFieldTag(theme *TerminalTheme, field any, tag string) string {
	structValue := reflect.ValueOf(theme).Elem()
	fieldValue := reflect.ValueOf(field).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		value := structValue.Field(i)
		if value.Addr().Interface() == fieldValue.Addr().Interface() {
			return structValue.Type().Field(i).Tag.Get(tag)
		}
	}

	return ""
}

func toTerminalTheme(vs *VSCodeTheme) *TerminalTheme {
	result := new(TerminalTheme)

	chooseColor := func(dest *string, defaultColor string, colors ...string) {
		for _, c := range colors {
			if c != "" {
				*dest = c
				return
			}
		}

		tag := getFieldTag(result, dest, "json")
		fmt.Fprintf(os.Stderr, "no color for %q using default %q\n", tag, defaultColor)

		*dest = defaultColor
	}

	c := &vs.Colors

	result.Name = vs.Name
	chooseColor(&result.Background, "#ff0000", c.TerminalBackground, c.EditorBackground)
	chooseColor(&result.Black, "#000000", c.TerminalAnsiBlack)
	chooseColor(&result.Blue, "#6182b8", c.TerminalAnsiBlue)
	chooseColor(&result.BrightBlack, "#90a4ae", c.TerminalAnsiBrightBlack)
	chooseColor(&result.BrightBlue, "#6182b8", c.TerminalAnsiBrightBlue)
	chooseColor(&result.BrightCyan, "#39adb5", c.TerminalAnsiBrightCyan)
	chooseColor(&result.BrightGreen, "#91b859", c.TerminalAnsiBrightGreen)
	chooseColor(&result.BrightPurple, "#7c4dff", c.TerminalAnsiBrightMagenta)
	chooseColor(&result.BrightRed, "#e53935", c.TerminalAnsiBrightRed)
	chooseColor(&result.BrightWhite, "#ffffff", c.TerminalAnsiBrightWhite)
	chooseColor(&result.BrightYellow, "#ffb62c", c.TerminalAnsiBrightYellow)
	chooseColor(&result.CursorColor, "#ff0000", c.TerminalCursorForeground, c.EditorCursorForeground)
	chooseColor(&result.Cyan, "#39adb5", c.TerminalAnsiCyan)
	chooseColor(&result.Foreground, "#ff0000", c.TerminalForeground, c.EditorForeground)
	chooseColor(&result.Green, "#91b859", c.TerminalAnsiGreen)
	chooseColor(&result.Purple, "#7c4dff", c.TerminalAnsiMagenta)
	chooseColor(&result.Red, "#e53935", c.TerminalAnsiRed)
	chooseColor(&result.SelectionBackground, "#ff0000", c.TerminalSelectionBackground, c.EditorSelectionBackground)
	chooseColor(&result.White, "#ffffff", c.TerminalAnsiWhite)
	chooseColor(&result.Yellow, "#ffb62c", c.TerminalAnsiYellow)

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
