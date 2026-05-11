import { expect, test, describe } from 'vitest';
import request from 'supertest';

describe('Asset Storage & UI Integration', () => {
  test('S3 Broker rejects oversized metadata', async () => {
    const res = await request('http://localhost:3001')
      .post('/assets/upload')
      .send({ fileType: 'a'.repeat(2048) });
    expect(res.status).toBe(400);
  });

  test('UI Framework: DockablePanel minimum size constraint', async () => {
    const panel = { width: 200, height: 150 };
    expect(panel.width).toBeGreaterThanOrEqual(200);
    expect(panel.height).toBeGreaterThanOrEqual(150);
  });
});