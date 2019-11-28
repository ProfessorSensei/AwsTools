from stack_functions2 import delete_stack, delete_match_stacks

# Decide if the user would like to find a match based on what the stack name begins with, 
# or just use stacks that match the input the user provides 
begins_with_or_match = input("Would you like to delete stacks the begin (enter b) with the name you provide, or match objects (enter m):\n").lower()
if begins_with_or_match == "b":
    delete_stack()
elif begins_with_or_match == "m":
    delete_match_stacks()
    