import { Box, Flex, Text } from "@jszymanowski/breeze-react";

interface ErrorFallbackProps {
  error: Error;
}

export const ErrorFallback = ({ error }: ErrorFallbackProps) => {
  return (
    <Flex direction="col" align="center" className="mt-8">
      <Box variant="muted" width="full" className="mb-8 p-4">
        <Text variant="inherit" family="mono" align="center" size="xs">
          {error.message}
        </Text>
      </Box>
    </Flex>
  );
};
