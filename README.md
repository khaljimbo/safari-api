# safari-api

![](https://github.com/khaljimbo/safari-api/workflows/Docker%20Image%20CI/badge.svg)
![](https://github.com/khaljimbo/safari-api/workflows/Go%20Build/badge.svg)
![](https://github.com/khaljimbo/safari-api/workflows/Go%20Test/badge.svg)

## About

This is a simple API written in Go with the values hard coded inside the code for now. I've been putting this together in conjuction with a some training courses that I have done about Go. There are plenty of things to improve on and add as I learn more about Go. You can see the TODO section at the bottom of this file for how I plan to improve on this. 

The idea behind this is to build a simple API that my family can use to record animal sightings when we're on our holidays in Africa.

## Tech Stack

Obviously this is built in Go. However outlined below is what I built this on. 

* Windows 10
* VSCode with various plugins
* Docker for Windows with Kubernetes
* Github Actions for CI\CD tasks
* Helm 3

Ordinarily I would use a CI server such as TeamCity (with the build manifest descriebd in Kotlin) or Azure DevOps (described in plain YAML) partnered with a CD tool such as Octopus Deploy or AzureDevOps (also described in YAML) however for this piece I have chosen GitHub actions for these duties.

### Accessing the API

When the TODOs are completed I would host this externally. Using a service like Azure Front Door.

## TODO

Plenty to do for this to expand on this as I use this to continue learning Go.

* Learn more about testing and add code covereage checkers
* Add some sort of Login system to protect access
* Migrate CI\CD tasks to Travis or Jenkins
* Add some sort of backend storage such as MySQL for persistant data storage
* Learn Go structured logging and send events to Seq for log ingestion and inspection