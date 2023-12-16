import base58
from solders.keypair import Keypair
import csv

WALLETS_AMOUNT = 100

with open("wallets.csv", 'w', newline='') as csvfile:
    csv_writer = csv.writer(csvfile)
    csv_writer.writerow(["WALLET", "PRIVATE KEY"])

    for x in range(WALLETS_AMOUNT):
        account = Keypair()
        privateKey = base58.b58encode(account.secret() + base58.b58decode(str(account.pubkey()))).decode('utf-8')

        csv_writer.writerow([account.pubkey(), privateKey])
