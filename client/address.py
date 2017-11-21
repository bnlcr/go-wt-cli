"""
Defines an Address Class
"""

class Address(object):
    """
    Constructor for a address
    """
    def __init__(self, line_one, line_two, zip_code, city, country):
        #TODO: Implement consistancy checks
        # Initialize properties
        self.line_one = line_one
        self.line_two = line_two
        self.zip_code = zip_code
        self.city = city
        self.country = country
