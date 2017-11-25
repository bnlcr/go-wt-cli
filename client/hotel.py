"""
Defines an Hotel Class representing an Hotel Winding Tree smartcontract
"""
from client.location import Location
from client.address import Address
from client.contract import Contract

class Hotel(Contract):
    """
    Constructor for an hotel
    """
    def __init__(self, name, description, timezone, address, location, contract_address=None, contract_owner=None):
        #TODO: Implement consistancy checks
        # Initialize properties
        self.name = name
        self.description = description
        self.timezone = timezone
        self.address = address
        self.location = location
        self.contract_owner = contract_owner
        self.contract_address = contract_address
        self.unittypes = []
        self.units = []

    def get_contract_payload(self):
        """
        Returns the hotel contract payload
        """
        return { 
            'name': self.name, 
            'description': self.description, 
            'timezone': self.timezone,
            'address': {
                'lineOne': self.address.line_one,
                'lineTwo': self.address.line_two,
                'zipCode': self.address.zip_code,
                'city': self.address.city,
                'country': self.address.country,
            },
            'location': {
                'latitude': self.location.latitude,
                'longitude': self.location.longitude,
            },
        }
