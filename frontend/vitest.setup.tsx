import { afterAll, afterEach, beforeAll } from "vitest";
import { server } from "./tests/mocks/server";

beforeAll(() => server.listen({ onUnhandledRequest: "error" }));

afterEach(() => {
  localStorage.clear();
  server.resetHandlers();
});

afterAll(() => server.close());
