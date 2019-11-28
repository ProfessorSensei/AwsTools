# AWS Interface for Python 
import boto3
# import regex
import re

# Cloudformation client to perform actions
cf_client = boto3.client("cloudformation")
# cloudformation resource 
cf_resource = boto3.resource("cloudformation")

def delete_stack():
    # Gather the name of the project to delete cloudformation stacks with 
    which_stacks = input("Please enter the project name for the Cloudformation stacks to be deleted:\n").lower()
    stacks_beginning_with = f"{which_stacks}-"

    # create a list with stacks from the project entered above 
    stacks_to_delete = [stack.name for stack in cf_resource.stacks.all() if stacks_beginning_with in stack.name]

    # need to add in a a message if there are no matches 
    if len(stacks_to_delete) < 1:
        print(f"Looks like there are no stacks starting with {which_stacks}")
    for stack in stacks_to_delete:
                x = re.search(f"^{stacks_beginning_with}", stack)
                if x:
    #                 Delete_Stack.delete_stack(stack)
                    try:
                        # print out to show the names of the stacks
                        print(f"{stack}...")
                        # delete_stacks = cf_client.delete_stack(
                        #     StackName=stack
                        # )
                    except:
                        print(f"{stack} could not be deleted. Please try deleting resources from the console first")
def delete_match_stacks():
    # Gather the name of the project to delete cloudformation stacks with 
    which_stacks = input("Please enter a value that the cloudformation stack(s) contain that need to be deleted:\n").lower()
    stacks_beginning_with = f"{which_stacks}-"

    # create a list with stacks from the project entered above 
    stacks_to_delete = [stack.name for stack in cf_resource.stacks.all() if stacks_beginning_with in stack.name]

    # need to add in a a message if there are no matches 
    if len(stacks_to_delete) < 1:
        print(f"Looks like there are no stacks starting with {which_stacks}")
    for stack in stacks_to_delete:
        try:
            # print out to show the names of the stacks
            print(f"{stack}...")
            # delete_stacks = cf_client.delete_stack(
            #     StackName=stack
            # )
        except:
            print(f"{stack} could not be deleted. Please try deleting resources from the console first")