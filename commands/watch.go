// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"flag"
	"fmt"
	"github.com/gofunct/common/pkg/logger/zap"
	"github.com/gokyle/fswatch"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	watchAuto bool
	watchDur  string
)

func init() {
	watchCmd.Flags().StringVar(&watchDur, "dur", "15s", "how often to watch for fs changes")
	watchCmd.Flags().BoolVar(&watchAuto, "auto", false, "auto add new files in directories")
}

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "watch filesystem changes",
	Run: func(cmd *cobra.Command, args []string) {
		var w *fswatch.Watcher
		paths := flag.Args()
		if watchAuto {
			w = fswatch.NewAutoWatcher(paths...)
		} else {
			w = fswatch.NewWatcher(paths...)
		}
		fmt.Println("[+] listening...")

		l := w.Start()
		go func() {
			for {
				n, ok := <-l
				if !ok {
					return
				}
				var status_text string
				switch n.Event {
				case fswatch.CREATED:
					status_text = "was created"
				case fswatch.DELETED:
					status_text = "was deleted"
				case fswatch.MODIFIED:
					status_text = "was modified"
				case fswatch.PERM:
					status_text = "permissions changed"
				case fswatch.NOEXIST:
					status_text = "doesn't exist"
				case fswatch.NOPERM:
					status_text = "has invalid permissions"
				case fswatch.INVALID:
					status_text = "is invalid"
				}
				fmt.Printf("[+] %s %s\n", n.Path, status_text)
			}
		}()
		go func() {
			dur, err := time.ParseDuration(watchDur)
			if err != nil {
				zap.LogE("parse flag duration", err)
				dur = 15 * time.Second
			}

			for {
				<-time.After(dur)
				if !w.Active() {
					fmt.Println("[!] not watching anything")
					os.Exit(1)
				}
				fmt.Printf("[-] watching: %+v\n", w.State())
			}
		}()
		time.Sleep(60 * time.Second)
		fmt.Println("[+] stopping...")
		w.Stop()
		time.Sleep(5 * time.Second)
	},
}
