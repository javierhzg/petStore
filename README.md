# petStore

Come up with standalone and/or code that can be deployed to server code along with unit test included to call all of the following PETSTORE APIs and have the information displayed to the caller 

GET /pets: The backend returns a page of available pets in the PetStore. This is an example of the HTTP integration type. The URL of the integration endpoint is http://petstore-demo-endpoint.execute-api.com/petstore/pets.

POST /pets: for write access to the API's /pets resource that is integrated with the backend /petstore/pets resource. Upon receiving a correct request, the backend adds the specified pet to the PetStore and returns the result to the caller. The integration is also HTTP.

GET /pets/{petId}: for read access to a pet as identified by a petId value as specified as a path variable of the incoming request URL. The backend returns the specified pet found in the PetStore. The URL of the backend HTTP endpoint is http://petstore-demo-endpoint.execute-api.com/petstore/pets/n, where n is an integer as the identifier of the queried pet.

For the overall assignment, come up with detailed README.md and have code checked in GitHub public repository for us to review it.
