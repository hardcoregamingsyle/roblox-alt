// ... existing code
    public void OnDataReceived(byte[] data) {
        // ... validation
        try {
            // ... deserialization
        } finally {
            _pool.Return(data, clearArray: true); // Securely zero out buffer
        }
    }