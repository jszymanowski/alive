import { beforeEach, describe, expect, test } from "vitest";

import { createUser, fetchCurrentUser } from "@/api/userApi";

describe("User API", () => {
  beforeEach(() => {
    localStorage.clear();
    localStorage.setItem("still-kicking-auth-token", "FAKE-TOKEN");
  });

  describe("fetchCurrentUser", () => {
    test("retrieves the current user", async () => {
      const result = await fetchCurrentUser();
      expect(result.email).to.equal("montgomery.burns@snpp.com");
      expect(result.name).to.equal("Monty");
    });

    test("fails to retrieve the current user without a token", async () => {
      localStorage.clear(); // Remove the token
      await expect(fetchCurrentUser()).rejects.toThrowError("An unexpected error occurred");
    });
  });

  describe("createUser", () => {
    test("creates the current user", async () => {
      const createParams = { name: "Waylon", email: "waylon.smithers@snpp.com" };
      const result = await createUser(createParams);

      expect(result.name).to.equal("Waylon");
      expect(result.email).to.equal("waylon.smithers@snpp.com");
    });

    test("fails to create a user with missing name", async () => {
      // biome-ignore lint/suspicious/noExplicitAny: Testing invalid params
      const createParams = { email: "waylon.smithers@snpp.com" } as any;
      await expect(createUser(createParams)).rejects.toThrow();
    });

    test("fails to create a user when unauthorized", async () => {
      localStorage.clear(); // Remove the token
      const createParams = { name: "Waylon", email: "waylon.smithers@snpp.com" };
      await expect(createUser(createParams)).rejects.toThrow();
    });
  });
});
