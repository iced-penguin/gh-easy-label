package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type LabelSet struct {
	name   string
	labels []Label
}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List defined label sets",
		Long:  "List defined label sets",
		RunE: func(cmd *cobra.Command, args []string) error {
			labelSetsMap, err := getLabelSets()
			if err != nil {
				return fmt.Errorf("failed to get label sets: %v", err)
			}
			if len(labelSetsMap) == 0 {
				return nil
			}

			// NOTE: 表示順を固定するためにスライスにしてソート
			labelSets := []LabelSet{}
			for name, labels := range labelSetsMap {
				labelSets = append(labelSets, LabelSet{name, labels})
			}
			sort.Slice(labelSets, func(i, j int) bool {
				return labelSets[i].name < labelSets[j].name
			})

			printLabelSets(labelSets)
			return nil
		},
	}
	return cmd
}

func printLabelSets(labelSets []LabelSet) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Label Set", "Label", "Description", "Color"})
	for _, labelSet := range labelSets {
		for _, label := range labelSet.labels {
			table.Append([]string{labelSet.name, label.Name, label.Description, label.Color})
		}
	}
	table.SetAutoMergeCells(true)
	table.Render()
}
