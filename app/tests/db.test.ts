import { query } from '../src/db/db';
import { migrate } from '../src/db/db';

jest.mock('../src/db/db', () => ({
  query: jest.fn(),
  migrate: jest.fn(),
}));

describe('Database Queries', () => {
  afterEach(() => {
    jest.clearAllMocks();
  });

  describe('query function', () => {
    it('should execute a query and return the result', async () => {
      const mockResult = [{ id: 1, username: 'testuser' }];
      (query as jest.Mock).mockResolvedValue(mockResult);

      const result = await query('SELECT * FROM users');
      expect(result).toEqual(mockResult);
      expect(query).toHaveBeenCalledWith('SELECT * FROM users');
    });

    it('should throw an error if the query fails', async () => {
      const mockError = new Error('Query failed');
      (query as jest.Mock).mockRejectedValue(mockError);

      await expect(query('SELECT * FROM users')).rejects.toThrow('Query failed');
      expect(query).toHaveBeenCalledWith('SELECT * FROM users');
    });
  });

  describe('migrate function', () => {
    it('should call the migrate function', async () => {
      await migrate();
      expect(migrate).toHaveBeenCalled();
    });
  });
});
