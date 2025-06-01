package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"silver-adventure/crypto"
	"silver-adventure/kms"
)

// Simulated in-memory storage
type Record struct {
	ID            string
	EncryptedData string
	EncryptedDEK  string
	KEKKeyID      string
	CreatedAt     time.Time
}

var storage = map[string]Record{}

func main() {
	const kekAlias = "alias/silver-adventure-key" // Replace with your KEK alias or KeyID in AWS KMS
	plaintext := []byte("this is top secret data")

	// Create AWS KMS client
	kmsClient, err := kms.NewClient()
	if err != nil {
		log.Fatalf("Failed to create KMS client: %v", err)
	}

	// Step 1: Generate a random DEK
	dek, err := crypto.GenerateDEK()
	if err != nil {
		log.Fatalf("Failed to generate DEK: %v", err)
	}

	// Step 2: Encrypt data locally with the DEK
	encryptedData, _, err := crypto.EncryptWithDEK(dek, plaintext)
	if err != nil {
		log.Fatalf("Failed to encrypt data: %v", err)
	}

	// Step 3: Encrypt DEK using AWS KMS
	encryptedDEK, err := kmsClient.EncryptDEK(kekAlias, dek)
	if err != nil {
		log.Fatalf("Failed to encrypt DEK with KMS: %v", err)
	}

	// Step 4: Save to "storage"
	id := uuid.NewString()
	storage[id] = Record{
		ID:            id,
		EncryptedData: base64.StdEncoding.EncodeToString(encryptedData),
		EncryptedDEK:  base64.StdEncoding.EncodeToString(encryptedDEK),
		KEKKeyID:      kekAlias,
		CreatedAt:     time.Now(),
	}

	fmt.Printf("✅ Encrypted record stored with ID: %s\n\n", id)

	// Step 5: Retrieve and decrypt
	record := storage[id]

	// Decode from base64
	encryptedDataBytes, _ := base64.StdEncoding.DecodeString(record.EncryptedData)
	encryptedDEKBytes, _ := base64.StdEncoding.DecodeString(record.EncryptedDEK)

	// Decrypt DEK with KMS
	decryptedDEK, usedKeyID, err := kmsClient.DecryptDEK(encryptedDEKBytes)
	if err != nil {
		log.Fatalf("Failed to decrypt DEK: %v", err)
	}

	fmt.Printf("🔐 DEK decrypted with KMS KeyID: %s\n", usedKeyID)

	// Decrypt data with DEK
	decryptedData, err := crypto.DecryptWithDEK(decryptedDEK, encryptedDataBytes)
	if err != nil {
		log.Fatalf("Failed to decrypt data: %v", err)
	}

	fmt.Printf("📦 Decrypted data: %s\n", decryptedData)
}
