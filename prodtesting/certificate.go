package prodtesting

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func GenPem() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	fmt.Println(err)

	SNLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	SN, err := rand.Int(rand.Reader, SNLimit)
	fmt.Println(err)

	tempplate := x509.Certificate{
		SerialNumber: SN,
		Subject: pkix.Name{
			Organization: []string{"test"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	tempplate.DNSNames = append(tempplate.DNSNames, "localhost")
	tempplate.EmailAddresses = append(tempplate.EmailAddresses, "test@test.com")

	certBytes, err := x509.CreateCertificate(rand.Reader, &tempplate, &tempplate, &privateKey.PublicKey, privateKey)
	fmt.Println(err)

	certFile, err := os.Create("cert.pem")
	fmt.Println(err)
	fmt.Println(pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}))
	fmt.Println(certFile.Close())

	keyFile, err := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	fmt.Println(err)
	// pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv.(*rsa.PrivateKey))})
	fmt.Println(pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}))
	fmt.Println(keyFile.Close())
}
