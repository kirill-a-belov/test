package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"awesomeProject1/test/models"
	"awesomeProject1/test/storages"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Config struct {
	InputFilePath string
}

var config Config

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze events overlap",
	Run: func(cmd *cobra.Command, args []string) {

		inputFile, err := os.Open(config.InputFilePath)
		if err != nil {
			log.Fatalf("error: open input file - %v", err)
		}

		stat, err := inputFile.Stat()
		if err != nil {
			log.Fatalf("error: input file bad statistic - %v", err)
		}
		if stat.Size() == 0 {
			log.Fatalf("error: input file empty - %v", err)
		}

		events := []*models.Event{}
		eventsStorage := storages.PrefixTree{}

		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := scanner.Text()
			if err := scanner.Err(); err != nil {
				log.Println("read line from file error", line)
				continue
			}

			event := &models.Event{}
			if err := event.FillFromCSVString(line); err != nil {
				log.Println("fill from line error", err, line)
				continue
			}

			events = append(events, event)
			eventsStorage.Add(event)

		}
		inputFile.Close()

		fmt.Println()
		fmt.Println("Simple solution with higher complexity:")
		count := 0
		for i := 0; i < len(events); i++ {
			for j := i + 1; j < len(events); j++ {
				if strings.Compare(events[i].EndTime, events[j].StartTime) >= 0 {
					fmt.Println(events[i], events[j])
					count++
				}
			}
		}
		fmt.Println("Total overlapsed pairs: ", count)

		fmt.Println()
		fmt.Println("Complicated solution with lower complexity:")
		count = 0
		eventNames := make(map[string]bool)
		for _, e := range events {
			eventNames[e.Name] = true
			for _, r := range eventsStorage.FindBefore(e.EndTime) {
				if !eventNames[r.Name] {
					fmt.Println(e, r)
					count++
				}
			}
		}
		fmt.Println("Total overlapsed pairs: ", count)

	},
}

func init() {
	// init config
	RootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().AddFlagSet(config.Flags())
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("TempAnalyzerConfig", pflag.PanicOnError)
	f.StringVar(&c.InputFilePath,
		"input_file",
		"./events.csv",
		"Path for input file")

	return f
}
