package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/kosmos.io/netdoctor/cmd/floater/app/options"
)

func NewFloaterCommand(ctx context.Context) *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Use:  "netdr-floater",
		Long: `Environment for executing commands`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := Run(ctx, opts); err != nil {
				return err
			}
			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	return cmd
}

func Run(_ context.Context, _ *options.Options) error {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8889"
	}
	fmt.Print("PORT: ", port)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			fmt.Print(fmt.Errorf("response writer error: %s", err))
		}
	})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Print(fmt.Errorf("launch server error: %s", err))
		panic(err)
	}

	return nil
}
