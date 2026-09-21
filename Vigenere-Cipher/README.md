# Vigenere Cipher
## Alur Program
### Vigenere Cipher Encryption
1. Input plaintext and key
2. Siapkan result string
3. Lakukan for loop untuk setiap karakter dalam plaintext dan dikurangi dengan int('A')
4. Lakukan for loop untuk setiap karakter dalam key dan dikurangi dengan int('A')
5. Hitung shift dengan rumus (plaintext_char + key_char) % 26
6. Ambil karakter dari alfabet dengan indeks shift
7. Tambahkan karakter ke result string
8. Ulangi proses untuk setiap karakter dalam plaintext
9. Return result string

### Vigenere Cipher Decryption
1. Input ciphertext and key
2. Siapkan result string
3. Lakukan for loop untuk setiap karakter dalam ciphertext dan dikurangi dengan int('A')
4. Lakukan for loop untuk setiap karakter dalam key dan dikurangi dengan int('A')
5. Hitung shift dengan rumus (ciphertext_char - key_char) % 26
6. Ambil karakter dari alfabet dengan indeks shift
7. Tambahkan karakter ke result string
8. Ulangi proses untuk setiap karakter dalam ciphertext
9. Return result string

### Main Function
1. Define plaintext and key
2. Panggil encrypt function dengan plaintext dan key
3. Print result string
4. Print Key
5. Panggil decrypt function dengan result string dan key
6. Print decrypted string
