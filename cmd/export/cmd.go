package export

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	flagOutputDir string
	flagDatadir   string
)

var Cmd = &cobra.Command{
	Use:   "export",
	Short: "exports all the data into json files",
	Run:   run,
}

func init() {
	Cmd.Flags().StringVar(&flagDatadir, "badger-db-directory", "",
		"directory that stores the protocol state")
	_ = Cmd.MarkFlagRequired("datadir")

	Cmd.Flags().StringVar(&flagOutputDir, "output-dir", "",
		"Directory to write new Execution State to")
	_ = Cmd.MarkFlagRequired("output-dir")
}

func run(*cobra.Command, []string) {

	log.Info().Msg("start exporting blocks")
	blockID, err := ExportBlocks(flagDatadir, flagOutputDir)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get export blocks")
	}

	log.Info().Msg("start exporting events")
	err = ExportEvents(blockID, flagDatadir, flagOutputDir)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get export events")
	}

	log.Info().Msg("start exporting transactions")
	err = ExportExecutedTransactions(blockID, flagDatadir, flagOutputDir)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get export transactions")
	}
}
