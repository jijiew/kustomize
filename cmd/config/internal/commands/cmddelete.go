package commands

import (
	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/cmd/config/ext"
	"sigs.k8s.io/kustomize/cmd/config/internal/generateddocs/commands"
	"sigs.k8s.io/kustomize/kyaml/openapi"
	"sigs.k8s.io/kustomize/kyaml/setters2/settersutil"
)

// NewDeleteRunner returns a command runner.
func NewDeleteRunner(parent string) *DeleteRunner {
	r := &DeleteRunner{}
	c := &cobra.Command{
		Use:     "delete DIR NAME",
		Args:    cobra.MinimumNArgs(2),
		Short:   commands.SetShort,
		Long:    commands.SetLong,
		Example: commands.SetExamples,
		PreRunE: r.preRunE,
		RunE:    r.runE,
	}
	fixDocs(parent, c)
	r.Command = c

	return r
}


func DeleteCommand(parent string) *cobra.Command {
	return NewDeleteRunner(parent).Command
}

type DeleteRunner struct {
	Command     *cobra.Command
	DeleteSetter settersutil.DeleterCreator
	OpenAPIFile string
}

func (r *DeleteRunner) preRunE(c *cobra.Command, args []string) error {
	var err error
	r.DeleteSetter.Name = args[1]

	r.OpenAPIFile, err = ext.GetOpenAPIFile(args)
	if err != nil {
		return err
	}

	if err := openapi.AddSchemaFromFile(r.OpenAPIFile); err != nil {
		return err
	}

	return nil
}

func (r *DeleteRunner) runE(c *cobra.Command, args []string) error {
	return handleError(c, r.delete(c, args))
}

func (r *DeleteRunner) delete(c *cobra.Command, args []string) error {
	return r.DeleteSetter.Delete(r.OpenAPIFile, args[0])
}