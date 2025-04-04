import { redirect } from "@tanstack/react-router";
import type { AuthContextType } from "@/context/AuthContext";

const routes = {
  login: '/auth/login',
  summary: '/summary'
}

const redirectIfAuthenticated = (auth: AuthContextType) => {
  if (auth.isAuthenticated) {
    throw redirect({
      to: routes.summary,
      search: {
        redirect: location.href,
      },
    });
  }
};

const redirectIfUnauthenticated = (auth: AuthContextType) => {
  if (!auth.isAuthenticated) {
    throw redirect({
      to: routes.login,
      search: {
        redirect: location.href,
      },
    });
  }
};

export default routes;
export { redirectIfAuthenticated, redirectIfUnauthenticated };