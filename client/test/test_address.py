"""
Unit Tests for the Address class
"""
import unittest
from client.address import Address

class AddressUnitTest(unittest.TestCase):
    """
    Tests address use cases
    """
    def test_init(self):
        """ Test a basic initialization """
        my_address = Address('My Line One', 'My Line Two', '75001', 'Paris', 'FR')
        self.assertEqual(my_address.line_one, 'My Line One')
        self.assertEqual(my_address.line_two, 'My Line Two')
        self.assertEqual(my_address.zip_code, '75001')
        self.assertEqual(my_address.city, 'Paris')
        self.assertEqual(my_address.country, 'FR')
