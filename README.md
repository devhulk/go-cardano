# Go Cardano

### Go Abstraction for CardanoCLI

Just beginning development. Contributors welcome. 

### Pre-Reqs
- Go installed
- Tested on Debian GNU/Linux 10 (buster) so should work on Debian based systems
- Cardano Node Must be Up and Running. Here is a [guide](https://developers.cardano.org/docs/get-started/installing-cardano-node)
- Tested using cardano-cli 1.30.0


### Functionality

In order to mint NFTs on cardano you need to go through a specific set of steps. Those steps can be accomplished on Cardano using the CardanoCLI. 

The tasks include :

1. Creating Wallet Keys
2. Creating Wallet Payment Address
3. Getting Protocol Params
4. Create Minting Policy (Keys and PolicyScript)
5. Generating PolicyID
6. Creating / Deriving NFT Metadata
7. Building Cardano Transaction To Be Submitted
8. Submit Transaction to the Blockchain


There are steps in between like calculating the fee and funding the wallet. A detailed walkthrough can be found [here](https://developers.cardano.org/docs/native-tokens/minting-nfts/). That walks you through minting a single NFT using the cardano-cli and various bash commands. The first iteration of this repo will implement this workflow then go from there. 
