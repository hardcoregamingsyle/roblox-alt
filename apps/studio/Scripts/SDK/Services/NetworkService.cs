using System;
using System.Buffers;
using System.IO;
using System.Security;
using ProtoBuf;

namespace NexusStudio.SDK.Services;

public class NetworkService
{
    private readonly ArrayPool<byte> _pool = ArrayPool<byte>.Shared;
    private const int MaxPacketSize = 8192;

    public void Send<T>(T payload)
    {
        byte[] buffer = _pool.Rent(MaxPacketSize);
        try {
            using var ms = new MemoryStream(buffer);
            Serializer.Serialize(ms, payload);
            if (ms.Position > MaxPacketSize) 
                throw new SecurityException("Payload exceeds 8KB limit");
            
            // Transport logic here
        } finally {
            _pool.Return(buffer);
        }
    }
}