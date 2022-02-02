# ScaleoTempLink
### Description:

>This script generates a link to log in to your personal dashboard using the scaleo service API. The script can be run on a single user or on all users in the users.json file. The link is sent to the telegram bot to which the user must be subscribed. The .env contains the link with the API key and the telegram bot token. users.json contains email and pass from scale and personal chat_id from telegram'''

>[a similar project with an active telegram bot](https://github.com/sarff/ScaleoAcriveBot_GetLink)

1) rename .env_example to .env
2) rename users.json_example to users.json
3) Fill these files with the correct information (users.json and .env)

example how to run:
1) ./ScaleoTempLink -f_email="example@example.com" -f_pass="PASSWD!" -f_chat=2911115192  # send to this one
2) ./ScaleoTempLink #send to all from file
