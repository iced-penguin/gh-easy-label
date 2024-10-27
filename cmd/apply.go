package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cli/go-gh/v2"
	"github.com/spf13/cobra"
)

func newApplyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply [label_set]",
		Short: "Apply label set",
		Long:  "Replace all existing labels with the labels from the specified label set.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			labelSetName := args[0]
			labelSets, err := getLabelSets()
			if err != nil {
				return fmt.Errorf("failed to get label sets: %v", err)
			}
			labelSet, ok := labelSets[labelSetName]
			if !ok {
				return fmt.Errorf("not found label set: %s", labelSetName)
			}

			// 既存のラベルを全て削除
			existingLabelNames, err := fetchExistingLableNames()
			if err != nil {
				return fmt.Errorf("failed to fetch files: %v", err)
			}
			for _, labelName := range existingLabelNames {
				if err := deleteLabel(labelName); err != nil {
					return fmt.Errorf("failed to delete label: %v", err)
				}
			}
			// ラベルをリポジトリに登録する
			for _, label := range labelSet {
				args := []string{"label", "create", label.Name, "-c", label.Color, "-d", label.Description, "-f"}
				if _, err := execGH(args); err != nil {
					return fmt.Errorf("failed to create label: %v", err)
				}
			}
			return nil
		},
	}
	return cmd
}

func execGH(args []string) (bytes.Buffer, error) {
	stdout, stderr, _ := gh.Exec(args...)
	if stderr.Len() > 0 {
		return bytes.Buffer{}, fmt.Errorf("%s", stderr.String())
	}
	return stdout, nil
}

func fetchExistingLableNames() ([]string, error) {
	args := []string{"label", "list", "--json", "name"}
	res, err := execGH(args)
	if err != nil {
		return nil, err
	}
	var labels []struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(res.Bytes(), &labels); err != nil {
		return nil, err
	}
	var labelNames []string
	for _, label := range labels {
		labelNames = append(labelNames, label.Name)
	}
	return labelNames, nil
}

func deleteLabel(name string) error {
	args := []string{"label", "delete", name, "--yes"}
	_, err := execGH(args)
	if err != nil {
		return err
	}
	return nil
}
