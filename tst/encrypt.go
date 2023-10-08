package tst

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"io"
	"math/big"

	"github.com/mr-tron/base58"
)

func TTT() {

	// RSA
	DoRSA()

	// ECDSA
	DoECDSA()

	// AES
	DoAES()
}

func toBigInt2(s string) (*big.Int, *big.Int, error) {
	bs, err := base58.Decode(s)
	if err != nil {
		return nil, nil, err
	}

	n0 := new(big.Int).SetBytes(bs[:1]).Uint64()

	fmt.Printf("Total Length is %d ,First length is %d \n", len(bs), n0)

	a := new(big.Int).SetBytes(bs[1 : n0+1])
	b := new(big.Int).SetBytes(bs[n0+1:])

	return a, b, nil
}

func DoECDSA() {
	// D私钥

	// X,Y
	x, y, err := toBigInt2("3qpWRFcPFcf9hMSJ6QSVU37ak2FqjDCASR776r16kwvijrBcC8ut53ACGrdQ34JxYLLixxKEcWpo6fhs2ye1T4Tck")
	if err != nil {
		fmt.Printf("Get BigInt XY Failure :: %s \n", err.Error())
		return
	}

	fmt.Printf("(X,Y) :: (%s,%s)\n", x.String(), y.String())

	pub := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// R,S
	r, s, err := toBigInt2("3spaeUhqUAw7RYT5XzHPar7nvqNSQabBGFiFSe13efkvqtudSXqvu6MKv8FMAnELxfHDnfNu4zBEZ9DMmorRMWM4v")
	if err != nil {
		fmt.Printf("Get BigInt RS Failure :: %s \n", err.Error())
		return
	}
	fmt.Printf("(R,S) :: (%s,%s)\n", r.String(), s.String())

	str := "www.nsmei.com :: I am from flutter !"
	h := sha256.New()
	io.WriteString(h, str)
	res := h.Sum(nil)

	fmt.Printf("%#v\n", ecdsa.Verify(pub, res, r, s))
}

func DoRSA() {

	/*

		// Generate Private Key
		priKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			fmt.Printf("rsa.GenerateKey() Failure :: %s\n", err.Error())
			return
		}

		// Encode Private Key to Base58
		pri58Key := base58.Encode(x509.MarshalPKCS1PrivateKey(priKey))

		fmt.Printf("Private Key %s\n", pri58Key)
		// 2SM2qooAi1HUDry7Bk76gDBM6MxU5so2jrSDoNLMoGiCBKiaEzSc8oCoQdv7y4qDM2p2f6Qon1u2MWMRv7vWJ96dRPEJwiDU8sda6h8DUDxBdWmMsPhr4khQy3KLGFuV4ErbUNe5p6VRb2yp5DyUy2kmuGJP7Z1Dh3Eeus9FHhzrDp2pLDV84YtM5piSYuhPiKmn4pKXqYmQp1zrtBrwVruMtdTts57EtYUUnAZbXEKVoJksUoGsLVkGhXyJ4A45AZaGX8vFnHhR26xVcX81HjzUHCgBKV8UYe1duYC1ckvvA2TuPPBPnFriTGLvZwnh7qDYmcZhkJgzziExXgWTXTJSpmBgkXg4hsKQuEqavxwgGkJ4GFT5GHmyNppJJFmDpSs1EWLu2igqNADwvNrf2W4kd4pi2wAdXCQecnUJVRNCQ1yxx7iaib6GMNvtXiGhNUmmXJugTEP8eDgJvxAQuhV28DTQMCN6P3gsEb9qczCPxfj22jQqYouDovzSxNReHtfY32RuMpV3jRjBLo3Cn261Y9PUREoUfj8PR7suFUKyqjjjn8ZXYQgWUdY84cechsLgDZFve9pcB4iTjRNRTZq6aFDtntNp2hQRbPZpjfTVdadzc7YS5v4ke5pi46ueHiJaxE9MspooRMsV33XzCjLmxpbH68MvBErKoWnqtHsnDd6UnfYYWKCYHHaKfr1Ya7MSF8t4wZKQ8sdwSV37Gk8soC8uchYrp4QHPQRe38EH8Z1dP4Up26x1Jz2yUs6c5sTZvvJ1VRKiVvwWik9f2LucCZ5yiebCmBKE98pf6ZM1DnEEASeTeSKqwo538oqfc234oSsfH4TQvQ6RpHvG1AprdC4iwe6zXVAPYCqnaAj3D4odavGMZzb9xxEqM3fDJxtQGBp1pbRCiUZ8tzWF9iBta4iuTe4W46CE2LkvgFURQCfcGKhGfqxQY4tjfoeLVDJpZSquC46rtvyyzngDeYXdiZoxZmt43GQaVmZRtUsbfVLTo2sWDz5fF2WHhWBEJKLTPoEUJaTy5zP8e2jJGBxLgMtNDH74EyK8B6qLZXjQK9Mejpypy3RcTxFphv2JzsTpPVUgiNtd15mjip8E6ibysz8H81gAb6KQeUAZEMo9vKEfvFWPWm47eetcpCDcn4DsBY4Xe4EiSEjGt8Qyuxs2arrqd3y78eCj3R3gGAhXZTMr755QQ1tReKRQfwgXaTYE6keSGWWGrZ89M1YFhu9gPVx7vgoYNd1mirxFPavoGwX963CKdzPzSpLRsFGnni1gD5xVKDps7P6SEW8XcpqtUEqbtQgsUoMnFb1WTzSkcR8dEitGh8X5zNpqkWHYyPq9baiQ4iTKRntUXScLJo7zAU13Azb2ZzZ6pd6AW8KG8aBoHfBiYoHAwCEb68qV6GnAuYWd6wnRPFDBycT4gPccBTv8V7EZ2gWHdNCuNffDhjnN5rNeZGJZMZJ8LppGyP7xyiFs3PFgmBC7A3XwHCrnowJkAKEu1nSSYD5nDHCHm5NkD7nd5VzGNdKHP46J6LjG8hR8N2ZHqwQj97SgQ1P3f1ujL9Hik29Dvgpes365iwPmcoVtYxgXHX4xf1NN5pKhjNk5z4eZsFct4L9z9UiVv5G

	*/

	// Key Private Key from Base58
	pri58Key := "2SM2qooAi1HUDry7Bk76gDBM6MxU5so2jrSDoNLMoGiCBKiaEzSc8oCoQdv7y4qDM2p2f6Qon1u2MWMRv7vWJ96dRPEJwiDU8sda6h8DUDxBdWmMsPhr4khQy3KLGFuV4ErbUNe5p6VRb2yp5DyUy2kmuGJP7Z1Dh3Eeus9FHhzrDp2pLDV84YtM5piSYuhPiKmn4pKXqYmQp1zrtBrwVruMtdTts57EtYUUnAZbXEKVoJksUoGsLVkGhXyJ4A45AZaGX8vFnHhR26xVcX81HjzUHCgBKV8UYe1duYC1ckvvA2TuPPBPnFriTGLvZwnh7qDYmcZhkJgzziExXgWTXTJSpmBgkXg4hsKQuEqavxwgGkJ4GFT5GHmyNppJJFmDpSs1EWLu2igqNADwvNrf2W4kd4pi2wAdXCQecnUJVRNCQ1yxx7iaib6GMNvtXiGhNUmmXJugTEP8eDgJvxAQuhV28DTQMCN6P3gsEb9qczCPxfj22jQqYouDovzSxNReHtfY32RuMpV3jRjBLo3Cn261Y9PUREoUfj8PR7suFUKyqjjjn8ZXYQgWUdY84cechsLgDZFve9pcB4iTjRNRTZq6aFDtntNp2hQRbPZpjfTVdadzc7YS5v4ke5pi46ueHiJaxE9MspooRMsV33XzCjLmxpbH68MvBErKoWnqtHsnDd6UnfYYWKCYHHaKfr1Ya7MSF8t4wZKQ8sdwSV37Gk8soC8uchYrp4QHPQRe38EH8Z1dP4Up26x1Jz2yUs6c5sTZvvJ1VRKiVvwWik9f2LucCZ5yiebCmBKE98pf6ZM1DnEEASeTeSKqwo538oqfc234oSsfH4TQvQ6RpHvG1AprdC4iwe6zXVAPYCqnaAj3D4odavGMZzb9xxEqM3fDJxtQGBp1pbRCiUZ8tzWF9iBta4iuTe4W46CE2LkvgFURQCfcGKhGfqxQY4tjfoeLVDJpZSquC46rtvyyzngDeYXdiZoxZmt43GQaVmZRtUsbfVLTo2sWDz5fF2WHhWBEJKLTPoEUJaTy5zP8e2jJGBxLgMtNDH74EyK8B6qLZXjQK9Mejpypy3RcTxFphv2JzsTpPVUgiNtd15mjip8E6ibysz8H81gAb6KQeUAZEMo9vKEfvFWPWm47eetcpCDcn4DsBY4Xe4EiSEjGt8Qyuxs2arrqd3y78eCj3R3gGAhXZTMr755QQ1tReKRQfwgXaTYE6keSGWWGrZ89M1YFhu9gPVx7vgoYNd1mirxFPavoGwX963CKdzPzSpLRsFGnni1gD5xVKDps7P6SEW8XcpqtUEqbtQgsUoMnFb1WTzSkcR8dEitGh8X5zNpqkWHYyPq9baiQ4iTKRntUXScLJo7zAU13Azb2ZzZ6pd6AW8KG8aBoHfBiYoHAwCEb68qV6GnAuYWd6wnRPFDBycT4gPccBTv8V7EZ2gWHdNCuNffDhjnN5rNeZGJZMZJ8LppGyP7xyiFs3PFgmBC7A3XwHCrnowJkAKEu1nSSYD5nDHCHm5NkD7nd5VzGNdKHP46J6LjG8hR8N2ZHqwQj97SgQ1P3f1ujL9Hik29Dvgpes365iwPmcoVtYxgXHX4xf1NN5pKhjNk5z4eZsFct4L9z9UiVv5G"
	genPri58Key, err := base58.Decode(pri58Key)
	if err != nil {
		fmt.Printf("base58.Decode() Failure :: %s\n", err.Error())
		return
	}

	// Generate Private Key
	genPriKey, err := x509.ParsePKCS1PrivateKey(genPri58Key)
	if err != nil {
		fmt.Printf("x509.ParsePKCS1PrivateKey() Failure :: %s\n", err.Error())
		return
	}

	// Get Public Key
	pubKey := &genPriKey.PublicKey

	// Encode Public Key to Base58
	pub58Key := base58.Encode(x509.MarshalPKCS1PublicKey(pubKey))
	fmt.Printf("Public Key %s\n", pub58Key)
	// 4e1BUTgGBfqVXCyqBn6cj9jpNhzFartL5sYe1Bdt4GRY3MBN1NzDyQu7bKaEKgeqRoiqScQ7BFZfkVQFv9nqnkYysqSJxBVeRLqhwEMM8S63mYwAXbDrqQgA7kS2Zoa37ddTQm2ktQrUDiMYwNvgJzM8aLw7FJreY4QnWpEvo4CS3wE6YCkgTDPZpzsKT8rcvbWV4PJXSeJrp7VaqychR3JLKi9c4yQGe3Hd7zKFRAYqGJFvo3tYUDSPbLc8w2swzRz8cTNHUUgTdCU3gt3HR2dK1KkCBgbEwgHoBjbNHnKBzKmco2Zuma58GnrJP3k3W48FBK6yNXZ6S6iKTiRzEvptDydzM6Mwrk9m7SHLFTYR4QiWc

	cipher58Text := "5MKcmdNYZdi3khS8PUbcrZJUhdpNn5KFMDWkC5J2URkXuoj8fvaKqkY7ka6CJkyZmwvy1ffJGqpdZmXFtKR6qg9xLiu3u1zEGU2MUq2io91dFFMZVFaBXyBqh3DUrAxK24YyJF3vMivZ4Vo21sZvTZDQ7bN9R4PDD8a4dNHHNLcQwxaMV3KDPb8zeBfvMSxjGLVWkR459WCpYPQeYCzo4UJqXUyEWKoNRdoBt7wgunPjweatcvcnwMM9dKTur62Gs794e5PyFu8YViXLPL1zXQeL9iAmhb3hJB3bAsYZ5MTpoTk3fSac3s3vXCwou53f7NHp3h4jL2vUCmCbawCKFmSHpD6mRu"
	cipherText, err := base58.Decode(cipher58Text)
	if err != nil {
		fmt.Printf("base58.Decode() Failure :: %s\n", err.Error())
		return
	}

	plainText, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, genPriKey, cipherText, nil)
	if err != nil {
		fmt.Printf("rsa.DecryptPKCS1v15() Failure :: %s\n", err.Error())
		return
	}

	fmt.Printf("Plain Text %s\n", string(plainText))
}

func DoAES() {
	key, err := base58.Decode("DGy5bu7zKUAuxcyHcWCeyg")
	if err != nil {
		fmt.Printf("base58.Decode(KEY) Failure :: %s\n", err.Error())
		return
	}

	iv, err := base58.Decode("GpnyUiCwv5u7ngjMWd8jci")
	if err != nil {
		fmt.Printf("base58.Decode(IV) Failure :: %s\n", err.Error())
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("aes.NewCipher() Failure :: %s\n", err.Error())
		return
	}

	plainText := []byte("ExamplePlainText")

	// Encrypt
	if len(plainText)%aes.BlockSize != 0 {
		fmt.Printf("Plain Text :: %s\n", "plaintext is not a multiple of the block size")
		return
	}

	cipherText := make([]byte, len(plainText))

	enMode := cipher.NewCBCEncrypter(block, iv)
	enMode.CryptBlocks(cipherText[:], plainText)
	fmt.Printf("AES Cipher Text :: %s\n", base58.Encode(cipherText))

	// Decrypt
	if len(cipherText) < aes.BlockSize || len(cipherText)%aes.BlockSize != 0 {
		fmt.Printf("Cipher Text :: %s\n", "ciphertext is not a multiple of the block size")
		return
	}

	orgPlainText := make([]byte, len(cipherText))

	deMode := cipher.NewCBCDecrypter(block, iv)
	deMode.CryptBlocks(orgPlainText, cipherText)

	fmt.Printf("AES Plain Text :: %s\n", orgPlainText)

	// MD5/HMAC
	mac := hmac.New(md5.New, key)
	mac.Write(plainText)
	h5Message := mac.Sum(nil)

	fmt.Printf("AES MD5/HMAC Message :: %s\n", base58.Encode(h5Message))
}
