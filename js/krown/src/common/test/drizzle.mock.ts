import { DrizzleDB } from '../../drizzle/types/drizzle';

export const setupMockDrizzle = (): Partial<DrizzleDB> => ({
  select: jest.fn().mockReturnValue({
    from: jest.fn().mockReturnValue({
      where: jest.fn().mockReturnValue({
        limit: jest.fn().mockResolvedValue(null),
      }),
    }),
  }),
  insert: jest.fn().mockReturnValue({
    values: jest.fn().mockReturnValue({
      returning: jest.fn().mockResolvedValue([null]),
    }),
  }),
  transaction: jest
    .fn()
    .mockImplementation(
      async <T>(
        callback: (tx: Partial<DrizzleDB>) => Promise<T>,
      ): Promise<T> => {
        const tx: Partial<DrizzleDB> = {
          insert: jest.fn().mockReturnValue({
            values: jest.fn().mockReturnValue({
              returning: jest.fn().mockResolvedValue([null]),
            }),
          }),
        };
        return await callback(tx);
      },
    ),
});

export const setMockSelect = (
  mockDrizzle: Partial<DrizzleDB>,
  returnValue: any,
) => {
  (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
    from: jest.fn().mockReturnValue({
      where: jest.fn().mockReturnValue({
        limit: jest.fn().mockResolvedValue(returnValue),
      }),
      innerJoin: jest.fn().mockReturnValue({
        where: jest.fn().mockResolvedValue(returnValue),
      }),
    }),
  });
};
