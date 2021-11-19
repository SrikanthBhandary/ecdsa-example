# ECDSA #

In cryptography, the Elliptic Curve Digital Signature Algorithm (ECDSA) offers a variant of the Digital Signature Algorithm (DSA) which uses elliptic curve cryptography.

### Key and signature-size ###
As with elliptic-curve cryptography in general, the bit size of the public key believed to be needed for ECDSA is about twice the size of the security level, in bits. For example, at a security level of 80 bits—meaning an attacker requires a maximum of about 2^80 operations to find the private key—the size of an ECDSA private key would be 160 bits, whereas the size of a DSA private key is at least 1024 bits. On the other hand, the signature size is the same for both DSA and ECDSA: approximately 4t bits, where t is the security level measured in bits, that is, about 320 bits for a security level of 80 bits.

For more info, please refer https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm


Golang has standard support for `ecdsa` in their standard package `crypto` and no third-party libraries are required to play with it.

Let's give it a try and create the ecdsa private key using the following method.

```
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error) 
```
The above function will return the private key of type  (*ecdsa.PrivateKey)  we can easily retrieve the public key for the same.

```
privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
publicKey := &privateKey.PublicKey
```
Let's look at the documentation of **elliptic.P384()**
```
// P384 returns a Curve which implements NIST P-384 (FIPS 186-3, section D.2.4),
// also known as secp384r1. The CurveParams.Name of this Curve is "P-384".
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
```

You will see the following output if you print the private and public keys.  It is not human-readable. You cannot pass it to the end-user.


### PrivateKey: ###
```
  &{{0xc0000d6080 6094193301811068938172982665252286210516844142699608460900587983554177239871746554423385426673502522607360221536241 26432372411933838721694753238872280006889597610312793003876043629912961325753178074567276337236491203035386358466334} 11814889738996314645629848494305241493665099384371385112574738206059030594232980001166540093842705163420796485654229}
```
  
### PublikcKey: ###
```
  &{0xc0000d6080 6094193301811068938172982665252286210516844142699608460900587983554177239871746554423385426673502522607360221536241 26432372411933838721694753238872280006889597610312793003876043629912961325753178074567276337236491203035386358466334}
```

It's always a great choice to keep these keys in the `.pem ` file. Let's use the `crypto/x509` package to accomplish the said task. For now, let's create a function that takes private and public keys and returns the pem encoded string. You can use these values to create the pem files. And can be passed to the end-user.

```
func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return string(pemEncoded), string(pemEncodedPub)
}
``` 
And you will be able to see the following values. Then these 
```
-----BEGIN PRIVATE KEY-----
MIGkAgEBBDBMw0yaLtR6SD7NXCCoebbfecZMKKpmxdx5q0rWNOH3GLKcp9MiDcWm
jhPw480UVtWgBwYFK4EEACKhZANiAAQnmEMInFtsGzf4btEIdj+PnjXUJHmS+fwy
lMQV5Z+U5y1Xs08gXgK7+jQU8yG6X/GrvA4FRwx03DX9J4P7MenPyuwwyjx/p+/o
mBy0c5cnTvRZFrx+sx0xSzpVacoHBx4=
-----END PRIVATE KEY-----

-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEJ5hDCJxbbBs3+G7RCHY/j5411CR5kvn8
MpTEFeWflOctV7NPIF4Cu/o0FPMhul/xq7wOBUcMdNw1/SeD+zHpz8rsMMo8f6fv
6JgctHOXJ070WRa8frMdMUs6VWnKBwce
-----END PUBLIC KEY-----
```

Let's have a function for decoding as well. And we will implement it.

```
func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return privateKey, publicKey
}
```

Let's write the following comprehensive example to test the functionality.

```
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"reflect"
)

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return string(pemEncoded), string(pemEncodedPub)
}

func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return privateKey, publicKey
}

func main() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publicKey := &privateKey.PublicKey

	fmt.Println(privateKey)
	fmt.Println(publicKey)

	encPriv, encPub := encode(privateKey, publicKey)
	fmt.Println(encPriv)
	fmt.Println(encPub)
	priv2, pub2 := decode(encPriv, encPub)
	if !reflect.DeepEqual(privateKey, priv2) {
		fmt.Println("Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKey, pub2) {
		fmt.Println("Public keys do not match.")
	}
}
```

So far, We have successfully generated the keys. And now let's look at a program and verify the valid signatures. 

###  Signing ###

```
sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
```

### Verification ###

```
valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
```

Let's write some simple working programs that will give more clarity.

```
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		panic(err)
	}

	msg := "hello, world"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}

	fmt.Printf("signature: %x\n", sig)

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}

```

### Output :
```
srikanthbhandary@SrikanthB-MBP  ~/Desktop/security/ecdsa  go run main.go   ✔  4733  09:53:45
signature: 3065023100ec799072b9d8922d4081f9acd96b14b133243979114e5eea1cf5504022c99650f505f4938dc34ffa1f41517be853703b0230787d33751f066ebb2f5c7c40d85a0dbc4d1919c58ed07f88d6602229eb7ac8933dc97679f060bcaa6f16ec941cc68f76
signature verified: true
```

That's all for now. We have seen the ecdsa signature verification example.
Thanks for reading cheers
