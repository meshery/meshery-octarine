// Copyright 2019 The Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package octarine

type supportedOperation struct {
	// a friendly name
	name string
	// the template file name
	templateName string
	// the app label
	appLabel string
	// // returnLogs specifies if the operation logs should be returned
	// returnLogs bool
}

const (
	customOpCommand        = "custom"
	runVet                 = "octarine_vet"
	installOctarineCommand = "octarine_install"
	installBookInfoCommand = "install_book_info"
)

var supportedOps = map[string]supportedOperation{
	installOctarineCommand: {
		name: "Install the latest version of Octarine's data plane",
		// templateName: "install_octarine.tmpl",
	},
	installBookInfoCommand: {
		name: "Install the canonical Book Info Application",
		// templateName: "install_bookinfo.tmpl",
	},
	runVet: {
		name: "Vet Ocatarine's deployment",
		// templateName: "octarine_vet.tmpl",
		// appLabel:     "octarine-vet",
		// returnLogs:   true,
	},
	customOpCommand: {
		name: "Custom YAML",
	},
}
