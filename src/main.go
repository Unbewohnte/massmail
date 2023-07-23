/*
massmailer - send an e-mail to a list of addresses
Copyright (C) 2023 Kasyanov Nikolay Alexeyevich (Unbewohnte)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"Unbewohnte/massmail/config"
	"Unbewohnte/massmail/logger"
	"bufio"
	"flag"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
)

const version = "v0.1.2"

const (
	configFilename string = "conf.json"
)

var (
	printVersion = flag.Bool(
		"version", false,
		"Print version and exit",
	)

	wDir = flag.String(
		"wdir", "",
		"Force set working directory",
	)

	configFile = flag.String(
		"conf", configFilename,
		"Configuration file name to create|look for",
	)

	workingDirectory string
	configFilePath   string
)

func init() {
	// set log output
	logger.SetOutput(os.Stdout)

	// parse and process flags
	flag.Parse()

	if *printVersion {
		fmt.Printf(
			"massmailer %s\n(c) 2023 Kasyanov Nikolay Alexeyevich (Unbewohnte)\n",
			version,
		)
		os.Exit(0)
	}

	// print logo
	logger.GetOutput().Write([]byte(
		`███╗   ███╗ █████╗ ███████╗███████╗███╗   ███╗ █████╗ ██╗██╗     
████╗ ████║██╔══██╗██╔════╝██╔════╝████╗ ████║██╔══██╗██║██║     
██╔████╔██║███████║███████╗███████╗██╔████╔██║███████║██║██║     
██║╚██╔╝██║██╔══██║╚════██║╚════██║██║╚██╔╝██║██╔══██║██║██║     
██║ ╚═╝ ██║██║  ██║███████║███████║██║ ╚═╝ ██║██║  ██║██║███████╗
╚═╝     ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝╚══════╝ `,
	))
	logger.GetOutput().Write([]byte(version + " by Unbewohnte\n\n"))

	// work out working directory path
	if *wDir != "" {
		workingDirectory = *wDir
	} else {
		wdir, err := os.Getwd()
		if err != nil {
			logger.Error("Failed to determine working directory path: %s", err)
			return
		}
		workingDirectory = wdir
	}

	logger.Info("Working in \"%s\"", workingDirectory)

	// global path to configuration file
	configFilePath = filepath.Join(workingDirectory, *configFile)
}

func main() {
	// open config
	logger.Info("Trying to open config \"%s\"", configFilePath)

	var conf *config.Conf
	conf, err := config.OpenConfigFile(configFilePath)
	if err != nil {
		logger.Error(
			"Failed to open configuration file: %s. Creating a new one with the same name instead...",
			err,
		)

		err = config.CreateConfigFile(*config.Default(), configFilePath)
		if err != nil {
			logger.Error("Could not create new configuration file: %s", err)
			return
		}
		logger.Info("Created new configuration file. Exiting...")
		return
	}
	logger.Info("Successfully opened configuration file")

	// Sanitize inputs
	if len(conf.From) == 0 {
		logger.Error("Configuration's from is not specified!")
	}

	if len(conf.Host) == 0 {
		logger.Error("Configuration's host is not specified!")
	}

	if len(conf.FromHostPassword) == 0 {
		logger.Error("Configuration's from_host_password is not specified!")
	}

	if len(conf.ToDBPath) == 0 {
		logger.Error("Configuration's to_db_path is not specified!")
	}

	if len(conf.MessageFilePath) == 0 {
		logger.Error("Configuration's message_file_path is not specified!")
	}

	// Retrieve email message
	messageFile, err := os.Open(conf.MessageFilePath)
	if err != nil {
		logger.Error("Failed to open a message file: %s", err)
		return
	}
	defer messageFile.Close()

	isHtml := strings.HasSuffix(conf.MessageFilePath, ".html")

	messageBytes, err := io.ReadAll(messageFile)
	if err != nil {
		logger.Error("Failed to read message contents: %s", err)
		return
	}

	// Work out email addresses file
	addressesFile, err := os.Open(conf.ToDBPath)
	if err != nil {
		logger.Error("Failed to open addresses file: %s", err)
		return
	}
	defer addressesFile.Close()

	// Authenticate!
	mailDialer := gomail.NewDialer(conf.Host, int(conf.HostSMTPPort), conf.From, conf.FromHostPassword)

	// Send message to every address
	logger.Info("Starting to send...")
	wg := &sync.WaitGroup{}
	scanner := bufio.NewScanner(addressesFile)
	for scanner.Scan() {
		to := scanner.Text()
		if len(to) == 0 {
			continue
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			msg := gomail.NewMessage()
			msg.SetHeader("From", conf.From)
			msg.SetHeader("To", to)
			msg.SetHeader("Subject", conf.MessageSubject)
			for _, attachmentPath := range conf.MessageAttachmentPaths {
				// Check if exists
				stats, err := os.Stat(attachmentPath)
				if err != nil || stats.IsDir() {
					logger.Warning("%s does not exist or it is a directory. Skipping it...", attachmentPath)
					continue
				}
				msg.Attach(attachmentPath)
			}
			if isHtml {
				msg.SetBody(mime.TypeByExtension(".html"), string(messageBytes))
			} else {
				msg.SetBody(mime.TypeByExtension(".txt"), string(messageBytes))
			}
			err = mailDialer.DialAndSend(msg)
			if err != nil {
				logger.Info("Failed --> %s: %s", to, err)
			} else {
				logger.Info("Sent --> %s", to)
			}
		}(wg)

		time.Sleep(time.Millisecond * time.Duration(conf.MessageSendDelayMS))
	}

	wg.Wait()
	logger.Info("Completed")
}
