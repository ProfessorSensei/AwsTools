import boto3
ssm_client = boto3.client('ssm')
f1 = open("current_parameters.yml", "w")

# Create a class to create management parameters
class SSM_Parameters(object):
    def __init__(self, parameter_name):
        self.parameter = parameter_name
    # Function for string parameters
    def search_string_parameters(self, parameter_name):
        self.parameter_name = parameter_name
        try:
            response = ssm_client.get_parameter(
                Name=f'{parameter_name}',
                WithDecryption=False
            )
            located_parameter = response['Parameter']['Name']
            located_value = response['Parameter']['Value']
            print(f'Parameter name: {located_parameter}, Parameter Value: {located_value}', file=f1)
        except:
            print(f'{parameter_name} NOT Found. Needs to be created')
    # Function for secure string parameters
    def search_secure_string_parameters(self, secure_param):
        self.secure_param = secure_param
        try:
            response = ssm_client.get_parameter(
                Name=f'{secure_param}',
                WithDecryption=True
            )
            located_parameter = response['Parameter']['Name']
            located_value = response['Parameter']['Value']
            print(f'Parameter name: {located_parameter}, Parameter Value: {located_value}', file=f1)
        except:
            print(f'{secure_param} NOT Found. Needs to be created')
    # Function to create parameters
    def create_string_parameters(self, parameter_name, parameter_value):
        response = ssm_client.put_parameter(
            Name=f'{parameter_name}',
            Value=f'{parameter_value}',
            Type='String',
            Overwrite=True
        )
    # Function to create secure string parameters
    def create_secure_string_parameters(self, parameter_name, parameter_value):
        response = ssm_client.put_parameter(
            Name=f'{parameter_name}',
            Value=f'{parameter_value}',
            Type='SecureString',
            Overwrite=True
        )
    def delete_parameter(self, parameter_name):
        response = ssm_client.delete_parameter(
            Name=f'{parameter_name}'
        )
