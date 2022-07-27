package cmd

import (
	"context"
	"fmt"

	consumerHandler "github.com/ilmimris/coinbit-test/internal/adapter/inbound/consumer/handler"
	"github.com/ilmimris/coinbit-test/internal/adapter/outbound/producer"
	"github.com/ilmimris/coinbit-test/internal/adapter/outbound/viewtable"
	"github.com/lovoo/goka"
	log "github.com/sirupsen/logrus"

	"github.com/ilmimris/coinbit-test/cmd/bootstrap"
	"github.com/ilmimris/coinbit-test/config"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/consumer"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest"
	"github.com/ilmimris/coinbit-test/internal/core/codec"
	"github.com/spf13/cobra"
)

var bst bootstrap.Bootstrapper

var rootCommand = &cobra.Command{
	Use: "account-svc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("root command")
	},
}

var (
	tmc            *goka.TopicManagerConfig
	depositTopic   string = "deposits"
	balanceGroup   string = "balance"
	thresholdGroup string = "aboveThreshold"
)

func Run() {
	rootCommand.Execute()

	var cfgFile string
	cobra.OnInitialize(func() {
		// init bootstrap
		bst = bootstrap.New()
		bst.Initialize(cfgFile)
		tmc = bst.GetTopicManagerConfig()
		bst.AddServices(initServices()...)
		bst.RunServices()
	})
	rootCommand.PersistentFlags().StringVar(&cfgFile, "config", "config.json", "config file (default is config.json)")

	rootCommand.AddCommand(restCommand)
	if err := rootCommand.Execute(); err != nil {
		log.WithContext(context.Background()).Panicf("%v", err)
	}
}

func initServices() (services []bootstrap.OptionService) {
	services = make([]bootstrap.OptionService, 0)

	var (
		// init codec for encode and decode
		depositWalletCodec = codec.NewDepositCodec()
		walletCodec        = codec.NewWalletCodec()
		thresholdCodec     = codec.NewThresholdCodec()
	)

	// init pub sub event
	depositWalletEventHandler := consumerHandler.NewDepositWalletEventHandler(bst.GetRegistryConsumer())
	processThresholdEventHandler := consumerHandler.NewProcessThresholdEventHandler(bst.GetRegistryConsumer())

	// deposit consumer balance group
	services[0] = bootstrap.NewGokaConsumerGroup(consumer.GokaConsumerGroupOpt{
		Brokers:            config.GlobalConfig.Kafka.Brokers,
		Group:              balanceGroup,
		Topic:              depositTopic,
		Handler:            depositWalletEventHandler,
		TopicManagerConfig: tmc,
		InputCodec:         depositWalletCodec,
		TableCodec:         walletCodec,
	})

	// process threshold consumer balance group
	services[1] = bootstrap.NewGokaConsumerGroup(consumer.GokaConsumerGroupOpt{
		Brokers:            config.GlobalConfig.Kafka.Brokers,
		Group:              balanceGroup,
		Topic:              depositTopic,
		Handler:            processThresholdEventHandler,
		TopicManagerConfig: tmc,
		InputCodec:         depositWalletCodec,
		TableCodec:         thresholdCodec,
	})

	services[2] = bootstrap.NewGokaProducer(producer.GokaProducerOpt{
		Brokers: config.GlobalConfig.Kafka.Brokers,
		Topic:   depositTopic,
		Codec:   depositWalletCodec,
	})

	services[3] = bootstrap.NewGokaViewTable(viewtable.GokaViewTableOpt{
		Brokers: config.GlobalConfig.Kafka.Brokers,
		Group:   balanceGroup,
		Codec:   walletCodec,
	})

	services[4] = bootstrap.NewGokaViewTable(viewtable.GokaViewTableOpt{
		Brokers: config.GlobalConfig.Kafka.Brokers,
		Group:   thresholdGroup,
		Codec:   thresholdCodec,
	})

	services[5] = bootstrap.NewServiceRest(rest.Options{
		Port: fmt.Sprintf(":%d", config.GlobalConfig.Rest.Port),
	})
	return
}
