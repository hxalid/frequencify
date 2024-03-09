/*
Copyright © 2019 Khalid Hasanov <xalid.h@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	freq "github.com/hxalid/frequencify/pkg/frequencify"
	"github.com/spf13/cobra"
)

const (
	baseURL = "https://en.wikipedia.org/w/api.php"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "frequencify",
	Short: "A Wikipedia web crawler",
	Long:  `Frequencify crawls Wikipedia by page id and prints top n words used on that page`,
	Run:   frequencify,
}

func frequencify(cmd *cobra.Command, args []string) {
	valid := true
	count, err := cmd.Flags().GetUint32("number")
	if err != nil || count == 0 {
		if err != nil {
			fmt.Printf("Count number should be a positive integer. %v\n", err)
		} else {
			fmt.Println("Count number should be a positive integer.")
		}
		valid = false
	}
	pid, err := cmd.Flags().GetUint32("pageid")
	if err != nil || pid == 0 {
		if err != nil {
			fmt.Printf("Page id should be a positive integer. %v\n", err)
		} else {
			fmt.Println("Page id should be a positive integer.")
		}
		valid = false
	}
	if !valid {
		os.Exit(1)
	}
	url := fmt.Sprintf("%s?action=query&prop=extracts&pageids=%d&explaintext&format=json",
		baseURL, pid)
	c := freq.NewClient(5*time.Second, url)
	c.Frequencify(count, pid)
}

func init() {
	rootCmd.Flags().Uint32P("number", "n", 1, "Top count - a positive integer")
	rootCmd.Flags().Uint32P("pageid", "p", 1, "Page id - a positive integer")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
