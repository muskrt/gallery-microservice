# Project 505: Microservices CI/CD Pipeline

## Description

This project aims to create full CI/CD Pipeline for microservice based application using [Golang Gallery  Microservices Application](https://github.com/muskrt/gallery-microservice.git). Jenkins Server deployed on Elastic Compute Cloud (EC2) Instance is used as CI/CD Server to build pipelines.



## Flow of Tasks for Project Realization

| Epic | Task  | Task #  | Task Definition   | Branch  |
| ---   | :---  | ---                  | :---              | :---    |
| Local Development Environment | Prepare GitHub Repository for the Project | MSP-1 | Clone the gallery app from the  repository [Gallery Microservices Application](https://github.com/muskrt/gallery-microservice.git) |
| Local Development Environment | Prepare GitHub Repository for the Project | MSP-2 | Prepare base branches namely `main` , `dev` , `release` for DevOps cycle. |
| Local Development Environment | Prepare Development Server Terraform Files | MSP-3 |  Prepare development server folder with Terraform File for developers, enabled with `Docker` , `Docker-Compose` , `Git` `golang v1.0.22`. | feature/msp_devops_3 |
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
| Nightly Build  | Docker Tag script | MSP-17  |  write a script to tag images under nigthly folder| feature/msp_devops_17 |
| Nightly Build  | Build and Unit Test script | MSP-18 | write a script to build and test the application under nigthly folder| feature/msp_devops_18 |
| Nightly Build  | Registry Push Script | MSP-19  | Write a Script to Push images to registry under nigthly folder | feature/msp_devops_19 |
| Nightly Build  | Prepare helm package  | MSP-20  | prepare helm package with kompose under nigthly/k8 folder| feature/msp_devops_20 |
| Nightly Build  | Prepare helm repo | MSP-21  | Create a script to prepare  helm repo with ecr(aws) under nightly folder| feature/msp_devops_21 |
| Nightly Build  | Prepare k8 set up files | MSP-22  | prepare k8 set up files with terraform and ansible | feature/msp_devops_22 |



| Testing Environment Setup | Implement Code Coverage | MSP-12  | Generate Code Coverage Report manually. | feature/msp-12 |
| Registry Setup for Development | Create Docker Registry for Dev Manually | MSP-14  | Create Docker Registry on Dockerhub  manually . | feature/msp-14 |
| CI Server Setup | Implement CI Pipeline | MSP-13 | Implement CI pipeline  for all `dev` , `feature` and `bugfix` branches. | feature/msp-13 |



| QA Automation Setup for Development | Create a QA Automation Environment - Part-1 | MSP-15  | Create a QA Automation Environment with Kubernetes. | feature/msp-15 |
| QA Automation Setup for Development | Create a QA Automation Environment - Part-2 | MSP-16  | Create a QA Automation Environment with Kubernetes. | feature/msp-16 |
| QA Automation Setup for Development | Prepare Petlinic Kubernetes YAML Files | MSP-17  | Prepare Petlinic Kubernetes YAML Files. | feature/msp-17 |
| QA Automation Setup for Development | Prepare a QA Automation Pipeline | MSP-18  | Prepare a QA Automation Pipeline on `dev` branch for Nightly Builds. | feature/msp-18 |
| QA Setup for Release | Create a QA Infrastructure with eksctl | MSP-19  | Create a Permanent QA Infrastructure for Kubernetes Cluster with eksctl. | feature/msp-19 |
| QA Setup for Release | Prepare Build Scripts for QA Environment | MSP-20  | Prepare Build Scripts for QA Environment | feature/msp-20 |
| QA Setup for Release | Build and Deploy App on QA Environment Manually | MSP-21  | Build and Deploy App for QA Environment Manually using Jenkins Jobs. | feature/msp-21 | 
| QA Setup for Release | Prepare a QA Pipeline | MSP-22  | Prepare a QA Pipeline using Jenkins on `release` branch for Weekly Builds. | feature/msp-22 |
| Staging and Production Setup | Prepare HA RKE Kubernetes Cluster | MSP-23  | Prepare High-availability RKE Kubernetes Cluster on AWS EC2 | feature/msp-23 |
| Staging and Production Setup | Install Rancher App on RKE K8s Cluster | MSP-24  | Install Rancher App on RKE Kubernetes Cluster| |
| Staging and Production Setup | Create Staging and Production Environment with Rancher | MSP-25  | Create Staging and Production Environment with Rancher by creating new cluster for Petclinic | |
| Staging Deployment Setup | Prepare and Configure Nexus Server | MSP-26  | Create and Configure Nexus Server for Pipelines | feature/msp-26|
| Staging Deployment Setup | Prepare a Staging Pipeline | MSP-27  | Prepare a Staging Pipeline on Jenkins Server | feature/msp-27|
| Production Deployment Setup | Prepare a Production Pipeline | MSP-28  | Prepare a Production Pipeline on Jenkins Server | feature/msp-28|
| Production Deployment Setup | Set Domain Name and TLS for Production | MSP-29  | Set Domain Name and TLS for Production Pipeline with Route 53 | feature/msp-29|
| Production Deployment Setup | Set Monitoring Tools | MSP-30  | Set Monitoring tools, Prometheus and Grafana | |







































