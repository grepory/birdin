// Copyright © 2016 Greg Poirier <greg.istehbest@gmail.com>
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

package cmd

import (
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/grepory/birdin/aviary"
	"github.com/grepory/birdin/birds"
	"github.com/grepory/birdin/birds/nms"
	"github.com/grepory/birdin/scheduler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func buildAviary() (a *aviary.Aviary, err error) {
	nmsBirdScheduler := &aviary.BirdScheduler{
		Bird: &nms.Bird{
			Tweeter: birds.NewAnaconda(viper.GetString("nms-access-token"), viper.GetString("nms-secret-token")),
		},
		Scheduler: &scheduler.TickerScheduler{
			Duration: 12 * time.Hour,
		},
	}

	a = aviary.New(nmsBirdScheduler)
	return a, nil
}

// aviaryCmd represents the aviary command
var aviaryCmd = &cobra.Command{
	Use:   "aviary",
	Short: "Start an aviary tweet scheduler",
	Long:  `Running the aviary will start a Tweet scheduler using all of the birds in the aviary.`,
	Run: func(cmd *cobra.Command, args []string) {
		anaconda.SetConsumerKey(viper.GetString("consumer-key"))
		anaconda.SetConsumerSecret(viper.GetString("consumer-secret"))
		a, err := buildAviary()
		if err != nil {
			panic(err)
		}

		a.Tweet()
	},
}

func init() {
	RootCmd.AddCommand(aviaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aviaryCmd.PersistentFlags().String("foo", "", "A help for foo")
	aviaryCmd.PersistentFlags().String("nms-access-token", "", "Twitter API Access Token")
	aviaryCmd.PersistentFlags().String("nms-secret-token", "", "Twitter API Access Secret Token")
	viper.BindPFlag("nms-access-token", aviaryCmd.PersistentFlags().Lookup("nms-access-token"))
	viper.BindPFlag("nms-secret-token", aviaryCmd.PersistentFlags().Lookup("nms-secret-token"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aviaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
