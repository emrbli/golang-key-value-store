## Golang Key Value Service REST-API
---
**Overview**

This document specifies the tools used in the Key-Value store and reorganizes how to use them. In this created service, In-Memory Key-Value Service was created and how to use the API is specified in the HTML file in the folder named "API-DOC". While writing the service, it was tried to comply with the "DDD" Design Pattern rules. In addition, the Go Documentation was followed and "The Uber Go-Style Guide" was used as a guide. It has been turned into a container with Docker. Licensed under the MIT License.


---
**Operation Guide**

In this service, the data is first pulled from the "mainData.json" file previously registered in the system, and the data recorded by the system in "/tmp/TIMESTAMP-data.json" is transferred to the "mainData.json" side every 1 minute. When the system stops and stands up again, the data is loaded into memory.

---

<h5> Endpoints</h5>

- **GET** for a key or all keys.

- **PUT** for set a key and value.

- **DELETE** for all.

#####Check out the "Api Doc" folder for a more detailed API Documentation.

---
<h3> For Using </h3>

- Install Docker what you use operating system.
- Go to the location where you uploaded the project
- Run "docker build -t go-key-value ."
- Run "docker run go-key-value"

