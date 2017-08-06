# Ethereum IPLD plugin

Add ethereum support to ipfs!

## Building
Make sure to rewrite the gx dependencies in the directory above first, then
Either run `make` or `go build -buildmode=plugin -o=ethereum.so`.

* *NOTE*: As of 2017.07.17 the `plugins` lib in Go only works in Linux.

## Installing
Move `ethereum.so` to `$IPFS_PATH/plugins/ethereum.so` and set it to be executable:

```sh
mkdir -p ~/.ipfs/plugins
mv ethereum.so ~/.ipfs/plugins/
chmod +x ~/.ipfs/plugins/ethereum.so
```

### I don't have linux but I want to do this somehow!

As stated above, the _plugin_ library only works in Linux. Bug the go team to support your system!

* ... Or use a linux virtualbox, and mount this directory.

* ... Or hack your way via docker-fu [with this short, unsupported guide](hacks/docker.md)

* ... Or, if you are in OSX, [use this handy script](hacks/osx.sh)

## Usage and Examples

Make sure you have the right version of ipfs installed and start up the ipfs daemon!

### Add an ethereum block written in JSON

You may want to take a block given by your favorite client's JSON RPC API.
We have a couple of those in the `test-data` directory.

```
cat ./test_data/eth-block-body-json-997522 | ipfs dag put --input-enc json --format eth-block
```

And get the CID of the block header back!

```
z43AaGEzuAXhWf9pWAm63QCERtFpqcc6gQX3QBBNaG1syxGGhg6
```

Now, you can get this block header

```
ipfs dag get z43AaGEzuAXhWf9pWAm63QCERtFpqcc6gQX3QBBNaG1syxGGhg6
```

Which will get you (with the right IPLD formatted for the other objects)

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0x4bb96091ee9d802ed039c4d1a5f6216f90f81b01","difficulty":11966502474733,"extra":"14MBBACER2V0aIdnbzEuNS4xhWxpbnV4","gaslimit":3141592,"gasused":21000,"mixdigest":"0x2565992ba4dbd7ab3bb08d1da34051ae1d90c79bc637a21aa2f51f6380bf5f6a","nonce":"0xf7a14147c2320b2d","number":997522,"parent":{"/":"z43AaGF24mjRxbn7A13gec2PjF5XZ1WXXCyhKCyxzYVBcxp3JuG"},"receipts":{"/":"z44vkPhjt2DpRokuesTzi6BKDriQKFEwe4Pvm6HLAK3YWiHDzrR"},"root":{"/":"z45oqTRunK259j6Te1e3FsB27RJfDJop4XgbAbY39rwLmfoVWX4"},"time":1455362245,"tx":{"/":"z443fKyLvyDQBBQRGMNnPb8oPhPerbdwUX2QsQCUKqte1hy4kwD"},"uncles":{"/":"z43c7o73GVAMgEbpaNnaruD3ZbF4T2bqHZgFfyWqCejibzvJk41"}}
```

#### Piping from the RPC

The astute reader will say "_Let's then pipe the output of my RPC directly to IPFS!_"

```
curl -s -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x1b4", true],"id":1}' https://mainnet.infura.io | ipfs dag put --input-enc json --format eth-block && echo
```

Will give you

```
z43AaGF7XiKhgVVcYxNJv3ZrebEkDE5yhna22N74AusBdMvi6pV
```


And call

```
ipfs dag get z43AaGF7XiKhgVVcYxNJv3ZrebEkDE5yhna22N74AusBdMvi6pV
```

To get

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0xbb7b8287f3f0a933474a79eae42cbca977791171","difficulty":21109876668,"extra":"R2V0aC9MVklWL3YxLjAuMC9saW51eC9nbzEuNC4y","gaslimit":5000,"gasused":0,"mixdigest":"0x4fffe9ae21f1c9e15207b1f472d5bbdd68c9595d461666602f2be20daf5e7843","nonce":"0x689056015818adbe","number":436,"parent":{"/":"z43AaGF8SkCtKoht2v1e3yC9DWHi4iV2dynyi3BTCP7sPs7HR2T"},"receipts":{"/":"z44vkPheUUg5HBpxkq5sFFz5d9ckigtBBW7WCJXQSZA1gV233Ap"},"root":{"/":"z45oqTS9WCLjMeLnFvTbWiqxXRi1PdwYtDjnNQy6PyWKokGD8r8"},"time":1438271100,"tx":{"/":"z443fKyJXGFJKPgzhha8eqpkEz3rHUL5M7cvcfJQVGzwt3MwcVn"},"uncles":{"/":"z43c7o74hjCAqnyneWetkyXU2i5KuGQLbYfVWZMvJMG4VTYABtz"}}
```

Or go even more _extreme_ with a single pipe

```
ipfs dag get $(curl -s -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x2DC6C0", true],"id":1}' https://mainnet.infura.io | ipfs dag put --input-enc json --format eth-block)

{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0xea674fdde714fd979de3edf0f56aa9716b898ec8","difficulty":103975266902792,"extra":"ZXRoZXJtaW5lIC0gRVUy","gaslimit":3996095,"gasused":269381,"mixdigest":"0xb1907b36cfc58e666ec1f2d2b60422fc222b0994739bfe0a4b10ba68960cf2ab","nonce":"0xe9d09233833686d4","number":3000000,"parent":{"/":"z43AaGEzw2GGeV87BKohV7ykpq6FssJjBagzmGufRvVzq2Zvz5o"},"receipts":{"/":"z44vkPhdBNQZZd4HQGaDT34oRmkLQYputo7hqnYHKPdMNroFe7R"},"root":{"/":"z45oqTS4Ad9gDmfoQL9zY1g35AVWdbGbap1cQ37wbBaDZ5EASCs"},"time":1484475035,"tx":{"/":"z443fKyHcB6hA7QNg1XEFXsyKyat4EkUgVinTnvnhZLtJJgu2sH"},"uncles":{"/":"z43c7o74hjCAqnyneWetkyXU2i5KuGQLbYfVWZMvJMG4VTYABtz"}}
```

### Add an ethereum block encoded in RLP

This plugin also supports whether your block is an RLP encoded block header or
a block body (that is: its header, transactions and uncle list).

Let's test it out

#### Adding an RLP encoded block header

Just

```
cat ./test-data/eth-block-header-rlp-999999 | ipfs dag put --input-enc raw --format eth-block
```

You will get your cid `z43AaGF4uHSY4waU68L3DLUKHZP7yfZoo6QbLmid5HomZ4WtbWw`. Checking it,

```
ipfs dag get z43AaGF4uHSY4waU68L3DLUKHZP7yfZoo6QbLmid5HomZ4WtbWw
```

And we get our header back

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5","difficulty":12555463106190,"extra":"14MBAwOER2V0aIdnbzEuNC4yhWxpbnV4","gaslimit":3141592,"gasused":231000,"mixdigest":"0x5b10f4a08a6c209d426f6158bd24b574f4f7b7aa0099c67c14a1f693b4dd04d0","nonce":"0xf491f46b60fe04b3","number":999999,"parent":{"/":"z43AaGF6wP6uoLFEauru5oLK5JS5MGfNuGDK1xWEpQK4BqkJkL3"},"receipts":{"/":"z44vkPhhDSTXPAswvC1rdDunzkgZ7FgAAnhGQtNDNDk9m9N2BZA"},"root":{"/":"z45oqTSAZvPiiPV8hMZDH5fi4NkaAkMYTJC6PmaeWBmYUpbMpoh"},"time":1455404037,"tx":{"/":"z443fKyHHMwVy13VXtD4fdRcUXSqkr79Q5E8hcmEravVBq3Dc51"},"uncles":{"/":"z43c7o74hjCAqnyneWetkyXU2i5KuGQLbYfVWZMvJMG4VTYABtz"}}
```

#### Adding an RLP encoded block body (header, txs and uncle list)

We should get similars result trying to parse an RLP encoded block body.
Add `./test-data/eth-block-body-rlp-997522`

```
cat eth-block-body-rlp-997522 | ipfs dag put --input-enc raw --format eth-block
```

```
z43AaGExMLxj6ujVVbx3j4LRc6QGMBiqYCrgot5hG8Vnxm7Tf9M
```

```
ipfs dag get z43AaGExMLxj6ujVVbx3j4LRc6QGMBiqYCrgot5hG8Vnxm7Tf9M
```

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0x4bb96091ee9d802ed039c4d1a5f6216f90f81b01","difficulty":11966502474733,"extra":"14MBBACER2V0aIdnbzEuNS4xhWxpbnV4","gaslimit":3141592,"gasused":21000,"mixdigest":"0x2565992ba4dbd7ab3bb08d1da34051ae1d90c79bc637a21aa2f51f6380bf5f6a","nonce":"0xf7a14147c2320b2d","number":997522,"parent":{"/":"z43AaGF24mjRxbn7A13gec2PjF5XZ1WXXCyhKCyxzYVBcxp3JuG"},"receipts":{"/":"z44vkPheUUg5HBpxkq5sFFz5d9ckigtBBW7WCJXQSZA1gV233Ap"},"root":{"/":"z45oqTRunK259j6Te1e3FsB27RJfDJop4XgbAbY39rwLmfoVWX4"},"time":1455362245,"tx":{"/":"z443fKyLvyDQBBQRGMNnPb8oPhPerbdwUX2QsQCUKqte1hy4kwD"},"uncles":{"/":"z43c7o73GVAMgEbpaNnaruD3ZbF4T2bqHZgFfyWqCejibzvJk41"}}

```

## Navigate to a block's parent (and parent of a parent...)

If you have a chain of blocks available, you can easily navigate to a block's parent and so on.

Import the following blocks

```
cat ./test_data/eth-block-body-json-999999 | ipfs dag put --input-enc json --format eth-block
cat ./test_data/eth-block-body-json-999998 | ipfs dag put --input-enc json --format eth-block
cat ./test_data/eth-block-header-rlp-999997 | ipfs dag put --input-enc raw --format eth-block
cat ./test_data/eth-block-header-rlp-999996 | ipfs dag put --input-enc raw --format eth-block
```

(Notice how we are using block headers and bodies in different encodings).

Now, let's see how this goes, so we have this block

```
ipfs dag get z43AaGF4uHSY4waU68L3DLUKHZP7yfZoo6QbLmid5HomZ4WtbWw
```

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5","difficulty":12555463106190,"extra":"14MBAwOER2V0aIdnbzEuNC4yhWxpbnV4","gaslimit":3141592,"gasused":231000,"mixdigest":"0x5b10f4a08a6c209d426f6158bd24b574f4f7b7aa0099c67c14a1f693b4dd04d0","nonce":"0xf491f46b60fe04b3","number":999999,"parent":{"/":"z43AaGF6wP6uoLFEauru5oLK5JS5MGfNuGDK1xWEpQK4BqkJkL3"},"parentHash":"0xd33c9dde9fff0ebaa6e71e8b26d2bda15ccf111c7af1b633698ac847667f0fb4","receiptHash":"0x7fa0f6ca2a01823208d80801edad37e3e3a003b55c89319b45eb1f97862ad229","receipts":{"/":"z44vkPhhDSTXPAswvC1rdDunzkgZ7FgAAnhGQtNDNDk9m9N2BZA"},"root":{"/":"z45oqTSAZvPiiPV8hMZDH5fi4NkaAkMYTJC6PmaeWBmYUpbMpoh"},"rootHash":"0xed98aa4b5b19c82fb35364f08508ae0a6dec665fa57663dca94c5d70554cde10","time":1455404037,"tx":{"/":"z443fKyHHMwVy13VXtD4fdRcUXSqkr79Q5E8hcmEravVBq3Dc51"},"txHash":"0x447cbd8c48f498a6912b10831cdff59c7fbfcbbe735ca92883d4fa06dcd7ae54","uncleHash":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","uncles":{"/":"z43c7o74hjCAqnyneWetkyXU2i5KuGQLbYfVWZMvJMG4VTYABtz"}}
```

... and we call its parent

```
ipfs dag get z43AaGF4uHSY4waU68L3DLUKHZP7yfZoo6QbLmid5HomZ4WtbWw/parent
```

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0xf8b483dba2c3b7176a3da549ad41a48bb3121069","difficulty":12561596698199,"extra":"2YMBAwKER2V0aIdnbzEuNC4yh3dpbmRvd3M=","gaslimit":3141592,"gasused":252000,"mixdigest":"0xcaf27314d80cb3e888d32646402d617d8f8379ca23a6b0255e974e407ffdd846","nonce":"0xbc7609306a77d0a2","number":999998,"parent":{"/":"z43AaGF67aUUDzGGimXySbgNzJJitkTVUTvpaf9jrqxe8BKuJL2"},"parentHash":"0xc6fd988b2d086a7b6eee3d25bad453830391014ba268cf6cc5d139741cb51273","receiptHash":"0xb0310e47b0cc7d3bb24c65ec21ec0ddf8dcf1672bc9866d6ba67e83d33215568","receipts":{"/":"z44vkPhkV1Tp7osq3p4yThA7EdE5ikvZUZTtDpvvpkMNGvxC9HZ"},"root":{"/":"z45oqTSAdVfS8g8n7NrSBKTeydujwoRgw52ZQehEZaVhCd4QNx6"},"rootHash":"0xee8306f6cebba17153516cb6586de61d6294b49bc5534eb9378acb848907b277","time":1455404013,"tx":{"/":"z443fKyKQh6b7HWVtYXuJi6shDUtUTsFaw4g3vToP5n9eEvb3Jn"},"txHash":"0x6414d72a4c223bce7d1309869332b148670eb66af4e3b3ba6d1a55aa0bb3fd4f","uncleHash":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","uncles":{"/":"z43c7o74hjCAqnyneWetkyXU2i5KuGQLbYfVWZMvJMG4VTYABtz"}}
```

Why not calling its "grandparent" (parent of a parent)?

```
ipfs dag get z43AaGF4uHSY4waU68L3DLUKHZP7yfZoo6QbLmid5HomZ4WtbWw/parent/parent
```

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5","difficulty":12567733286589,"extra":"14MBAwOER2V0aIdnbzEuNC4yhWxpbnV4","gaslimit":3141592,"gasused":189000,"mixdigest":"0xedd380b8b600469c89d763fbb73c1ed4128164c2b8ccc41ed73d3e16f8d2a8de","nonce":"0x7b9a013e3da652ca","number":999997,"parent":{"/":"z43AaGF26webNZ5MTwHkhjcQZEqGkBfx65gDXALsa8171tUf5tU"},"parentHash":"0x8b6535a0e3e346ee87e0194456d95971988d3981639bf7065f602d11b7adeab9","receiptHash":"0x85c15ea267eda062e4470a875f6fe3135d8d63f561e409f5c0732c5539c35d1b","receipts":{"/":"z44vkPhhdMZfdtKqVtLTzXcppkFsT3W4PgZtK5SmTgPbHZsTUKC"},"root":{"/":"z45oqTS1N7W29LJtqzpcFZie3cPgwAyf6BT73MetEePYD9RGFn3"},"rootHash":"0x64d912e03889ea4754dd1039bd38a19677335aacd3399a3c3a3a74314588d584","time":1455403990,"tx":{"/":"z443fKyFeSt9z2MYAdG8GU26oT882qxkwCLE7C61UNSZJ4RpYEQ"},"txHash":"0x2c2c26e1629b431ad5fa033d90f4ec5c2b59d437cf1a34082195f5f771b3735d","uncleHash":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","uncles":{"/":"z43c7o74hjCAqnyneWetkyXU2i5KuGQLbYfVWZMvJMG4VTYABtz"}}
```

Since we are there, let's see what happens with their parent in turn

```
ipfs dag get z43AaGF4uHSY4waU68L3DLUKHZP7yfZoo6QbLmid5HomZ4WtbWw/parent/parent/parent
```

```
{"bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","coinbase":"0xf8b483dba2c3b7176a3da549ad41a48bb3121069","difficulty":12573872872824,"extra":"2YMBAwKER2V0aIdnbzEuNC4yh3dpbmRvd3M=","gaslimit":3141592,"gasused":42000,"mixdigest":"0xb388de31f2a59ec8cacf363866101a6904545d4dffa5d69537427e0df6f3aa2f","nonce":"0x5645faf4502c64d9","number":999996,"parent":{"/":"z43AaGF5oDbc3A3yMSMAjGaWeVCGvA2Pgrbb1C4mqfpZSCBhQiC"},"parentHash":"0xc249891eb893a583be09d904b7d952988098fd8bdf5de09003f7a4811fd0c591","receiptHash":"0xc7ce189fbc688fd45b844288a4d6016ca6002d77b1fa9e741716622608fb9312","receipts":{"/":"z44vkPhn5Bj9VT3BVsuDaMZ87gBQfYGtzSuJnYiEzLc91jggZbP"},"root":{"/":"z45oqTS3Sk3JQMTG7W3nXFF5o8RVGvbNotPLpGP1tbNJ22RNRBZ"},"rootHash":"0x83c016016b084a6074ea327e9ede376501965a8a18141f1bb3aef7a7c732bfec","time":1455403968,"tx":{"/":"z443fKyPdU3PSLUXTNtUhQ8BefUipk3CQf3n4xvRMf8LLRbiRKN"},"txHash":"0xa2c9608d6d1083b677012732bf149d232f02d32d465423b1ccb6306938bad451","uncleHash":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","uncles":{"/":"z43c7o74hjCAqnyneWetkyXU2i5KuGQLbYfVWZMvJMG4VTYABtz"}}

```

... And so on...

## Navigate though the transactions of a block

(WIP)

## TODO

* `[0x90]` - `eth-block` input:
  * Support the input of `eth-tx-trie` (`[0x95]`), when we are adding a block body.
  * Checkout navigation from a link provided in an obtained block.
  * Can we get the `eth-tx` (`[0x91]`) pointed by the `eth-tx-trie` leaf?

* `[0x92]` - `eth-tx-receipt`:
  * Propose a script to get all receipts from a block and make a JSON array of them.
  * Support the input of this JSON array to form the `eth-tx-receipt-trie` (`[0x96]`) leaves, and the `eth-tx-receipt` objects.

* `[0x97]` - `eth-state-trie`. Support input for RLP encoded state trie elements.
  * HINT: We get them from the Parity IPFS API.

* The rest of the IPLD ETH Types:
  * `[0x93]` - `eth-account-snapshot`
  * `[0x94]` - `eth-block-list`
  * `[0x98]` - `eth-storage-trie`
