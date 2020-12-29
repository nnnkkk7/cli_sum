package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func sumCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sum",
		Short: "sum calcu",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			itemOne, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
				return nil
			}

			itemTwo, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
				return nil
			}

			if len(args) == 3 {
				itemThree, err := strconv.Atoi(args[2])
				if err != nil {
					log.Fatal(err)
					return nil
				}
				fmt.Println(itemOne + itemTwo + itemThree)
			} else {
				fmt.Println(itemOne + itemTwo)
			}

			return nil
		},
	}

	return cmd
}
