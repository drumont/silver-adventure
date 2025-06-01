package kms

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type Client struct {
	svc *kms.Client
}

func NewClient() (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		return nil, err
	}

	return &Client{
		svc: kms.NewFromConfig(cfg),
	}, nil
}

func (c *Client) EncryptDEK(kekAlias string, dek []byte) ([]byte, error) {
	out, err := c.svc.Encrypt(context.TODO(), &kms.EncryptInput{
		KeyId:     &kekAlias,
		Plaintext: dek,
	})
	if err != nil {
		return nil, err
	}
	return out.CiphertextBlob, nil
}

func (c *Client) DecryptDEK(blob []byte) ([]byte, string, error) {
	out, err := c.svc.Decrypt(context.TODO(), &kms.DecryptInput{
		CiphertextBlob: blob,
	})
	if err != nil {
		return nil, "", err
	}
	return out.Plaintext, *out.KeyId, nil
}
