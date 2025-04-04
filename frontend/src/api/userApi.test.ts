import { beforeEach, describe, expect, test } from "vitest";

import { createUser, fetchCurrentUser } from "@/api/userApi";

describe("User API", () => {
  beforeEach(() => {
    localStorage.clear();
    localStorage.setItem("still-kicking-auth-token", "FAKE-TOKEN");
  });

  describe.skip("fetchCurrentUser", () => {
    test("retrieves the current user", async () => {
      localStorage.setItem("still-kicking-auth-token", "FAKE-TOKEN");

      const result = await fetchCurrentUser();
      expect(result.email).to.equal("montgomery.burns@snpp.com");
      expect(result.name).to.equal("Monty");
    });

    test("fails to retrieve the current user without a token", async () => {
      localStorage.clear();
      
      try {
        await fetchCurrentUser();
      } catch (error) {
        if (error instanceof Error) {
          expect(error.message).to.equal("An unexpected error occurred");
        } else {
          throw error;
        }
      }
    });
  });

  describe("createUser", () => {
    test("creates the current user", async () => {
      const createParams = { name: "Waylon", email: "waylon.smithers@snpp.com" };
      const result = await createUser(createParams);

      expect(result.name).to.equal("Waylon");
      expect(result.email).to.equal("waylon.smithers@snpp.com");
    });
  });
});
