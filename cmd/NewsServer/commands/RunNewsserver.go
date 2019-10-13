package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	models "github.com/jpramirez/go-newsapi/pkg/models"

	utils "github.com/jpramirez/go-newsapi/pkg/utils"

	feed "github.com/jpramirez/go-newsapi/pkg/api"
	constants "github.com/jpramirez/go-newsapi/pkg/constants"
)

var rootCmd = &cobra.Command{
	Use:   "NewsAPI",
	Short: "Client for newsapi.org notification Service",
	Long:  `A Fast and Flexible on push platform`,
	RunE:  runWebServer,
}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var config models.Config
var cfgFile string
var projectBase string
var monitorMode string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/config.json)")
	rootCmd.PersistentFlags().Bool("default", true, "Use default configuration")
}

func runWebServer(cmd *cobra.Command, args []string) error {
	config, err := utils.LoadConfiguration(cfgFile)
	if err != nil {
		log.Fatalln("Error and Exiting")
	}
	f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	fmt.Println(config.LogFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	_feed, err := feed.NewFeedAPI(config)

	//Testing Feed
	var parameters models.NewsQueryParameter
	parameters.Category = "technology"
	parameters.PageSize = 10
	parameters.Language = "en"
	ret, err := _feed.GetFeed(constants.TopHeadLinesEndpoint, parameters)

	fmt.Println(ret.TotalResults)
	fmt.Println(ret.Status)
	fmt.Println(ret.Articles)

	//Testing Sources
	var sourcePrameters models.NewsQueryParameter
	sourcePrameters.Category = "technology"
	retSources, err := _feed.GetSources(sourcePrameters)

	fmt.Println(retSources.Status)
	fmt.Println(retSources.Sources)

	return err
}
