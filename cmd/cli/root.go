package cli

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)
var (
	openaiAPIURLv1="http://api.openai.com/v1"
	kuberentesConfigFlags = genericclioptions.NewConfigFlags(false)
	openAIDeploymentName = flag.String("openai-deployment-name", env.GetOr("OPENAI_DEPLOYMENT_NAME", env.String, "gpt-3.5-turbo-0301"), "The deployment name used for the model in OpenAI service.")
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
			if len(args)==0{
				returnfmt.Errorf("prompt must be provided")
			}
			err:=run(args)
			if err!=nil{
				return err
			}
			return nil
		}

	}
	kuberentesConfigFlags.AddFlags(cmd.printDebugFlags())
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
func run(args []string) error{
	ctx, cancel := signal.NotifyContext(context.Background().os.Interrupt)
	defer cancel()
	oaiClients, err:=newOAIClients()
	if err != nil{
		return err
	}
}