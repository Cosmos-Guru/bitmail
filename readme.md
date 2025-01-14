# Bitmail: A Decentralized, Permissionless Messaging Protocol Leveraging Cosmos SDK

## Abstract

The evolution of blockchain technology has transcended beyond simple token transfers, addressing more complex use-cases such as secure, decentralized messaging. This abridged white paper introduces **Bitmail**, an open source messaging protocol built on the Cosmos SDK, designed to eliminate spam and enhance privacy through decentralized verification and advanced encryption mechanisms. Unlike first-generation blockchain applications, which primarily facilitate numerical transactions between parties, Bitmail leverages second-generation blockchain capabilities to enable secure data exchanges. This document delves into the limitations of early blockchain messaging solutions—particularly those impacting scalability and feasibility—and illustrates how Bitmail rectifies these shortcomings via asymmetric encryption and decentralized infrastructure. By enabling users to share encrypted hash links on a permissionless, decentralized network, Bitmail presents a viable alternative to traditional email systems, which are notoriously susceptible to spam.

---
### Important Links
[Running a Full Node](https://github.com/Cosmos-Guru/bitmail/blob/main/docs/tutorials/BootstrapNode.md)

[Cosmos Journey Website - Cosmos SDK Tutorials](https://www.cosmos-journey.io/)


## Why Bitmail?

### Limitations of Traditional Email and Early Blockchain Attempts

Traditional email systems are riddled with **spam**—over half of all global email traffic—which burdens users with cluttered inboxes and exposes them to phishing and malware. The centralized nature of email further exacerbates security concerns, as malicious actors can easily bypass conventional sender-authentication mechanisms. 

Early attempts to develop blockchain-based messaging addressed some of these issues but faced two primary obstacles:

1. **On-Chain Data Storage**: Embedding large messages directly on a blockchain inflated storage requirements and transaction costs, severely limiting scalability.
2. **High Transaction Fees**: Networks with limited throughput and high gas costs made sending and storing messages prohibitively expensive for everyday use.

### How Bitmail Addresses These Problems

**Bitmail** overcomes these challenges by integrating decentralized storage (e.g., IPFS) with a **Cosmos SDK**-based blockchain. Messages and files are stored off-chain—ensuring scalability and low costs—while a minimal on-chain record prevents spam and helps verify senders’ identities.

- **Spam Elimination**: Bitmail transactions require valid blockchain identities, making it costly for malicious actors to spam the network.
- **Privacy & Security**: Asymmetric encryption ensures only intended recipients can decrypt messages, preventing unauthorized access or data exposure.
- **Scalability & Cost-Effectiveness**: Off-chain storage and an ample token supply (1 trillion Bitmail tokens) keep transaction fees negligible.

---

## What Is Bitmail and How Does It Work?

**Bitmail** is a decentralized messaging protocol that harnesses asymmetric encryption, decentralized storage, and the Cosmos SDK to provide secure, spam-free, and privacy-centric communication. Below is a concise overview, showcasing how users interact through an open-source Bitmail app. A Cosmos SDK based Bitmail Custom Transation is made possible with Bitmail Networks proprietry EHL(Encrypted Hash Link) Module. 

A Bitmail message is based the code for the `tx.proto` File below:

```
message MsgCreateHashCid {
  string creator  = 1;
  string receiver = 2;
  string hashlink = 3;
  string vaultid  = 4;
}
```

An example of this is below, but for brevity, Bitmail message is used to write the hashcid(Encryted Hash Link) to the Bitmail Network Blockchain. Where in the case below `creator` is Alice. `receiver` is Bob, `hashlink` is the encrypted hashlink Alice shares with bob, and `vaultid`, is chosen by the Bitmail app, which allows for mulitple vaults to be used, with the default vaultid as IPFS. Please Note: the Bitmail Network is agnotic to which vault(file storage) should be used. By add this parameter, we let Application developers decide not only which valut to use, but also this allows for multiple vaults or back up vaults. 


1. **Alice Composes a Bitmail**  
   Alice writes a love letter to Bob in her Bitmail app.

2. **Encryption**  
   - When she hits “Send,” the Bitmail app encrypts her Bitmail using Bob’s **public key**.  
   - This means only Bob—who holds the **matching private key**—can decrypt the contents of the Bitmail.

3. **Decentralized Storage**  
   - The app uploads the encrypted Bitmail to a decentralized storage network (e.g., IPFS).  
   - IPFS returns a **Content Identifier (CID)**, a unique alphanumeric locator pointing to the encrypted file on its network.

4. **CID Encryption and On-Chain Notification**  
   - For added security, the Bitmail app encrypts the CID again using Bob’s public key.  
   - It then submits a **custom Bitmail transaction** to the Cosmos SDK-based Bitmail Network, embedding this encrypted CID.  
   - Each transaction costs **1 Bitmail Token**, effectively thousands of times cheaper than a penny due to the large token supply.

5. **Message Retrieval and Decryption**  
   - Bob’s Bitmail app notifies him of a new message from Alice.  
   - He uses his **Bitmail Network private key** to decrypt the CID and retrieve the encrypted message from IPFS.  
   - Finally, Bob decrypts the message itself, allowing him—and only him—to read Alice’s letter.

---

## Conclusion

Bitmail emerges as a pioneering solution for decentralized messaging by exploiting the strengths of blockchain technology—particularly the Cosmos SDK—and decentralized storage systems. By relegating the bulk of data to IPFS (or similar networks) and requiring minimal on-chain transactions, Bitmail achieves low operational costs, high scalability, and robust spam prevention. Its reliance on asymmetric encryption and permissionless networks ensures that privacy and security remain paramount. 

As blockchain technologies continue to evolve, Bitmail’s architecture anticipates growing demands for secure, tamper-resistant, and spam-free communication channels in both personal and enterprise environments. In addressing the limitations of traditional email systems and early blockchain messaging solutions, Bitmail positions itself as a robust protocol poised to shape the future of decentralized communications.

---

## References

1. **Bitcoin Whitepaper**: Nakamoto, S. (2008). *Bitcoin: A Peer-to-Peer Electronic Cash System.* [https://bitcoin.org/bitcoin.pdf](https://bitcoin.org/bitcoin.pdf)  
2. **Ethereum Whitepaper**: Buterin, V. (2014). *Ethereum Whitepaper.* [https://ethereum.org/en/whitepaper/](https://ethereum.org/en/whitepaper/)  
3. **Cosmos SDK Documentation**: Cosmos Network. [https://docs.cosmos.network/](https://docs.cosmos.network/)  
4. **IPFS Documentation**: Protocol Labs. [https://docs.ipfs.io/](https://docs.ipfs.io/)  
5. **Asymmetric Encryption Techniques**: Stallings, W. (2017). *Cryptography and Network Security: Principles and Practice.* Pearson.  
6. **Decentralized Storage Systems**: Benet, J. (2014). *IPFS - Content Addressed, Versioned, P2P File System.* [https://ipfs.io/ipfs/QmYwAPJzv5CZsnAzt8auVTL38WELu6mQXxA3oFZgR5m6v6](https://ipfs.io/ipfs/QmYwAPJzv5CZsnAzt8auVTL38WELu6mQXxA3oFZgR5m6v6)  
7. **Spam Statistics**: Cisco. (2020). *Cybersecurity Report.* [https://www.cisco.com/c/en/us/products/security/security-reports.html](https://www.cisco.com/c/en/us/products/security/security-reports.html)
