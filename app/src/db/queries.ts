import {query} from './db';
import type {User} from '../models/user';
import type {Tweet} from '../models/tweet';
import type {Follow} from '../models/follow';

export const createUser = async (username: string, email: string, password: string): Promise<User> => {
  const sql = 'INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *';
  const [user] = await query<User>(sql, [username, email, password]);
  return user;
};

export const getUserByUsername = async (username: string): Promise<User | null> => {
  const sql = 'SELECT id, username, email, password FROM users WHERE username = $1';
  const [user] = await query<User>(sql, [username]);
  return user || null;
};

export const getUserById = async (id: string): Promise<User | null> => {
  const sql = 'SELECT id, username, email, password FROM users WHERE id = $1';
  const [user] = await query<User>(sql, [id]);
  return user || null;
};

export const createTweet = async (tweet: Tweet): Promise<Tweet> => {
  const sql = 'INSERT INTO tweets (user_id, content) VALUES ($1, $2) RETURNING *';
  const [newTweet] = await query<Tweet>(sql, [tweet.userId, tweet.content]);
  return newTweet;
};

export const getUserTweets = async (userId: string): Promise<Tweet[]> => {
  const sql = 'SELECT id, user_id, content, created_at FROM tweets WHERE user_id = $1';
  return await query<Tweet>(sql, [userId]);
};

export const createFollow = async (followerId: string, followedId: string): Promise<Follow> => {
  const sql = 'INSERT INTO follows (follower_id, following_id) VALUES ($1, $2) RETURNING *';
  const [follower] = await query<Follow>(sql, [followerId, followedId]);
  return follower;
};

export const deleteFollow = async (followerId: string, followedId: string): Promise<void> => {
  const sql = 'DELETE FROM follows WHERE follower_id = $1 AND following_id = $2';
  await query(sql, [followerId, followedId]);
};
