<p style="text-align:center;" align="center"><a href="https://layer5.io/meshery"><img align="center" style="margin-bottom:20px;" src="https://raw.githubusercontent.com/layer5io/layer5/master/assets/images/meshery/meshery-logo-tag-light-text-side.png"  width="70%" /></a><br /><br /></p>

[![Docker Pulls](https://img.shields.io/docker/pulls/layer5/meshery-octarine.svg)](https://hub.docker.com/r/layer5/meshery-octarine)
[![Go Report Card](https://goreportcard.com/badge/github.com/layer5io/meshery-octarine)](https://goreportcard.com/report/github.com/layer5io/meshery-octarine)
[![Build Status](https://github.com/layer5io/meshery-octarine/workflows/Meshery-Octarine/badge.svg)](https://github.com/layer5io/meshery-octarine/actions)
[![GitHub](https://img.shields.io/github/license/layer5io/meshery-octarine.svg)](LICENSE)
[![GitHub issues by-label](https://img.shields.io/github/issues/layer5io/meshery-octarine/help%20wanted.svg)](https://github.com/layer5io/meshery-octarine/issues?q=is%3Aopen+is%3Aissue+label%3A"help+wanted")
[![Website](https://img.shields.io/website/https/layer5.io/meshery.svg)](https://layer5.io/meshery/)
[![Twitter Follow](https://img.shields.io/twitter/follow/layer5.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=mesheryio)
[![Slack](https://img.shields.io/badge/slack-380-lightgrey)](http://slack.layer5.io)

# meshery-octarine
Meshery adapter for Octarine

><p align="center"><i>If you’re using Meshery or if you like the project, please ★ star this repository to show your support! 🤩</i></p>

## How the Meshery Adapter for Octarine Works
The Octarine adapter relys on an existing control plane already being up. Octarine provides a shared control plane for all Meshery users.
In order to deploy Octarine's dataplane into a target cluster the adapter performs the follwing operations:
* A new account is provisioned in the Octarine Control Plane.
* A domain is registered in that account. A domain in Octarine identified a k8s cluster.
* The YAML files required for deploying the data plane on the target cluster are generated.
* The YAML files are applied on the `octarine-dataplane` namespace in the target cluster.

Once the Octarine's data plane services are deployed, the adapter can be used to deploy Bookinfo. The steps here are:
* Enable the target namespace for automatic sidecar injection.
* Deploy Bookinfo to the target namespace.

## Environement Variables
In order to connect to the Octarine Control Plane the adapter requires the follwing environment variables to be set:
* OCTARINE_DOCKER_USERNAME: The docker username needed to pull Octarine's images to the target cluster. Do not use your own docker credentials. Use the ones supplies by Octarine.
* OCTARINE_DOCKER_EMAIL: The docker username needed to pull Octarine's images to the target cluster.
* OCTARINE_DOCKER_PASSWORD: The docker username needed to pull Octarine's images to the target cluster.
* OCTARINE_ACC_MGR_PASSWD : The password that will be assigned to the user 'meshery' in the new account.
* OCTARINE_CREATOR_PASSWD : The password needed to create an account in Octarine.
* OCTARINE_DELETER_PASSWD : The password needed to delete the account in Octarine.
* OCTARINE_CP : The address of the Octarine Control Plane. Example: meshery-cp.octarinesec.com
* OCTARINE_DOMAIN : The name that will be assigned to the target cluster in Octarine. Example: meshery:domain
---
## [Meshery](https://layer5.io/meshery)

A service mesh playground to faciliate learning about functionality and performance of different service meshes. Meshery incorporates the collection and display of metrics from applications running in the playground.

## Contributing
Please do! Contributions, updates, [discrepancy reports](/../../issues) and [pull requests](/../../pulls) are welcome. This project is community-built and welcomes collaboration. Contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

Not sure where to start? See the [newcomers welcome guide](https://docs.google.com/document/d/17OPtDE_rdnPQxmk2Kauhm3GwXF1R5dZ3Cj8qZLKdo5E/edit.) for how, where and why to contribute. Or grab an open issue with the [help-wanted label](https://github.com/issues?utf8=✓&q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+label%3A%22help+wanted%22+) and jump in.

## License

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).

#### About Layer5
[Layer5.io](https://layer5.io) is a service mesh community, serving as a repository for information pertaining to the surrounding technology ecosystem (service meshes, api gateways, edge proxies, ingress and egress controllers) of microservice management in cloud native environments.
