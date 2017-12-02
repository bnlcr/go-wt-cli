"""
Defines a generic class to handle contracts features
"""
from datetime import datetime
import http.client
import urllib.parse
import json

class ContractException(Exception):
    """ Exception raised by the Contract Class """
    pass

class ContractAlreadyDeployedException(ContractException):
    """ Exception raised when attempting to deploy a contract that is already deployed """
    pass

class ContractNotDeployedException(ContractException):
    """ Exception raised when attempting actions on a contract not deployed """
    pass

class ContractConnexionError(ContractException):
    """ Error raised by the Contract Class when connexion failed """
    pass

class Contract(object):
    """
    Defines a Contract Class for generic contract related functions
    """
    def __init__(self, contract_address=None, contract_owner=None):
        """
        Constructor for the contract

        Keyword arguments:
        contract_address -- Address of the smartcontract on the bblockchain (default None)
        contract_owner -- Address of the owner on the bblockchain (default None)
        """
        self.contract_address = contract_address
        self.contract_owner = contract_owner
        self.last_syncronization = None
        self.api_endpoint = 'localhost:8181'

    def get_contract_rest_payload(self):
        """
        Returns the contract payload (Abstract)
        """
        raise NotImplementedError

    def get_contract_rest_ressource(self):
        """
        Returns the contract ressource (Abstract)
        """
        raise NotImplementedError

    def execute_api_call(self, method, api_path, data=None):
        """ Execute the call to the REST API """
        # Build Payload
        request_data = {}
        if data:
            request_data.update(data)
        post_data = urllib.parse.urlencode(request_data)

        # Create Headers
        headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        }

        # Build Connection
        connexion = http.client.HTTPSConnection(self.api_endpoint)
        connexion.request(method, api_path, post_data, headers)
        response = connexion.getresponse()

        # Check the HTTP Status code
        if response.status not in [200, 201, 202]:
            raise ContractConnexionError('HTTP error: ' + str(response.status))

        # Interpret output as JSON
        try:
            json_string = response.read().decode()
            data = json.loads(json_string)
        except json.decoder.JSONDecodeError:
            raise ContractConnexionError('API Endpoint response is not JSON')

        return data

    def deploy(self):
        """
        Deploy a smartcontract on the blockchain
        """
        # Check if the contract is already deployed
        if self.contract_address:
            raise ContractAlreadyDeployedException('Contract has already been deployed')

        # Call the REST API
        api_result = self.execute_api_call(
            'POST',
            self.get_contract_rest_ressource(),
            self.get_contract_rest_payload())

        # Check if the address has been returned
        if 'address' not in api_result:
            raise ContractConnexionError('No contract address returned by API Endpoint')

        # Update adress and return it
        self.contract_address = api_result['address']
        return self.contract_address

    def syncronize(self):
        """
        Retrieves the smartcontract from the blockchain
        """
        # Check if the contract is already deployed
        if not self.contract_address:
            raise ContractNotDeployedException('Contract has not yet been deployed')

        #TODO: Implement actual call to REST API
        self.last_syncronization = datetime.now()
