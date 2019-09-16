import json
import boto3
import time 
from datetime import date
from datetime import datetime, timedelta

# This function checks the Over Limit Resource file for instances that are a week ol
# It will delete instances only if they are over a week old
def hello(event, context):
    ec2_client = boto3.client('ec2')
    date_filter = datetime.strftime(datetime.today(), "%Y-%m-%d %H:%M")
    week_back = datetime.today() - timedelta(days=6) 
    print week_back
    week_back_date = datetime.strftime(week_back, "%Y-%m-%d %H:%M")
    print week_back_date

    s3 = boto3.client('s3')
    script = s3.get_object(Bucket='<Bucket-containing-file-with-Resource-limits>', Key='out_of_bounds_resources.json')
    test_script = script['Body'].read()

    open_file = open('/tmp/tmp_over_resource_limit_amount.json', 'w') 
    open_file.write(test_script)
    open_file.close()

    with open('/tmp/tmp_over_resource_limit_amount.json', 'r+') as file: 
        decoded = json.load(file)

        OOB_Instances = decoded['Over_Limit_Resources']['ec2']['Over_Limit_Instances']
        print OOB_Instances

        delete_ids = []
        for name, date in OOB_Instances.iteritems():
            # If the instance is a week old and in over limit resource file, add it to be deleted
            if date <= week_back_date:
                print "yes"
                print date
                delete_ids.append(name)

        for id in delete_ids:
            try:
                ec2_client.terminate_instances(InstanceIds=[id])
                print "%s wants to thank you all for everything, they had no complaints whatsoever"%(id)
            # Incase the list is empty or termination protection is on 
            except:
                print "Thank you all for everything, I have no complaints whatsoever"
        # Clear list of out of bound instances after they are deleted
        decoded['Over_Limit_Resources']['ec2']['Over_Limit_Instances'] = {}
        file.seek(0)
        json.dump(decoded, file, indent=2, sort_keys=True)
        file.truncate()
    with open('/tmp/tmp_over_resource_limit_amount.json', 'rb') as f:
        s3.upload_fileobj(f, '<Bucket-containing-file-with-Resource-limits>', 'out_of_bounds_resources.json')
