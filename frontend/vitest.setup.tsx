import { afterAll, afterEach, beforeAll } from "vitest";
import { server } from "./tests/mocks/server";

beforeAll(() => server.listen({ onUnhandledRequest: "error" }));
afterAll(() => server.close());

afterEach(() => {
  localStorage.clear();
});
