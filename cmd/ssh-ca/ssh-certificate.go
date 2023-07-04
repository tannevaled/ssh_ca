package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type SSHCertificate struct {
	private_key ed25519.PrivateKey
	public_key  ssh.PublicKey
	certificate *ssh.Certificate
}

type SSHCertificateInterface interface {
	getPrivateKey() ed25519.PrivateKey
	getPublicKey() ssh.PublicKey
	getCertificate() *ssh.Certificate
	setCertificate(*ssh.Certificate)
	publicKeyFingerprintSHA256() string
	caPublicKeyFingerprintSHA256() string
}

func NewSSHCertificate() SSHCertificateInterface {
	_, __ed25519_private_key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Errorf("failed to create private key: %s\n", err)
		os.Exit(1)
	}
	__public_key, err := ssh.NewPublicKey(__ed25519_private_key.Public())
	if err != nil {
		fmt.Errorf("failed to create public key: %s\n", err)
		os.Exit(1)
	}
	__self := &SSHCertificate{
		private_key: __ed25519_private_key,
		public_key:  __public_key,
	}

	return __self
}

func (self *SSHCertificate) getPrivateKey() ed25519.PrivateKey {
	return self.private_key
}

func (self *SSHCertificate) getPublicKey() ssh.PublicKey {
	return self.public_key
}

func (self *SSHCertificate) getCertificate() *ssh.Certificate {
	return self.certificate
}

func (self *SSHCertificate) setCertificate(certificate *ssh.Certificate) {
	self.certificate = certificate
}

func (self *SSHCertificate) publicKeyFingerprintSHA256() string {
	return ssh.FingerprintSHA256(self.public_key)
}

func (self *SSHCertificate) caPublicKeyFingerprintSHA256() string {
	return ssh.FingerprintSHA256(self.certificate.SignatureKey)
}
