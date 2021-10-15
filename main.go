//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
    "os"

    "github.com/portapps/portapps/v3"
    "github.com/portapps/portapps/v3/pkg/log"
    "github.com/portapps/portapps/v3/pkg/utl"
)

var app *portapps.App

func init() {
    var err error

    // Init app
    if app, err = portapps.New("dolphin-portable", "Dolphin"); err != nil {
        log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
    }
}

func main() {
    preLaunch()

    defer app.Close()

    app.Process = utl.PathJoin(app.AppPath, "Dolphin.exe")
    app.Launch(os.Args[1:])
}

func preLaunch() {
    utl.CreateFolder(app.DataPath)

    configurationPath := utl.PathJoin(app.AppPath, "User")

    if fileInfo, err := os.Lstat(configurationPath); err == nil {
        if fileInfo.Mode() & os.ModeSymlink != os.ModeSymlink {
            if err = os.RemoveAll(configurationPath); err != nil {
                log.Fatal().Err(err).Msg("Unable to remove the configuration directory.")
            }
        }
    }

    if _, err := os.Lstat(configurationPath); err != nil {
        if err = os.Symlink(app.DataPath, configurationPath); err != nil {
            log.Fatal().Err(err).Msg("Unable to create a symlink to the configuration directory.")
        }
    }
}
