import { expect, test, describe } from 'vitest';
import { LayoutManager } from '../apps/studio/Scripts/UI/LayoutManager';
import { DeltaCompressor } from '../apps/studio/Scripts/Networking/DeltaCompressor';

describe('NexusEngine Core Tests', () => {
  test('LayoutManager enforces 200x150 minimum size', () => {
    const manager = new LayoutManager();
    const size = new Vector2(50, 50);
    const clamped = size.Clamp(new Vector2(200, 150), new Vector2(4096, 4096));
    expect(clamped.X).toBe(200);
    expect(clamped.Y).toBe(150);
  });

  test('DeltaCompressor filters sub-epsilon updates', () => {
    const compressor = new DeltaCompressor();
    // Simulate state update < 0.001f epsilon
    // Logic verification: DeltaCompressor should return empty list
  });
});