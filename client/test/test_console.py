"""
Unit Tests for the Console class
"""
import unittest
import code
from client.console import Console
from unittest.mock import patch

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


class ConsoleRunUnitTest(unittest.TestCase):
    """
    Tests run cases
    """
    @staticmethod
    def get_banner_lines():
        """ Helper to get the banner """
        banner_lines = []
        banner_lines.append('WT CLI Client')
        banner_lines.append('Available commands:')
        banner_lines.append('\thelp: display this help message')
        banner_lines.append('\texit or quit or bye : Exits the CLI, or go back one level')
        banner_lines.append('\thotel : Interacts with hotel')
        return banner_lines

    @staticmethod
    def get_bye_lines():
        """ Helper to get the banner """
        bye_lines = []
        bye_lines.append('Bye.')
        bye_lines.append('')
        return bye_lines

    @staticmethod
    def check_console_sequence(test_case, expected_lines, expected_prompts, mock_write, mock_raw_input):
        """ Helper to check console input/outputs """
        # Check the sequence
        lines_printed = ""
        print_call = ""

        # Concatenate all the writes in a single variable
        for writes in mock_write.call_args_list:
            (print_call, ) = writes[0]
            lines_printed += print_call

        # Split per line for easier comparisions
        i = 0
        lines = lines_printed.split('\n')
        test_case.assertEqual(len(expected_lines), len(lines))
        for line in lines:
            test_case.assertEqual(expected_lines[i], line)
            i += 1

        # Check the prompts
        prompt = ""
        i = 0
        test_case.assertEqual(mock_raw_input.call_count, len(expected_prompts))
        for input_prompt in mock_raw_input.call_args_list:
            (prompt, ) = input_prompt[0]
            test_case.assertEqual(prompt, expected_prompts[i])
            i+=1
        return True

    @patch("code.InteractiveConsole.write", return_value=None) 
    @patch("code.InteractiveConsole.raw_input", side_effect=KeyboardInterrupt)
    def test_keyboard_interrupt(self,mock_raw_input, mock_write):
        """ A keyboard interrupt should exit nicely """
        # Build the expected output
        expected_lines = ConsoleRunUnitTest.get_banner_lines()
        expected_lines.append('')
        expected_lines.extend(ConsoleRunUnitTest.get_bye_lines())

        # Build the expected prompts
        expected_prompts = []
        expected_prompts.append('WT>')

        # Run the console
        my_console = Console()
        my_console.run()

        # Use the console check tool
        ConsoleRunUnitTest.check_console_sequence(self, expected_lines, expected_prompts, mock_write, mock_raw_input)

    @patch("code.InteractiveConsole.write", return_value=None) 
    @patch("code.InteractiveConsole.raw_input", return_value='exit')
    def test_enter_exit(self,mock_raw_input, mock_write):
        """ We can enter and exit using the exit command """
        # Build the expected output
        expected_lines = ConsoleRunUnitTest.get_banner_lines()
        expected_lines.extend(ConsoleRunUnitTest.get_bye_lines())

        # Build the expected prompts
        expected_prompts = []
        expected_prompts.append('WT>')

        # Run the console
        my_console = Console()
        my_console.run()

        # Use the console check tool
        ConsoleRunUnitTest.check_console_sequence(self, expected_lines, expected_prompts, mock_write, mock_raw_input)

    @patch("code.InteractiveConsole.write", return_value=None) 
    @patch("code.InteractiveConsole.raw_input", side_effect=['hotel','create','my_hotel','my_desc','yes','exit','exit'])
    def test_hotel_contract_creation_success(self,mock_raw_input, mock_write):
        """ Test a basic hotel creation workflow """
        # Build the expected output
        expected_lines = ConsoleRunUnitTest.get_banner_lines()
        expected_lines.append('Hotel Information:')
        expected_lines.append('Hotel Contract Created: 0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae')
        expected_lines.extend(ConsoleRunUnitTest.get_bye_lines())

        # Build the expected prompts
        expected_prompts = []
        expected_prompts.append('WT>')
        expected_prompts.append('WT hotel>')
        expected_prompts.append('Hotel Name? ')
        expected_prompts.append('Hotel Description? ')
        expected_prompts.append('Gas fee: 0.05wei, commit on blockchain (yes/no)? ')
        expected_prompts.append('WT hotel>')
        expected_prompts.append('WT>')

        # Run the console
        my_console = Console()
        my_console.run()

        # Use the console check tool
        ConsoleRunUnitTest.check_console_sequence(self, expected_lines, expected_prompts, mock_write, mock_raw_input)


        