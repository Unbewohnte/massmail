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

package logger

import (
	"io"
	"log"
	"os"
)

// 3 basic loggers in global space
var (
	// neutral information logger
	infoLog *log.Logger
	// warning-level information logger
	warningLog *log.Logger
	// error-level information logger
	errorLog *log.Logger
)

func init() {
	infoLog = log.New(os.Stdout, "[INFO] ", log.Ltime)
	warningLog = log.New(os.Stdout, "[WARNING] ", log.Ltime)
	errorLog = log.New(os.Stdout, "[ERROR] ", log.Ltime)
}

// Set up loggers to write to the given writer
func SetOutput(writer io.Writer) {
	if writer == nil {
		writer = io.Discard
	}
	infoLog.SetOutput(writer)
	warningLog.SetOutput(writer)
	errorLog.SetOutput(writer)
}

// Get current logger's output writer
func GetOutput() io.Writer {
	return infoLog.Writer()
}

// Log information
func Info(format string, a ...interface{}) {
	infoLog.Printf(format, a...)
}

// Log warning
func Warning(format string, a ...interface{}) {
	warningLog.Printf(format, a...)
}

// Log error
func Error(format string, a ...interface{}) {
	errorLog.Printf(format, a...)
}
