import type {
  User,
  DatabaseIdentifier,
} from "@/types";

export const createDatabaseIdentifier = (): DatabaseIdentifier => Math.floor(Math.random() * 100_000);

export const createUser = (overrides?: Partial<User>): User => ({
  userId: createDatabaseIdentifier(),
  email: "montgomery.burns@snpp.com",
  name: "Monty",
  ...overrides,
});