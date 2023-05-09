from os import environ

# Example using async credentials.
from azure.identity.aio import DefaultAzureCredential
from kiota_authentication_azure.azure_identity_authentication_provider import AzureIdentityAuthenticationProvider


class MSGraphAPI:
    def __init__(self):
        scopes = environ["MSGRAPH_USER_SCOPES"]
        self._graph = self._auth(scopes)

    def _auth(scopes):
        # Create authentication provider object. Used to authenticate request
        credential = ClientSecretCredential(
            tenant_id='TENANT_ID',
            client_id='CLIENT_ID',
            client_secret='CLIENT_SECRET'
        )
        scopes = ['https://graph.microsoft.com/.default']
        auth_provider = AzureIdentityAuthenticationProvider(credential, scopes=scopes)

        # Initialize a request adapter with the auth provider.
        request_adapter = GraphRequestAdapter(auth_provider)

        # Create an API client with the request adapter.
        client = GraphServiceClient(request_adapter)
