import json 
import boto3
import time 
from datetime import date
from datetime import datetime


def hello(event, context):
    s3 = boto3.client('s3')
    date_filter = datetime.strftime(datetime.today(), "%Y-%m-%d %H:%M")
    script = s3.get_object(Bucket='<Bucket-containing-file-with-Resource-limits>', Key='<allowed_resource_amount_file>.json')
    test_script = script['Body'].read()

    open_file = open('/tmp/tmp_over_resource_limit_amount.json', 'w') 
    open_file.write(test_script)
    open_file.close()

    decoded = json.loads(test_script)
    # Grab allowed amount of instances
    allowed_amount = decoded['Resources']['ec2']['ec2_maximum']

    aws_instances = []
    all_instances = boto3.resource('ec2') 
    ec2_client = boto3.client('ec2')
    # Find Running instances 
    live_instances = all_instances.instances.filter(Filters=[{'Name': 'instance-state-name', 'Values': ['running']}]) 
    for instance in live_instances:
        aws_instances.append(instance.id)
    with open('/tmp/tmp_over_resource_limit_amount.json', 'r+') as file:
        json_data = json.load(file)
        # Change instance limit to how many are currently running
        json_data['Resources']['ec2']['ec2_maximum'] = int(len(aws_instances))
        
        file.seek(0)
        json.dump(json_data, file, indent=2, sort_keys=True)
        file.truncate()
    with open('/tmp/tmp_over_resource_limit_amount.json', 'rb') as f:
            s3.upload_fileobj(f, '<Bucket-containing-file-with-Resource-limits>', '<allowed_resource_amount_file>.json')
    
    script2 = s3.get_object(Bucket='<Bucket-containing-file-with-Resource-limits>', Key='over_resource_limit_amount.json')
    test_script2 = script2['Body'].read()
    # Now clear Over_Limit file
    
    open_file2 = open('/tmp/over_resource_limit_amount.json', 'w') 
    open_file2.write(test_script2)
    open_file2.close()
    
    with open('/tmp/tmp_over_resource_limit_amount.json', 'r+') as Over_Limit_file:
        json_data = json.load(Over_Limit_file)
        json_data['Over_Limit_Resources']['Over_Limit_ec2']['Over_Limit_Instances'] = {}
        
        Over_Limit_file.seek(0)
        json.dump(json_data, Over_Limit_file, indent=2, sort_keys=True)
        Over_Limit_file.truncate()
    with open('/tmp/tmp_over_resource_limit_amount.json', 'rb') as f:
            s3.upload_fileobj(f, '<Bucket-containing-file-with-Resource-limits>', 'over_resource_limit_amount.json')
    
    
    
    print "Allowed Instance amount has been updated"

