# Hotel API

The Hotel API allows to interact with an hotel for both Hoteliers and Travel Agents.

## Retrieve an Hotel

`GET /hotels/<hotel address>`

Retrieves the details from a given Hotel Smartcontract.

### Parameters

* `<hotel address>`: The address of the Hotel Smartcontract

### Output

* `name`: Name of the Hotel Property 
* `description`: Description of the Hotel Property
* `timezone`: Timezone where the hotel is located
* `address`: Physical address of the Hotel Property
  * `lineOne`: First line of the postal address
  * `lineTwo`: Second line of the postal address
  * `zipCode`: ZIP Code of the postal address
  * `city`: City Name
  * `country`: Country of the Hotel
* `location`: location of the Hotel Property
  * `latitude`: Latitude of the location of the Hotel Property
  * `longitude`: Longitude of the location of the Hotel Property
* `owner`: Ethereum address of the owner of the hotel
* `unitTypes`: Array of `address` related to the unit types in this hotel.

### Status Code List

* `200`: Success
* `404`: Hotel Contract not found

## Create an Hotel

`POST /hotels`

Allows an Hotelier to create an Hotel on the Blockchain using the Winding Tree contract. The created Smartcontract is automatically registered on the Winding Tree index.

### Input

* `name`: Name of the Hotel Property
* `description`: Description of the Hotel Property
* `timezone`: Timezone where the hotel is located
* `address`: Physical address of the Hotel Property
  * `lineOne`: First line of the postal address
  * `lineTwo`: Second line of the postal address
  * `zipCode`: ZIP Code of the postal address
  * `city`: City Name
  * `country`: Country of the Hotel
* `location`: location of the Hotel Property
  * `latitude`: Latitude of the location of the Hotel Property
  * `longitude`: Longitude of the location of the Hotel Property

### Output

* `address`: Address of the created smartcontract on the Ethereum blockchain

### Status Code List

* `200`: Success
* `400`: Invalid request or missing fields

## Modify an Hotel

`PATCH /hotels/<address>`

Allows an Hotelier to modify an existing Hotel on the Blockchain using the Winding Tree contract.

### Parameters

* `<address>`: The address of the Hotel Smartcontract

### Input

* `name`: Name of the Hotel Property
* `description`: Description of the Hotel Property
* `timezone`: Timezone where the hotel is located
* `address`: Physical address of the Hotel Property
  * `lineOne`: First line of the postal address
  * `lineTwo`: Second line of the postal address
  * `zipCode`: ZIP Code of the postal address
  * `city`: City Name
  * `country`: Country of the Hotel
* `location`: location of the Hotel Property
  * `latitude`: Latitude of the location of the Hotel Property
  * `longitude`: Longitude of the location of the Hotel Property

### Output

* `address`: Address of the created Hotel smartcontract on the Ethereum blockchain

### Status Code List

* `200`: Success
* `400`: Invalid request or missing fields
* `401`: Not Authorized to delete the Hotel, operation restricted to the owner.
* `404`: Hotel Contract not found

## Delete an Hotel

`DELETE /hotels/<hotel address>`

Allows an Hotelier to delete an existing Hotel on the Blockchain using the Winding Tree contract. This operation delete in cascade all sub-smartcontracts related to this hotel.

### Parameters

* `<hotel address>`: The address of the Hotel Smartcontract

### Status Code List

* `200`: Success
* `401`: Not Authorized to delete the Hotel, operation restricted to the owner.
* `404`: Hotel Contract not found