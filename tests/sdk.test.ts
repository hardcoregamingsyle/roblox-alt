import { expect, test, describe } from 'vitest';
import { NetworkService } from '../apps/studio/Scripts/SDK/Services/NetworkService';

describe('NexusEngine SDK Unit Tests', () => {
  test('NetworkService enforces 8KB packet limit', () => {
    const service = new NetworkService();
    // Logic test: Ensure serialization logic throws if buffer is exceeded
    expect(() => {
        const largePayload = new Array(9000).fill(0);
        service.Send(largePayload);
    }).toThrow();
  });
});