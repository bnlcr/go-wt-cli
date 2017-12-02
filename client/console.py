#!/usr/bin/env python
"""
Class to handle Console actions
"""
import code
from client.hotel import Hotel
from client.address import Address
from client.location import Location

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

        # Set the namespace
        self.namespace = None
        self.interactive_console = code.InteractiveConsole()

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
        """ Interactive menu to create an hotel """
        # Basic Information
        self.interactive_console.write("Hotel Information:\n")
        hotel = Hotel()
        hotel.name = self.interactive_console.raw_input("Hotel Name? ")
        hotel.description = self.interactive_console.raw_input("Hotel Description? ")
        hotel.timezone = 'Europe/Berlin'

        #TODO: Remove hardcode and use GeoIP to recommend a location
        hotel.location = Location('-669.556', '7878.545')

        #TODO: Remove hardcode and use an API to propose an adress
        hotel.address = Address('MyLineOne', 'MyLineTwo', 'MyZipCode', 'MyCity', 'MyCountry')
        
        # Ask confirmation before deploying
        # TODO: Compute actual gas price
        should_deploy = self.interactive_console.raw_input("Gas fee: 0.05wei, deploy on blockchain (yes/no)? ")
        if(should_deploy == 'yes'):
            hotel_contract_address = hotel.deploy()
            self.interactive_console.write("Hotel Contract Created: " + hex(hotel_contract_address) + "\n")
    
    def command_help(self, command=None):
        """ Display an help message """
        help_message = ""
        help_message += "Available commands:\n"
        help_message += "\thelp: display this help message\n"
        help_message += "\texit or quit or bye : Exits the CLI, or go back one level\n"
        
        # Commands for hotel
        if command == 'hotel':
            help_message += "\tcreate : Create an hotel\n"

        # If no command is set, we are at top level help
        else:
            help_message += "\thotel : Interacts with hotel\n"
        
        # Write help message in console
        self.interactive_console.write(help_message)

    def command_exit(self):
        """ Display an exit message and terminates """
        self.interactive_console.write("Bye.\n")


    def run(self):
        """
        Run the main console loop.
        Waits for event on the command line and answer to it
        """
        self.interactive_console.write("WT CLI Client\n")
        self.command_help()
        command = []

        # Loop until the user asks to exit
        try:
            while not((len(command)==1) and (command[0] in ['exit', 'quit', 'bye'])):
                # Command stack
                #self.interactive_console.write('Stack:' + ','.join(command) + '\n')

                # Prepare prompt
                prompt = 'WT'

                # Handle commands
                if(len(command)>0):

                    # Handle exit
                    if command[-1] in ['exit', 'quit', 'bye']:
                        command.pop()
                        command.pop()

                    # Handle help
                    elif command[0] == 'help':
                        self.command_help()
                        command.pop()

                    # Handle Hotel
                    elif command[0] == 'hotel':
                        # Handle namespace
                        if(len(command) == 1):
                            prompt += ' hotel'

                        # Handle help
                        elif((len(command) == 2) and (command[1] == 'help')):
                            prompt += ' hotel'
                            self.command_help(command[0])
                            command.pop()

                        # Handle creation
                        elif((len(command) == 2) and (command[1] == 'create')):
                            prompt += ' hotel'
                            self.command_hotel_create()
                            command.pop()

                        # Otherwise we need help
                        elif(len(command) == 2):
                            prompt += ' hotel'
                            self.command_help(command[0])
                            command.pop()
                    
                    # Command not knowned we delete the stack
                    else:
                        command = []
                        self.command_help()

                
                # Read from keyboard
                prompt += '>'
                reads = self.interactive_console.raw_input(prompt)
                if(len(reads) > 0):
                    command.extend(reads.split(' '))

        # Handle Keyboard exceptions, to avoid showing traceback
        except KeyboardInterrupt:
            self.interactive_console.write('\n')
        except EOFError:
            self.interactive_console.write('\n')
        
        # Show Exit message
        self.command_exit()
