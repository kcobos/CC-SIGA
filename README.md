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