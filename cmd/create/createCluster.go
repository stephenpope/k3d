/*
Copyright © 2019 Thorsten Klein <iwilltry42@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package create

import (
	"github.com/spf13/cobra"

	"github.com/rancher/k3d/pkg/types"

	log "github.com/sirupsen/logrus"
)

// NewCmdCreateCluster returns a new cobra command
func NewCmdCreateCluster() *cobra.Command {

	// create new command
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Create a new k3s cluster in docker",
		Long:  `Create a new k3s cluster with containerized nodes (k3s in docker).`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Debugln("create cluster called")
		},
	}

	// add flags
	cmd.Flags().StringP("name", "n", types.DefaultClusterName, "Set a name for the cluster")
	cmd.Flags().StringP("api-port", "a", "6443", "Specify the Kubernetes cluster API server port (Format: `--api-port [host:]port`")
	cmd.Flags().IntP("workers", "w", 0, "Specify how many workers you want to create")

	// add subcommands

	// done
	return cmd
}