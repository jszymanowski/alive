export type DatabaseIdentifier = number;

export interface User {
  userId: DatabaseIdentifier;
  email: string;
  name: string;
}
