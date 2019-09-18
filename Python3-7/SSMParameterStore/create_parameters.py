import boto3
import json
from  parameter_functions import SSM_Parameters

ssm_client = boto3.client('ssm')
s3 = boto3.client('s3')

# Ask user to set project and environment
project_search = input("Please enter the project for this parameter: \n").strip()

# Set to lowercase to create/update parameter
project_param = project_search.lower()

print("-------------------------------------- String Parameters -----------------------------------------------------")
existent_string = []
nonexistent_string = []

# Production Parameters
with open('productionParameters.json') as f:
    file = json.load(f)
    production_param = file['parameters']['production_string_parameters']
    secure_production_param = file['parameters']['production_secure_string_parameters']

# Development Parameters
with open('developmentParameters.json') as e:
    file = json.load(e)
    development_param = file['parameters']['development_string_parameters']
    secure_development_param = file['parameters']['development_secure_string_parameters']

# Production Parameters
print('--------------------------------------------- Production Parameters -----------------------------------')
for parameter in production_param:
    name = production_param[parameter]["Name"]
    parameter_name = f'/{project_param}{name}'
    value = production_param[parameter]["Value"]
    parameter_value = f'{value}'
    p = SSM_Parameters(parameter_name)
    p.create_string_parameters(parameter_name, parameter_value)
# Secure Production Parameters
for parameter in secure_production_param:
    name = secure_production_param[parameter]["Name"]
    secure_param = f'/{project_param}{name}'
    value = secure_production_param[parameter]['Value']
    secure_param_value = f'{value}'
    ps = SSM_Parameters(secure_param)
    ps.create_secure_string_parameters(secure_param, secure_param_value)
print('--------------------------------------------- Development Parameters -----------------------------------')
for parameter in development_param:
    name = development_param[parameter]["Name"]
    parameter_name = f'/{project_param}{name}'
    value = development_param[parameter]["Value"]
    parameter_value = f'{value}'
    d = SSM_Parameters(parameter_name)
    d.create_string_parameters(parameter_name, parameter_value)
for parameter in secure_development_param:
    name = secure_development_param[parameter]["Name"]
    secure_param = f'/{project_param}{name}'
    value = secure_development_param[parameter]['Value']
    secure_param_value = f'{value}'
    ds = SSM_Parameters(secure_param)
    ds.create_secure_string_parameters(secure_param, secure_param_value)
