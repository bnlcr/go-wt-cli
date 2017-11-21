"""
Unit Tests for the Location class
"""
import unittest
from client.location import Location

class LocationUnitTest(unittest.TestCase):
    """
    Tests location use cases
    """
    def test_init(self):
        """ Test a basic initialization """
        my_location = Location(48.8737793, -73.989308)
        self.assertEqual(my_location.latitude, 48.8737793)
        self.assertEqual(my_location.longitude, -73.989308)
        