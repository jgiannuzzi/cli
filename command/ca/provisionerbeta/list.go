package provisionerbeta

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/smallstep/certificates/pki"
	"github.com/smallstep/cli/errs"
	"github.com/smallstep/cli/flags"
	"github.com/urfave/cli"
)

// nolint
func listCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Action: cli.ActionFunc(listAction),
		Usage:  "list provisioners configured in the CA",
		UsageText: `**step beta ca provisioner list**
[**--ca-url**=<uri>] [**--root**=<file>]`,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "ca-url",
				Usage: "<URI> of the targeted Step Certificate Authority.",
			},
			cli.StringFlag{
				Name:  "root",
				Usage: "The path to the PEM <file> used as the root certificate authority.",
			},
		},
		Description: `**step beta ca provisioner list** lists the provisioners configured
in the CA.

## EXAMPLES

Prints a JSON list with active provisioners:
'''
$ step beta ca provisioner list
'''`,
	}
}

// nolint
func listAction(ctx *cli.Context) error {
	if err := errs.NumberOfArguments(ctx, 0); err != nil {
		return err
	}

	root := ctx.String("root")
	caURL, err := flags.ParseCaURL(ctx)
	if err != nil {
		return err
	}

	provisioners, err := pki.GetProvisioners(caURL, root)
	if err != nil {
		return errors.Wrap(err, "error getting the provisioners")
	}

	b, err := json.MarshalIndent(provisioners, "", "   ")
	if err != nil {
		return errors.Wrap(err, "error marshaling provisioners")
	}

	fmt.Println(string(b))
	return nil
}
