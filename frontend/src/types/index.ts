export type DatabaseIdentifier = number;

export interface User {
  id: DatabaseIdentifier;
  email: string;
  name: string;
}
