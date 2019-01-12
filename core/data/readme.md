### CMX/core/data
###### Provides abstractions / ORM support.

The goal of CMX/core/data is to provide a layer where applications using it don't know if they are using SQL, NOSQL or
whatever you want to use. as long as it can be written to and read back from (sorry /dev/null and https://devnull-as-a-service.com/)


##### Planned Features
- Basic implementation uses GORM on the backend
- 