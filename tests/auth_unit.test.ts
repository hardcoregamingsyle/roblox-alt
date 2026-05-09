import { expect, test, describe } from 'vitest';
import { ValidateUsername } from '../packages/auth/utils';
import { HashPassword, VerifyPassword } from '../packages/auth/password';

describe('Auth Unit Tests', () => {
  test('Username validation (NFKC normalization & regex)', () => {
    expect(ValidateUsername('user123')).toBe(true);
    expect(ValidateUsername('us')).toBe(false); // Too short
    expect(ValidateUsername('user_name!')).toBe(false); // Invalid chars
  });

  test('Password hashing and verification', async () => {
    const pass = 'SuperSecurePass123!';
    const hash = await HashPassword(pass);
    expect(VerifyPassword(pass, hash)).toBe(true);
    expect(VerifyPassword('wrongpass', hash)).toBe(false);
  });
});