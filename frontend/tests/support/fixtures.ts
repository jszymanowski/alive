import type { DatabaseIdentifier, User } from "@/types";

export const createDatabaseIdentifier = (): DatabaseIdentifier =>
  Math.floor(Math.random() * 100_000);

export const createUser = (overrides?: Partial<User>): User => ({
  id: createDatabaseIdentifier(),
  email: "montgomery.burns@snpp.com",
  name: "Monty",
  ...overrides,
});
