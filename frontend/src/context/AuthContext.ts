import { createContext } from "react";
import type { User } from "@/types";

export interface AuthContextType {
  isAuthenticated: boolean | null;
  user: User | null;
  login: (token: string) => void;
  logout: () => void;
  loading: boolean;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);
