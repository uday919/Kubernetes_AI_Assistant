package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)
var (
	version = "dev"
	openAIAPIKey = flag.String("openai-api-key", env.GetOr("OPEN_API_KEY",env.String), "OpenAI API key for the OpenAI service. This is required" )
	debug = flag.Bool("debug",env,GetOr("DEBUG",strconv.ParseBool,false), "Enable debug logging")
)
func InitAndExecute() {
	if *openAIAPIKey == "" {
		fmt.Println("OpenAI API key is required")
		os.Exit(1)
	}
	if err := RootCmd().Execute(); err != nil {
		os.exit(1)
	}
}
func RootCmd() *cobra.Command {
	cmd := &cobra.command{
		Use:"Kubernetes-ai-assistant",
		Short:"Kubernetes-ai-assistant",
		Long:"Kubernetes-ai-assistant is a plugin for kubectl (command line tool for kubernetes) gives you the power of open AI API",
		Version: version,
		SilenceUsage:true,
		PersistentPreRun: func(cmd *cobra.Command,args []string){
			if *debug{
				log.SetLevel(log.DebugLevel)
				printDebugFlags()
			}

		},
		RunE:func(_ *cobra.Command, args []string) error{

		}

	}
	return cmd
}
func printDebugFlags(){
	log.Debugf("openai-endpoint: %s", *openAIEndpoint)
	log.Debugf("openai-deployment-name: %s", *openAIDeploymentName)
	log.Debugf("azure-openai-map: %s", *azureModelMap)
	log.Debugf("temperature: %f", *temperature)
	log.Debugf("use-k8s-api:%t",*usek8sAPI)
	log.Debugf("k8s-openapi-url:%s",*k8sOpenAPIURL)
}