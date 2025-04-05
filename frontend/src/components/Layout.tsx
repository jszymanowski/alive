import { Box, Flex } from "@jszymanowski/breeze-react";

interface Props {
  children: React.ReactNode;
}

const Layout = ({ children }: Props) => (
  <Flex direction="col" className="min-h-screen w-full max-w-6xl">
    <Box as="main" width="full">
      {children}
    </Box>
  </Flex>
);

export default Layout;
