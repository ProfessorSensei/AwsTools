import boto3
import json
from  parameter_functions import SSM_Parameters

ssm_client = boto3.client('ssm')
s3 = boto3.client('s3')


# Ask user to set project and environment
project_search = input("Please enter the project for this parameter: \n").strip()

# Set to lowercase to create/update parameter
project_param = project_search.lower()

with open('productionParameters.json') as f:
    file = json.load(f)
    production_param = file['parameters']['production_string_parameters']
    secure_production_param = file['parameters']['production_secure_string_parameters']

with open('developmentParameters.json') as e:
    file = json.load(e)
    development_param = file['parameters']['development_string_parameters']
    secure_development_param = file['parameters']['development_secure_string_parameters']


# Production Parameters
print('--------------------------------------------- Production Parameters -----------------------------------')
for parameter in production_param:
    name = production_param[parameter]["Name"]
    parameter_name = f'/{project_param}{name}'
    p = SSM_Parameters(parameter_name)
    try:
        p.delete_parameter(parameter_name)
    except:
        print(f"Parameter {parameter_name} does not exist, no need to delete")
# Secure Production Parameters
for parameter in secure_production_param:
    name = secure_production_param[parameter]["Name"]
    secure_param = f'/{project_param}{name}'
    ps = SSM_Parameters(secure_param)
    try:
        ps.delete_parameter(secure_param)
    except:
        print(f"Parameter {secure_param} does not exist, no need to delete")

# Development Parameters
print('--------------------------------------------- Development Parameters -----------------------------------')
for parameter in development_param:
    name = development_param[parameter]["Name"]
    parameter_name = f'/{project_param}{name}'
    d = SSM_Parameters(parameter_name)
    try:
        d.delete_parameter(parameter_name)
    except:
        print(f"Parameter {parameter_name} does not exist, no need to delete")
for parameter in secure_development_param:
    name = secure_development_param[parameter]["Name"]
    secure_param = f'/{project_param}{name}'
    ds = SSM_Parameters(secure_param)
    try:
        ds.delete_parameter(secure_param)
    except:
        print(f"Parameter {secure_param} does not exist, no need to delete")