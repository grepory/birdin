// Package birds houses all of our Birds (tweet generators). A Bird uses a
// Tweeter to post a generated Twitter status message. This example uses
// the anaconda Twitter API client to create a Tweeter for the nms
// Bird.
//
// TODO(greg): move this over to a test package and add an example.
//
// anaconda.SetConsumerKey(viper.GetString("consumer-key"))
// anaconda.SetConsumerSecret(viper.GetString("consumer-secret"))
// bird := nms.Bird{
//   Tweeter: birds.NewAnaconda(viper.GetString("nms-access-token"), viper.GetString("nms-secret-token")),
// }
// status, err := bird.Tweet()
// if err != nil {
//   panic(err)
// }
// fmt.Printf("Tweeting:\n%s\n", status)
package birds
