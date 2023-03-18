import sys
import json

from eth_typing import Address
from eth_utils import decode_hex

from staking_deposit.credentials import Credential
from staking_deposit.settings import ALL_CHAINS

args = sys.argv

sign_sk = args[1]
withdraw_sk = args[2]
index = args[3]
chain = args[4]
eth1_address = args[5]


result = Credential.make_from_pk(sign_sk,
                               withdraw_sk,
                               int(index),
                               32 * 10 ** 18,
                               ALL_CHAINS[chain],
                               eth1_address
                               )

result = result.get_bls_to_execution_change_dict(0)

print(json.dumps(result, sort_keys=True, indent=4))
