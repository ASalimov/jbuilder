package cmd

import (
	"fmt"
	"github.com/ASalimov/jbuilder/cmd/jb"
	"github.com/spf13/cobra"
)

func init() {
	useCmd := &cobra.Command{
		Use:                   "use ENV_NAME",
		DisableFlagsInUseLine: true,
		Short:                 "Makes jenkins environment by default",
		Run: func(cmd *cobra.Command, args []string) {
			use(cmd)
		},
		Args: cobra.ExactArgs(1),
	}
	rootCmd.AddCommand(useCmd)
}

func use(cmd *cobra.Command) {
	jb.Init(cmd.Flags().Args()[0])
	jb.SetDef(cmd.Flags().Args()[0])
	fmt.Println(cmd.Flags().Args()[0] + " have been set by default")
}
