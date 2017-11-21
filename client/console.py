#!/usr/bin/env python
"""
Class to handle Console actions
"""
from code import InteractiveConsole

class Console(object):
    """
    Defines the console that is fired up at when starting the CLI
    """

    def __init__(self, args=None):
        """
        Constructor for the console.
        Allow to pass arguments from command line
        """
        # Initialize values
        self.server_host = 'localhost'
        self.server_port = 8456

        # Override values based on arguments from command line
        if args != None:
            # Override the Host
            if 'host' in args:
                self.server_host = args['host']

            # Override the Port
            if 'port' in args:
                self.server_port = args['port']


    def command_hotel_list(self):
        print( "(TODO: start list hotel shell)")

    def command_hotel_create(self):
        print( "(TODO: start create hotel shell)")
    
    def command_hotel(self):
        """
        Run the hotel interactive shell
        """
        # Display helper message
        print('Available commands: create, list, exit')

        # Start main loop
        keyboard_input = ""
        while keyboard_input != 'exit':
            # Ask for a command
            keyboard_input = InteractiveConsole.raw_input("(hotel)>>")

            # Handle command
            if keyboard_input == 'create':
                self.command_hotel_create()
            elif keyboard_input == 'list':
                self.command_hotel_list()
            else:
                print('Incorrect command. Available commands: create, list, exit')



    def run(self):
        """
        Run the main console loop.
        Waits for event on the command line and answer to it
        """
        # Display helper message
        print('Winding Tree CLI Client')
        print('Available commands: hotel, exit')

        # Start main loop
        keyboard_input = ""
        while keyboard_input != 'exit':
            # Ask for a command
            keyboard_input = InteractiveConsole.raw_input(">>")

            # Handle command
            if keyboard_input == 'hotel':
                self.command_hotel()

            else:
                print('Incorrect command. Available commands: hotel, exit')


