package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit labels",
	Long:  `Edit labels with the default editor.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, err := getConfigFilename()
		if err != nil {
			return err
		}
		editor, err := selectEditor()
		if err != nil {
			return err
		}
		if err := openEditor(editor, filename); err != nil {
			return err
		}
		return nil
	},
}

func selectEditor() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = os.Getenv("VISUAL")
	}
	if editor == "" {
		// ユーザーにエディタを選択させる
		fmt.Println("The environment variable EDITOR or VISUAL is not set. Select an editor:")
		fmt.Println("1. vim")
		fmt.Println("2. nano")
		fmt.Println("3. code")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			editor = "vim"
		case 2:
			editor = "nano"
		case 3:
			editor = "code"
		default:
			return "", fmt.Errorf("invalid selection")
		}
	}
	return editor, nil
}

func openEditor(editor, filename string) error {
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
