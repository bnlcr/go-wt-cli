# REST JSON API Documentation

# Hotel
## Create an Hotel
### URI and Ressource 
PUT /hotel

### Input:
* name: Name of the Hotel Property
* description: Description of the Hotel Property
* timezone: Timezone where the hotel is located
* address: Physical address of the Hotel Property
* adress.lineOne: First line of the physical address
* adress.lineTwo: Second line of the physical address
* adress.zipCode: ZIP Code of the physical address
* location: location of the Hotel Property
* location.latitude: Latitude of the location of the Hotel Property
* location.longitude: Longitude of the location of the Hotel Property

### Output
HTTP Status code: 200 - Success
* address: Address of the created smartcontract on the Ethereum blockchain

HTTP Status code: 4XX-5XX
* errorDescription: Text of the error

## Retrieve an Hotel
### URI
GET /hotel/<smart contract address>

###Input:
* Smart Contract address: The address of the Hotel Smartcontract 

### Output:
* name: Name of the Hotel Property
* description: Description of the Hotel Property
* timezone: Timezone where the hotel is located
* address: Physical address of the Hotel Property
* adress.lineOne: First line of the physical address
* adress.lineTwo: Second line of the physical address
* adress.zipCode: ZIP Code of the physical address
* location: location of the Hotel Property
* location.latitude: Latitude of the location of the Hotel Property
* location.longitude: Longitude of the location of the Hotel Property
* manager: Ethereum address of the owner of the hotel

