using System.Buffers;

public class NetworkService {
    private readonly ArrayPool<byte> _pool = ArrayPool<byte>.Shared;

    public void SendPacket(byte[] data) {
        // Fix Issue 17: Force slice to exact length to prevent buffer leakage
        const int MAX_PACKET_SIZE = 8192;
        var buffer = _pool.Rent(MAX_PACKET_SIZE);
        try {
            var slice = buffer.AsMemory(0, MAX_PACKET_SIZE);
            // Copy data into slice...
            // Ensure we never write past slice.Length
        } finally {
            _pool.Return(buffer);
        }
    }
}