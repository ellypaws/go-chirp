import {Pool} from 'pg';
import dotenv from "dotenv";

dotenv.config();

const pool = new Pool({
  user: process.env.DB_USER,
  host: process.env.DB_HOST,
  database: process.env.DB_DBNAME,
  password: process.env.DB_PASSWORD,
  port: parseInt(process.env.DB_PORT || '5432', 10),
  ssl: process.env.DB_SSL === 'true' ? {rejectUnauthorized: false} : undefined,
});

export const query = async <T>(text: string, params?: any[]): Promise<T[]> => {
  const client = await pool.connect();
  try {
    const result = await client.query(text, params);
    return result.rows;
  } catch (error) {
    console.error(`Database query failed. Query: ${text}, Params: ${params}, Error: ${error}`);
    console.log(`Connected to user: ${process.env.DB_USER}, host: ${process.env.DB_HOST}, database: ${process.env.DB_DBNAME}, port: ${process.env.DB_PORT}`);
    throw error
  } finally {
    client.release();
  }
};

export const migrate = async () => {
  const createTablesSQL = `
    CREATE TABLE IF NOT EXISTS users (
      id SERIAL PRIMARY KEY,
      username TEXT UNIQUE NOT NULL,
      email TEXT UNIQUE NOT NULL,
      password TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS tweets (
      id SERIAL PRIMARY KEY,
      user_id INTEGER NOT NULL REFERENCES users(id),
      content TEXT NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    CREATE TABLE IF NOT EXISTS follows (
      id SERIAL PRIMARY KEY,
      follower_id INTEGER NOT NULL REFERENCES users(id),
      following_id INTEGER NOT NULL REFERENCES users(id)
    );
  `;
  try {
    await query(createTablesSQL);
    console.log('Tables created successfully');
  } catch (error) {
    console.error('Table creation failed', error);
  }
};
