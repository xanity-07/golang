package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func IntroHashing() {
	password := "password123"
	// //* Hash the password takes a byte slice as an arg
	// hash256 := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(hash256)
	// fmt.Println(hash512)

	// //* Print hex value
	// fmt.Printf("SHA-256 Hash hex value: %x\n", hash256)
	// fmt.Printf("SHA-512 Hash hex value: %x\n", hash512)

	//? Generate salt
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// ? Hash password with salt
	signupHash := hashPassword(password, salt)
	// ? salt string
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt string:", saltStr) // * simulate storing in db
	fmt.Println("Hash:", signupHash)     // * simulate storing in db

	//* Verify Password
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}
	loginHash := hashPassword("password123", decodedSalt)
	if loginHash == signupHash {
		fmt.Println("Logged in")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}

}

// TODO Make salt func return byte slice and err
// TODO Make salt slice
// TODO Read return salt and err
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, err
}

// TODO Make hashPassword func takes a password and a salt byte slice
// TODO returns a string
// TODO merge password and salt make them the same type
// TODO hash it and return a string with Encoding to string
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
