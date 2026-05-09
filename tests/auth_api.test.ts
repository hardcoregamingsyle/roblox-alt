import { expect, test, describe } from 'vitest';
import request from 'supertest';

describe('Auth API Integration', () => {
  const api = 'http://localhost:3000';

  test('POST /auth/register - success', async () => {
    const res = await request(api)
      .post('/auth/register')
      .send({ username: 'testuser', password: 'password123' });
    expect(res.status).toBe(201);
  });

  test('POST /auth/login - invalid credentials', async () => {
    const res = await request(api)
      .post('/auth/login')
      .send({ username: 'testuser', password: 'wrongpassword' });
    expect(res.status).toBe(401);
  });

  test('POST /auth/register - payload limit (16KB)', async () => {
    const largePayload = { username: 'a', password: 'a'.repeat(20000) };
    const res = await request(api).post('/auth/register').send(largePayload);
    expect(res.status).toBe(400);
  });
});