// Gd3
// Copyright (C) 2017-2018+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package lib

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/purpleidea/gd3/gapi"

	mgmt "github.com/purpleidea/mgmt/lib"
	"github.com/urfave/cli"
)

// Flags is an alias for the mgmt Flags struct which is used to pass in DEBUG.
type Flags mgmt.Flags

// run is the main run target.
func run(c *cli.Context) error {

	obj := &mgmt.Main{}

	obj.Program = c.App.Name
	obj.Version = c.App.Version
	if val, exists := c.App.Metadata["flags"]; exists {
		if flags, ok := val.(mgmt.Flags); ok {
			obj.Flags = flags
		}
	}

	if h := c.String("hostname"); c.IsSet("hostname") && h != "" {
		obj.Hostname = &h
	}

	if s := c.String("prefix"); c.IsSet("prefix") && s != "" {
		obj.Prefix = &s
	}
	obj.TmpPrefix = c.Bool("tmp-prefix")
	obj.AllowTmpPrefix = c.Bool("allow-tmp-prefix")

	obj.Noop = c.Bool("noop")

	obj.Seeds = c.StringSlice("seeds")
	obj.ClientURLs = c.StringSlice("client-urls")
	obj.ServerURLs = c.StringSlice("server-urls")
	obj.IdealClusterSize = c.Int("ideal-cluster-size")
	obj.NoServer = c.Bool("no-server")

	obj.ConvergedTimeout = -1 // disabled
	obj.NoPgp = true          // TODO: hardcoded for now
	obj.Sema = c.Int("sema")
	if c.Bool("graphviz") {
		obj.Graphviz = obj.Program // pick a sensible filename
		obj.GraphvizFilter = "dot" // pick a sensible default
	}

	obj.GAPI = &gapi.Gd3GAPI{ // graph API
		// TODO: add more parameters here!
		Program: obj.Program,
		Version: obj.Version,
	}

	if err := obj.Init(); err != nil {
		return err
	}

	// install the exit signal handler
	exit := make(chan struct{})
	defer close(exit)
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt) // catch ^C
		//signal.Notify(signals, os.Kill) // catch signals
		signal.Notify(signals, syscall.SIGTERM)

		select {
		case sig := <-signals: // any signal will do
			if sig == os.Interrupt {
				log.Println("Interrupted by ^C")
				obj.Exit(nil)
				return
			}
			log.Println("Interrupted by signal")
			obj.Exit(fmt.Errorf("killed by %v", sig))
			return
		case <-exit:
			return
		}
	}()

	if err := obj.Run(); err != nil {
		return err
		//return cli.NewExitError(err.Error(), 1) // TODO: ?
		//return cli.NewExitError("", 1) // TODO: ?
	}
	return nil
}

// CLI is the entry point for using gd3 normally from the CLI.
func CLI(program, version string, flags Flags) error {

	// test for sanity
	if program == "" || version == "" {
		return fmt.Errorf("program was not compiled correctly, see Makefile")
	}
	app := cli.NewApp()
	app.Name = program // App.name and App.version pass these values through
	app.Version = version
	app.Usage = "next generation glusterfs management"
	app.Metadata = map[string]interface{}{ // additional flags
		"flags": flags,
	}
	//app.Action = ... // without a default action, help runs

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run",
			Action:  run,
			Flags: []cli.Flag{
				// useful for testing multiple instances on same machine
				cli.StringFlag{
					Name:  "hostname",
					Value: "",
					Usage: "hostname to use",
				},

				cli.StringFlag{
					Name:   "prefix",
					Usage:  "specify a path to the working prefix directory",
					EnvVar: "GD3_PREFIX",
				},
				cli.BoolFlag{
					Name:  "tmp-prefix",
					Usage: "request a pseudo-random, temporary prefix to be used",
				},
				cli.BoolFlag{
					Name:  "allow-tmp-prefix",
					Usage: "allow creation of a new temporary prefix if main prefix is unavailable",
				},
				cli.BoolFlag{
					Name:  "noop",
					Usage: "globally force all resources into no-op mode",
				},
				cli.IntFlag{
					Name:  "sema",
					Value: -1,
					Usage: "globally add a semaphore to all resources with this lock count",
				},
				cli.BoolFlag{
					Name:  "graphviz",
					Usage: "generate a graphviz graph",
				},

				// if empty, it will startup a new server
				cli.StringSliceFlag{
					Name:   "seeds, s",
					Value:  &cli.StringSlice{}, // empty slice
					Usage:  "default etc client endpoint",
					EnvVar: "GD3_SEEDS",
				},
				// port 2379 and 4001 are common
				cli.StringSliceFlag{
					Name:   "client-urls",
					Value:  &cli.StringSlice{},
					Usage:  "list of URLs to listen on for client traffic",
					EnvVar: "GD3_CLIENT_URLS",
				},
				// port 2380 and 7001 are common
				cli.StringSliceFlag{
					Name:   "server-urls, peer-urls",
					Value:  &cli.StringSlice{},
					Usage:  "list of URLs to listen on for server (peer) traffic",
					EnvVar: "GD3_SERVER_URLS",
				},
				cli.IntFlag{
					Name:   "ideal-cluster-size",
					Value:  -1,
					Usage:  "ideal number of server peers in cluster; only read by initial server",
					EnvVar: "GD3_IDEAL_CLUSTER_SIZE",
				},
				cli.BoolFlag{
					Name:  "no-server",
					Usage: "do not let other servers peer with me",
				},
			},
		},
	}
	app.EnableBashCompletion = true
	return app.Run(os.Args)
}
