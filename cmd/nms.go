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
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/grepory/birdin/birds"
	"github.com/grepory/birdin/birds/nms"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// nmsCmd represents the nms command
var nmsCmd = &cobra.Command{
	Use:   "nms",
	Short: "No One's Sky animal tweeter",
	Long: `Running the nms command will Tweet a random No One's Sky animal description. For example:

λ ./birdin nms
Using config file: /Users/grepory/.birdin.yaml
Tweeting:
Age: eternal
Gender: alpha
Temperament: grumpy`,
	Run: func(cmd *cobra.Command, args []string) {
		anaconda.SetConsumerKey(viper.GetString("consumer-key"))
		anaconda.SetConsumerSecret(viper.GetString("consumer-secret"))
		bird := nms.Bird{
			Tweeter: birds.NewAnaconda(viper.GetString("nms-access-token"), viper.GetString("nms-access-token-secret")),
		}

		status, err := bird.Tweet()
		fmt.Printf("Tweeting:\n%s\n", status)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(nmsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nmsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nmsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
