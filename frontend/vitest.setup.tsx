import { beforeAll, afterAll, afterEach } from "vitest";
import { server } from "./tests/mocks/server";

beforeAll(() => server.listen({ onUnhandledRequest: "error" }));
afterAll(() => server.close());

afterEach(() => {
  localStorage.clear();
});