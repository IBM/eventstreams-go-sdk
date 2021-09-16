[![Build Status](https://travis-ci.com/IBM/eventstreams-go-sdk.svg?&branch=main)](https://travis-ci.com/IBM/eventstreams-go-sdk)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
# IBM Cloud Eventstreams Go SDK Version 1.1.0

## Introduction

IBM Event Streams for IBM Cloud™ is a high-throughput message bus built with Apache Kafka. 
It is optimized for event ingestion into IBM Cloud and event stream distribution between your services and applications.

Event Streams provides a REST API to help connect your existing systems to your Event Streams Kafka cluster. 
Using the API, you can integrate Event Streams with any system that supports RESTful APIs.

Documentation [IBM Cloud Eventstreams Service APIs](https://cloud.ibm.com/apidocs/event-streams).

This is the Eventstreams Software Development Kit for `Go`
It includes a library of functions used to access a Eventstreams cluster.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    + [`go get` command](#go-get-command)
    + [Go modules](#go-modules)
    + [`dep` dependency manager](#dep-dependency-manager)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Eventstreams Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Admin Rest](https://cloud.ibm.com/apidocs/event-streams) | pkg/adminrestv1
[Schema Registry](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_schema_registry&locale=en#schema_registry_rest_endpoints) | pkg/schemaregistryv1

## Prerequisites

* An [IBM Cloud](https://cloud.ibm.com/registration) account.
* The [IBM Cloud CLI.](https://cloud.ibm.com/docs/cli?topic=cli-getting-started)
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* A IBM Cloud Eventstreams Instance Create one [here](https://cloud.ibm.com/registration?target=/catalog/services/event-streams)
* Go version 1.14 or above.

## Installation
The current version of this SDK: 1.1.0

There are a few different ways to download and install the Eventstreams Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the SDK to allow your Go application to
use it:

```
go get -u github.ibm.com/IBM/eventstreams-go-sdk 
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
	"github.ibm.com/IBM/eventstreams-go-sdk/pkg/adminrestv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.ibm.com/IBM/eventstreams-go-sdk"
  version = "1.1.0"

```

then run `dep ensure`.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md).

For details on using the SDK for administering Kafka topics, please see [kafka_topic_operations.md](./kafka_topic_operations.md).

For details on using the SDK for schema operations, please see [schema_operations.md](./schema_operations.md)

See [examples](./examples) for examples on using adminrest and schema SDKs.

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/eventstreams-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).