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

package config

import (
	"encoding/json"
	"io"
	"os"
)

type Conf struct {
	From                   string   `json:"from"`
	Host                   string   `json:"host"`
	HostSMTPPort           uint     `json:"host_SMTP_port"`
	FromHostPassword       string   `json:"from_host_password"`
	ToDBPath               string   `json:"to_db_path"`
	MessageFilePath        string   `json:"message_file_path"`
	MessageSubject         string   `json:"message_subject"`
	MessageAttachmentPaths []string `json:"message_attachment_paths"`
	MessageSendDelayMS     uint     `json:"message_send_delay_miliseconds"`
}

// Default configuration file structure
func Default() *Conf {
	return &Conf{
		From:                   "me@example.com",
		Host:                   "smtp.example.com",
		HostSMTPPort:           587,
		FromHostPassword:       "me@example.com_password",
		ToDBPath:               "mail_list.txt",
		MessageFilePath:        "message.html",
		MessageSubject:         "Message subject",
		MessageAttachmentPaths: []string{""},
		MessageSendDelayMS:     50,
	}
}

// Write current configuration to w
func (c *Conf) WriteTo(w io.Writer) error {
	jsonData, err := json.MarshalIndent(c, " ", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// Read configuration from r
func (c *Conf) ReadFrom(r io.Reader) error {
	jsonData, err := io.ReadAll(r)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(jsonData, c)
	if err != nil {
		return err
	}

	return nil
}

// Creates configuration file at path
func CreateConfigFile(conf Conf, path string) error {
	confFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer confFile.Close()

	err = conf.WriteTo(confFile)
	if err != nil {
		return err
	}

	return nil
}

// Tries to open configuration file at path. If it fails - returns default configuration
func OpenConfigFile(path string) (*Conf, error) {
	confFile, err := os.Open(path)
	if err != nil {
		return Default(), err
	}
	defer confFile.Close()

	var conf Conf
	err = conf.ReadFrom(confFile)
	if err != nil {
		return Default(), err
	}

	return &conf, nil
}
