# hello-merchant
Hello Merchant is a sound (Speech) announcement / notification engine for Merchant to receive payment alerts / actions on their choice of devices without having the need to purchase expensive sound box or checking their phones.

## Brief introduction on hello merchant
https://youtu.be/Qzk3yiCNDnY?si=lc9BJwxXJRSqRHGo

## Pre-requisite for Installation
1. Pangea Account
2. TiDB Serverless Cloud Account
2. Pusher Account
3. Vercel Account

## Installation Steps
```
git clone https://github.com/godwinpinto/hello-merchant
```

## TiDB Setup
1. Visit TiDB Cloud https://tidbcloud.com/ and create an account
2. Create a cluster
3. Execute the script from directory script/script.sql
4. This will create a new database UPN_DB in the cluster
5. Choose connect button on top right corner and note the DB credentials (you may need to set a password at this step too)
6. details to note: Endpoint is public, db_url, db_password, db_port, db_user

## Vercel Setup
1. Visit Vercel https://vercel.com/ and setup an account
2. Install vercel cli on your machine (https://vercel.com/docs/cli)


## Pusher Setup
1. Visit Pusher and register https://pusher.com/
2. Create a channel on dashboard and note the app keys

## Pangea Setup
1. Visit Pangea and register https://pangea.cloud/
2. You need to setup the next set of services
  - 2.1 AuthN (Add Redirect URL to web-box url and admin console url that you would setup in next steps)
  - 2.2 Secure Audit Log
  - 2.3. Redact (Create custom rule, Account-> starts with r)
  - 2.4. Embargo
  - 2.5. Vault
  - 2.5.1. Create A JWT KEY as symmetric
  - 2.5.2. Create a Random Secret of 32 characters
  - 2.6. IP Intel
  - 2.7. User Intel


## Project Components
The project is a mono repo consisting of various modules as listed below;
1. admin (The browser console for users to register and be onboarded to Hello Merchants)
2. connector-xrpl (The program that listen for Ripple Ledger transactions)
3. web-box (The browser based application for receiving announcements)
4. sound-box (TODO, esp32 micro controller based sound device)
5. mobile-box (TODO, mobile app to receive announcements)
6. kaios-box (TODO, Kai OS based app for feature phones)
7. connector-square (In Progress, application that listens for Square app transactions)
8. script (the TiDB HTAP compatible script for storing and retreiving data)

### Environment variables for admin and web-box module
Set the below environment variables in Vercel settings-> environment variables
1. VITE_PANGEA_VAULT_SECRET_AC_NO  (Pangea Vault secret for account no)
2. VITE_PANGEA_SECURE_AUDIT_LOG_TOKEN  (Pangea Secure Audit Token)
3. VITE_SERVER_URL=/api
4. VITE_PANGEA_VAULT_JWT_ID  (Pangea JWT ID from Vault)
5. VITE_PANGEA_VAULT_TOKEN  (Pangea vault Token)
6. VITE_PANGEA_USERINTEL_TOKEN  (Pangea User Intel Token)
7. VITE_PANGEA_IPINTEL_TOKEN  (Pangea ipintel Token)
8. VITE_PANGEA_EMBARGO_TOKEN  (Pangea embargo Token)
9. VITE_PANGEA_AUTHN_URL  (Pangea AuthN hosted URL)
10. VITE_PANGEA_REDACT_TOKEN (Pangea Redact Token)
11. PANGEA_AUTHN_TOKEN (Pangea AuthN Token)
12. PANGEA_DOMAIN  (Pangea Domain)
13. DB_URL (TiDB URL)
14. DB_PORT=4000
15. DB_NAME=UPN_DB
16. DB_USER (TiDB User)
17. DB_PASSWORD (TiDB password)
18. VITE_PUSHER_APP_CHANNEL=HELLO_MERCHANT_BOX
19. VITE_RIPPLE_ECHO_URL=/api
20. VITE_REDIRECT_URI (The URL to redirect to)
21. VITE_PUSHER_APP_CLUSTER=ap2
22. VITE_PUSHER_APP_KEY (App Key from Pusher)

### admin
1. Go to admin directory and type "vercel --prod" (post installing and configuring vercel cli)
2. This will host the application for you on vercel cloud
3. Set the default redirect in PANGEA to <DOMAIN_NAME>/ (slash)
4. Don't forget to update the redirect URL in your AuthN Pangea panel once application is hosted

### web-box
1. Go to web-box directory and type "vercel --prod" (post installing and configuring vercel cli)
2. This will host the application for you on vercel cloud
3. Set the redirect in PANGEA to <DOMAIN_NAME>/ (slash)
4. Don't forget to update the redirect URL in your AuthN Pangea panel once application is hosted

### connector-xrpl
1. Go to connector-xrpl and run "npm i" to install dependencies
2. Set below environent variables in .env file in root directory
  - 2.1. PORT=3001
  - 2.2. DATABASE_URL=mysql://USER:PASSWORD@URL:PORT/UPN_DB?sslmode=require&sslcert=/etc/ssl/certs/ca-certificates.crt
  - 2.3. PUSHER_APP_ID
  - 2.4. PUSHER_KEY
  - 2.5. PUSHER_SECRET
  - 2.6. PUSHER_CLUSTER=ap2
  - 2.7. PUSHER_CHANNEL=HELLO_MERCHANT_BOX
3. Run "npm run dev" to start for run as docker image
4. sudo docker build -t pangea-xrpl-connector .
5. sudo docker run -e PORT=3010 -e DATABASE_URL="mysql://USER:PASSWORD@URL:PORT/UPN_DB?sslmode=require&sslcert=/etc/ssl/certs/ca-certificates.crt" -e PUSHER_APP_ID=XXXX -e PUSHER_KEY=XXXXX -e PUSHER_SECRET=XXXXX -e PUSHER_CLUSTER=ap2 -e PUSHER_CHANNEL=HELLO_MERCHANT_BOX -e HELLO_ADMIN_URL=https://hello-merchant-admin.coauth.dev -p 3310:3010 -d pangea-xrpl-connector



### Limitation
1. Admin and web-box both run on vercel only, since its a mix of vuejs (for frontend) and golang application (for services)

### Testing
1. Signup and Login with hello.merchant@coauth.dev on the admin portal
2. Under channels you can Enter XRPL Account Number: rDKH6NniQpqoAJNBh4bTf7y9rXigwkyZHa and save
3. For testing you need to download XUMM wallet OR use my custom site https://ripplepay.coauth.dev/ to transfer 2 XRP amount to the account number mentioned in point 2
  - 3.1. If using XUMM Wallet or your custom account number then you need to create top up money on the account / wallet address in Testnet. Use this link to generate custom wallet https://xrpl.org/xrp-testnet-faucet.html  
4. Login to web-box url and play sound once and now make a transaction from point 3 to get announcement

##Future Modules:
This repository is a work in progress initiated at the Pangea hackathon. However, new development will continue and future aims to support below integrations
1. XRPL
2. SquareUp
3. Paytm
4. Google Pay
5. Stripe
6. Coinbase
7. Paypal
8. Phonepe

In addition a lot of refactor will take place as the application is a prototype and future developments will include breaking changes...

### Notes:
While I understand the setup is a bit more complicated. Going forward the plan is to create a one click deploy (with Terraform) on a cloud stack without having the husle to install all this manually 











