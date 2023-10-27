# Ledger Description

This is a REST API that can easily be used with cURL or postman. 

# Instructions

1. Run the program in your favorite editor (hopefully you're using VS Code) 
e.g. run in terminal: go run ledger_api.go

2. When you run the program you will see a popup window that asks for you to allow the program to run as localhost:3000.  Click "Allow" in the window. 

3. Once the program is running, open your favorite browser (hopefully you're using Brave) and go to http://localhost:3000

    -you should see the welcome message on the page. 

4. Use http://localhost:3000/account or http://localhost:3000/account/{id} to show a list of accounts in JSON format:

        It's suggested to have a JSON viewer installed as a 
        browser extension so the JSON is displayed with a 
        pretty format as intended.  e.g. in Brave, 
        install JSON Viewer

    *View Transactions: 

        GET transactions as JSON in the browser is the same command but usning /transactions as the endpoint
        http://localhost:3000/transaction
        or
        http://localhost:3000/transaction/{id}
        to GET a transaction by ID
    

# Create a New Account

5. To create a new account use the following cURL command: 
    
    curl POST script as follows:  curl -d '{"Id":"5","Name":"Example","Direction":"Debit","Balance":"0"}' -X POST http://localhost:3000/createAccount

    -Note that that "Name":"Example" can be renamed to whatever you want but "Direciont" should only be "Debit" or "Credit" and "Balance" is always "0" for a new account.  Ex:

        "Name":"Superman", "Direction":"Credit", "Balance":"0"
        We've changed the name and "debit" to "credit" while keeping "Balance" as "0"

# Create a New Transaction
        
6. To create a new account use the following cURL command: 
    
    curl POST script as follows:
    
        curl -d '{"Id":"Test","Amount":"Example","Direction":"Debit","Account_ID":"0"}' -X POST http://localhost:3000/createTransaction  

    -Note that that "Account":"Example" can be a denomination but "Direciont" should only be "Debit" or "Credit" and "Balance" is always "0" for a new account.  Ex:

    
        "Amount":"100", "Direction":"Credit", "Account_Id":"0x56fh49"
        We've changed the name and "debit" to "credit" while using an existing account_Id and 

# Supporting Info

        
    For questions, bugs, and general feedback contact me directly on the project link on github: 
    https://github.com/ialexander28/ledger_api.git
