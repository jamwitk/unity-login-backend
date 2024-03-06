using System;
using System.IO;
using System.Security.Cryptography;
using System.Text;

namespace Helpers
{
    public static class Crypto
    {
        // Adjust these for your desired security level
        private const int KeySize = 256;  // 128, 192, or 256 bits
        private const int DerivationIterations = 1000; // Slows down brute-force attacks

        public static string Encrypt(string plainText, string passPhrase)
        {
            // Salt and IV (initialization vector) are randomly generated each time
            byte[] saltBytes = GenerateRandomBytes(KeySize / 8); // 32 bytes for 256-bit key
            byte[] ivBytes = GenerateRandomBytes(16); // AES block size is 16 bytes

            using (var password = new Rfc2898DeriveBytes(passPhrase, saltBytes, DerivationIterations))
            {
                byte[] keyBytes = password.GetBytes(KeySize / 8);

                using (var symmetricKey = Aes.Create())
                {
                    symmetricKey.BlockSize = 128; // AES default
                    symmetricKey.Mode = CipherMode.CBC; // Common, requires IV
                    symmetricKey.Padding = PaddingMode.PKCS7; // Adds padding if needed

                    using (var encryptor = symmetricKey.CreateEncryptor(keyBytes, ivBytes))
                    {
                        using (var memoryStream = new MemoryStream())
                        {
                            using (var cryptoStream = new CryptoStream(memoryStream, encryptor, CryptoStreamMode.Write))
                            {
                                using (var streamWriter = new StreamWriter(cryptoStream))
                                {
                                    streamWriter.Write(plainText);
                                }
                                byte[] cipherTextBytes = memoryStream.ToArray();
                                byte[] combinedBytes = CombineArrays(saltBytes, ivBytes, cipherTextBytes);
                                return Convert.ToBase64String(combinedBytes);
                            }
                        }
                    }
                }
            }
        }

        public  static string Decrypt(string cipherText, string passPhrase)
        {
            byte[] combinedBytes = Convert.FromBase64String(cipherText);
            byte[] saltBytes = ExtractBytes(combinedBytes, 0, KeySize / 8);
            byte[] ivBytes = ExtractBytes(combinedBytes, KeySize / 8, 16);
            byte[] cipherTextBytes = ExtractBytes(combinedBytes, (KeySize / 8) + 16, combinedBytes.Length - ((KeySize / 8) + 16));

            using (var password = new Rfc2898DeriveBytes(passPhrase, saltBytes, DerivationIterations))
            {
                byte[] keyBytes = password.GetBytes(KeySize / 8);

                using (var symmetricKey = Aes.Create())
                {
                    symmetricKey.BlockSize = 128;
                    symmetricKey.Mode = CipherMode.CBC;
                    symmetricKey.Padding = PaddingMode.PKCS7;

                    using (var decryptor = symmetricKey.CreateDecryptor(keyBytes, ivBytes))
                    {
                        using (var memoryStream = new MemoryStream(cipherTextBytes))
                        {
                            using (var cryptoStream = new CryptoStream(memoryStream, decryptor, CryptoStreamMode.Read))
                            {
                                using (var streamReader = new StreamReader(cryptoStream))
                                {
                                    return streamReader.ReadToEnd();
                                }
                            }
                        }
                    }
                }
            }
        }

        // Helper methods (omitted for brevity; see below for explanations)
        private static byte[] GenerateRandomBytes(int length)
        {
            using var rng = RandomNumberGenerator.Create();
            var bytes = new byte[length];
            rng.GetBytes(bytes);
            return bytes;
        }
        private static byte[] CombineArrays(byte[] a1, byte[] a2, byte[] a3)
        {
            byte[] combined = new byte[a1.Length + a2.Length + a3.Length];
            Buffer.BlockCopy(a1, 0, combined, 0, a1.Length);
            Buffer.BlockCopy(a2, 0, combined, a1.Length, a2.Length);
            Buffer.BlockCopy(a3, 0, combined, a1.Length + a2.Length, a3.Length);
            return combined;
        }
        private static byte[] ExtractBytes(byte[] source, int startIndex, int length)
        {
            byte[] result = new byte[length];
            Buffer.BlockCopy(source, startIndex, result, 0, length);
            return result;
        }
    }
}