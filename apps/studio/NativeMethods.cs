public class NativeMethods {
    // Fix: Buffer Overflow/Length validation
    public void ExecuteNative(string input) {
        if (string.IsNullOrEmpty(input) || input.Length > 2048) {
            throw new ArgumentException("Input too long");
        }
        // ... process ...
    }
}