using System.Security.Cryptography;
using System.IO;

public class SceneBundle {
    private static readonly byte[] SecretKey = System.Text.Encoding.UTF8.GetBytes(System.Environment.GetEnvironmentVariable("BUNDLE_SECRET"));

    public static byte[] LoadAndVerify(byte[] data, byte[] signature) {
        using (var hmac = new HMACSHA256(SecretKey)) {
            byte[] computedHash = hmac.ComputeHash(data);
            if (!CryptographicOperations.FixedTimeEquals(computedHash, signature)) {
                throw new SecurityException("Invalid bundle signature");
            }
        }
        return data; // Proceed to safe deserialization
    }
}