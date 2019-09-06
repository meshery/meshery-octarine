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

import "github.com/layer5io/meshery-octarine/meshes"

type supportedOperation struct {
	// a friendly name
	name string
	// the template file name
	templateName string
	opType meshes.OpCategory
}

const (
	customOpCommand        = "custom"
	runVet                 = "octarine_vet"
	installOctarineCommand = "octarine_install"
	installBookInfoCommand = "install_book_info"
)

var supportedOps = map[string]supportedOperation{
	installOctarineCommand: {
		name: "Latest version of Octarine's data plane",
		// templateName: "install_octarine.tmpl",
		opType: meshes.OpCategory_INSTALL,
	},
	installBookInfoCommand: {
		name: "Sample application BookInfo",
		// templateName: "install_bookinfo.tmpl",
		opType: meshes.OpCategory_SAMPLE_APPLICATION,
	},
	runVet: {
		name: "Vet Ocatarine's deployment",
		// templateName: "octarine_vet.tmpl",
		opType: meshes.OpCategory_VALIDATE,
	},
	customOpCommand: {
		name: "Apply custom configuration (YAML)",
		opType: meshes.OpCategory_CUSTOM,
	},
}
