# silver-adventure

**silver-adventure** is a Go example application demonstrating envelope encryption with AWS KMS. Sensitive data is encrypted locally using a Data Encryption Key (DEK), which is itself encrypted using a Key Encryption Key (KEK) managed by AWS KMS.

## Features

- Generates a random Data Encryption Key (DEK)
- Encrypts data locally with the DEK
- Encrypts the DEK with AWS KMS (KEK)
- Simulated in-memory storage of encrypted data and encrypted DEK
- Decrypts the DEK via AWS KMS and data via the DEK

## Project Structure

```
.
├── main.go                # Example usage of envelope encryption
├── crypto/
│   └── envelope.go        # DEK generation and encryption/decryption functions
├── kms/
│   └── client.go          # AWS KMS client for DEK encryption/decryption
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites

- Go 1.20 or higher
- AWS account with KMS access and a key created (alias or KeyID)
- AWS environment variables set (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`)

## Installation

```bash
git clone https://github.com/your-username/silver-adventure.git
cd silver-adventure
go mod tidy
```

## Usage

1. Edit the `kekAlias` variable in `main.go` to match your KMS key alias or KeyID.
2. Run the program:

```bash
go run main.go
```

You will see output showing storage, encryption, and decryption of data.

## Security

- Data Encryption Keys (DEK) are never stored in plaintext.
- Only the encrypted DEK is stored; the Key Encryption Key (KEK) remains in AWS KMS.
- Ensure your AWS credentials are secured and have the least privilege necessary for KMS operations.

