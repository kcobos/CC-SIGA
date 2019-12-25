[![Build Status](https://travis-ci.org/kcobos/CC-SIGA.svg?branch=master)](https://travis-ci.org/kcobos/CC-SIGA)
[![codecov](https://codecov.io/gh/kcobos/CC-SIGA/branch/master/graph/badge.svg)](https://codecov.io/gh/kcobos/CC-SIGA)
[![GPL Licence](https://badges.frapsoft.com/os/gpl/gpl.png?v=103)](https://opensource.org/licenses/GPL-3.0/)

# CC-SIGA
This is a repository to design SIGA infrastructure on server side considering all parts involved in SIGA (sensors, APPS, storing, large tasks...).

SIGA (Sistema Integral de Gestión de Aparcamientos (Integral Parking Management System)) is designed for reserved parkings in public area like authorities, ambulances or, its main aim, disabled parkings. It is based on the parking lot state data is received on the system automatically when the state changes by sensors disposed on parking lot.

On the other hand, users of this type of parking have an mobile application to view and navigate to free parkings. Moreover, users can rate parking lots have been used and report if a parking lot is used without accreditation or if there is any kind of problem in any parking lot. Also, there is an application to manage parking lots, users and authorizations used by public administration and police who is informed if a parking lot is been used without authorization.

## Infrastructure overview
The infrastructure on the server side it is composed by several services to be ready to scale the system and include more functionalities if it's needed in a future. Besides, separating functionalities we get improve the security of the system.

The first approximation of the infrastructure consists of multiple services as we describe below:
 * API Gateway: as a cloud system, an API Gateway is needed to be able to change the cloud without changing applications which connect to this cloud.
 * Parking: manage all parking lots and them states. This service is called by users and sensors so it has to be as quick as possible responding all requests. When a parking is occupied this API receives the ID of the parking lot and the ID of the user, if any, and check if the user ID exists. To do that it is needed a relation between this microservice an user microservice. In addition, any change of parking lot status must be recorded in other database (ParkingHist). 
 * ParkingHist: it is the parking historical states. Every change of parking lot status must be recorded to be able to know parking lot occupancy or incidences for reporting or improve the system.
 * User: manage users and permissions. These users could be administrators or clients. Every connection to this cloud have to be authenticated. 
 * ImageProcess: if sensors get a parking lot occupied without an accreditation (user ID), they take a picture of the vehicle and this picture is been had to process to get information like a car identification. This process could be take a lot of time so data come through a queue and after the process ends, the system has to upgrade the parking status and the parking historical.
 * Configuration: this cloud system needs a configuration service to be able to make changes as quick as possible.
 * Logging: to improve all system and check possible problems of the system a logging system is needed. This service is connected to all services, as configuration, to store and process all logs.

All services have to be in a same entry point. To do that, the system needs an API gateway which handles the outside requests without have to change URLs. Depend on demand and users or sensors number, we could divide this API gateway to two APIs, one for users and the other for sensors.

To sum that, this first approximation could be like:

![Overview scheme](./docs/CC_overview.jpeg)

## Specifying the infrastructure
It is time to specify in more detail the different parts of the project. This specification deals with the choice of programming languages to be used in the different microservices, the type and engine of the databases to be used as well as the API Gateway to be used, the logging and configuration systems and the queue to be used.

### API Gateway
The API Gateway we could choose is NGINX API Gateway. NGINX is a powerful load balancer, web server and reverse proxy but, also, it has a specific solution aimed at microservices‑based applications and it has a lot of [documentation](https://www.nginx.com/blog/deploying-nginx-plus-as-an-api-gateway-part-1/) . However, this solution is not free neither open source.

Looking for a free and open source solution [Krakend](https://www.krakend.io/) could be the best option for an API Gateway and also it has a lot of [documentation](https://www.krakend.io/docs/overview/introduction/).

### Configuration system
A fundamental part of a microservices‑based application is configure all distributed microservices as easy as possible. Here we could face to a several configuration and service discovery tools. Due to this is a open project, we want all tools which compose the project open as well.

To facilitate the decision of which configuration tool is the best for this project, is better to see a [comparative](https://stackshare.io/service-discovery) where [Consul](https://www.consul.io/) could be one of the best free and open source tool. Also it has a lot of [documentation](https://www.consul.io/docs/index.html) to read. Another could be [etcd](https://etcd.io/) installed by default in cloud OS.

### Logging system
Another fundamental part of a microservices‑based application is to log all microservices errors or information in only one site. This site could be centralized or distributed but developers wont have to see logs microservice per microservice.

As the before decision, here we could see a [comparative](https://stackshare.io/log-management) where [Logstash](https://www.elastic.co/products/logstash) could be one of the best tools to centralize all logs. Moreover, we a lot of [information](https://www.elastic.co/guide/en/logstash/current/index.html) about this tool.

### Microservices
#### Parking
This is a microservice which knows where are all parking lots and it states (free, busy or bad occupied). It has to response very quick to clients which can be users, administrators or sensor. For this reason and the concurrency of it functionality [Go Lang](https://golang.org/) could be one of the best option.

On the database side, all data are going to be structured and it has to be a parking location. Due to that there are going to be geographic objects and, probably, there will be operation with these objects. For this reason, we think, [PostgreSQL](https://www.postgresql.org/) with his [PostGIS](https://postgis.net/) plugin could be the best free option.

#### ParkingHist
ParkingHist is like a log system which only saves parking lot state related data. Here we only save data when a parking lot state has changed. Maybe to do that we could chose a logging tool or made it from scratch but the solution has to be as quick as possible.

About it database, all data is going to be structured so a relational DB will be the best option.

#### User
This microservice will process and store all user data like personal data or favorites, used parking lots. So this has to be a very secured and fast microservice so [Python](https://www.python.org/) and a API Rest in [Django](https://www.djangoproject.com/) could be the best option.

About it database, Django [officially supports](https://docs.djangoproject.com/en/3.0/ref/databases/) PostgreSQL, MariaDB, MySQL, Oracle and SQLite. Oracle is ruled out because it is not free and SQLite too because is not for a cloud, it is for small applications or clients.

About PostgreSQL, MariaDB or MySQL, MariaDB is a community-developed, commercially supported fork of the MySQL so between the other two, [in this comparative](https://db-engines.com/en/system/MariaDB%3bPostgreSQL), [MariaDB](https://mariadb.org/) could be better to do that than PostgreSQL.

#### ImageProcess
ImageProcess is a batch microservice to process images to find a vehicle identification, if any, in the picture taken when a parking lot is bad occupied. This is only called when sensors don't get any identification or get an identification which not in users and then they take a picture of the parking lot.

This picture is sent to *Parking* microservice, it saves the picture in a data storage and send the picture name and a ID to a queue. ImageProcess receive these requests, process the image and reply if an identification has been found. This reply is direct and via API to *Parking* microservice.

Python could be the best option to process the image due to it all [image processing libraries](https://www.datasciencecentral.com/profiles/blogs/9-python-libraries-which-can-help-you-in-image-processing). 

##### Queue
This queue is to send from *Parking* to *ImageProcess* microservices which image has going to be processed. So this queue has to be accesible from Go Lang and Python.

[In this ranking](https://stackshare.io/message-queue) of message queue tools [RabittMQ](https://www.rabbitmq.com/) is good placed and it could be accessed from Go Lang and Python through [amqp](https://github.com/streadway/amqp) and [pika](https://github.com/pika/pika).
