// Copyright © 2016 Yieldbot <devops@yieldbot.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/yieldbot/sensuplugin/sensuutil"
)

var warnThreshold int
var critThreshold int
var checkKey string

var checkChronyCmd = &cobra.Command{
	Use:   "checkChrony",
	Short: "Check various values in chrony to ensure all is well",
	Long: `This will use 'chronyc tracking' to build a map of keys allowing the
  user to check against any of the values to ensure they are within tolerated
  limits for their environment.`,
	Run: func(cmd *cobra.Command, args []string) {

		chronyCheck := exec.Command("chronyc", "tracking")

		out, err := chronyCheck.Output()
		if err != nil {
			sensuutil.EHndlr(err)
		}

		chronyCheck.Start()
		data := createMap(string(out))

		if debug {
			for k, v := range data {
				fmt.Println("Key: ", k, "Current value: ", v)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(checkChronyCmd)

	checkChronyCmd.Flags().IntVarP(&warnThreshold, "warn", "", 4, "the alert warning threshold")
	checkChronyCmd.Flags().IntVarP(&critThreshold, "crit", "", 8, "the alert critical threshold")
	checkChronyCmd.Flags().StringVarP(&checkKey, "checkKey", "", "", "the key to check")

}