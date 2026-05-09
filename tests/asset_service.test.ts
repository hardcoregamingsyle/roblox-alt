import { expect, test, describe } from 'vitest';
import request from 'supertest';

describe('Asset Storage Service Integration', () => {
  const API_URL = 'http://localhost:3001';

  test('Upload URL Request: Should reject invalid file types', async () => {
    const res = await request(API_URL).post('/assets/upload').send({ fileType: 'exe' });
    expect(res.status).toBe(400);
  });

  test('Upload URL Request: Should accept valid file types', async () => {
    const res = await request(API_URL).post('/assets/upload').send({ fileType: 'model' });
    expect(res.status).toBe(200);
    expect(res.body).toHaveProperty('url');
    expect(res.body).toHaveProperty('assetId');
  });

  test('Upload URL Request: Should enforce payload size limits', async () => {
    const largeBody = { fileType: 'a'.repeat(2048) };
    const res = await request(API_URL).post('/assets/upload').send(largeBody);
    expect(res.status).toBe(400);
  });
});