# safari-api

![](https://github.com/khaljimbo/safaripark-api/workflows/Docker%20Image%20CI/badge.svg)
![](https://github.com/khaljimbo/safaripark-api/workflows/Go%20Build/badge.svg)
![](https://github.com/khaljimbo/safaripark-api/workflows/Go%20Test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/khaljimbo/safari-api)](https://goreportcard.com/report/github.com/khaljimbo/safari-api)

## About

This is a simple API written in Go with the values hard coded inside the code for now. I've been putting this together in conjuction with a some training courses that I have done and am still doing on Go. There are plenty of things to improve on and add as I learn more about Go. You can see the TODO section at the bottom of this file for how I plan to improve on this. 

The idea behind this is to build a simple API that my family can use to record animal sightings when we're on our holidays in Africa.

### Tech Stack

Obviously this is built in Go. However outlined below is what I built this on. 

* Windows 10
* VSCode with various plugins
* Docker for Windows with Kubernetes for testing (AKS for more testing once my Terraform is migrated from Azure DevOps)
* Github Actions for CI\CD tasks
* Helm 3

Ordinarily I would use a CI server such as TeamCity (with the build manifest described in Kotlin) or Azure DevOps (described in plain YAML) partnered with a CD tool such as Octopus Deploy or AzureDevOps (also described in YAML) however for this piece I have chosen GitHub actions for these duties. Mainly to learn how to use GitHub Actions.

### Access

When the TODOs are completed I would host this externally using a service like Azure Front Door for DNS and WAF duties. Inter-POD connectivity will be handled by Consul or Istio. 

## Build and hacking

Using Go Get

```powershell
go get github.com/khaljimbo/safaripark-api/api
```

Building the Dockerfile

```powershell
docker build . -f deployment/Dockerfile -t khaljimbo/safaripark-go:latest
```

## Test

I have been using the standard "go test" framework. I have learned how to write simple tests in Go by reading some blog posts. I can either execute directory in VSCode or by running "go test".

## Running the API

### Locally via Docker Run

```powershell
docker run -p 8000:8000/tcp -it khaljimbo/safaripark-api:latest
```

### With Kubernetes and Helm

```powershell
helm install ./deployment/helm/safari-go --name safaripark-api --name-space safaripark-api
```

## Using the API

Using a cli tool like curl or PowerShell you are able to query the endpoints and return the results in the file.

```powershell
Invoke-WebRequest -Uri "http://localhost:8000/park" -Method GET
```

## TODO

Plenty to do for this to expand on this as I use this to continue learning Go.

* Learn more about testing and add code covereage
* Add some sort of Login system to protect access
* Migrate CI\CD tasks to Travis or Jenkins
* Add some sort of backend storage such as MySQL for persistant data storage
* Learn Go structured logging and send events to Seq for log ingestion and inspection using structured logging like the Logrus Go package.

### Following the 12 Factor App Principles

Following the [12 Factor App ](https://12factor.net/) principles. Below overview copied from [12 Factor App ](https://12factor.net/)

The twelve-factor app is a methodology for building software-as-a-service apps that:

1) Use declarative formats for setup automation, to minimize time and cost for new developers joining the project;
Have a clean contract with the underlying operating system, offering maximum portability between execution environments;

2) Are suitable for deployment on modern cloud platforms, obviating the need for servers and systems administration;
Minimize divergence between development and production, enabling continuous deployment for maximum agility;

3) And can scale up without significant changes to tooling, architecture, or development practices.

The below table is the 12 Factors and how the app and environment can be moulded to suit them.

| ✓   | Factor | Status |
|-----|---------|--------|
|1    |Code Base: One codebase tracked in revision control, many deploys | This project is isolated to one code base for everything, build, test and deploy.
| 2   | Dependencies: Explicitly declare and isolate dependencies | All dependancies are declared inside the package
| 3   | Config: Store config in the environment | There is little config here for the API itself here but all deployment config is stored within this repo and can be variabalised easily
| 4   | Backing services: Treat backing services as attached resources | Won't lie here, I'm not sure what this means
| 5   | Build, release, run: Strictly separate build and run stages | This can be done by carefully thought out CI. I have used GitHub actions to achieve this seperation of build/test and release stages. If using TeamCity or Azure DevOps this can be acheived by using dependant stages in the build and release. 
| 6   | Processes: Execute the app as one or more stateless processes | This is quite stateful at the moment since I have not introduced a backend such as MySQL to hold the data. 
| 7   | Port binding: Export services via port binding | The container and K8s are listening for HTTP on port 8000, this can be changed dynamically.
| 8   | Concurrency: Scale out via the process model | This can be done via the K8s Deployment.yml to add more nodes and the helm chart to deploy to more nodes.
| 9   | Disposability: Maximize robustness with fast startup and graceful shutdown | Unsure about how to achieve this in the Go code but K8s goes along way to helping achieve this by minimizing cold starts and allowing nodes to be spun up before being deployed to.
| 10  | Dev/prod parity: Keep development, staging, and production as similar as possible | In this example the environments would be identical, similarly if more environments were at play I would keep them as identical as possible, traditionally it is Data that is usually the only differing factor so keeping anonymized production like data should go a long way to help achieve this.
| 11   | Logs: Treat logs as event streams | While there is no logging on this API for now, log streams are vital to understanding what your code is doing in a production environments. Apps like Seq and Honeycomb can be used for this. 
| 12   | Admin processes: Run admin/management tasks as one-off processes | In production all admin tasks will be documented and automated to allow for repeatability.

## Cloud Nativeness

By being containerized and built for deployment to K8s using Helm this fits in the CNCF guidelines with being ready for scale. As this is containerized it is also cloud agnostic since the container image can be lifted and shifted to multiple cloud providers.
