"""
Defines a generic class to handle contracts features
"""
from datetime import datetime

class ContractException(Exception):
    """ Exception raised by the Contract Class """
    pass

class ContractAlreadyDeployedException(ContractException):
    """ Exception raised when attempting to deploy a contract that is already deployed """
    pass

class ContractNotDeployedException(ContractException):
    """ Exception raised when attempting actions on a contract not deployed """
    pass

class Contract(object):
    """
    Defines a Contract Class for generic contract related functions
    """
    def __init__(self, address=None, owner=None):
        """
        Constructor for the contract

        Keyword arguments:
        address -- Address of the smartcontract on the bblockchain (default None)
        address -- Address of the owner on the bblockchain (default None)
        """
        self.address = address
        self.owner = owner
        self.last_syncronization = None

    def get_contract_payload(self):
        """
        Returns the contract payload (Abstract)
        """
        raise NotImplementedError

    def deploy(self):
        """
        Deploy a smartcontract on the blockchain
        """
        # Check if the contract is already deployed
        if(self.address):
            raise ContractAlreadyDeployedException('Contract has already been deployed')

        #TODO: Implement actual call to REST API
        self.address = 0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae

    def syncronize(self):
        """
        Retrieves the smartcontract from the blockchain
        """
        # Check if the contract is already deployed
        if(not(self.address)):
            raise ContractNotDeployedException('Contract has not yet been deployed')

        #TODO: Implement actual call to REST API
        self.last_syncronization = datetime.now()
