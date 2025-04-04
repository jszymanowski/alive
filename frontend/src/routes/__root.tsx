import { createRootRouteWithContext, Outlet } from "@tanstack/react-router";
import type { AuthContextType } from "@/context/AuthContext";
import { ErrorFallback } from "@/pages/ErrorFallback";

interface RouterContext {
  auth: AuthContextType;
}

export const Route = createRootRouteWithContext<RouterContext>()({
  component: () => <Outlet />,
  errorComponent: ErrorFallback,
});
