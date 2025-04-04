import { createRootRouteWithContext, Outlet } from "@tanstack/react-router";

import { ErrorFallback } from "@/pages/ErrorFallback";
import type { AuthContextType } from "@/context/AuthContext";

interface RouterContext {
  auth: AuthContextType;
}

export const Route = createRootRouteWithContext<RouterContext>()({
  component: () => <Outlet />,
  errorComponent: ErrorFallback,
});
