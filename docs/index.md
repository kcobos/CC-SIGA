# CC-SIGA
## Infrastructure and deployment for SIGA

This is a repository to design SIGA infrastructure on server side considering all parts involved in SIGA (sensors, APPS, storing, large tasks...).

## Overview
In a global form, we have to divide all functionalities to be able to get a distributed and scalable system. The new functionalities which could be needed in this system will be easier to include in the project or even it is possible to make them standalone by this way of designing systems in microservices.

To begin with, the system will consist of a part that will serve as an interface for communication with the sensors. In this communication the server can receive images for its vehicle identification treatment. Because this treatment could take more time, a queue will be used to prevent the interface will stop listening to clients requests (sensors).

Continuing for client requests on the server side, the system needs two more interfaces, one for system administration and the other for final user. The first has to be more securized for obvious reasons and the second has to respond very quickly to final users requests.

The entire system needs a database or several databases to store data. You could divide data storage by layers to have precalculated data to respond faster to users. In addition, it is necessary to store all histories of system states and requests for security reasons and to improve it.