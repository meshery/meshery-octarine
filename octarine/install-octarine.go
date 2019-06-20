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

import (
	"io/ioutil"
	"math/rand"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	nsMgrUsername      = "meshery"
	nsMgrPassword      = ""
	octarineDeployment = ""
)

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (oClient *OctarineClient) createCpObjects() error {
	cmd := exec.Command("octactl", "login", "creator@octarine", oClient.octarineControlPlane, "--password", "creator-password")
	logrus.Debugf("Login to namespace %s", oClient.octarineNamespace)
	err := cmd.Run()
	if err != nil {
		logrus.Errorf("Command finished with error: %v", err)
		return err
	}
	oClient.octarineNamespace = "meshery-" + randSeq(6)
	cmd = exec.Command("octactl", "namespace", "create", oClient.octarineNamespace, nsMgrUsername, nsMgrPassword)
	logrus.Debugf("Creating namespace %s", oClient.octarineNamespace)
	err = cmd.Run()
	if err != nil {
		logrus.Errorf("Command finished with error: %v", err)
		return err
	}
	cmd = exec.Command("octactl", "login", nsMgrUsername + "@" + oClient.octarineNamespace, oClient.octarineControlPlane, "--password", nsMgrPassword)
	logrus.Debugf("Login to namespace %s", oClient.octarineNamespace)
	err = cmd.Run()
	if err != nil {
		logrus.Errorf("Command finished with error: %v", err)
		return err
	}
	cmd = exec.Command("octactl", "deployment", "create", oClient.octarineNamespace, nsMgrUsername, nsMgrPassword)
	logrus.Debugf("Creating deployment %s in namespace %s", octarineDeployment, oClient.octarineNamespace)
	err = cmd.Run()
	if err != nil {
		logrus.Errorf("Command finished with error: %v", err)
		return err
	}
	return nil
}

func (oClient *OctarineClient) deleteCpObjects() error {
	cmd := exec.Command("octactl", "login", "deleter@octarine", oClient.octarineControlPlane, "--password", "deleter-password")
	logrus.Debugf("Login as deleter to namespace %s", oClient.octarineNamespace)
	err := cmd.Run()
	if err != nil {
		logrus.Errorf("Command finished with error: %v", err)
		return err
	}
	cmd = exec.Command("octactl", "namespace", "delete", oClient.octarineNamespace, "--force")
	logrus.Debugf("Creating namespace %s", oClient.octarineNamespace)
	err = cmd.Run()
	if err != nil {
		logrus.Errorf("Command finished with error: %v", err)
		return err
	}
	return nil
}

func (oClient *OctarineClient) getOctarineDataplaneYAML(namespace string) (string, error) {
	cmd := exec.Command("octactl", "dataplane", "install", "--k8s-namespace", namespace, octarineDeployment)
	logrus.Debugf("Creating dataplane yaml for deployment %s in namespace %s", octarineDeployment, namespace)
	dp, err := cmd.Output()
	if err != nil {
		logrus.Errorf("Command finished with error: %v", err)
		return "", err
	}
	return string(dp), nil
}

const (
	bookInfoInstallFile        = "/bookinfo.yaml"
)

func (oClient *OctarineClient) getOctarineYAMLs(namespace string) (string, error) {
	dp, err := oClient.getOctarineDataplaneYAML(namespace)
	if err != nil {
		err = errors.Wrap(err, "unable to create dataplane yaml")
		logrus.Error(err)
		return "", err
	}
	return dp, nil
}

func (oClient *OctarineClient) getBookInfoAppYAML() (string, error) {
	b, err := ioutil.ReadFile(bookInfoInstallFile)
    if err != nil {
		err = errors.Wrap(err, "Failed to read bookinfo.yaml")
		logrus.Error(err)
		return "", err
    }
	str := string(b)
	return str, nil
}
