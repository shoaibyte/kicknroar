import { z } from 'zod';

// Auth Validators
export const loginSchema = z.object({
  email: z.string().email('Invalid email address'),
  password: z.string().min(6, 'Password must be at least 6 characters'),
});

export const signupSchema = z.object({
  full_name: z.string().min(2, 'Name must be at least 2 characters').max(100, 'Name must be at most 100 characters'),
  email: z.string().email('Invalid email address'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
  phone: z.string().min(10, 'Phone number is required').max(20, 'Phone number must be at most 20 characters'),
});

// Match Validators
export const createMatchSchema = z.object({
  title: z.string().min(3, 'Title must be at least 3 characters').max(100, 'Title must be at most 100 characters'),
  venue_id: z.string().min(1, 'Venue is required'),
  match_date: z.string().min(1, 'Match date is required'),
  start_time: z.string().min(1, 'Start time is required'),
  duration_hours: z.number().min(0.5, 'Duration must be at least 0.5 hours').max(4, 'Duration must be at most 4 hours'),
  max_players: z.number().min(2, 'Must have at least 2 players').max(22, 'Maximum 22 players allowed'),
  cost_per_player: z.number().min(1, 'Cost per player must be at least 1'),
  match_type: z.enum(['casual', 'competitive', 'tournament'], {
    required_error: 'Match type is required',
  }),
  visibility: z.enum(['public', 'private', 'friends_only'], {
    required_error: 'Visibility is required',
  }),
  skill_level_required: z.enum(['beginner', 'intermediate', 'advanced', 'professional']).optional(),
  description: z.string().optional(),
  rules_notes: z.string().optional(),
});

export type LoginInput = z.infer<typeof loginSchema>;
export type SignupInput = z.infer<typeof signupSchema>;
export type CreateMatchInput = z.infer<typeof createMatchSchema>;

