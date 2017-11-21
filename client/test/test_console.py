"""
Unit Tests for the Console class
"""
import unittest
from client.console import Console

class ConsoleInitUnitTest(unittest.TestCase):
    """
    Tests init use cases
    """
    def test_init_no_arguments(self):
        """ Test a basic initialization """
        console = Console()
        self.assertEqual(console.server_host, 'localhost')
        self.assertEqual(console.server_port, 8456)

    def test_init_override_server_name(self):
        """ Initialize with remote host using DNS Name """
        args = {'host':'wt.example.net', 'port': 12345}
        console = Console(args)
        self.assertEqual(console.server_host, 'wt.example.net')
        self.assertEqual(console.server_port, 12345)

    def test_init_override_server_ip(self):
        """ Initialize with remote host using DNS Name """
        args = {}
        args = {'host':'127.0.0.1', 'port': 98765}
        console = Console(args)
        self.assertEqual(console.server_host, '127.0.0.1')
        self.assertEqual(console.server_port, 98765)
