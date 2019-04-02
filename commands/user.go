package commands

import (
	"errors"
	"fmt"

	"github.com/MichaelMure/git-bug/cache"
	"github.com/MichaelMure/git-bug/util/interrupt"
	"github.com/spf13/cobra"
)

func runUser(cmd *cobra.Command, args []string) error {
	backend, err := cache.NewRepoCache(repo)
	if err != nil {
		return err
	}
	defer backend.Close()
	interrupt.RegisterCleaner(backend.Close)

	if len(args) > 1 {
		return errors.New("only one identity can be displayed at a time")
	}

	var id *cache.IdentityCache
	if len(args) == 1 {
		id, err = backend.ResolveIdentityPrefix(args[0])
	} else {
		id, err = backend.GetUserIdentity()
	}

	if err != nil {
		return err
	}

	fmt.Printf("Id: %s\n", id.Id())
	fmt.Printf("Name: %s\n", id.Name())
	fmt.Printf("Login: %s\n", id.Login())
	fmt.Printf("Email: %s\n", id.Email())
	fmt.Printf("Last modification: %s (lamport %d)\n",
		id.LastModification().Time().Format("Mon Jan 2 15:04:05 2006 +0200"),
		id.LastModificationLamport())
	fmt.Println("Metadata:")
	for key, value := range id.ImmutableMetadata() {
		fmt.Printf("    %s --> %s\n", key, value)
	}
	// fmt.Printf("Protected: %v\n", id.IsProtected())

	return nil
}

var userCmd = &cobra.Command{
	Use:     "user [<user-id>]",
	Short:   "Display or change the user identity.",
	PreRunE: loadRepo,
	RunE:    runUser,
}

func init() {
	RootCmd.AddCommand(userCmd)
	userCmd.Flags().SortFlags = false
}