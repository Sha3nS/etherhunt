package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/etherhunt/config"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// StartCMD start by command for init params
func InvokeCMD() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "invoke",
		Short: "evm invoke",
		Long:  "evm invoke",
		RunE: func(cmd *cobra.Command, _ []string) error {
			err := config.InitConfig()
			if err != nil {
				fmt.Printf("init config error: %s\n", err)
			}
			invoke()
			fmt.Println("supervisor stopped")
			return nil
		},
	}
	attachFlags(startCmd, []string{flagPrivateKey, flagContract, flagMethod, flagArgs, flagAfterBlockHeight})

	return startCmd
}

// start this is real start function
func invoke() {
	// init server
	ethServer := server.NewEvmClient()
	if err := ethServer.Start(); err != nil {
		fmt.Errorf("server start failed, err: %s", err.Error())
		return
	}
	// new an error channel to receive errors
	errorC := make(chan error, 1)

	// handle exit signal in separate go routines
	go handleExitSignal(errorC)

	// listen error signal in main function
	select {
	case err := <-errorC:
		if err != nil {
			_ = fmt.Errorf("Service terminated")
			return
		}
		err = proxyServer.Stop()
		if err != nil {
			cliLog.Error("Stop err: ", err)
		}
		cliLog.Info("All is stopped!")
	}
}

// handleExitSignal listen exit signal for process stop
func handleExitSignal(exitC chan<- error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	defer signal.Stop(signalChan)

	for sig := range signalChan {
		cliLog.Infof("received exit signal: %d (%s)", sig, sig)
		exitC <- nil
	}
}

// migrate
func migrate() error {
	sqlConfig := conf.Config.StorageConfig
	if sqlConfig == nil {
		cliLog.Errorf("bad config, check your sql configuration")
		return fmt.Errorf("bad config, check your sql configuration")
	}

	db, err := gorm.Open(mysql.Open(sqlConfig.Dsn), &gorm.Config{
		Logger: gorm_logger.Default.LogMode(gorm_logger.Info),
	})
	if err != nil {
		cliLog.Errorf("failed to connect database")
		return fmt.Errorf("failed to connect database")
	}
	err = schema.Migrate(db)
	if err != nil {
		cliLog.Errorf("failed to migrate database")
		return fmt.Errorf("failed to migrate database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		cliLog.Errorf("failed to get sqlDB")
		return fmt.Errorf("failed to get sqlDB")
	}
	return sqlDB.Close()
}

func printLogo() {
	log := logger.GetLogger(logger.ModuleCli)
	log.Infof(logo())
}

func logo() string {
	fig := figure.NewFigure("TBiS-Gateway", "slant", true)
	s := fig.String()
	fragment := "================================================================================================="
	versionInfo := "::TBiS-Gateway::  version(" + conf.CurrentVersion + ")"
	return fmt.Sprintf("\n%s\n%s%s\n%s\n", fragment, s, fragment, versionInfo)
}