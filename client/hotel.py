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
    def __init__(self, contract_address=None, contract_owner=None):
        # Initialize properties
        Contract.__init__(self, contract_address, contract_owner)
        self.name = None
        self.description = None
        self.timezone = None
        self.address = None
        self.location = None
        self.unittypes = []
        self.units = []

    def get_contract_rest_ressource(self):
        """
        Returns the contract ressource
        """
        return '/hotels'

    def get_contract_rest_payload(self):
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
