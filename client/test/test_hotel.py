"""
Unit Tests for the Hotel class
"""
import unittest
from unittest.mock import patch
import json
from client.hotel import Hotel
from client.location import Location
from client.address import Address

class HotelUnitTest(unittest.TestCase):
    """ Unit tests for Hotel Contacts """
    @patch("client.contract.Contract.execute_api_call", return_value={'address': "0x46506d900559ad005feb4645dcbb2dbbf65e19cc"})
    def test_create_hotel_success(self, mock_execute_api_call):
        """ Tests a successful hotel creation """
        my_hotel_dict = {"name": "My Hotel Name", "description": "My Hotel Description", "timezone": "Europe/Berlin", "address": {"lineOne": "MyLineOne", "lineTwo": "MyLineTwo", "zipCode": "MyZipCode", "city": "MyCity", "country": "MyCountry"}, "location": {"latitude": "-669.556", "longitude": "7878.545"}}
        my_hotel = Hotel()
        my_hotel.name = "My Hotel Name"
        my_hotel.description = "My Hotel Description"
        my_hotel.timezone = 'Europe/Berlin'
        my_hotel.location = Location('-669.556', '7878.545')
        my_hotel.address = Address('MyLineOne', 'MyLineTwo', 'MyZipCode', 'MyCity', 'MyCountry')
        my_hotel_address = my_hotel.deploy()
        self.assertEqual(my_hotel_address,"0x46506d900559ad005feb4645dcbb2dbbf65e19cc")
        self.assertEqual(mock_execute_api_call.call_args_list[0][0],('POST', '/hotels', my_hotel_dict))
