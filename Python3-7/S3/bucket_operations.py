# import bucket functions 
from bucket_functions import delete_begin, delete_match

# prompt user if they want to delete buckets with beginning with or match
begin_match = input("Would you like to delete buckets beginning (b) with a name or matching(m)?:b or m\n").lower()
if begin_match == "b":
    delete_begin()
elif begin_match == "m":
    delete_match()
else:
    print(f"I'm sorry {begin_match} is not a valid option. Please try again")