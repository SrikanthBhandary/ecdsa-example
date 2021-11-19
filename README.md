# ecdsa-example

In cryptography, the Elliptic Curve Digital Signature Algorithm (ECDSA) offers a variant of the Digital Signature Algorithm (DSA) which uses elliptic curve cryptography.

### Key and signature-size ###
As with elliptic-curve cryptography in general, the bit size of the public key believed to be needed for ECDSA is about twice the size of the security level, in bits.For example, at a security level of 80 bits—meaning an attacker requires a maximum of about 2^80 operations to find the private key—the size of an ECDSA private key would be 160 bits, whereas the size of a DSA private key is at least 1024 bits. On the other hand, the signature size is the same for both DSA and ECDSA: approximately 4t bits, where t is the security level measured in bits, that is, about 320 bits for a security level of 80 bits.

For more info, please refer https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm


Golang has standard support for `ecdsa` in thier standard package `crypto` and no third party libraries are required to play with it.

Let's give a try and create the ecdsa private key using the following method.

```
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error) 
```
The above function will return the privatekey of type  (*ecdsa.PrivateKey)  and we can easily retrive the public key for the same.

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

If you try to print the values of privateKey and public key, you are going to see the following output and it is not readable and you cannot pass it to the end user.


### PrivateKey: ###
```
  &{{0xc0000d6080 6094193301811068938172982665252286210516844142699608460900587983554177239871746554423385426673502522607360221536241 26432372411933838721694753238872280006889597610312793003876043629912961325753178074567276337236491203035386358466334} 11814889738996314645629848494305241493665099384371385112574738206059030594232980001166540093842705163420796485654229}
```
  
### PublikcKey: ###
```
  &{0xc0000d6080 6094193301811068938172982665252286210516844142699608460900587983554177239871746554423385426673502522607360221536241 26432372411933838721694753238872280006889597610312793003876043629912961325753178074567276337236491203035386358466334}
```
