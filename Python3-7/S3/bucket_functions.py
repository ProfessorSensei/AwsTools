# Import boto3
import boto3
# S3 Client
buckets = boto3.client("s3")
# import regex
import re 

def delete_begin():
    # Ask user if they want to use begins with, or match for buckets to delete
    bucket_begins = input("Please enter the name the bucket/s begin with that you would like to delete:\n")
    buckets_beginning_with = f"{bucket_begins}-"
    # create a list of buckets that begin with user input 
    buckets_to_delete = [bucket["Name"] for bucket in buckets.list_buckets()["Buckets"] if buckets_beginning_with in bucket["Name"]]
    for bucket in buckets_to_delete:
        x = re.search(f"^{buckets_beginning_with}", bucket)
        if x:
            try:
                # commenting out to avoid an issue 
                # buckets.delete_bucket(
                #     Bucket=bucket,
                # )
                print(f"'{bucket}' would have been deleted (uncomment delete portion)")
            except: 
                print(f"{bucket} could not be deleted. please try again")

def delete_match():
    bucket_match = input("Please enter the name the bucket/s begin with that you would like to delete:\n")
    buckets_match_with = f"{bucket_match}-"
    # create a list of buckets that begin with user input 
    buckets_to_delete = [bucket["Name"] for bucket in buckets.list_buckets()["Buckets"] if buckets_match_with in bucket["Name"]]
    for bucket in buckets_to_delete:
        try:
            print(f"'{bucket}' would have been deleted (uncomment function)")
            # commenting out to avoid an accident
            # buckets.delete_bucket(
            #     Bucket=bucket,
            # )
        except:
            print(f"{bucket} could not be deleted. Please try again")