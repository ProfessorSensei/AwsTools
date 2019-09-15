import json # To read and write json files (allowed resources and out of bounds resources)
import boto3 # Python's AWS API
import time 
from datetime import date
from datetime import datetime


def instancerun(event, context):
    # Date filtered for comparing date of instance 
    date_filter = datetime.strftime(datetime.today(), "%Y-%m-%d %H:%M")
    # This will be used to modify the resource files in the S3 Bucket
    s3 = boto3.client('s3')
    # This will be used to publish the message to SNS topic
    sns = boto3.client('sns')
    # Grab file from bucket
    script = s3.get_object(Bucket='<Bucket-containing-file-with-Resource-limits>', Key='allowed_resources.json')
    test_script = script['Body'].read()
    # Create tmp file for Lambda functin
    open_file = open('/tmp/allowed_resources.json', 'w') 
    # Write contents to tmp file 
    open_file.write(test_script)
    open_file.close()

    decoded = json.loads(test_script)

    # This is the allowed amount of Instances from the Allowed Resource file
    allowed_ec2_amount = decoded['Resources']['ec2']['permitted_ec2_maximum']
    # Dictionary for instances and creation date
    aws_instances = {}
    # Use EC2 resource to gather information on the instances
    all_instances = boto3.resource('ec2') 
    ec2_client = boto3.client('ec2')
    # Find Running running instances in the default region
    live_instances = all_instances.instances.filter(Filters=[{'Name': 'instance-state-name', 'Values': ['running']}]) 
    for instance in live_instances:
        instance_launch_date = datetime.strftime(instance.launch_time, "%Y-%m-%d %H:%M")
        # Update instance dictionary with instance launch date and instance ID
        aws_instances.update({'%s'%(instance.id): '%s'%(instance_launch_date)})
    print len(aws_instances)
    print aws_instances
        # If the number of instances running is higher than the specified amount
    if len(aws_instances) > allowed_ec2_amount: 
        print "Success"
        # Grabs Out of Bounds file to be written to 
        oob_file = s3.get_object(Bucket='<Bucket-containing-file-with-Resource-limits>', Key='out_of_bounds_resources.json')
        oob_file_body = oob_file['Body'].read()
        # Creates tmp file for lambda function. This will be used to write the new file in the bucket replacing the old file
        open_oob = open('/tmp/tmp_over_resource_limit_amount.json', 'w') 
        open_oob.write(oob_file_body)
        open_oob.close()
        # Add the instance ID's and lauch date to the Out of Bounds file
        with open('/tmp/tmp_over_resource_limit_amount.json', 'r+') as file:
            test_list = []
            json_data = json.load(file)
            # Changes dictionary to have instance names 
            json_data['Over_Limit_Resources']['Over_Limit_ec2']['Over_Limit_Instances'] = aws_instances
            file.seek(0)
            json.dump(json_data, file, indent=2, sort_keys=True)
            file.truncate()
        with open('/tmp/tmp_over_resource_limit_amount.json', 'rb') as f:
            s3.upload_fileobj(f, '<Bucket-containing-file-with-Resource-limits>', 'out_of_bounds_resources.json')
        print "email to thank you all for everything, I have no complaints whatsoever has been sent"
        # Sends email to user, alerting them of changes that need to be approved
        # sns.publish(
        #     TopicArn='arn:aws:sns:us-east-1:637047564439:Threshold-Above-Limit',
        #     Message="These Instances's: %s, are above the set limit. if you would like to keep them, please allow at:\n"
        #         "<link_to_increase_function_to_allow_instances>"%(aws_instances),
        #     Subject='EC2 Machines')
    # If the amount of dunning instances is equal to or below the allowed amount empyt over resource  list
    else:
        oob_file = s3.get_object(Bucket='<Bucket-containing-file-with-Resource-limits>', Key='out_of_bounds_resources.json')
        oob_file_body = oob_file['Body'].read()
        # Grab over resource file from bucket
        open_oob = open('/tmp/tmp_over_resource_limit_amount.json', 'w') 
        open_oob.write(oob_file_body)
        open_oob.close()
        with open('/tmp/tmp_over_resource_limit_amount.json', 'r+') as file:
            test_list = []
            json_data = json.load(file)
            # Empties dictionary so nothing will be in that section of over resource limit file
            json_data['Over_Limit_Resources']['Over_Limit_ec2']['Over_Limit_Instances'] = {}
            print "Thank you all for everything, I have no complaints whatsoever"
            file.seek(0)
            json.dump(json_data, file, indent=2, sort_keys=True)
            file.truncate()
        # Upload new file in place of the old one
        with open('/tmp/tmp_over_resource_limit_amount.json', 'rb') as f:
            s3.upload_fileobj(f, '<Bucket-containing-file-with-Resource-limits>', 'out_of_bounds_resources.json')
            


