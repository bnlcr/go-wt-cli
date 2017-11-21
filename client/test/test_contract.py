"""
Unit Tests for the Contract class
"""
import unittest
from datetime import datetime
from client.contract import Contract, ContractAlreadyDeployedException, ContractNotDeployedException

class ContractInitUnitTest(unittest.TestCase):
    """
    Tests contract initialization use cases
    """
    def test_init_without_address(self):
        """ Test a basic initialization without address"""
        my_contract = Contract()
        self.assertEqual(my_contract.address, None)

    def test_init_with_address(self):
        """ Test a basic initialization with address"""
        my_contract = Contract(0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae)
        self.assertEqual(my_contract.address, 0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae)

class ContractAbstractsUnitTest(unittest.TestCase):
    """
    Tests calls on abstracts functions
    """
    def test_get_contract_payload_abstract_call(self):
        """ Getting a payload on the abstract class should raise an exception """
        my_contract = Contract()
        with self.assertRaises(NotImplementedError):
            my_contract.get_contract_payload()


class ContractDeploymentUnitTest(unittest.TestCase):
    """
    Tests on contracts deployment
    """
    def test_deploy_success(self):
        """ Deploy a contract with success """
        my_contract = Contract()
        my_contract.deploy()
        self.assertEqual(my_contract.address, 0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae)

    def test_deploy_already_deployed(self):
        """ Attempting to deploy a contract already deployed should raise an exception """
        my_contract = Contract(0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae)
        with self.assertRaises(ContractAlreadyDeployedException):
            my_contract.deploy()


class ContractSyncronizeUnitTest(unittest.TestCase):
    """
    Tests on contracts syncronization
    """
    def test_syncronize_success(self):
        """ Contract syncronization success """
        my_contract = Contract(0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae)
        self.assertEqual(my_contract.last_syncronization, None)
        my_contract.syncronize()
        self.assertNotEqual(my_contract.last_syncronization, None)
        self.assertTrue(my_contract.last_syncronization < datetime.now())

    def test_syncronize_when_contract_not_deployed(self):
        """ Attempting to sync a contract not deployed should raise an exception """
        my_contract = Contract()
        self.assertEqual(my_contract.last_syncronization, None)
        with self.assertRaises(ContractNotDeployedException):
            my_contract.syncronize()