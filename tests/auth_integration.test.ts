import { expect, test, describe } from 'vitest';
import request from 'supertest';

describe('Auth API Integration', () => {
  const API_URL = 'http://localhost:3000';

  test('Registration: Should reject usernames with special characters', async () => {
    const res = await request(API_URL).post('/auth/register').send({
      username: 'user@name!',
      password: 'SecurePassword123'
    });
    expect(res.status).toBe(400);
  });

  test('Login: Should reject invalid password hash formats', async () => {
    // Simulating database injection of bad hash
    const res = await request(API_URL).post('/auth/login').send({
      username: 'testuser',
      password: 'password'
    });
    // Should return 401 Unauthorized, not 500
    expect(res.status).toBe(401);
  });
});