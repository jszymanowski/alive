import AuthProvider from "@/context/AuthProvider";
import ThemeProvider from "@/context/ThemeProvider";

// import { useAuth } from "@/context/use-auth";

interface AppProvidersProps {
  children: React.ReactNode;
}

const AppProviders = ({ children }: AppProvidersProps) => {
  return (
    <ThemeProvider defaultTheme="light" storageKey="still-kicking-theme">
      <AuthProvider>{children}</AuthProvider>
    </ThemeProvider>
  );
};

const InnerApp = () => {
  // const auth = useAuth();
  return <>hi</>;
};

const App = () => (
  <AppProviders>
    <InnerApp />
  </AppProviders>
);

export default App;
