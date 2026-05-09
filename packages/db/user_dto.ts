import { z } from 'zod';

// FIX: Mass Assignment Prevention
export const UserUpdateSchema = z.object({
  username: z.string().min(3).max(32).optional(),
  password: z.string().min(8).optional(),
}).strict();