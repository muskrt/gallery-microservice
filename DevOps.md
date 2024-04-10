# Project Gallery: Microservices CI/CD Pipeline

## Description

This project aims to create full CI/CD Pipeline for microservice based application using [Golang Gallery  Microservices Application](https://github.com/muskrt/gallery-microservice.git). Jenkins Server deployed on Elastic Compute Cloud (EC2) Instance is used as CI/CD Server to build pipelines.



## Flow of Tasks for Project Realization

| Epic | Task  | Task #  | Task Definition   | Branch  |
| ---   | :---  | ---                  | :---              | :---    |
| Local Development Environment | Prepare GitHub Repository for the Project | MSP-1 | Clone the gallery app from the  repository [Gallery Microservices Application](https://github.com/muskrt/gallery-microservice.git) |
| Local Development Environment | Prepare GitHub Repository for the Project | MSP-2 | Prepare base branches namely `main` , `dev` , `release` for DevOps cycle. |
| Local Development Environment | Prepare Development Server Terraform Files | MSP-3 |  Prepare development server folder with Terraform File for developers, enabled with `Docker` , `Docker-Compose` , `Git` ,`Golang 1.0.22`. | feature/msp_devops_3 |
| CI Server Setup | Install Jenkins Server | MSP-4 | Install Jenkins Server for CI/CD Pipeline. | feature/msp_devops_4 |
| Local Development Environment | Check the Golang Build Setup on Dev Branch | MSP-5 | Check the Golang builds for `test` , `package` , and `install` phases on `dev` branch |
| Local Development Environment | Prepare a Script for Building  Packaging the Application | MSP-6 |  Prepare a script to build and package the application with Golang  | feature/msp_devops_6 |
| Local Development Build | Prepare Dockerfiles for Microservices | MSP-7 | Prepare Dockerfiles for each microservices. | feature/msp_devops_7 |
| Local Development Environment | Prepare Script for Building Docker Images | MSP-8 |  Prepare a script to package and build the docker images for all microservices. | feature/msp_devops_8 |
| Local Development Build | Create Docker Compose File for Local Development | MSP-9 |  Prepare docker compose file to deploy the application locally. | feature/msp_devops_9 |
| Local Development Build | Create Docker Compose File for Local Development | MSP-10 |  Prepare a script to test the deployment of the app locally. | feature/msp_devops_10 |
| Local Development Environment | Prepare dockerhub registry | MSP-11 |  Prepare dockerhub registry |
| Local Development Environment | Prepare images  | MSP-12 |  write a script to tag images for the registry for dev environment | feature/msp_devops_12|
| Local Development Environment | Prepare image push script | MSP-13 |  write a script to push images | feature/msp_devops_13|
| CI Server Setup | Configure Jenkins Server for Project | MSP-13  | install necessary plugins on  Jenkins  for Project Setup. | 
| Testing Environment Setup | Implement Unit Tests | MSP-14  | Implement  Unit Tests for all the services. | feature/msp_devops_14 |
| Testing Environment Setup | report code coverage | MSP-15  | write a script to make a coverage report | feature/msp_devops_15 |
| CI Pipeline Script | prepare CI Script | MSP-16  | Implement CI pipeline| feature/msp_devops_16 |
| Nightly Build  | Docker Tag script | MSP-17  |  write a script to build and tag images under nigthly folder| feature/msp_devops_17 |
| Nightly Build  | Build and Unit Test script | MSP-18 | write a script to build and test the application under nigthly folder| feature/msp_devops_18 |
| Nightly Build  | Registry Push Script | MSP-19  | Write a Script to Push images to registry under nigthly folder | feature/msp_devops_19 |
| Nightly Build  | Prepare helm package with HPA  | MSP-20  | prepare helm package with kompose under nigthly/k8 folder| feature/msp_devops_20 |
| Nightly Build  | Prepare k8 set up files | MSP-21  | prepare k8 set up files with terraform and ansible | feature/msp_devops_21 |
| Nightly Build  | prepare helm repo in s3(aws) | MSP-22  | write a script to create helm repo | feature/msp_devops_22 |
| Nightly Build  | create ansible playbook | MSP-23  | prepate deployment playbook | feature/msp_devops_23 |
| Nightly Build  | Unit Test | MSP-24  | create a script to run unit tests | feature/msp_devops_24 |



















