using System;
using System.Security;

public class NetworkManager {
    private byte[] _fixedBuffer = new byte[8192];

    public void ProcessPacket(byte[] data) {
        if (data.Length > _fixedBuffer.Length) {
            throw new SecurityException("Buffer overflow attempt detected");
        }
        Array.Copy(data, _fixedBuffer, data.Length);
    }
}