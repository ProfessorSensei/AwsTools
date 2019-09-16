# this code was heavily inspired from a course I bought. 
# I can't remember which course it was, so I do not remember 
# who to credit
# sorry about that if you end up seeing it. 
import boto3

ec2 = boto3.client('ec2')


# start instances
def instance_start(event, context):
    instances_ids = gather_Instance_Ids()
    response = ec2.start_instances(
        InstanceIds=instances_ids,
        # set to false if you plan on actually 
        # using the function
        DryRun=True
    )
    return response


# stop instances
def instance_stop(event, context):
    instances_ids = gather_Instance_Ids()
    response = ec2.stop_instances(
        InstanceIds=ec2_instances,
        # set to false if you plan on actually 
        # using the function
        DryRun=True
    )
    return response


# instance ID's
def gather_Instance_Ids():
    response = ec2.describe_instances(DryRun=False)
    instances = []
    for reservation in response["Reservations"]:
        for instance in reservation["Instances"]:
            # take instance ID from describe call
            instances.append(instance["InstanceId"])
    return instances
