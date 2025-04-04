import { createFileRoute } from "@tanstack/react-router";
import HomePage from "@/pages/HomePage";

import { redirectIfAuthenticated } from "@/route-helpers";

export const Route = createFileRoute("/")({
  beforeLoad: ({ context }) => {
    redirectIfAuthenticated(context.auth);
  },
  component: RouteComponent,
});

function RouteComponent() {
  return <HomePage />;
}
