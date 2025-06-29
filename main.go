package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	re "regexp"
	"time"

	"github.com/urfave/cli/v3"
)

func initialize() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dirpath := filepath.Join(home_dir, ".lifelog")
	// config := filepath.Join(dirpath,"config.json")
	current_data := filepath.Join(dirpath, "lifelog.txt")
	err = os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return current_data, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	filepath, err := initialize()
	check(err)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	check(err)
	defer file.Close()
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:      "log",
				UsageText: "Add a log to the file",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var out string
					for i := 0; i < cmd.Args().Len(); i++ {
						out = out + fmt.Sprintf(" %v", cmd.Args().Get(i))
					}
					current_time := time.Now().Local()
					_, err := fmt.Fprintf(file, "%s:%s\n", current_time.Format("2006-01-02 15:04:05"), out)
					check(err)
					err = file.Sync()
					check(err)
					return nil
				},
			},
			{
				Name: "show",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					filter := false
					if cmd.Args().Len() > 0 {
						fmt.Println(cmd.Args().Get(0))
						filter = true
					}
					scanner := bufio.NewScanner(file)
					buf := make([]byte, 1024)
					for {
						n, err := file.Read(buf)
						if err == io.EOF {
							break
						}
						check(err)
						if n > 0 {
							if filter {
								filter_text := fmt.Sprintf("@%v", cmd.Args().Get(0))
								fmt.Println(filter_text)
								match, err := re.Match(filter_text, buf[:n])
								check(err)
								if match {
									fmt.Println(string(buf[:n]))
								} else {
									fmt.Println("Nope")
								}
							} else {
								fmt.Println(string(buf[:n]))
							}
						}
					}
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// func search_tags(tag string) string {

// }
