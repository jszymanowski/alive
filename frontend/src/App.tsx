<<<<<<< HEAD
function App() {
  return <>Hi</>;
}

export default App;
=======
import { RouterProvider, createRouter } from "@tanstack/react-router";

import AuthProvider from "@/context/AuthProvider";
import ThemeProvider from "@/context/ThemeProvider";

import NotFoundErrorPage from "@/pages/NotFoundErrorPage";

import { routeTree } from "@/routeTree.gen";
import { useAuth } from "@/context/useAuth";

const router = createRouter({
  routeTree,
  context: {
    auth: undefined!, // This will be set after we wrap the app in an AuthProvider
  },
  defaultNotFoundComponent: NotFoundErrorPage,
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

interface AppProvidersProps {
  children: React.ReactNode;
}

const AppProviders = ({ children }: AppProvidersProps) => {
  return (
    <ThemeProvider defaultTheme="light" storageKey="still-kicking-theme">
      <AuthProvider>{children}</AuthProvider>
    </ThemeProvider>
  );
}

const InnerApp = () => {
  const auth = useAuth();
  return <RouterProvider router={router} context={{ auth }} />;
}

const App = () => (
  <AppProviders>
    <InnerApp />
  </AppProviders>
);

export default App
>>>>>>> 018f6aa (Add AppProviders)
