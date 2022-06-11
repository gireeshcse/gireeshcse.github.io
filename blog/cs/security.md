### RSA

* Rivest Shamir Adleman.
* 1977
* Most common Asymmetric Encryption Algorithm
* Creates a pair of "commutative" keys.
    - encrypt with one and decrypt with other.

#### Step 1

- Generating Keys:
    - Select two prime numbers. (P,Q)
    - Caliculate Product (P*Q) = N
    - Caliculate Totient (P-1) * (Q-1) = T
    - Select Public Key (E)
        - Must be Prime
        - Must be less than Totient
        - Must not be a factor of the Totient
    - Select a Private Key (D)
        * Product of D and E, divided by T must result in a remainder of 1
        * (D * E) MOD T = 1

Example: 

    - Primes (P,Q) = (7,19)
    - Product N = 133
    - Totient T = 108
    - Public Key E = 29
    - Private Key D = 41

- Encryption and Decryption
    - Encryption
        - Message^E MOD N = Cipher Text
    - Decryption:
        - Cipher^D MOD N = Message


Example:
    - Message = 60
    - Encrypt with Public key (60^29) MOD 133 = 86
    - Decrypt with Private key 86^41 MOD 133 = 60

    - Message = 60
    - Encrypt with Privatekey (60^41) MOD 133 = 72
    - Decrypt with Public  key 72^29 MOD 133 = 60

* 1024 bit RSA keys was recommended standard since 2002.
* 2048 bit RSA keys is recommended standard since 2015.

### Diffie-Hellman

* Asymmetric Encryption Algorithm
* Allows two parties to establish a shared secret over an unsecured medium
* Shared Secret is never transmitted only values used to derive secret.
    - Shared secret is then used to generate Symmetric keys.
* Security of DH is dependent on Discrete Logarithm problem.
    - Only method to break this is trying every possible combination (i.e brute force)

#### Method 

- Agree upon two numbers 
    - P Prime Number = 13
    - G Generator of P = 6
- Randomly generate a Private Key
    - for Bob = 5
    - For Alice = 4
- Caliculate Public Key: (G^Private) Mod P
    - For Bob (6^5) MOD 13 = 2
    - For Alice (6^4) MOD 13 = 9
- Exchange Public Keys.
- Calculate the shared secret: (Shared_Public^Private) MOD P
    - For Bob (9^5) MOD 13 = 3
    - For Alice (2^4) MOD 13 = 3


### DSA - Digital Signature Algorithm

* Used for two operations
    - Signature Generation
        - Takes Data and a Private Key and produces output a **Signature**
    - Signature Verification (Output either True or False)
        - Data
        - Public Key (Correlating to the Private Key which created the Signature)
        - Signature

* If signature is valid, it proves Integrity and Authentication of data.